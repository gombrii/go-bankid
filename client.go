package bankid

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
)

type Client struct {
	HTTP *http.Client
	URL  string
}

type Config struct {
	CA         []byte
	ClientCert []byte
	URL        string
}

func New(cfg Config) (Client, error) { //TODO: Denna är nog halvfärdig
	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(cfg.CA); !ok {
		return Client{}, errors.New("error parsing CA cert from PEM")
	}

	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:            certPool,
				InsecureSkipVerify: false,
				Certificates: []tls.Certificate{
					{Certificate: [][]byte{cfg.ClientCert}},
				},
			},
		},
	}

	return Client{
		HTTP: &client,
		URL:  cfg.URL,
	}, nil
}

func (c Client) Auth(ctx context.Context, endUserIP string, opts *AuthOpts) (AuthResp, error) {
	//TODO: Glöm inte validering här!! :)

	if opts == nil {
		opts = &AuthOpts{}
	}

	req := authReq{
		EndUserIP:             endUserIP,
		App:                   opts.App,
		ReturnRisk:            opts.ReturnRisk,
		ReturnURL:             opts.ReturnURL,
		UserNonVisibleData:    opts.UserNonVisibleData,
		UserVisibleData:       opts.UserVisibleData,
		UserVisibleDataFormat: opts.UserVisibleDataFormat,
		Web:                   opts.Web,
		Requirement:           opts.Requirement,
	}

	httpResp, err := c.sendReq(ctx, fmt.Sprint(c.URL, "%s/rp/v6.0/auth"), req)
	if err != nil {
		return AuthResp{}, fmt.Errorf("bankid: %v", err)
	}

	resp := AuthResp{}
	if err := unmarshalResp(httpResp, resp); err != nil {
		return AuthResp{}, fmt.Errorf("bankid: %v", err)
	}

	return resp, nil
}

func (c Client) Sign(ctx context.Context, endUserIP string, userVisibleData string, opts *SignOpts) (SignResp, error) {
	//TODO: Glöm inte validerin här!! :)

	if opts == nil {
		opts = &SignOpts{}
	}

	req := signReq{
		EndUserIP:             endUserIP,
		App:                   opts.App,
		ReturnRisk:            opts.ReturnRisk,
		ReturnURL:             opts.ReturnURL,
		UserNonVisibleData:    opts.UserNonVisibleData,
		UserVisibleData:       userVisibleData,
		UserVisibleDataFormat: opts.UserVisibleDataFormat,
		Web:                   opts.Web,
		Requirement:           opts.Requirement,
	}

	httpResp, err := c.sendReq(ctx, fmt.Sprint(c.URL, "/rp/v6.0/sign"), req)
	if err != nil {
		return SignResp{}, fmt.Errorf("bankid: %v", err)
	}

	resp := SignResp{}
	if err := unmarshalResp(httpResp, resp); err != nil {
		return SignResp{}, fmt.Errorf("bankid: %v", err)
	}

	return resp, nil
}

func (c Client) Payment() {
	panic("UNIMPLEMENTED")
}

func (c Client) PhoneAuth() {
	panic("UNIMPLEMENTED")
}

func (c Client) PhoneSign() {
	panic("UNIMPLEMENTED")
}

func (c Client) Collect(ctx context.Context, orderRef string) (CollectResp, error) {
	//TODO: Glöm inte validerin här!! :)
	httpResp, err := c.sendReq(ctx, fmt.Sprint(c.URL, "/rp/v6.0/collect"), collectReq{OrderRef: orderRef})
	if err != nil {
		return CollectResp{}, fmt.Errorf("bankid: %v", err)
	}

	resp := CollectResp{}
	if err := unmarshalResp(httpResp, resp); err != nil {
		return CollectResp{}, fmt.Errorf("bankid: %v", err)
	}

	return resp, nil
}

func (c Client) Cancel(ctx context.Context, orderRef string) error {
	//TODO: Glöm inte validerin här!! :)
	httpResp, err := c.sendReq(ctx, fmt.Sprint(c.URL, "/rp/v6.0/cancel"), cancelReq{OrderRef: orderRef})
	if err != nil {
		return fmt.Errorf("bankid: %v", err)
	}

	resp := CollectResp{}
	if err := unmarshalResp(httpResp, resp); err != nil {
		return fmt.Errorf("bankid: %v", err)
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

	resp, err := c.HTTP.Do(httpReq)
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
