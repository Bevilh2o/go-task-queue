package processor

import (
	"crypto/sha256"
	"errors"
)

func Process(payload string) error {
	// Simula lavoro CPU-bound (coerente con i progetti precedenti)
	data := []byte(payload)

	for i := 0; i < 1000; i++ {
		hash := sha256.Sum256(data)
		data = hash[:]
	}

	// Simula errore casuale (per retry)
	if len(payload)%3 == 0 {
		return errors.New("forced failure")
	}

	return nil
}
