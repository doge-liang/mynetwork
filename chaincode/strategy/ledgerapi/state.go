package ledgerapi

import "strings"

// SplitKey splits a key on colon
func SplitKey(key string) []string {
	return strings.Split(key, ":")
}

// MakeKey joins key parts using colon
func MakeKey(keyParts ...string) string {
	return strings.Join(keyParts, ":")
}

type StateInterface interface {
	GetSplitKey() []string
	Serialize() ([]byte, error)
}
