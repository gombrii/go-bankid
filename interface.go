// © 2025 Simon Oscar Gombrii. Released under the MIT License.

package bid

import "context"

type Client interface {
	Auth(ctx context.Context, endUserIP string, opts *AuthOpts) (AuthResp, error)
	Sign(ctx context.Context, endUserIP string, userVisibleData string, opts *SignOpts) (SignResp, error)
	Payment(ctx context.Context, endUserIP string, userVisibleTransaction UserVisibleTransaction, opts *PaymentOpts) (PaymentResp, error)
	PhoneAuth(ctx context.Context, callInitiator CallInitiator, opts *PhoneAuthOpts) (PhoneAuthResp, error)
	PhoneSign(ctx context.Context, callInitiator CallInitiator, userVisibleData string, opts *PhoneSignOpts) (PhoneSignResp, error)
	Collect(ctx context.Context, orderRef string) (CollectResp, error)
	Cancel(ctx context.Context, orderRef string) error
}
