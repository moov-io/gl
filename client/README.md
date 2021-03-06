# Go API client for openapi

Moov Accounts is an HTTP service which represents both a general ledger and chart of accounts for customers. The service is designed to abstract over various core systems and provide a uniform API for developers.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://groups.google.com/forum/#!forum/moov-users](https://groups.google.com/forum/#!forum/moov-users)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./openapi"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:8085*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountsApi* | [**CreateAccount**](docs/AccountsApi.md#createaccount) | **Post** /accounts | Create Account
*AccountsApi* | [**CreateTransaction**](docs/AccountsApi.md#createtransaction) | **Post** /accounts/transactions | Create Transaction
*AccountsApi* | [**GetAccountTransactions**](docs/AccountsApi.md#getaccounttransactions) | **Get** /accounts/{accountID}/transactions | Get Account transactions
*AccountsApi* | [**Ping**](docs/AccountsApi.md#ping) | **Get** /ping | Ping Accounts service
*AccountsApi* | [**ReverseTransaction**](docs/AccountsApi.md#reversetransaction) | **Post** /accounts/transactions/{transactionID}/reversal | Reverse a transaction
*AccountsApi* | [**SearchAccounts**](docs/AccountsApi.md#searchaccounts) | **Get** /accounts/search | Search for Accounts


## Documentation For Models

 - [Account](docs/Account.md)
 - [AccountAddress](docs/AccountAddress.md)
 - [CreateAccount](docs/CreateAccount.md)
 - [CreateAccountAddress](docs/CreateAccountAddress.md)
 - [CreatePhone](docs/CreatePhone.md)
 - [CreateTransaction](docs/CreateTransaction.md)
 - [Error](docs/Error.md)
 - [Phone](docs/Phone.md)
 - [Transaction](docs/Transaction.md)
 - [TransactionLine](docs/TransactionLine.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Author



