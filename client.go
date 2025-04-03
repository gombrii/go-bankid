package bankid

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"net/http"
)

const SupportedVersion = "v6.0"

type BankIDClient struct {
	http *http.Client
	url  string
}

type Config struct {
	URL        string
	RootCA     []byte
	ClientCert []byte
	ClientKey  []byte
}

//TODO: Either test this package from the inside or create two separate constructor functions, one of which takes in the http.Client from the outside so that it's mockable.
//TODO: Add documentation to all functions and types. The type documentation can be copied straight from developers.bankid.com
//TODO: Versionsnummer bör vara HELT separata från CAVAs versionsnummer, men vilken cavaversion som stöds bör vara uppenbart i README

func New(cfg Config) (BankIDClient, error) {
	rootCAs := x509.NewCertPool()
	if ok := rootCAs.AppendCertsFromPEM(cfg.RootCA); !ok {
		return BankIDClient{}, errors.New("error parsing CA cert from PEM")
	}

	cert, err := tls.X509KeyPair(cfg.ClientCert, cfg.ClientKey)
	if err != nil {
		return BankIDClient{}, fmt.Errorf("failed to load client certificate/key: %w", err)
	}

	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      rootCAs,
				Certificates: []tls.Certificate{cert},
				MinVersion:   tls.VersionTLS13,
			},
		},
	}

	return BankIDClient{
		http: &client,
		url:  cfg.URL,
	}, nil
}

func (c BankIDClient) Auth(ctx context.Context, endUserIP string, opts *AuthOpts) (AuthResp, error) {
	if endUserIP == "" {
		return AuthResp{}, errors.New("endUserIP is empty")
	}

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
	err := c.send(ctx, fmt.Sprintf("%s/rp/%s/auth", c.url, SupportedVersion), req, resp)
	if err != nil {
		return AuthResp{}, fmt.Errorf("bankid: %v", err)
	}

	return resp, nil
}

func (c BankIDClient) Sign(ctx context.Context, endUserIP string, userVisibleData string, opts *SignOpts) (SignResp, error) {
	if endUserIP == "" {
		return SignResp{}, errors.New("endUserIP is empty")
	}
	if userVisibleData == "" {
		return SignResp{}, errors.New("userVisibleData is empty")
	}

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
	err := c.send(ctx, fmt.Sprintf("%s/rp/%s/sign", c.url, SupportedVersion), req, resp)
	if err != nil {
		return SignResp{}, fmt.Errorf("bankid: %v", err)
	}

	return resp, nil
}

func (c BankIDClient) Payment(ctx context.Context, endUserIP string, userVisibleTransaction UserVisibleTransaction, opts *PaymentOpts) (PaymentResp, error) {
	if endUserIP == "" {
		return PaymentResp{}, errors.New("endUserIP is empty")
	}
	if userVisibleTransaction.TransactionType == "" {
		return PaymentResp{}, errors.New("userVisibleTransaction.TransactionType is empty")
	}
	if userVisibleTransaction.Recipient.Name == "" {
		return PaymentResp{}, errors.New("userVisibleTransaction.Recipient.Name is empty")
	}

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
	err := c.send(ctx, fmt.Sprintf("%s/rp/%s/payment", c.url, SupportedVersion), req, resp)
	if err != nil {
		return PaymentResp{}, fmt.Errorf("bankid: %v", err)
	}

	return resp, nil
}

func (c BankIDClient) PhoneAuth(ctx context.Context, callInitiator CallInitiator, opts *PhoneAuthOpts) (PhoneAuthResp, error) {
	if callInitiator == "" {
		return PhoneAuthResp{}, errors.New("callInitiator is empty")
	}

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
	err := c.send(ctx, fmt.Sprintf("%s/rp/%s/phone/auth", c.url, SupportedVersion), req, resp)
	if err != nil {
		return PhoneAuthResp{}, fmt.Errorf("bankid: %v", err)
	}

	return resp, nil
}

func (c BankIDClient) PhoneSign(ctx context.Context, callInitiator CallInitiator, userVisibleData string, opts *PhoneSignOpts) (PhoneSignResp, error) {
	if callInitiator == "" {
		return PhoneSignResp{}, errors.New("callInitiator is empty")
	}
	if userVisibleData == "" {
		return PhoneSignResp{}, errors.New("userVisibleData is empty")
	}

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
	err := c.send(ctx, fmt.Sprintf("%s/rp/%s/phone/sign", c.url, SupportedVersion), req, resp)
	if err != nil {
		return PhoneSignResp{}, fmt.Errorf("bankid: %v", err)
	}

	return resp, nil
}

func (c BankIDClient) Collect(ctx context.Context, orderRef string) (CollectResp, error) {
	if orderRef == "" {
		return CollectResp{}, errors.New("orderRef is empty")
	}

	resp := CollectResp{}
	err := c.send(ctx, fmt.Sprintf("%s/rp/%s/collect", c.url, SupportedVersion), collectReq{OrderRef: orderRef}, resp)
	if err != nil {
		return CollectResp{}, fmt.Errorf("bankid: %v", err)
	}

	return resp, nil
}

func (c BankIDClient) Cancel(ctx context.Context, orderRef string) error {
	if orderRef == "" {
		return errors.New("orderRef is empty")
	}

	err := c.send(ctx, fmt.Sprintf("%s/rp/%s/cancel", c.url, SupportedVersion), cancelReq{OrderRef: orderRef}, nil)
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
