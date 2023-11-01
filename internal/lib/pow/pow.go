package pow

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func ProofOfWorkIsValid(proof string) bool {
	return strings.HasPrefix(proof, "0000")
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

		if ProofOfWorkIsValid(hash) {
			proof = hash
			break
		}
	}
	return proof
}
