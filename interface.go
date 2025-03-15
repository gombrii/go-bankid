package bankid

import "context"

type Client interface {
	Auth(context.Context, string, *AuthOpts) (AuthResp, error)
	Sign(context.Context, string, string, *SignOpts) (SignResp, error)
	Payment(context.Context, string, UserVisibleTransaction, *PaymentOpts) (PaymentResp, error)
	PhoneAuth(context.Context, CallInitiator, *PhoneAuthOpts) (PhoneAuthResp, error)
	PhoneSign(context.Context, CallInitiator, string, *PhoneSignOpts) (PhoneSignResp, error)
	Collect(context.Context, string) (CollectResp, error)
	Cancel(context.Context, string) error
}
