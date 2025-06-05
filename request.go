// © 2025 Simon Oscar Gombrii. Released under the MIT License.

package bid

// AuthOpts contains options to augument your identification order.
//
// If ReturnRisk is set to true a risk indicator will be included in the collect response when the order completes. If a risk indicator is required for the order to complete, for example, if a risk requirement is applied, the returnRisk property is ignored, and a risk indicator is always included; otherwise a default value of false is used. The risk indication requires that the endUserIp is correct. Please note that the assessed risk will not be returned if the order was blocked, which may happen if a risk requirement is set.
type AuthOpts struct {
	// Additional data included when creating an order from your app.
	App *App
	// Whether a risk indicator should be included in the collect response when the order completes. If a risk indicator is required for the order to complete, for example, if a risk requirement is applied, the returnRisk property is ignored, and a risk indicator is always included; otherwise a default value of false is used. The risk indication requires that the endUserIp is correct. Please note that the assessed risk will not be returned if the order was blocked, which may happen if a risk requirement is set.
	ReturnRisk bool
	// Orders started on the same device as where the user's BankID is stored (started with autostart token) will call this URL when the order is completed.
	//
	// Any return URL provided in the start URL when the BankID app was launched will be ignored. If the user has a version of the BankID app that does not support getting the returnUrl from the server, the order will be cancelled and the user will be asked to update their app.
	//
	// The return URL you provide should include a nonce to the session. When the user returns to your app or web page, your service should verify that the order was completed successfully and that the device receiving the returnUrl is the same device that started the order.
	//
	// Using this parameter will make your service more secure and strengthen the channel binding between you and the user.
	//
	// Ensure that the cookie or user IP address has not changed from the starting page to the returnUrl page.
	ReturnURL string
	// Data that you wish to include but not display to the user.
	//
	// The value must be base 64-encoded.
	UserNonVisibleData string
	// Text displayed to the user during the order.
	//
	// The purpose is to provide context, thereby enabling the user to detect identification errors and avert fraud attempts.
	//
	// The text can be formatted using CR, LF and CRLF for new lines. The text must be encoded as UTF-8 and then base 64 encoded.
	UserVisibleData string
	// If present and set to "simpleMarkdownV1", this parameter indicates that userVisibleData holds formatting characters.
	UserVisibleDataFormat UserVisibleDataFormat
	// Additional data included when creating an order from your web page.
	Web *Web
	// Requirements on how the authentication order must be performed.
	Requirement *Requirement
}

// SignOpts contains options to augument your signing order.
type SignOpts struct {
	// Additional data included when creating an order from your app.
	App *App
	// Whether a risk indicator should be included in the collect response when the order completes. If a risk indicator is required for the order to complete, for example, if a risk requirement is applied, the returnRisk property is ignored, and a risk indicator is always included; otherwise a default value of false is used. The risk indication requires that the endUserIp is correct. Please note that the assessed risk will not be returned if the order was blocked, which may happen if a risk requirement is set.
	ReturnRisk bool
	// Orders started on the same device as where the user's BankID is stored (started with autostart token) will call this URL when the order is completed.
	//
	// Any return URL provided in the start URL when the BankID app was launched will be ignored. If the user has a version of the BankID app that does not support getting the returnUrl from the server, the order will be cancelled and the user will be asked to update their app.
	//
	// The return URL you provide should include a nonce to the session. When the user returns to your app or web page, your service should verify that the order was completed successfully and that the device receiving the returnUrl is the same device that started the order.
	//
	// Using this parameter will make your service more secure and strengthen the channel binding between you and the user.
	//
	// Ensure that the cookie or user IP address has not changed from the starting page to the returnUrl page.
	ReturnURL string
	// Data that you wish to include but not display to the user.
	//
	// The value must be base 64-encoded.
	UserNonVisibleData string
	// If present and set to "simpleMarkdownV1", this parameter indicates that userVisibleData holds formatting characters.
	UserVisibleDataFormat UserVisibleDataFormat
	// Additional data included when creating an order from your web page.
	Web *Web
	// Requirements on how the authentication order must be performed.
	Requirement *Requirement
}

