package utils

import (
	"fmt"
	"gopkg.in/resty.v1"
)

func MayReturnErrorForHTTPResponse(err error, resp *resty.Response, triedTo string) error {
	commonString := fmt.Sprintf(" while trying to " + triedTo + " ")
	if err != nil {
		return fmt.Errorf("error"+commonString+"- error: '%+v'", err)
	} else if resp.StatusCode() >= 400 {
		return fmt.Errorf("unexpected statuscode"+commonString+"- statuscode: '%d' - response: '%+v'", resp.StatusCode(), resp)
	}

	return nil
}
