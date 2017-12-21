package scalar

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/matrix-org/gomatrix"
)

//MakeRequest does the same as cli.MakeRequest but with a string as body instead of a struct
func MakeRequest(cli *gomatrix.Client, method string, httpURL string, reqBody string, resBody interface{}) ([]byte, error) {
	var req *http.Request
	var err error
	if reqBody != "" {
		byteString := []byte(reqBody)
		req, err = http.NewRequest(method, httpURL, bytes.NewBuffer(byteString))
	} else {
		req, err = http.NewRequest(method, httpURL, nil)
	}

	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := cli.Client.Do(req)
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	contents, err := ioutil.ReadAll(res.Body)
	if res.StatusCode/100 != 2 { // not 2xx
		var wrap error
		var respErr gomatrix.RespError
		if _ = json.Unmarshal(contents, &respErr); respErr.ErrCode != "" {
			wrap = respErr
		}

		// If we failed to decode as RespError, don't just drop the HTTP body, include it in the
		// HTTP error instead (e.g proxy errors which return HTML).
		msg := "Failed to " + method + " JSON to " + req.URL.Path
		if wrap == nil {
			msg = msg + ": " + string(contents)
		}

		return contents, gomatrix.HTTPError{
			Code:         res.StatusCode,
			Message:      msg,
			WrappedError: wrap,
		}
	}
	if err != nil {
		return nil, err
	}

	if resBody != nil {
		if err = json.Unmarshal(contents, &resBody); err != nil {
			return nil, err
		}
	}

	return contents, nil
}