// PaymentOpts contains options to augument your payment order.
type PaymentOpts struct {
	// Additional data included when creating an order from your app.
	App *App
	// Whether a risk indicator should be included in the collect response when the order completes. If a risk indicator is required for the order to complete, for example, if a risk requirement is applied, the returnRisk property is ignored, and a risk indicator is always included; otherwise a default value of false is used. The risk indication requires that the endUserIp is correct. Please note that the assessed risk will not be returned if the order was blocked, which may happen if a risk requirement is set.
	ReturnRisk bool
	// Orders started on the same device as where the user's BankID is stored (started with autostart token) will call this URL when the order is completed.
	//
	// Any return URL provided in the start URL when the BankID app was launched will be ignored. If the user has a version of the BankID app that does not support getting the returnUrl from the server, the order will be cancelled and the user will be asked to update their app.
	//
	// The return URL you provide should include a nonce to the session. When the user returns to your app or web page, your service should verify that the order was completed successfully and that the device receiving the returnUrl is the same device that started the order.
	//
	// Using this parameter will make your service more secure and strengthen the channel binding between you and the user.
	//
	// Ensure that the cookie or user IP address has not changed from the starting page to the returnUrl page.
	ReturnURL string
	// Indicate to the risk assessment system that the payment has a higher risk or is unusual for the user.
	RiskFlags []RiskFlag
	// Data that you wish to include but not display to the user.
	//
	// The value must be base 64-encoded.
	UserNonVisibleData string
	// Text displayed to the user during the order.
	//
	// The purpose is to provide context, thereby enabling the user to detect identification errors and avert fraud attempts.
	//
	// The text can be formatted using CR, LF and CRLF for new lines. The text must be encoded as UTF-8 and then base 64 encoded.
	UserVisibleData string
	// If present and set to "simpleMarkdownV1", this parameter indicates that userVisibleData holds formatting characters.
	UserVisibleDataFormat UserVisibleDataFormat
	// Additional data included when creating an order from your web page.
	Web *Web
	// Requirements on how the authentication order must be performed.
	Requirement *Requirement
}

// PhoneAuthOpts contains options to augument your phone identification order.
type PhoneAuthOpts struct {
	// The ID number of the user.
	//
	// The ID number is a Swedish national identification number (12 digits).
	PersonalNumber string
	// Data that you wish to include but not display to the user.
	//
	// The value must be base 64-encoded.
	UserNonVisibleData string
	// Text displayed to the user during the order.
	//
	// The purpose is to provide context, thereby enabling the user to detect identification errors and avert fraud attempts.
	//
	// The text can be formatted using CR, LF and CRLF for new lines. The text must be encoded as UTF-8 and then base 64 encoded.
	UserVisibleData string
	// If present and set to "simpleMarkdownV1", this parameter indicates that userVisibleData holds formatting characters.
	UserVisibleDataFormat UserVisibleDataFormat
	// Requirements on how the authentication order must be performed.
	Requirement *PhoneRequirement
}

// PhoneSignOpts contains options to augument your phone signing order.
type PhoneSignOpts struct {
	// The ID number of the user. The ID number is a Swedish national identification number (12 digits).
	PersonalNumber string
	// Data that you wish to include but not display to the user.
	//
	// The value must be base 64-encoded.
	UserNonVisibleData string
	// If present and set to "simpleMarkdownV1", this parameter indicates that userVisibleData holds formatting characters.
	UserVisibleDataFormat UserVisibleDataFormat
	// Requirements on how the authentication order must be performed.
	Requirement *PhoneRequirement
}

// App contains additional data included when creating an order from your app.
//
// When starting an order from your app client this data may be included in the request.
//
// You can send the parameter web or app, not both. If providing this parameter, at least one of its
// members must be specified.
type App struct {
	// The identifier of your application.
	//
	// This is the package name on Android and the bundle identifier on iOS.
	//
	// It is vital to use the correct value. If your service does not supply the correct value legitimate orders might be blocked.
	AppIdentifier string `json:"appIdentifier,omitempty"`
	// The device operating system where your app is running.
	DeviceOS string `json:"deviceOS,omitempty"`
	// The identifier of the device your client is running on
	//
	// This is used to uniquely identify the device and should be a value that is not tied to a single user of the device. Preferably, it should remain the same even if your app is reinstalled.
	DeviceIdentifier string `json:"deviceIdentifier,omitempty"`
	// The model of the device your app is running on.
	DeviceModelName string `json:"deviceModelName,omitempty"`
}

// Web contains additional data included when creating an order from your web page.
//
// When starting an order from your web page this data may be included in the request.
//
// You can send the parameter web or app, not both. If providing this parameter, at least one of
// its members must be specified.
type Web struct {
	// The identifier of the device running your client.
	//
	// Do not use a session cookie. Use a separate cookie or the hash of one.
	//
	// This value should be unique to the user's browser and persist across sessions.
	DeviceIdentifier string `json:"deviceIdentifier,omitempty"`
	// The domain that starts the BankID app.
	//
	// This should generally be your domain name followed by the public suffix, which will generally be the top level domain.
	//
	// Only the digits 0 to 9, the letters a to z, dot (".") and dash ("-") are allowed. When using an International Domain Name, the string must be punycode encoded.
	ReferringDomain string `json:"referringDomain,omitempty"`
	// The user agent of the user interacting with your web page.
	UserAgent string `json:"userAgent,omitempty"`
}

// Requirement contains requirements on how the authentication order must be performed.
type Requirement struct {
	CardReader          CardReader          `json:"cardReader,omitempty"`          // Whether the user needs to complete the order using a card reader for the signature.
	CertificatePolicies []CertificatePolicy `json:"certificatePolicies,omitempty"` // The OID in certificate policies in the user certificate.
	MRTD                bool                `json:"mrtd,omitempty"`                // Whether the user needs to confirm their identity with a valid Swedish passport or national ID card to complete the order.
	PersonalNumber      string              `json:"personalNumber,omitempty"`      // The personal identity number allowed to confirm the identification.
	PinCode             bool                `json:"pinCode,omitempty"`
}

