package utils

import (
	"fmt"
	"net/http"

	"github.com/hermsi1337/go-magento2/pkg/magento2"
	"github.com/pkg/errors"
	"gopkg.in/resty.v1"
)

var ErrBadRequest = fmt.Errorf("%s", "bad request")
var ErrInternalServer = fmt.Errorf("%s", "internal server error")
var ErrExecution = fmt.Errorf("%s", "failed while calling endpoint")

func wrapError(err error, triedTo string, response ...map[string]interface{}) error {
	if len(response) == 0 {
		return fmt.Errorf("error while trying to %w - %s", err, triedTo)
	}
	return fmt.Errorf("error while trying to %w - %s. %+v", err, triedTo, response)
}

func MayReturnErrorForHTTPResponse(err error, resp *resty.Response, triedTo string) error {
	if err != nil {
		err = wrapError(err, triedTo)
	} else if resp.StatusCode() == http.StatusNotFound {
		err = magento2.ErrNotFound
	} else if resp.StatusCode() >= http.StatusInternalServerError {
		err = wrapError(ErrInternalServer, triedTo)
	} else if resp.StatusCode() >= http.StatusBadRequest {
		additional := map[string]interface{}{
			"statusCode": resp.StatusCode(),
			"response":   string(resp.Body()),
		}
		err = wrapError(ErrBadRequest, triedTo, additional)
	}

	return errors.WithStack(err)
}
