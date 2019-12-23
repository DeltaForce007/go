package processors

import (
	"testing"

	"github.com/stellar/go/services/horizon/internal/db2/history"
	. "github.com/stellar/go/services/horizon/internal/test/transactions"
	"github.com/stretchr/testify/assert"
)

func TestAccountCreatedEffects(t *testing.T) {
	assert := assert.New(t)

	transaction := BuildLedgerTransaction(
		t,
		TestTransaction{
			Index:         1,
			EnvelopeXDR:   "AAAAAGL8HQvQkbK2HA3WVjRrKmjX00fG8sLI7m0ERwJW/AX3AAAAZAAAAAAAAAAaAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAoZftFP3p4ifbTm6hQdieotu3Zw9E05GtoSh5MBytEpQAAAACVAvkAAAAAAAAAAABVvwF9wAAAEDHU95E9wxgETD8TqxUrkgC0/7XHyNDts6Q5huRHfDRyRcoHdv7aMp/sPvC3RPkXjOMjgbKJUX7SgExUeYB5f8F",
			ResultXDR:     "AAAAAAAAAGQAAAAAAAAAAQAAAAAAAAABAAAAAAAAAAA=",
			MetaXDR:       "AAAAAQAAAAIAAAADAAAAOQAAAAAAAAAAYvwdC9CRsrYcDdZWNGsqaNfTR8bywsjubQRHAlb8BfcLGrZY9dZxbAAAAAAAAAAZAAAAAAAAAAEAAAAAYvwdC9CRsrYcDdZWNGsqaNfTR8bywsjubQRHAlb8BfcAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAABAAAAOQAAAAAAAAAAYvwdC9CRsrYcDdZWNGsqaNfTR8bywsjubQRHAlb8BfcLGrZY9dZxbAAAAAAAAAAaAAAAAAAAAAEAAAAAYvwdC9CRsrYcDdZWNGsqaNfTR8bywsjubQRHAlb8BfcAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAABAAAAAwAAAAMAAAA5AAAAAAAAAABi/B0L0JGythwN1lY0aypo19NHxvLCyO5tBEcCVvwF9wsatlj11nFsAAAAAAAAABoAAAAAAAAAAQAAAABi/B0L0JGythwN1lY0aypo19NHxvLCyO5tBEcCVvwF9wAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAEAAAA5AAAAAAAAAABi/B0L0JGythwN1lY0aypo19NHxvLCyO5tBEcCVvwF9wsatlahyo1sAAAAAAAAABoAAAAAAAAAAQAAAABi/B0L0JGythwN1lY0aypo19NHxvLCyO5tBEcCVvwF9wAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAA5AAAAAAAAAAChl+0U/eniJ9tObqFB2J6i27dnD0TTka2hKHkwHK0SlAAAAAJUC+QAAAAAOQAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAA==",
			FeeChangesXDR: "AAAAAgAAAAMAAAA3AAAAAAAAAABi/B0L0JGythwN1lY0aypo19NHxvLCyO5tBEcCVvwF9wsatlj11nHQAAAAAAAAABkAAAAAAAAAAQAAAABi/B0L0JGythwN1lY0aypo19NHxvLCyO5tBEcCVvwF9wAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAEAAAA5AAAAAAAAAABi/B0L0JGythwN1lY0aypo19NHxvLCyO5tBEcCVvwF9wsatlj11nFsAAAAAAAAABkAAAAAAAAAAQAAAABi/B0L0JGythwN1lY0aypo19NHxvLCyO5tBEcCVvwF9wAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAA==",
			Hash:          "0e5bd332291e3098e49886df2cdb9b5369a5f9e0a9973f0d9e1a9489c6581ba2",
		},
	)

	expected := []map[string]interface{}{
		map[string]interface{}{
			"address":     "GCQZP3IU7XU6EJ63JZXKCQOYT2RNXN3HB5CNHENNUEUHSMA4VUJJJSEN",
			"operationID": int64(240518172673),
			"details": map[string]interface{}{
				"starting_balance": "1000.0000000",
			},
			"effectType": history.EffectType(0),
			"order":      uint32(1),
		},
		map[string]interface{}{
			"address":     "GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H",
			"operationID": int64(240518172673),
			"details": map[string]interface{}{
				"amount":     "1000.0000000",
				"asset_type": "native",
			},
			"effectType": history.EffectType(3),
			"order":      uint32(2),
		},
		map[string]interface{}{
			"address":     "GCQZP3IU7XU6EJ63JZXKCQOYT2RNXN3HB5CNHENNUEUHSMA4VUJJJSEN",
			"operationID": int64(240518172673),
			"details": map[string]interface{}{
				"public_key": "GCQZP3IU7XU6EJ63JZXKCQOYT2RNXN3HB5CNHENNUEUHSMA4VUJJJSEN",
				"weight":     1,
			},
			"effectType": history.EffectType(10),
			"order":      uint32(3),
		},
	}

	op := transactionOperationWrapper{
		index:          0,
		transaction:    transaction,
		operation:      transaction.Envelope.Tx.Operations[0],
		ledgerSequence: uint32(56),
	}

	effects, err := op.Effects()
	assert.NoError(err)
	assert.Equal(expected, effects)
}
