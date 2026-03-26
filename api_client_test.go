package magento2

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// tokenClaims mirrors the structure of a Magento 2 admin JWT.
type tokenClaims struct {
	UID    int   `json:"uid"`
	UTypID int   `json:"utypid"`
	IAT    int64 `json:"iat"`
	EXP    int64 `json:"exp"`
}

func buildTestJWT(t *testing.T, claims tokenClaims) string {
	t.Helper()

	header := base64.RawURLEncoding.EncodeToString([]byte(`{"kid":"1","alg":"HS256"}`))

	claimsJSON, err := json.Marshal(claims)
	require.NoError(t, err, "failed to marshal claims")

	payload := base64.RawURLEncoding.EncodeToString(claimsJSON)

	return fmt.Sprintf("%s.%s.fakesignature", header, payload)
}

// --- Suite ---

type TokenRefreshSuite struct {
	suite.Suite
	validToken    string
	validClaims   tokenClaims
	expiredToken  string
	expiredClaims tokenClaims
}

func (s *TokenRefreshSuite) SetupTest() {
	now := time.Now()

	// Valid token: issued now, expires in 1 hour.
	s.validClaims = tokenClaims{
		UID:    36,
		UTypID: 2,
		IAT:    now.Unix(),
		EXP:    now.Add(1 * time.Hour).Unix(),
	}
	s.validToken = buildTestJWT(s.T(), s.validClaims)

	// Expired token: issued 2 hours ago, expired 1 hour ago.
	s.expiredClaims = tokenClaims{
		UID:    36,
		UTypID: 2,
		IAT:    now.Add(-2 * time.Hour).Unix(),
		EXP:    now.Add(-1 * time.Hour).Unix(),
	}
	s.expiredToken = buildTestJWT(s.T(), s.expiredClaims)
}

func TestTokenRefreshSuite(t *testing.T) {
	suite.Run(t, new(TokenRefreshSuite))
}

// --- Valid token tests ---

func (s *TokenRefreshSuite) TestValidToken_ParsesSuccessfully() {
	refreshAt := parseTokenRefreshAt(s.validToken)
	assert.False(s.T(), refreshAt.IsZero(), "valid token should return non-zero refreshAt")
}

func (s *TokenRefreshSuite) TestValidToken_RefreshAtThreeQuarters() {
	now := time.Now()
	lifetime := time.Duration(s.validClaims.EXP-s.validClaims.IAT) * time.Second
	expectedRefresh := now.Add(time.Until(time.Unix(s.validClaims.EXP, 0)) * 3 / 4)

	refreshAt := parseTokenRefreshAt(s.validToken)

	diff := refreshAt.Sub(expectedRefresh).Abs()
	assert.Less(s.T(), diff, 2*time.Second,
		"refreshAt should be at 3/4 of lifetime (%s), got diff %s", lifetime, diff)
}

func (s *TokenRefreshSuite) TestValidToken_RefreshIsBeforeExpiry() {
	expiry := time.Unix(s.validClaims.EXP, 0)
	refreshAt := parseTokenRefreshAt(s.validToken)

	assert.True(s.T(), refreshAt.Before(expiry),
		"refreshAt (%s) should be before expiry (%s)", refreshAt, expiry)
}

func (s *TokenRefreshSuite) TestValidToken_RefreshIsAfterNow() {
	refreshAt := parseTokenRefreshAt(s.validToken)
	assert.True(s.T(), refreshAt.After(time.Now()),
		"refreshAt should be in the future")
}

// --- Expired token tests ---

func (s *TokenRefreshSuite) TestExpiredToken_ReturnsZero() {
	refreshAt := parseTokenRefreshAt(s.expiredToken)
	assert.True(s.T(), refreshAt.IsZero(),
		"expired token should return zero refreshAt, got %s", refreshAt)
}

// --- Non-JWT tests ---

func (s *TokenRefreshSuite) TestNonJWT_EmptyString() {
	assert.True(s.T(), parseTokenRefreshAt("").IsZero())
}

func (s *TokenRefreshSuite) TestNonJWT_PlainBearerToken() {
	assert.True(s.T(), parseTokenRefreshAt("abc123def456ghi789").IsZero())
}

func (s *TokenRefreshSuite) TestNonJWT_TwoPartsOnly() {
	assert.True(s.T(), parseTokenRefreshAt("header.payload").IsZero())
}

func (s *TokenRefreshSuite) TestNonJWT_InvalidBase64Payload() {
	assert.True(s.T(), parseTokenRefreshAt("header.!!!invalid!!!.sig").IsZero())
}

func (s *TokenRefreshSuite) TestNonJWT_ValidBase64ButNoJSON() {
	payload := base64.URLEncoding.EncodeToString([]byte("not json"))
	token := fmt.Sprintf("header.%s.sig", payload)
	assert.True(s.T(), parseTokenRefreshAt(token).IsZero())
}

func (s *TokenRefreshSuite) TestNonJWT_ValidJWTWithoutExpClaim() {
	payload := base64.RawURLEncoding.EncodeToString([]byte(`{"uid":36,"utypid":2}`))
	token := fmt.Sprintf("eyJhbGciOiJIUzI1NiJ9.%s.sig", payload)
	assert.True(s.T(), parseTokenRefreshAt(token).IsZero())
}

// --- Different lifetimes ---

func (s *TokenRefreshSuite) TestFourHourLifetime_RefreshAtThreeHours() {
	now := time.Now()
	claims := tokenClaims{
		UID:    36,
		UTypID: 2,
		IAT:    now.Unix(),
		EXP:    now.Add(4 * time.Hour).Unix(),
	}
	token := buildTestJWT(s.T(), claims)

	refreshAt := parseTokenRefreshAt(token)
	expectedRefresh := now.Add(3 * time.Hour)

	diff := refreshAt.Sub(expectedRefresh).Abs()
	assert.Less(s.T(), diff, 2*time.Second,
		"4h token should refresh at ~3h, diff was %s", diff)
}

func (s *TokenRefreshSuite) TestShortLifetime_FiveMinutes() {
	now := time.Now()
	claims := tokenClaims{
		UID:    36,
		UTypID: 2,
		IAT:    now.Unix(),
		EXP:    now.Add(5 * time.Minute).Unix(),
	}
	token := buildTestJWT(s.T(), claims)

	refreshAt := parseTokenRefreshAt(token)
	expectedRefresh := now.Add(3*time.Minute + 45*time.Second)

	diff := refreshAt.Sub(expectedRefresh).Abs()
	assert.Less(s.T(), diff, 2*time.Second,
		"5m token should refresh at ~3m45s, diff was %s", diff)
}
