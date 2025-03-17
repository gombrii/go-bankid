package validator

//TODO:TODO:TODO: Utslagsfrågan. BÖR ens ett klientbibliotek som detta validera input?
// Nackdelar: 
// - Mer underhåll
// - Måste troligen vara HELT uttömmande
// Fördelar:
// - Användaren får snabbare sad-path (försummbar fördel)
// - Biblioteket kan testas offline

//TODO: Validatorn skall validera indata:
// - måste den finnas skall den finnas
// - krävs ett visst format skall det kontrolleras
// - om två parametrar har en relation, t.ex. att de är ömsesidigt uteslutande, då skall det kontrolleras
// - Parametrar som använder konstanter bör kolla att värdet tillhör något av de tillåtna värdena
// 
// NOTE: Glöm inte att allt detta, och allt annat, skall testas

//TODO: Specifika grejer att validera
// - endUserIP måste finnas och är rättformaterad IPv4 eller IPv6
// - Lite orsäkert, men om returnUrl finns måste den innehålla en "nonce to the session"
// - Endast en av Web och App kan förses. Förses en av dem måste naturligtvis något av dess fält vara initierade
// - ReferringDomain "Only the digits 0 to 9, the letters a to z, dot (".") and dash ("-") are allowed. When using an International Domain Name, the string must be punycode encoded."
// - cardReader, om satt måste en kort-relaterad certifiatePolocy också vara satt, men inte vice versa
// - personalNumber måste givetvis ha rätt format
// - Money får inte sättas i kombination med transactionType "npa"
// - Amount måste vara en siffersträng och kan innehålla en decimalseparator i form av ett kommatecken ","
// 
// 