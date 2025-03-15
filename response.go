package bankid

import "fmt"

//TODO: Bör vissa fält i response ha egna typer för att kunna jämföra, t.ex. Status och HintCode

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
	Status         string          `json:"status"`
	HintCode       string          `json:"hintCode,omitempty"`
	CompletionData *CompletionData `json:"completionData,omitempty"`
}

type CompletionData struct {
	User            string `json:"user"`
	Device          string `json:"device"`
	BankIDIssueDate string `json:"bankIDIssueDate"`
	StepUp          string `json:"stepUp"`
	Signature       string `json:"signature"`
	OCSPResponse    string `json:"ocspResponse"`
	Risk            string `json:"risk"`
}

type err400 struct {
	ErrorCode string `json:"errorCode"`
	Details   string `json:"details"`
}

func (err err400) Error() string {
	return fmt.Sprintf("errorCode=%s, details=%s", err.ErrorCode, err.Details)
}
