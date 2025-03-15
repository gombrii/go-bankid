package bankid

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

type UserVisibleDataFormat string

const (
	Plaintext        UserVisibleDataFormat = "plaintext"
	SimpleMarkdownV1 UserVisibleDataFormat = "simpleMarkdownV1"
)

type CardReader string

const (
	Class1 CardReader = "class1"
	Class2 CardReader = "class2"
)

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

type TransactionType string

const (
	Card TransactionType = "card"
	NPA  TransactionType = "npa"
)

type CallInitiator string

const (
	User CallInitiator = "user"
	RP   CallInitiator = "RP"
)
