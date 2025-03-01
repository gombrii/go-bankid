package client

import "github.com/gomsim/go-bankid/bankid"

type authReq struct {
	App                   *bankid.App         `json:"app,omitempty"`
	ReturnRisk            bool                `json:"returnRisk,omitempty"`
	ReturnURL             string              `json:"returnURL,omitempty"`
	UserNonVisibleData    string              `json:"userNonVisibleData,omitempty"`
	UserVisibleData       string              `json:"userVisibleData,omitempty"`
	UserVisibleDataFormat string              `json:"userVisibleDataFormat,omitempty"`
	Web                   *bankid.Web         `json:"web,omitempty"`
	Requirement           *bankid.Requirement `json:"requirement,omitempty"`
}

func FromOpts(opts bankid.AuthOpts) authReq {
	return authReq{
		App:                   opts.App,
		ReturnRisk:            opts.ReturnRisk,
		ReturnURL:             opts.ReturnURL,
		UserNonVisibleData:    opts.UserNonVisibleData,
		UserVisibleData:       opts.UserVisibleData,
		UserVisibleDataFormat: opts.UserVisibleDataFormat,
		Web:                   opts.Web,
		Requirement:           opts.Requirement,
	}
}
