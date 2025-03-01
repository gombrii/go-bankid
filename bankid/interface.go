package bankid

type Service interface {
	Auth()
	Sign()
	Payment()
	PhoneAuth()
	PhoneSign()
	Collect()
	Cancel()
}
