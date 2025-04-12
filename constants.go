// © 2025 Simon Oscar Gombrii. Released under the MIT License.

package bankid

// Recommended user messages
const (
	RFA1_SWE   = "Starta BankID-appen."
	RFA2_SWE   = "Du verkar inte ha BankID-appen. Installera den och skaffa ett BankID."
	RFA3_SWE   = "Åtgärden avbröts. Försök igen."
	RFA4_SWE   = "En identifiering eller underskrift pågår redan för ditt personnummer. Försök igen."
	RFA5_SWE   = "Något gick fel. Försök igen."
	RFA6_SWE   = "Åtgärden avbröts."
	RFA8_SWE   = "BankID-appen svarar inte. Kontrollera att den är startad och att du har internetanslutning. Försök sedan igen."
	RFA9_SWE   = "Skriv in din säkerhetskod i BankID-appen och välj Identifiera eller Skriv under."
	RFA13_SWE  = "Försöker starta BankID-appen."
	RFA15A_SWE = "Söker efter BankID. Säkerställ att du har ett giltigt BankID på den här datorn. Om du har ett BankID på kort, sätt in kortet i kortläsaren."
	RFA15B_SWE = "Söker efter BankID. Säkerställ att du har ett gitligt BankID på den här enheten."
	RFA16_SWE  = "Ditt BankID är för gammalt eller spärrat. Använd ett annat BankID eller skaffa ett nytt hos din bank."
	RFA17A_SWE = "Du verkar inte ha BankID-appen/programmet. Installera den och skaffa ett BankID hos din bank."
	RFA17B_SWE = "Misslyckades att läsa av QR-koden. Starta BankID-appen och läs av QR-koden."
	RFA19_SWE  = "Vill du använda BankID på den här datorn eller ett Mobilt BankID?"
	RFA20_SWE  = "Vill du använda BankID på den här enheten eller på en annan enhet?"
	RFA21_SWE  = "En identifiering eller underskrift pågår."
	RFA22_SWE  = "Något gick fel. Försök igen."
	RFA23_SWE  = "Fotografera och läs av din ID-handling med BankID-appen."

	RFA1_ENG   = "Start your BankID app. "
	RFA2_ENG   = "You don't seem to have the BankID app. Install the app and get a BankID."
	RFA3_ENG   = "The action was cancelled. Please try again."
	RFA4_ENG   = "An identification or signature is already in progress for your personal identity number. Please try again."
	RFA5_ENG   = "Something went wrong. Please try again."
	RFA6_ENG   = "The action was cancelled."
	RFA8_ENG   = "The BankID app is not responding. Please check that it’s started and that you have internet access. Try again."
	RFA9_ENG   = "Enter your security code in the BankID app and select Identify or Sign."
	RFA13_ENG  = "Trying to start your BankID app."
	RFA15A_ENG = "Searching for BankID. Make sure you have a valid BankID on this computer. If you have a BankID on card, please insert the card into your card reader."
	RFA15B_ENG = "Searching for BankID. Make sure you have a valid BankID on this device."
	RFA16_ENG  = "Your BankID is blocked or too old. Please use another BankID or get a new one from your bank."
	RFA17A_ENG = "You don't seem to have the BankID app/program. Please install it and get a BankID from your bank."
	RFA17B_ENG = "Failed to scan the QR code. Start the BankID app and scan the QR code."
	RFA19_ENG  = "Would you like to use BankID on this computer, or a Mobile BankID?"
	RFA20_ENG  = "Do you want to use BankID on this device or another device?"
	RFA21_ENG  = "An identification or signing is in progress."
	RFA22_ENG  = "Something went wrong. Please try again."
	RFA23_ENG  = "Take a photo of, and scan, you ID document with the BankID app."
)

const (
	ProdURL = "https://appapi2.bankid.com"
	TestURL = "https://appapi2.test.bankid.com"
)

// A CertificatePolicy restricts the method with which an identification or signing can be performed.
// It matches the OID in certificate policies in the user certificate.
//
// When using one of the BankID on card policies, the cardReader requirement can be used to further restrict the type of card reader allowed. If no cardReader requirement is passed, all supported kinds of card readers are permitted.
type CertificatePolicy string

const (
	ProdBankIDOnFile             CertificatePolicy = "1.2.752.78.1.1"
	ProdBankIDOnCard             CertificatePolicy = "1.2.752.78.1.2"
	ProdMobileBankID             CertificatePolicy = "1.2.752.78.1.5"
	TestBankIDOnFile             CertificatePolicy = "1.2.3.4.5"
	TestBankIDOnCard             CertificatePolicy = "1.2.3.4.10"
	TestMobileBankID             CertificatePolicy = "1.2.3.4.25"
	TestBankIDForSomeBankIDBanks CertificatePolicy = "1.2.752.60.1.6"
)

