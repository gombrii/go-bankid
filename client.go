// © 2025 Simon Oscar Gombrii. Released under the MIT License.

// TODO: Document package
package bankid

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"net/http"
)

// The currently supported version of the BankID API. See https://developers.bankid.com/api-references/auth--sign/overview for more information.
const SupportedAPIVersion = "v6.0"

type BankIDClient struct {
	*http.Client
	URL string
}

// A Config is used to configure a [bankIDClient].
type Config struct {
	RootCA     []byte // The root certificate authority of the BankID service.
	ClientCert []byte // The certificate of your client. Ordered from the bank you have your BankID agreement with.
	ClientKey  []byte // Your client certificate's private key.
}

// New creates a new [bankIDClient] that can be used with production BankIDs.
func NewProd(cfg Config) (BankIDClient, error) {
	return newClient(cfg, ProdURL)
}	

// New creates a new [bankIDClient] that can be used with test BankIDs.
func NewTest(cfg Config) (BankIDClient, error) {
	return newClient(cfg, TestURL)
}

func newClient(cfg Config, url string) (BankIDClient, error) {
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
		Client: &client,
		URL:    url,
	}, nil
}

// Auth initiates an identification order.
//
// endUserIP is mandatory and is the user IP address as it is seen by your service. IPv4 and IPv6 are allowed. Make sure that the IP address you include as endUserIp is the address of your end user's device, not the internal address of any reverse proxy between you and the end user. In use cases where the IP address is not available, e. g. for voice-based services, the internal representation of those systems' IP address is ok to use.
//
// Use opts to augument your identification order. Otherwise pass nil.
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
	err := c.send(ctx, fmt.Sprintf("%s/rp/%s/auth", c.URL, SupportedAPIVersion), req, resp)
	if err != nil {
		return AuthResp{}, fmt.Errorf("bankid: %v", err)
	}

	return resp, nil
}

// Sign initiates an signing order.
//
// endUserIP is mandatory and is the user IP address as it is seen by your service. IPv4 and IPv6 are allowed. Make sure that the IP address you include as endUserIp is the address of your end user's device, not the internal address of any reverse proxy between you and the end user. In use cases where the IP address is not available, e. g. for voice-based services, the internal representation of those systems' IP address is ok to use.
//
// userVisibleData is mandatory and is the text displayed to the user during the order. The purpose is to provide context, thereby enabling the user to detect identification errors and avert fraud attempts. The text can be formatted using CR, LF and CRLF for new lines. The text must be encoded as UTF-8 and then base 64 encoded.
//
// Use opts to augument your identification order. Otherwise pass nil.
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
	err := c.send(ctx, fmt.Sprintf("%s/rp/%s/sign", c.URL, SupportedAPIVersion), req, resp)
	if err != nil {
		return SignResp{}, fmt.Errorf("bankid: %v", err)
	}

	return resp, nil
}

// Payment initiates an payment order.
//
// endUserIP is mandatory and is the user IP address as it is seen by your service. IPv4 and IPv6 are allowed. Make sure that the IP address you include as endUserIp is the address of your end user's device, not the internal address of any reverse proxy between you and the end user. In use cases where the IP address is not available, e. g. for voice-based services, the internal representation of those systems' IP address is ok to use.
//
// userVisibleTransaction is mandatory and contains information about the transaction being approved.
//
// Use opts to augument your identification order. Otherwise pass nil.
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
	err := c.send(ctx, fmt.Sprintf("%s/rp/%s/payment", c.URL, SupportedAPIVersion), req, resp)
	if err != nil {
		return PaymentResp{}, fmt.Errorf("bankid: %v", err)
	}

	return resp, nil
}

// PhoneAuth initiates an phone identification order.
//
// callInitiator is mandatory and indicates if the user or your organization initiated the phone call.
//
// Use opts to augument your identification order. Otherwise pass nil.
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
	err := c.send(ctx, fmt.Sprintf("%s/rp/%s/phone/auth", c.URL, SupportedAPIVersion), req, resp)
	if err != nil {
		return PhoneAuthResp{}, fmt.Errorf("bankid: %v", err)
	}

	return resp, nil
}

// PhoneSign initiates an phone signing order.
//
// callInitiator is mandatory and indicates if the user or your organization initiated the phone call.
//
// userVisibleData is mandatory and is the text displayed to the user during the order. The purpose is to provide context, thereby enabling the user to detect identification errors and avert fraud attempts. The text can be formatted using CR, LF and CRLF for new lines. The text must be encoded as UTF-8 and then base 64 encoded.
//
// Use opts to augument your identification order. Otherwise pass nil.
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
	err := c.send(ctx, fmt.Sprintf("%s/rp/%s/phone/sign", c.URL, SupportedAPIVersion), req, resp)
	if err != nil {
		return PhoneSignResp{}, fmt.Errorf("bankid: %v", err)
	}

	return resp, nil
}

// Collect collects information and status of an ongoing order.
//
// Collects the result of an order using orderRef as reference.
//
// Your service should continue calling collect every two seconds if the status reported is is pending. Your service must abort if the status is failed.
//
// The user identity is returned when complete.
func (c BankIDClient) Collect(ctx context.Context, orderRef string) (CollectResp, error) {
	if orderRef == "" {
		return CollectResp{}, errors.New("orderRef is empty")
	}

	resp := CollectResp{}
	err := c.send(ctx, fmt.Sprintf("%s/rp/%s/collect", c.URL, SupportedAPIVersion), collectReq{OrderRef: orderRef}, resp)
	if err != nil {
		return CollectResp{}, fmt.Errorf("bankid: %v", err)
	}

	return resp, nil
}

// Cancel cancels an ongoing signature, authentication or payment order.
//
// This is typically used if the user cancels the order in your service or app.
func (c BankIDClient) Cancel(ctx context.Context, orderRef string) error {
	if orderRef == "" {
		return errors.New("orderRef is empty")
	}

	err := c.send(ctx, fmt.Sprintf("%s/rp/%s/cancel", c.URL, SupportedAPIVersion), cancelReq{OrderRef: orderRef}, nil)
	if err != nil {
		return fmt.Errorf("bankid: %v", err)
	}

	return nil
}

func (c BankIDClient) send(ctx context.Context, url string, req any, resp any) error {
	httpResp, err := sendReq(ctx, c.Client, url, req)
	if err != nil {
		return err
	}

	return unmarshalResp(httpResp, resp)
}
