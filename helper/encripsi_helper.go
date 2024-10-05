package helper

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
)

func Sha1ToString(input string) string {
	h := sha1.New()
	h.Write([]byte(input))
	hash := h.Sum(nil)
	return hex.EncodeToString(hash)
}

func VerifySHA1Hash(input, hash string) (bool, error) {
	h := sha1.New()

	_, err := h.Write([]byte(input))
	if err != nil {
		return false, err
	}

	calculatedHash := hex.EncodeToString(h.Sum(nil))
	if calculatedHash != hash {
		return false, errors.New("Enkripsi tidak sesuai.")
	}

	return true, nil
}
