package validator

// NOTE: Detta packet bör om det finns kvar ligga under /internal och bör endast validera grundläggande saker såsom:
// - existensen av obligatoriska parametrar
// - konfliktfrihet mellan parametrar som är ömsesidigt uteslutande
// Paketet skall INTE:
// - detaljerat validera format av parametrar
// - att parametrar håller tillåtna värden.

//TODO: Validatorn skall validera indata:
// - måste den finnas skall den finnas
// - om två parametrar har en relation, t.ex. att de är ömsesidigt uteslutande, då skall det kontrolleras
// - EN grej som faktiskt är bra att validera är skapandet av uppkopplingen, t.ex. validering av certifikaten och sånt
// 
// NOTE: Glöm inte att allt detta, och allt annat, skall testas

//TODO: Specifika grejer att validera
// - endUserIP måste finnas
// - Endast en av Web och App kan förses. Förses en av dem måste naturligtvis något av dess fält vara initierade
// - cardReader, om satt måste en kort-relaterad certifiatePolocy också vara satt, men inte vice versa
// - Money får inte sättas i kombination med transactionType "npa" 