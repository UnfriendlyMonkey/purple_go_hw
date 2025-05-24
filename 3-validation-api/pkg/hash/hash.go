package hash

import (
	"crypto/md5"
	"encoding/hex"
)

func Hash(toHash string) (string, error) {
	hasher := md5.New()
	hasher.Write([]byte(toHash))
	hashStr := hex.EncodeToString(hasher.Sum(nil))
	return hashStr, nil
}
