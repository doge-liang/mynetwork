package ledgerapi

import "strings"

func MakeKey(keyParts ...string) string {
	return strings.Join(keyParts, ":")
}

func SplitKey(key string) []string {
	return strings.Split(key, ":")
}

type StateInterface interface {
	GetSplitKey() []string
	Serialize() ([]byte, error)
}
