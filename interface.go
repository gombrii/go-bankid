package bankid

import "context"

type Client interface {
	Auth(context.Context, string, *AuthOpts) (AuthResp, error)
	Sign(context.Context, string, string, *SignOpts) (SignResp, error)
	//Payment(context.Context, client.AuthReq) (client.SignResp, error)
	//PhoneAuth(context.Context, client.AuthReq) (client.SignResp, error)
	//PhoneSign(context.Context, client.AuthReq) (client.SignResp, error)
	Collect(context.Context, string) (CollectResp, error)
	Cancel(context.Context, string) error
}