// PhoneRequirement contains requirements on how the authentication order must be performed.
type PhoneRequirement struct {
	CardReader          CardReader          `json:"cardReader,omitempty"`          // Whether the user needs to complete the order using a card reader for the signature.
	CertificatePolicies []CertificatePolicy `json:"certificatePolicies,omitempty"` // The OID in certificate policies in the user certificate.
	PinCode             bool                `json:"pinCode,omitempty"`             // Users are required to confirm the order with their security code even if they have biometrics activated.
}

// UserVisibleTransaction contains information about the transaction being approved.
type UserVisibleTransaction struct {
	TransactionType TransactionType `json:"transactionType"`       // The type of a transaction.
	Recipient       Recipient       `json:"recipient"`             // The recipient of the payment.
	Money           *Money          `json:"money,omitempty"`       // Object that sets monetary amount for the payment.
	RiskWarning     string          `json:"riskWarning,omitempty"` // Indicate to the user that the payment has higher risk or is unusual for the user.
}

// Recipient contains informtion about the recipient of the payment.
type Recipient struct {
	// The name of the recipient of the payment.
	//
	// For the transaction type "card", this is the merchant name.
	Name string `json:"name"`
}

// Money sets monetary amount for the payment.
// If the transactionType is "npa" this isn't allowed to be set.
type Money struct {
	// The monetary amount of the payment.
	//
	// The string can contain one decimal separator which must be ",". The rest of the input must be numbers.
	//
	// Examples: "1000,00", "100,000", "100", "0".
	Amount string `json:"amount"`
	// The currency of the payment.
	//
	// This must be an ISO 4217 alphabetic currency code.
	Currency string `json:"currency"`
}

type authReq struct {
	App                   *App                  `json:"app,omitempty"`
	EndUserIP             string                `json:"endUserIp"`
	ReturnRisk            bool                  `json:"returnRisk,omitempty"`
	ReturnURL             string                `json:"returnUrl,omitempty"`
	UserNonVisibleData    string                `json:"userNonVisibleData,omitempty"`
	UserVisibleData       string                `json:"userVisibleData,omitempty"`
	UserVisibleDataFormat UserVisibleDataFormat `json:"userVisibleDataFormat,omitempty"`
	Web                   *Web                  `json:"web,omitempty"`
	Requirement           *Requirement          `json:"requirement,omitempty"`
}

type signReq struct {
	App                   *App                  `json:"app,omitempty"`
	EndUserIP             string                `json:"endUserIp"`
	ReturnRisk            bool                  `json:"returnRisk,omitempty"`
	ReturnURL             string                `json:"returnUrl,omitempty"`
	UserNonVisibleData    string                `json:"userNonVisibleData,omitempty"`
	UserVisibleData       string                `json:"userVisibleData"`
	UserVisibleDataFormat UserVisibleDataFormat `json:"userVisibleDataFormat,omitempty"`
	Web                   *Web                  `json:"web,omitempty"`
	Requirement           *Requirement          `json:"requirement,omitempty"`
}

type paymentReq struct {
	App                    *App                   `json:"app,omitempty"`
	EndUserIP              string                 `json:"endUserIp"`
	ReturnRisk             bool                   `json:"returnRisk,omitempty"`
	ReturnURL              string                 `json:"returnUrl,omitempty"`
	RiskFlags              []RiskFlag             `json:"riskFlags,omitempty"`
	UserNonVisibleData     string                 `json:"userNonVisibleData,omitempty"`
	UserVisibleData        string                 `json:"userVisibleData,omitempty"`
	UserVisibleDataFormat  UserVisibleDataFormat  `json:"userVisibleDataFormat,omitempty"`
	UserVisibleTransaction UserVisibleTransaction `json:"userVisibleTransaction"`
	Web                    *Web                   `json:"web,omitempty"`
	Requirement            *Requirement           `json:"requirement,omitempty"`
}

type phoneAuthReq struct {
	CallInitiator         CallInitiator         `json:"callInitiator"`
	PersonalNumber        string                `json:"personalNumber,omitempty"`
	UserNonVisibleData    string                `json:"userNonVisibleData,omitempty"`
	UserVisibleData       string                `json:"userVisibleData,omitempty"`
	UserVisibleDataFormat UserVisibleDataFormat `json:"userVisibleDataFormat,omitempty"`
	Requirement           *PhoneRequirement     `json:"requirement,omitempty"`
}

type phoneSignReq struct {
	CallInitiator         CallInitiator         `json:"callInitiator"`
	PersonalNumber        string                `json:"personalNumber,omitempty"`
	UserNonVisibleData    string                `json:"userNonVisibleData,omitempty"`
	UserVisibleData       string                `json:"userVisibleData"`
	UserVisibleDataFormat UserVisibleDataFormat `json:"userVisibleDataFormat,omitempty"`
	Requirement           *PhoneRequirement     `json:"requirement,omitempty"`
}

type collectReq struct {
	OrderRef string `json:"orderRef"`
}

type cancelReq struct {
	OrderRef string `json:"orderRef"`
}
