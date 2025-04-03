package bankid

type AuthOpts struct {
	App                   *App
	ReturnRisk            bool
	ReturnURL             string
	UserNonVisibleData    string
	UserVisibleData       string
	UserVisibleDataFormat UserVisibleDataFormat
	Web                   *Web
	Requirement           *Requirement
}

type SignOpts struct {
	App                   *App
	ReturnRisk            bool
	ReturnURL             string
	UserNonVisibleData    string
	UserVisibleDataFormat UserVisibleDataFormat
	Web                   *Web
	Requirement           *Requirement
}

type PaymentOpts struct {
	App                   *App
	ReturnRisk            bool
	ReturnURL             string
	RiskFlags             []RiskFlag
	UserNonVisibleData    string
	UserVisibleData       string
	UserVisibleDataFormat UserVisibleDataFormat
	Web                   *Web
	Requirement           *Requirement
}

type PhoneAuthOpts struct {
	PersonalNumber        string
	UserNonVisibleData    string
	UserVisibleData       string
	UserVisibleDataFormat UserVisibleDataFormat
	Requirement           *PhoneRequirement
}

type PhoneSignOpts struct {
	PersonalNumber        string
	UserNonVisibleData    string
	UserVisibleDataFormat UserVisibleDataFormat
	Requirement           *PhoneRequirement
}

type App struct {
	AppIdentifier    string `json:"appIdentifier,omitempty"`
	DeviceOS         string `json:"deviceOS,omitempty"`
	DeviceIdentifier string `json:"deviceIdentifier,omitempty"`
	DeviceModelName  string `json:"deviceModelName,omitempty"`
}

type Web struct {
	DeviceIdentifier string `json:"deviceIdentifier,omitempty"`
	ReferringDomain  string `json:"referringDomain,omitempty"` //NOTE: Speciell validering
	UserAgent        string `json:"userAgent,omitempty"`
}

type Requirement struct {
	CardReader          CardReader          `json:"cardReader,omitempty"` //NOTE: Speciell validering
	CertificatePolicies []CertificatePolicy `json:"certificatePolicies,omitempty"`
	MRTD                bool                `json:"mrtd,omitempty"`
	PersonalNumber      string              `json:"personalNumber,omitempty"` //NOTE: Speciell validering
	PinCode             bool                `json:"pinCode,omitempty"`
}

type PhoneRequirement struct {
	CardReader          CardReader          `json:"cardReader,omitempty"`
	CertificatePolicies []CertificatePolicy `json:"certificatePolicies,omitempty"`
	PinCode             bool                `json:"pinCode,omitempty"`
}

type UserVisibleTransaction struct {
	TransactionType TransactionType `json:"transactionType"`
	Recipient       Recipient       `json:"recipient"`
	Money           *Money          `json:"money,omitempty"` //NOTE: Speciell validering
	RiskWarning     string          `json:"riskWarning,omitempty"`
}

type Recipient struct {
	Name string `json:"name"`
}

type Money struct {
	Amount   string `json:"amount"` //NOTE: Speciell validering
	Currency string `json:"currency"`
}

type authReq struct {
	App                   *App                  `json:"app,omitempty"` //NOTE: Speciell validering
	EndUserIP             string                `json:"endUserIp"`     //NOTE: Speciell validering
	ReturnRisk            bool                  `json:"returnRisk,omitempty"`
	ReturnURL             string                `json:"returnUrl,omitempty"` //NOTE: Speciell validering
	UserNonVisibleData    string                `json:"userNonVisibleData,omitempty"`
	UserVisibleData       string                `json:"userVisibleData,omitempty"`
	UserVisibleDataFormat UserVisibleDataFormat `json:"userVisibleDataFormat,omitempty"`
	Web                   *Web                  `json:"web,omitempty"` //NOTE: Speciell validering
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
