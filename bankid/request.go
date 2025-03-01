package bankid

const (
	AppIdentifier    = "appIdentifier"
	DeviceOS         = "deviceOS"
	DeviceIdentifier = "deviceIdentifier"
	DeviceModelName  = "deviceModelName"
	ReferringDomain  = "referringDomain"
	UserAgent        = "userAgent"
)

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
