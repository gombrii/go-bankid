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

type CollectResp struct {
	OrderRef       string          `json:"orderRef"`
	Status         string          `json:"status"`
	HintCode       string          `json:"hintCode,omitempty"` //TODO: Bör vara pekare?
	CompletionData *CompletionData `json:"completionData,omitempty"`
}

type CompletionData struct {
	User            string `json:"user"`
	Device          string `json:"device"`
	BankIDIssueDate string `json:"bankIDIssueDate"`
	StepUp          string `json:"stepUp"`
	Signature       string `json:"signature"`    //TODO: []byte istället?
	OcspResponse    string `json:"ocspResponse"` //TODO: OCSPResponse istället?
	Risk            string `json:"risk"`
}

type Err400 struct {
	ErrorCode string `json:"errorCode"`
	Details   string `json:"details"`
}

func (err Err400) Error() string {
	return fmt.Sprintf("errorCode=%s, details=%s", err.ErrorCode, err.Details)
}
