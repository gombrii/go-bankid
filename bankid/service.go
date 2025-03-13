package bankid

import (
	"context"
	"fmt"

	"github.com/gomsim/go-bankid/internal/client"
)

type Client interface {
	Auth(context.Context, client.AuthReq) (client.AuthResp, error)
	Sign(context.Context, client.SignReq) (client.SignResp, error)
	//Payment(context.Context, client.AuthReq) (client.SignResp, error)
	//PhoneAuth(context.Context, client.AuthReq) (client.SignResp, error)
	//PhoneSign(context.Context, client.AuthReq) (client.SignResp, error)
	Collect(context.Context, string) (client.CollectResp, error)
	Cancel(context.Context, string) error
}

type BankID struct {
	client Client
}

func New( /*certifikat och nån config-fil*/ ) *BankID {
	return &BankID{}
}

func (b BankID) Auth(ctx context.Context, endUserIP string, opts *AuthOpts) (AuthResp, error) {
	req := client.AuthReq{
		EndUserIP: endUserIP, //TODO: Om App är identisk mellan service och client, kan man då helt enkelt typecasta dem?
		App: &client.App{
			AppIdentifier:    opts.App.AppIdentifier,
			DeviceOS:         opts.App.DeviceOS,
			DeviceIdentifier: opts.App.DeviceIdentifier,
			DeviceModelName:  opts.App.DeviceModelName,
		},
		ReturnRisk:            opts.ReturnRisk,
		ReturnURL:             opts.ReturnURL,
		UserNonVisibleData:    opts.UserNonVisibleData,
		UserVisibleData:       opts.UserVisibleData,
		UserVisibleDataFormat: opts.UserVisibleDataFormat,
		Web: &client.Web{
			DeviceIdentifier: opts.Web.DeviceIdentifier,
			ReferringDomain:  opts.Web.ReferringDomain,
			UserAgent:        opts.Web.UserAgent,
		},
		Requirement: &client.Requirement{
			CardReader:          opts.Requirement.CardReader,
			CertificatePolicies: opts.Requirement.CertificatePolicies,
			Mrtd:                opts.Requirement.Mrtd,
			PersonalNumber:      opts.Requirement.PersonalNumber,
			PinCode:             opts.Requirement.PinCode,
		},
	}

	resp, err := b.client.Auth(ctx, req)
	if err != nil {
		return AuthResp{}, fmt.Errorf("BankID: %v", err)
	}

	return AuthResp{
		OrderRef:       resp.OrderRef,
		AutoStartToken: resp.AutoStartToken,
		QRStartToken:   resp.QRStartToken,
		QRStartSecret:  resp.QRStartSecret,
	}, nil
}

func (b BankID) Sign(ctx context.Context) {

}

func (b BankID) Payment(ctx context.Context) {

}

func (b BankID) PhoneAuth(ctx context.Context) {

}
func (b BankID) PhoneSign(ctx context.Context) {

}

func (b BankID) Collect(ctx context.Context, orderRef string) (CollectResp, error) {
	resp, err := b.client.Collect(ctx, orderRef)
	if err != nil {
		return CollectResp{}, fmt.Errorf("BankID: %v", err)
	}

	return CollectResp{
		OrderRef:       resp.OrderRef,
		Status:         resp.Status,
		HintCode:       resp.HintCode,
		CompletionData: (*CompletionData)(resp.CompletionData), //TODO: Kan man göra såhär?
	}, nil

}
func (b BankID) Cancel(ctx context.Context, orderRef string) error {
	if err := b.client.Cancel(ctx, orderRef); err != nil {
		return fmt.Errorf("BankID: %v", err)
	}

	return nil
}
