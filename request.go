package bankid

//TODO: Kontrollera att alla jsontaggar som är frivilliga har omitempty och de andra inte har det

//TODO: Bör alla fält i Opts-structar vara värden och låta pekarna ligga i Req-strictarna? Det hade gjort användningen lite trevligrae.

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
	RiskFlags             []RiskFlag //TODO: Göra till egen typ, liksom andra
	UserNonVisibleData    string
	UserVisibleData       string //TODO: Är det jag som skall byte64-koda t.ex. detta?
	UserVisibleDataFormat UserVisibleDataFormat
	Web                   *Web
	Requirement           *Requirement
}

type PhoneAuthOpts struct {
	PersonalNumber        string                `json:"personalNumber"`
	UserNonVisibleData    string                `json:"userNonVisibleData,omitempty"`
	UserVisibleData       string                `json:"userVisibleData"` //TODO: Är det jag som skall byte64-koda t.ex. detta?
	UserVisibleDataFormat UserVisibleDataFormat `json:"userVisibleDataFormat,omitempty"`
	Requirement           *PhoneRequirement     `json:"requirement,omitempty"`
}

type PhoneSignOpts struct {
	PersonalNumber        string                `json:"personalNumber"`
	UserNonVisibleData    string                `json:"userNonVisibleData,omitempty"`
	UserVisibleDataFormat UserVisibleDataFormat `json:"userVisibleDataFormat,omitempty"`
	Requirement           *PhoneRequirement     `json:"requirement,omitempty"`
}

type App struct {
	AppIdentifier    string `json:"AppIdentifier,omitempty"`
	DeviceOS         string `json:"DeviceOS,omitempty"`
	DeviceIdentifier string `json:"DeviceIdentifier,omitempty"`
	DeviceModelName  string `json:"DeviceModelName,omitempty"`
}

type Web struct {
	DeviceIdentifier string `json:"deviceIdentifier,omitempty"`
	ReferringDomain  string `json:"referringDomain,omitempty"`
	UserAgent        string `json:"userAgent,omitempty"`
}

// TODO: Ser ut som att det finns olika varianter på denna för de olika ändpunkterna. Skapa flera
type Requirement struct {
	CardReader          CardReader          `json:"cardReader,omitempty"`
	CertificatePolicies []CertificatePolicy `json:"certificatePolicies,omitempty"`
	Mrtd                bool                `json:"mrtd,omitempty"`
	PersonalNumber      string              `json:"personalNumber,omitempty"`
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
	Money           Money           `json:"money,omitempty"`
	RiskWarning     string          `json:"riskWarning,omitempty"`
}

type Recipient struct {
	Name string `json:"name"`
}

type Money struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type authReq struct {
	App                   *App                  `json:"app,omitempty"`
	EndUserIP             string                `json:"endUserIp"`
	ReturnRisk            bool                  `json:"returnRisk,omitempty"`
	ReturnURL             string                `json:"returnURL,omitempty"`
	UserNonVisibleData    string                `json:"userNonVisibleData,omitempty"`
	UserVisibleData       string                `json:"userVisibleData,omitempty"` //TODO: Är det jag som skall byte64-koda t.ex. detta?
	UserVisibleDataFormat UserVisibleDataFormat `json:"userVisibleDataFormat,omitempty"`
	Web                   *Web                  `json:"web,omitempty"`
	Requirement           *Requirement          `json:"requirement,omitempty"`
}

type signReq struct {
	App                   *App                  `json:"app,omitempty"`
	EndUserIP             string                `json:"endUserIp"`
	ReturnRisk            bool                  `json:"returnRisk,omitempty"`
	ReturnURL             string                `json:"returnURL,omitempty"`
	UserNonVisibleData    string                `json:"userNonVisibleData,omitempty"`
	UserVisibleData       string                `json:"userVisibleData"` //TODO: Är det jag som skall byte64-koda t.ex. detta?
	UserVisibleDataFormat UserVisibleDataFormat `json:"userVisibleDataFormat,omitempty"`
	Web                   *Web                  `json:"web,omitempty"`
	Requirement           *Requirement          `json:"requirement,omitempty"`
}

type paymentReq struct {
	App                    *App                   `json:"app,omitempty"`
	EndUserIP              string                 `json:"endUserIp"`
	ReturnRisk             bool                   `json:"returnRisk,omitempty"`
	ReturnURL              string                 `json:"returnURL,omitempty"`
	RiskFlags              []RiskFlag             `json:"riskFlags,omitempty"`
	UserNonVisibleData     string                 `json:"userNonVisibleData,omitempty"`
	UserVisibleData        string                 `json:"userVisibleData"` //TODO: Är det jag som skall byte64-koda t.ex. detta?
	UserVisibleDataFormat  UserVisibleDataFormat  `json:"userVisibleDataFormat,omitempty"`
	UserVisibleTransaction UserVisibleTransaction `json:"userVisibleTransaction"`
	Web                    *Web                   `json:"web,omitempty"`
	Requirement            *Requirement           `json:"requirement,omitempty"`
}

type phoneAuthReq struct {
	CallInitiator         CallInitiator         `json:"callInitiator"`
	PersonalNumber        string                `json:"personalNumber"`
	UserNonVisibleData    string                `json:"userNonVisibleData,omitempty"`
	UserVisibleData       string                `json:"userVisibleData"` //TODO: Är det jag som skall byte64-koda t.ex. detta?
	UserVisibleDataFormat UserVisibleDataFormat `json:"userVisibleDataFormat,omitempty"`
	Requirement           *PhoneRequirement     `json:"requirement,omitempty"`
}

type phoneSignReq struct {
	CallInitiator         CallInitiator         `json:"callInitiator"`
	PersonalNumber        string                `json:"personalNumber"`
	UserNonVisibleData    string                `json:"userNonVisibleData,omitempty"`
	UserVisibleData       string                `json:"userVisibleData"` //TODO: Är det jag som skall byte64-koda t.ex. detta?
	UserVisibleDataFormat UserVisibleDataFormat `json:"userVisibleDataFormat,omitempty"`
	Requirement           *PhoneRequirement     `json:"requirement,omitempty"`
}

type collectReq struct {
	OrderRef string `json:"orderRef"`
}

type cancelReq struct {
	OrderRef string `json:"orderRef"`
}
