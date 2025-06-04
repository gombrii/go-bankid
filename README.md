[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![GitHub tag](https://img.shields.io/github/v/tag/your-user/your-repo)](https://github.com/your-user/your-repo/tags)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/user/repo)

Go-bankid is a client supporting BankID's API for performing identifications, digital signatures and much more. BankID is probably norhtern Europe's largest player in the digital ID industry. It allows the user to carry a fully trusted digital ID to prove their identidy online, opening up many possibilities in digital infrastructure.

This client is written using BankID's public documentation available at [developers.bankid.com](https://developers.bankid.com). Please explore to know more. The majority of the documentation in this library is taken straight from their page, only with some rephrasing to fit the context.

BankID has two environments, production and test. The former is used in real world application while the latter is used for testing integrations with their systems. This library supports both via [`NewProd`](./client.go#l32) and [`NewTest`](./client.go#l36) respectively.

The main package (bankid) is a work in progress, but all main functionality should be done. It is still not tested. packages ocsp, qr and signature will come in the future and contain nice-to-have functionality with treating some of the data returned by BankID.