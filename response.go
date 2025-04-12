// © 2025 Simon Oscar Gombrii. Released under the MIT License.

package bankid

import "fmt"

// An AuthResp is a successful response from the creation of an identification order.
type AuthResp struct {
	OrderRef       string `json:"orderRef"`       // A reference ID for an order. This is used to query the status of the order or to cancel it.
	AutoStartToken string `json:"autoStartToken"` // Used to compile the start URL.
	QRStartToken   string `json:"qrStartToken"`   // Used to compute the animated QR code.
	QRStartSecret  string `json:"qrStartSecret"`  // Used to compute the animated QR code.
}

// A SignResp is a successful response from the creation of a signing order.
type SignResp struct {
	OrderRef       string `json:"orderRef"`       // A reference ID for an order. This is used to query the status of the order or to cancel it.
	AutoStartToken string `json:"autoStartToken"` // Used to compile the start URL.
	QRStartToken   string `json:"qrStartToken"`   // Used to compute the animated QR code.
	QRStartSecret  string `json:"qrStartSecret"`  // Used to compute the animated QR code.
}

// A PaymentResp is a successful response from the creation of a payment order.
type PaymentResp struct {
	OrderRef       string `json:"orderRef"`       // A reference ID for an order. This is used to query the status of the order or to cancel it.
	AutoStartToken string `json:"autoStartToken"` // Used to compile the start URL.
	QRStartToken   string `json:"qrStartToken"`   // Used to compute the animated QR code.
	QRStartSecret  string `json:"qrStartSecret"`  // Used to compute the animated QR code.
}

// A PhoneAuthResp is a successful response from the creation of a phone identification order.
type PhoneAuthResp struct {
	OrderRef string `json:"orderRef"` // A reference ID for an order. This is used to query the status of the order or to cancel it.
}

// A PhoneSignResp is a successful response from the creation of a phone signing order.
type PhoneSignResp struct {
	OrderRef string `json:"orderRef"` // A reference ID for an order. This is used to query the status of the order or to cancel it.
}

// A CollectResp is a successful response from the querying the order status with a collect request.
type CollectResp struct {
	OrderRef       string          `json:"orderRef"`                 // A reference ID for an order. This is used to query the status of the order or to cancel it.
	Status         Status          `json:"status"`                   // The current status of the order.
	HintCode       HintCode        `json:"hintCode,omitempty"`       // Used to provide the user with details and instructions.
	CompletionData *CompletionData `json:"completionData,omitempty"` // Information about the user and the completed order.
}

// CompletionData contains information about the user and the completed order.
//
// The user has completed the order. completionData includes the signature, user information and the OCSP response. You should verify user information to proceed. You should retain completion data for future reference, compliance and audit purposes..
type CompletionData struct {
	// Information related to the user.
	User *User `json:"user,omitempty"`
	// Information related to the device.
	Device *Device `json:"device,omitempty"`
	// Information about additional verifications that were part of the order.
	StepUp *StepUp `json:"stepUp,omitempty"`
	// The date the BankID was issued to the user.
	//
	// The issue date is expressed using ISO 8601 date format with a UTC time zone offset.
	BankIDIssueDate string `json:"bankIdIssueDate,omitempty"`
	// The signature that is the result of the order.
	//
	// This is a base 64 encoded XML signature string.
	Signature string `json:"signature,omitempty"`
	// The OCSP response.
	//
	// This is a base 64 encoded OCSP response.
	//
	// The OCSP response is signed by a certificate that has the same issuer as the certificate being verified, and it has a nonce extension. The nonce is calculated as:
	// 	- SHA-1 hash over the base 64 XML signature encoded as UTF-8.
	// 	- 12 random bytes added after the hash.
	//
	// The nonce is 32 bytes (20 + 12).
	OCSPResponse string `json:"ocspResponse,omitempty"`
	// Indicates the risk level of the order based on data available in the order.
	Risk Risk `json:"risk,omitempty"`
}

// User contains information related to the user.
type User struct {
	PersonalNumber string `json:"personalNumber,omitempty"` // The ID number of the user. The ID number is a Swedish national identification number (12 digits).
	Name           string `json:"name,omitempty"`           // The first and last name of the user.
	GivenName      string `json:"givenName,omitempty"`      // The first name of the user.
	Surname        string `json:"surname,omitempty"`        // The surname of the user.
}

// Device contains information related to the device.
type Device struct {
	// The IP address of the user agent as the BankID server sees it.
	//
	// When an order is started with autoStartToken you can check that it matches the IP you service observes to ensure session fixation.
	IPAddress string `json:"ipAddress,omitempty"`
	// Unique hardware identifier for the user's device.
	UHI string `json:"uhi,omitempty"`
}

// StepUp contains information about additional verifications that were part of the order.
type StepUp struct {
	MRTD bool `json:"mrtd,omitempty"` // Whether an MRTD check was performed before the order was completed.
}

type err400 struct {
	ErrorCode string `json:"errorCode"`
	Details   string `json:"details"`
}

func (err err400) Error() string {
	return fmt.Sprintf("errorCode=%s, details=%s", err.ErrorCode, err.Details)
}
