package bankid

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"net/http"
)

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
//TODO: Make it configurable wether or not validation shoul be performed.
//TODO: Add documentation to all functions and types. The type documentation can be copied straight from developers.bankid.com
//TODO: Hur bör man tänka kring versionsnummer. En logisk versionering hade ju varit att exakt följa API:ts versionering, men frågan är om det blir för strikt.
// 		Risken FINNS ju att en icke-brytande förändring i API:t leder till en brytande förändring i detta bibliotek, t.ex. en ny obligatorisk parameter.
// 		Kika på General rules > Breaking changes för att resonera kring detta designval

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
	//TODO: Glöm inte validering här!! :)

	if opts == nil {
		opts = &AuthOpts{}
	}

	req := authReq{
		EndUserIP:             endUserIP,
		App:                   opts.App,
		ReturnRisk:            opts.ReturnRisk,
		ReturnURL:             opts.ReturnURL,
		UserNonVisibleData:    opts.UserNonVisibleData, //TODO: byte64-koda åt användaren
		UserVisibleData:       opts.UserVisibleData,    //TODO: byte64-koda åt användaren
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
		UserNonVisibleData:    opts.UserNonVisibleData, //TODO: byte64-koda åt användaren
		UserVisibleData:       userVisibleData,         //TODO: byte64-koda åt användaren
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
		UserNonVisibleData:     opts.UserNonVisibleData, //TODO: byte64-koda åt användaren
		UserVisibleData:        opts.UserVisibleData,    //TODO: byte64-koda åt användaren
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
		UserNonVisibleData:    opts.UserNonVisibleData, //TODO: byte64-koda åt användaren
		UserVisibleData:       opts.UserVisibleData,    //TODO: byte64-koda åt användaren
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
		UserNonVisibleData:    opts.UserNonVisibleData, //TODO: byte64-koda åt användaren
		UserVisibleData:       userVisibleData,         //TODO: byte64-koda åt användaren
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
