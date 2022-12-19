package dbo

import (
	"database/sql"
	"time"

	"crypto/sha256"
	"encoding/hex"
)

// Blockchain represents a row in the blockchain table.
type Blockchain struct {
	// The unique ID of the blockchain entry.
	ID int

	// The timestamp when the entry was created.
	Timestamp time.Time

	// The data associated with the entry.
	Data string

	// The hash of the previous block in the blockchain.
	PreviousBlockHash string

	// The hash of this block.
	BlockHash string
}

func InsertNewBlock(db *sql.DB, data string) error {
	// Obtain a full database lock.
	_, err := db.Exec("BEGIN")
	if err != nil {
		return err
	}

	// Query for the last entry in the database.
	var previousBlockHash string
	err = db.QueryRow("SELECT block_hash FROM blockchain ORDER BY id DESC LIMIT 1").Scan(&previousBlockHash)
	if err != nil {
		return err
	}

	// Compute the block_hash from the previous_block_hash and the data.
	blockHash := computeBlockHash(previousBlockHash, data)

	// Insert the new entry into the database.
	_, err = db.Exec("INSERT INTO blockchain (data, previous_block_hash, block_hash) VALUES ($1, $2, $3)", data, previousBlockHash, blockHash)
	if err != nil {
		return err
	}

	// Release the database lock.
	_, err = db.Exec("COMMIT")
	if err != nil {
		return err
	}

	return nil
}

func computeBlockHash(previousHash, data string) string {
	newstring := previousHash + data
	return computeSHA256(newstring)
}

func computeSHA256(input string) string {
	// Compute the SHA-256 hash of the input.
	hash := sha256.Sum256([]byte(input))

	// Return the hash as a hex-encoded string.
	return hex.EncodeToString(hash[:])
}