// UserVisibleDataFormat inticates If present and set to "simpleMarkdownV1", this parameter indicates that userVisibleData holds formatting characters.
type UserVisibleDataFormat string

const (
	Plaintext        UserVisibleDataFormat = "plaintext"        // userVisibleData contains base 64 encoded text using a sub-set of UTF-8 and CR, LF or CRLF for line breaks.
	SimpleMarkdownV1 UserVisibleDataFormat = "simpleMarkdownV1" // userVisibleData contains Simple Markdown version 1.
)

// Whether the user needs to complete the order using a card reader for the signature.
//
// This condition should always be combined with a certificatePolicies for a smart card to avoid undefined behaviour.
type CardReader string

const (
	Class1 CardReader = "class1" // The order must be confirmed with a card reader where the PIN code is entered on a computer keyboard, or a card reader of higher class.
	Class2 CardReader = "class2" // The order must be confirmed with a card reader where the PIN code is entered on the reader.
)

// A RiskFlag indicates to the risk assessment system that the payment has a higher risk or is unusual for the user.
type RiskFlag string

const (
	NewCard                  RiskFlag = "newCard"
	NewCustomer              RiskFlag = "newCustomer"
	NewRecipient             RiskFlag = "newRecipient"
	HighRiskRecipient        RiskFlag = "highRiskRecipient"
	LargeAmount              RiskFlag = "largeAmount"
	ForeignCurrency          RiskFlag = "foreignCurrency"
	CryptoCurrencyPurchase   RiskFlag = "cryptoCurrencyPurchase"
	MoneyTransfer            RiskFlag = "moneyTransfer"
	OverseasTransaction      RiskFlag = "overseasTransaction"
	RecurringPayment         RiskFlag = "recurringPayment"
	SuspiciousPaymentPattern RiskFlag = "suspiciousPaymentPattern"
	Other                    RiskFlag = "other"
)

// A TransactionType is the type of a transaction.
type TransactionType string

const (
	Card TransactionType = "card" // Card payment.
	NPA  TransactionType = "npa"  // Non-payment authentication.
)

// CallInitiator indicates if the user or your organization initiated the phone call.
type CallInitiator string

const (
	UserInitiator CallInitiator = "user" // The user called your organization.
	RPInitiator   CallInitiator = "RP"   // Your organization called the user.
)

// Status indicates the current status of the order.
type Status string

const (
	Pending  Status = "pending"  // The order is being processed. hintCode describes the status of the order.
	Complete Status = "complete" // The order is complete. completionData holds user information.
	Failed   Status = "failed"   // Something went wrong with the order. hintCode describes the error.
)

// A HintCode is used to provide the user with details and instructions.
type HintCode string

// These hintCides declare the state when an order is pending. You should use the hintCode to provide the user with details and instructions and keep calling collect until order fails or is complete.
const (
	OutstandingTransaction HintCode = "outstandingTransaction"
	NoClient               HintCode = "noClient"
	Started                HintCode = "started"
	UserMRTD               HintCode = "userMrtd"
	UserCallConfirm        HintCode = "userCallConfirm"
	UserSign               HintCode = "userSign"
)

// These hintCodes declare the final state when an order fails. You should use the hintCode to provide the user with details and instructions. The same orderRef must not be used for additional collect requests.
const (
	ExpiredTransaction     HintCode = "expiredTransaction"
	CertificateErr         HintCode = "certificateErr"
	UserCancel             HintCode = "userCancel"
	Cancelled              HintCode = "cancelled"
	StartFailed            HintCode = "startFailed"
	UserDeclinedCall       HintCode = "userDeclinedCall"
	NotSupportedByUserApp  HintCode = "notSupportedByUserApp"
	TransactionRiskBlocked HintCode = "transactionRiskBlocked"
)

// Risk indicates the risk level of the order based on data available in the order.
//
// This is only returned if requested in the order, and it may be absent if the risk could not be calculated.
//
// If you have sent the correct endUserIp and additional data, a risk indication with the value "high" means there are signs of the channel binding being compromised, or other highly concerning circumstances..
type Risk string

const (
	Low      Risk = "low"      // No or low risk identified in the available order data.
	Moderate Risk = "moderate" // Might require further action, investigation or follow-up by you based on the order data.
	High     Risk = "high"     // The order should be blocked or cancelled by you and needs further action, investigation or follow-up. This value will only be returned if you have requested to have the risk assement to be provided, but not supplied a risk condition.
)
