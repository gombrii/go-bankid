package bankid

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func sendReq(ctx context.Context, client *http.Client, url string, req any) (*http.Response, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshalling request body: %v", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("preparing request: %v", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("sending request: %v", err)
	}

	return resp, nil
}

func unmarshalResp(resp *http.Response, dst any) error {
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %v", err)
	}
	defer resp.Body.Close()

	switch {
	case resp.StatusCode == 400:
		respErr := err400{}
		if err = json.Unmarshal(respBody, &respErr); err != nil {
			return fmt.Errorf("unmarshalling response body: %v", err)
		}
		return respErr
	case resp.StatusCode > 400:
		return errors.New(string(respBody))
	case resp.ContentLength == 0: // for cancel response
		return nil
	default:
		if err = json.Unmarshal(respBody, dst); err != nil {
			return fmt.Errorf("unmarshalling response body: %v", err)
		}
		return nil
	}
}
