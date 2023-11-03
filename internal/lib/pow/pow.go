package pow

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"math"
	"math/rand"
	"strings"
	"time"
)

const HashPrefix = "0000"

func ProofOfWorkIsValid(nonce, proof string) bool {
	steps := 0
	hash := ""
	for i := 0; i < math.MaxInt; i++ {
		attempt := fmt.Sprintf("%s%d", nonce, steps)
		hash = sha256Hash(attempt)
		if IsPrefixValid(hash) {
			break
		}
		steps++
	}

	return proof[:len(hash)] == hash
}

func IsPrefixValid(proof string) bool {
	return strings.HasPrefix(proof, HashPrefix)
}

func GetNonce() string {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	minNum := 9999
	maxNum := 100000
	n := random.Intn(maxNum-minNum+1) + minNum
	result := fmt.Sprintf("%d", n)

	return result
}

func FindProofOfWork(nonce string) string {
	proof := ""
	for i := 0; i < math.MaxInt; i++ {
		fmt.Print("\rSteps: ", i)
		attempt := fmt.Sprintf("%s%d", nonce, i)
		// Using SHA-256 for hashing.
		h := sha256.New()
		h.Write([]byte(attempt))
		hashBytes := h.Sum(nil)
		hash := hex.EncodeToString(hashBytes)

		if IsPrefixValid(hash) {
			proof = hash
			break
		}
	}
	return proof
}

func sha256Hash(input string) string {
	hash := sha256.New()
	io.WriteString(hash, input)
	return hex.EncodeToString(hash.Sum(nil))
}
