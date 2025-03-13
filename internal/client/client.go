package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
	bidURL     string
}

type Config struct {
	ca         []byte
	clientCert []byte
	Timeout    time.Duration
}

func New(cfg Config) (Client, error) { //TODO: Denna är nog halvfärdig
	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(cfg.ca); !ok {
		return Client{}, errors.New("error parsing CA cert from PEM")
	}

	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:            certPool,
				InsecureSkipVerify: false,
				Certificates: []tls.Certificate{{
					Certificate: [][]byte{cfg.clientCert},
				}},
			},
		},
		Timeout: cfg.Timeout,
	}

	return Client{httpClient: &client}, nil
}

func (c Client) Auth(ctx context.Context, req AuthReq) (AuthResp, error) {
	httpResp, err := c.sendReq(ctx, fmt.Sprint(c.bidURL, "%s/rp/v6.0/auth"), req)
	if err != nil {
		return AuthResp{}, err
	}

	resp := AuthResp{}
	if err := unmarshalResp(httpResp, resp); err != nil {
		return AuthResp{}, err
	}

	return resp, nil
}

func (c Client) Sign(ctx context.Context, req SignReq) (SignResp, error) {
	httpResp, err := c.sendReq(ctx, fmt.Sprint(c.bidURL, "/rp/v6.0/sign"), req)
	if err != nil {
		return SignResp{}, err
	}

	resp := SignResp{}
	if err := unmarshalResp(httpResp, resp); err != nil {
		return SignResp{}, err
	}

	return resp, nil
}

func (c Client) Payment() {

}

func (c Client) PhoneAuth() {

}

func (c Client) PhoneSign() {

}

func (c Client) Collect(ctx context.Context, orderRef string) (CollectResp, error) {
	httpResp, err := c.sendReq(ctx, fmt.Sprint(c.bidURL, "/rp/v6.0/collect"), CollectReq{OrderRef: orderRef})
	if err != nil {
		return CollectResp{}, err
	}

	resp := CollectResp{}
	if err := unmarshalResp(httpResp, resp); err != nil {
		return CollectResp{}, err
	}

	return resp, nil
}

func (c Client) Cancel(ctx context.Context, orderRef string) error {
	httpResp, err := c.sendReq(ctx, fmt.Sprint(c.bidURL, "/rp/v6.0/cancel"), CancelReq{OrderRef: orderRef})
	if err != nil {
		return err
	}

	resp := CollectResp{}
	if err := unmarshalResp(httpResp, resp); err != nil {
		return err
	}

	return nil
}

func (c Client) sendReq(ctx context.Context, url string, req any) (*http.Response, error) { //TODO: Slå ihop sendReq och unmarshalResp????
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshalling request body: %v", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(body)) //TODO: NewReader eller NewBuffer?
	if err != nil {
		return nil, fmt.Errorf("preparing request: %v", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("sending request: %v", err)
	}

	return resp, nil
}

// TODO: Kontrollera att alla anrops response kan använda denna funktion. Det hänger på att alla har samma format på fel-responsen (>= 400)
func unmarshalResp(resp *http.Response, dst any) error { //TODO: Slå ihop sendReq och unmarshalResp????
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %v", err)
	}
	defer resp.Body.Close()

	switch {
	case resp.StatusCode == 400:
		respErr := Err400{}
		if err = json.Unmarshal(body, &respErr); err != nil {
			return fmt.Errorf("unmarshalling response body: %v", err)
		}
		return respErr
	case resp.StatusCode > 400:
		return errors.New(string(body))
	default:
		if err = json.Unmarshal(body, dst); err != nil {
			return fmt.Errorf("unmarshalling response body: %v", err)
		}
		return nil
	}
}
