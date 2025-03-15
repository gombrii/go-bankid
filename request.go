package bankid

type AuthOpts struct {
	App                   *App
	ReturnRisk            bool
	ReturnURL             string
	UserNonVisibleData    string
	UserVisibleData       string
	UserVisibleDataFormat string
	Web                   *Web
	Requirement           *Requirement
}

type SignOpts struct {
	App                   *App
	ReturnRisk            bool
	ReturnURL             string
	UserNonVisibleData    string
	UserVisibleDataFormat string
	Web                   *Web
	Requirement           *Requirement
}

type PaymenyOpts struct {
	App                   *App
	ReturnRisk            bool
	ReturnURL             string
	RiskFlags             []RiskFlag //TODO: Göra till egen typ, liksom andra
	UserNonVisibleData    string
	UserVisibleData       string
	UserVisibleDataFormat string
	Web                   *Web
	Requirement           *Requirement
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

type Requirement struct {
	CardReader          string   `json:"cardReader,omitempty"`
	CertificatePolicies []string `json:"certificatePolicies,omitempty"`
	Mrtd                bool     `json:"mrtd,omitempty"`
	PersonalNumber      string   `json:"personalNumber,omitempty"`
	PinCode             bool     `json:"pinCode,omitempty"`
}

type UserVisibleTransaction struct {
	TransactionType string    `json:"transactionType"`
	Recipient       Recipient `json:"recipient"`
	Money           Money     `json:"money,omitempty"`
	RiskWarning     string    `json:"riskWarning,omitempty"`
}

type Recipient struct {
	Name string `json:"name"`
}

type Money struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type authReq struct {
	App                   *App         `json:"app,omitempty"`
	EndUserIP             string       `json:"endUserIp"`
	ReturnRisk            bool         `json:"returnRisk,omitempty"`
	ReturnURL             string       `json:"returnURL,omitempty"`
	UserNonVisibleData    string       `json:"userNonVisibleData,omitempty"`
	UserVisibleData       string       `json:"userVisibleData,omitempty"`
	UserVisibleDataFormat string       `json:"userVisibleDataFormat,omitempty"`
	Web                   *Web         `json:"web,omitempty"`
	Requirement           *Requirement `json:"requirement,omitempty"`
}

type signReq struct {
	App                   *App         `json:"app,omitempty"`
	EndUserIP             string       `json:"endUserIp"`
	ReturnRisk            bool         `json:"returnRisk,omitempty"`
	ReturnURL             string       `json:"returnURL,omitempty"`
	UserNonVisibleData    string       `json:"userNonVisibleData,omitempty"`
	UserVisibleData       string       `json:"userVisibleData"`
	UserVisibleDataFormat string       `json:"userVisibleDataFormat,omitempty"`
	Web                   *Web         `json:"web,omitempty"`
	Requirement           *Requirement `json:"requirement,omitempty"`
}

type paymentReq struct {
	App                    *App                   `json:"app,omitempty"`
	EndUserIP              string                 `json:"endUserIp"`
	ReturnRisk             bool                   `json:"returnRisk,omitempty"`
	ReturnURL              string                 `json:"returnURL,omitempty"`
	RiskFlags              []RiskFlag             `json:"riskFlags,omitempty"`
	UserNonVisibleData     string                 `json:"userNonVisibleData,omitempty"`
	UserVisibleData        string                 `json:"userVisibleData"`
	UserVisibleDataFormat  string                 `json:"userVisibleDataFormat,omitempty"`
	UserVisibleTransaction UserVisibleTransaction `json:"userVisibleTransaction"`
	Web                    *Web                   `json:"web,omitempty"`
	Requirement            *Requirement           `json:"requirement,omitempty"`
}

type collectReq struct {
	OrderRef string `json:"orderRef"`
}

type cancelReq struct {
	OrderRef string `json:"orderRef"`
}
