package dictionary

import (
	"bytes"
	"encoding/gob"
	"time"

	"github.com/dgraph-io/badger"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func (d *Dictionary) Add(word string, definition string) error {
	caser := cases.Title(language.Und)

	entry := Entry{
		Word:       caser.String(word),
		Definition: definition,
		CreatedAt:  time.Now(),
	}

	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(entry)
	if err != nil {
		return err
	}
	return d.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(word), buffer.Bytes())
	})
}
