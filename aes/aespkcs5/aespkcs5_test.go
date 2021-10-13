package aespkcs5

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Encrypt(t *testing.T) {
	cases := []struct {
		Key       string
		Str       string
		Encrypted string
	}{
		{
			Key:       "BFE8916F13D0A7E7",
			Str:       "123456789",
			Encrypted: "59C2D8621A035DFC3307AF12ECD5AC65",
		},
		{
			Key:       "BFE8916F13D0A7E7",
			Str:       "1234567890123456",
			Encrypted: "7743B07B6BEF2EF9D20D30536C7F2DAAB2CAF8B0EB3DF340A676A63A17434D20",
		},
		{
			Key:       "BFE8916F13D0A7E7",
			Str:       "1234567890123456789",
			Encrypted: "7743B07B6BEF2EF9D20D30536C7F2DAA1C7DC1AA8889044F7F339B50CF1071B4",
		},
	}
	for _, item := range cases {
		t.Run(item.Str, func(t *testing.T) {
			encrypted, err := Encrypt(item.Key, item.Str)
			if err != nil {
				t.Failed()
			}
			actual := strings.ToUpper(hex.EncodeToString(encrypted))
			assert.Equal(t, item.Encrypted, actual)
		})
	}
}
