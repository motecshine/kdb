package kdb

import (
	"log"
	"testing"

	"github.com/dgraph-io/badger/v3"
)

func TestInitBadger(t *testing.T) {
	// Open the Badger database located in the /tmp/badger directory.
	// It will be created if it doesn't exist.
	db, err := badger.Open(badger.DefaultOptions("badger-data"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Start a writable transaction.
	txn := db.NewTransaction(true)
	defer txn.Discard()

	// Use the transaction...
	err = txn.Set([]byte("answer"), []byte("42"))
	if err != nil {
		t.Fatal(err)
	}

	// Commit the transaction and check for error.
	if err := txn.Commit(); err != nil {
		t.Fatal(err)
	}
}
