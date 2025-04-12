// © 2025 Simon Oscar Gombrii. Released under the MIT License.

package bankid

import "context"

type Client interface { // TODO: Should I even document this??
	// Create an identification order.
	Auth(ctx context.Context, endUserIP string, opts *AuthOpts) (AuthResp, error)
	// Create a signature order.
	Sign(ctx context.Context, endUserIP string, userVisibleData string, opts *SignOpts) (SignResp, error)
	// Create a payment order.
	Payment(ctx context.Context, endUserIP string, userVisibleTransaction UserVisibleTransaction, opts *PaymentOpts) (PaymentResp, error)
	// Create an identification order for a phone call.
	PhoneAuth(ctx context.Context, callInitiator CallInitiator, opts *PhoneAuthOpts) (PhoneAuthResp, error)
	// Create a signature order for a phone call.
	PhoneSign(ctx context.Context, callInitiator CallInitiator, userVisibleData string, opts *PhoneSignOpts) (PhoneSignResp, error)
	// Gather order information and status.
	Collect(ctx context.Context, orderRef string) (CollectResp, error)
	// Cancel a signature, authentication or payment order.
	Cancel(ctx context.Context, orderRef string) error
}
