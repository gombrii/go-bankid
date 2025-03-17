package bankid

import "fmt"

type AuthResp struct {
	OrderRef       string `json:"orderRef"`
	AutoStartToken string `json:"autoStartToken"`
	QRStartToken   string `json:"qrStartToken"`
	QRStartSecret  string `json:"qrStartSecret"`
}

type SignResp struct {
	OrderRef       string `json:"orderRef"`
	AutoStartToken string `json:"autoStartToken"`
	QRStartToken   string `json:"qrStartToken"`
	QRStartSecret  string `json:"qrStartSecret"`
}

type PaymentResp struct {
	OrderRef       string `json:"orderRef"`
	AutoStartToken string `json:"autoStartToken"`
	QRStartToken   string `json:"qrStartToken"`
	QRStartSecret  string `json:"qrStartSecret"`
}

type PhoneAuthResp struct {
	OrderRef string `json:"orderRef"`
}

type PhoneSignResp struct {
	OrderRef string `json:"orderRef"`
}

type CollectResp struct {
	OrderRef       string          `json:"orderRef"`
	Status         Status          `json:"status"`
	HintCode       HintCode        `json:"hintCode,omitempty"`
	CompletionData *CompletionData `json:"completionData,omitempty"`
}

type CompletionData struct {
	User            *User   `json:"user,omitempty"`
	Device          *Device `json:"device,omitempty"`
	StepUp          *StepUp `json:"stepUp,omitempty"`
	BankIDIssueDate string  `json:"bankIdIssueDate,omitempty"`
	Signature       string  `json:"signature,omitempty"`
	OCSPResponse    string  `json:"ocspResponse,omitempty"`
	Risk            Risk    `json:"risk,omitempty"`
}

type User struct {
	PersonalNumber string `json:"personalNumber,omitempty"`
	Name           string `json:"name,omitempty"`
	GivenName      string `json:"givenName,omitempty"`
	Surname        string `json:"surname,omitempty"`
}

type Device struct {
	IPAddress string `json:"ipAddress,omitempty"`
	UHI       string `json:"uhi,omitempty"`
}

type StepUp struct {
	MRTD bool `json:"mrtd,omitempty"`
}

type err400 struct {
	ErrorCode string `json:"errorCode"`
	Details   string `json:"details"`
}

func (err err400) Error() string {
	return fmt.Sprintf("errorCode=%s, details=%s", err.ErrorCode, err.Details)
}
