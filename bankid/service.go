package bankid

type Client interface {
	Auth()
	Sign()
	Payment()
	PhoneAuth()
	PhoneSign()
	Collect()
	Cancel()
}

type BankID struct {
	client Client
}

func New( /*certifikat och nån config-fil*/ ) *BankID {
	return &BankID{}
}

func (b BankID) Auth(enUserIP string, opts *AuthOpts) (AuthResp, error) {
	b.client.Auth()

	return AuthResp{
		OrderRef:       "",
		AutoStartToken: "",
		QRStartToken:   "",
		QRStartSecret:  "",
	}, nil
}

func (b BankID) Sign() {

}

func (b BankID) Payment() {

}

func (b BankID) PhoneAuth() {

}
func (b BankID) PhoneSign() {

}

func (b BankID) Collect() {

}
func (b BankID) Cancel() {

}
