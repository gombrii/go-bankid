package bankid

type AuthOpts struct {
	App                   *App         `json:"app,omitempty"`
	ReturnRisk            bool         `json:"returnRisk,omitempty"`
	ReturnURL             string       `json:"returnURL,omitempty"`
	UserNonVisibleData    string       `json:"userNonVisibleData,omitempty"`
	UserVisibleData       string       `json:"userVisibleData,omitempty"`
	UserVisibleDataFormat string       `json:"userVisibleDataFormat,omitempty"`
	Web                   *Web         `json:"web,omitempty"`
	Requirement           *Requirement `json:"requirement,omitempty"`
}

type SignOpts struct {
	App                   *App         `json:"app,omitempty"`
	ReturnRisk            bool         `json:"returnRisk,omitempty"`
	ReturnURL             string       `json:"returnURL,omitempty"`
	UserNonVisibleData    string       `json:"userNonVisibleData,omitempty"`
	UserVisibleDataFormat string       `json:"userVisibleDataFormat,omitempty"`
	Web                   *Web         `json:"web,omitempty"`
	Requirement           *Requirement `json:"requirement,omitempty"`
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

type authReq struct {
	EndUserIP             string       `json:"endUserIp"`
	App                   *App         `json:"app,omitempty"`
	ReturnRisk            bool         `json:"returnRisk,omitempty"`
	ReturnURL             string       `json:"returnURL,omitempty"`
	UserNonVisibleData    string       `json:"userNonVisibleData,omitempty"`
	UserVisibleData       string       `json:"userVisibleData,omitempty"`
	UserVisibleDataFormat string       `json:"userVisibleDataFormat,omitempty"`
	Web                   *Web         `json:"web,omitempty"`
	Requirement           *Requirement `json:"requirement,omitempty"`
}

type signReq struct {
	EndUserIP             string       `json:"endUserIp"`
	App                   *App         `json:"app,omitempty"`
	ReturnRisk            bool         `json:"returnRisk,omitempty"` //TODO: Lite allmän todo här. Bör strängar och booleaener som inte är obligatoriska vara pekare???
	ReturnURL             string       `json:"returnURL,omitempty"`
	UserNonVisibleData    string       `json:"userNonVisibleData,omitempty"`
	UserVisibleData       string       `json:"userVisibleData"`
	UserVisibleDataFormat string       `json:"userVisibleDataFormat,omitempty"`
	Web                   *Web         `json:"web,omitempty"`
	Requirement           *Requirement `json:"requirement,omitempty"`
}

type collectReq struct {
	OrderRef string `json:"orderRef"`
}

type cancelReq struct {
	OrderRef string `json:"orderRef"`
}
