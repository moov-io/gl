// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"strings"

	"github.com/go-kit/kit/log"
)

type transactionRepository interface {
	Ping() error
	Close() error

	createTransaction(tx transaction, opts createTransactionOpts) error
	getAccountTransactions(accountID string) ([]transaction, error) // TODO(adam): limit and/or pagination params
}

type createTransactionOpts struct {
	// AllowOverdraft is an option on creating a transaction where GL will let the account 'go negative'
	// and extend credit from the FI to the customer.
	AllowOverdraft bool
}

func initTransactionStorage(logger log.Logger, name string) (transactionRepository, error) {
	switch strings.ToLower(name) {
	case "qledger":
		return setupQLedgerTransactionStorage(os.Getenv("QLEDGER_ENDPOINT"), os.Getenv("QLEDGER_AUTH_TOKEN"))
	case "sqlite":
		return setupSqliteTransactionStorage(logger, getSqlitePath())
	}
	return nil, nil
}

// grabAccountIds returns an []string of each accountId from an array of transactionLines.
// We do this to query transactions that have been posted against an account.
func grabAccountIds(lines []transactionLine) []string {
	var out []string
	for i := range lines {
		out = append(out, lines[i].AccountId)
	}
	return out
}
