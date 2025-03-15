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

type BankIDClient struct {
	http *http.Client
	url  string
}

type Config struct {
	CA         []byte
	ClientCert []byte
	URL        string
}

//TODO: Either test this package from the inside or create two separate constructor functions, one of which takes in the http.Client from the outside so that it's mockable.
//TODO: Make it configurable wether or not validation shoul be performed.

func New(cfg Config) (BankIDClient, error) { //TODO: Denna är nog halvfärdig
	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(cfg.CA); !ok {
		return BankIDClient{}, errors.New("error parsing CA cert from PEM")
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

	return BankIDClient{
		http: &client,
		url:  cfg.URL,
	}, nil
}

func (c BankIDClient) Auth(ctx context.Context, endUserIP string, opts *AuthOpts) (AuthResp, error) {
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

	resp := AuthResp{}
	err := c.send(ctx, fmt.Sprint(c.url, "/rp/v6.0/auth"), req, resp)
	if err != nil {
		return AuthResp{}, fmt.Errorf("bankid: %v", err)
	}

	return resp, nil
}

func (c BankIDClient) Sign(ctx context.Context, endUserIP string, userVisibleData string, opts *SignOpts) (SignResp, error) {
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

	resp := SignResp{}
	err := c.send(ctx, fmt.Sprint(c.url, "/rp/v6.0/sign"), req, resp)
	if err != nil {
		return SignResp{}, fmt.Errorf("bankid: %v", err)
	}

	return resp, nil
}

func (c BankIDClient) Payment(ctx context.Context, endUserIP string, userVisibleTransaction UserVisibleTransaction, opts *PaymentOpts) (PaymentResp, error) {
	//TODO: Glöm inte validerin här!! :)

	if opts == nil {
		opts = &PaymentOpts{}
	}

	req := paymentReq{
		EndUserIP:              endUserIP,
		App:                    opts.App,
		ReturnRisk:             opts.ReturnRisk,
		ReturnURL:              opts.ReturnURL,
		UserNonVisibleData:     opts.UserNonVisibleData,
		UserVisibleData:        opts.UserVisibleData,
		UserVisibleDataFormat:  opts.UserVisibleDataFormat,
		UserVisibleTransaction: userVisibleTransaction,
		Web:                    opts.Web,
		Requirement:            opts.Requirement,
		RiskFlags:              opts.RiskFlags,
	}

	resp := PaymentResp{}
	err := c.send(ctx, fmt.Sprint(c.url, "/rp/v6.0/payment"), req, resp)
	if err != nil {
		return PaymentResp{}, fmt.Errorf("bankid: %v", err)
	}

	return resp, nil
}

func (c BankIDClient) PhoneAuth(ctx context.Context, callInitiator CallInitiator, opts *PhoneAuthOpts) (PhoneAuthResp, error) {
	//TODO: Glöm inte validerin här!! :)

	if opts == nil {
		opts = &PhoneAuthOpts{}
	}

	req := phoneAuthReq{
		CallInitiator:         callInitiator,
		PersonalNumber:        opts.PersonalNumber,
		UserNonVisibleData:    opts.UserNonVisibleData,
		UserVisibleData:       opts.UserVisibleData,
		UserVisibleDataFormat: opts.UserVisibleDataFormat,
		Requirement:           opts.Requirement,
	}

	resp := PhoneAuthResp{}
	err := c.send(ctx, fmt.Sprint(c.url, "/rp/v6.0/phone/auth"), req, resp)
	if err != nil {
		return PhoneAuthResp{}, fmt.Errorf("bankid: %v", err)
	}

	return resp, nil
}

func (c BankIDClient) PhoneSign(ctx context.Context, callInitiator CallInitiator, userVisibleData string, opts *PhoneSignOpts) (PhoneSignResp, error) {
	//TODO: Glöm inte validerin här!! :)

	if opts == nil {
		opts = &PhoneSignOpts{}
	}

	req := phoneSignReq{
		CallInitiator:         callInitiator,
		PersonalNumber:        opts.PersonalNumber,
		UserNonVisibleData:    opts.UserNonVisibleData,
		UserVisibleData:       userVisibleData,
		UserVisibleDataFormat: opts.UserVisibleDataFormat,
		Requirement:           opts.Requirement,
	}

	resp := PhoneSignResp{}
	err := c.send(ctx, fmt.Sprint(c.url, "/rp/v6.0/phone/sign"), req, resp)
	if err != nil {
		return PhoneSignResp{}, fmt.Errorf("bankid: %v", err)
	}

	return resp, nil
}

func (c BankIDClient) Collect(ctx context.Context, orderRef string) (CollectResp, error) {
	//TODO: Glöm inte validerin här!! :)

	resp := CollectResp{}
	err := c.send(ctx, fmt.Sprint(c.url, "/rp/v6.0/collect"), collectReq{OrderRef: orderRef}, resp)
	if err != nil {
		return CollectResp{}, fmt.Errorf("bankid: %v", err)
	}

	return resp, nil
}

func (c BankIDClient) Cancel(ctx context.Context, orderRef string) error {
	//TODO: Glöm inte validerin här!! :)

	err := c.send(ctx, fmt.Sprint(c.url, "/rp/v6.0/cancel"), cancelReq{OrderRef: orderRef}, nil)
	if err != nil {
		return fmt.Errorf("bankid: %v", err)
	}

	return nil
}

func (c BankIDClient) send(ctx context.Context, url string, req any, resp any) error {
	httpResp, err := sendReq(ctx, c.http, url, req)
	if err != nil {
		return err
	}

	return unmarshalResp(httpResp, resp)
}

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
