package magento2

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"gopkg.in/resty.v1"
)

var ErrBadRequest = fmt.Errorf("%s", "bad request")

func wrapError(err error, triedTo string, response ...map[string]interface{}) error {
	if len(response) == 0 {
		return fmt.Errorf("error while trying to %w - %s", err, triedTo)
	}
	return fmt.Errorf("error while trying to %w - %s. %+v", err, triedTo, response)
}

func mayReturnErrorForHTTPResponse(err error, resp *resty.Response, triedTo string) error {
	if err != nil {
		err = wrapError(err, triedTo)
	} else if resp.StatusCode() == http.StatusNotFound {
		err = ErrNotFound
	} else if resp.StatusCode() >= http.StatusBadRequest {
		additional := map[string]interface{}{
			"statusCode": resp.StatusCode(),
			"response":   string(resp.Body()),
		}
		err = wrapError(ErrBadRequest, triedTo, additional)
	}

	return errors.WithStack(err)
}

func mayTrimSurroundingQuotes(s string) string {
	minQuotes := 2
	if len(s) >= minQuotes {
		if s[0] == '"' && s[len(s)-1] == '"' {
			return s[1 : len(s)-1]
		}
	}
	return s
}
