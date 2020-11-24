package utils

import (
	"fmt"
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	cipherText := GeneratePassword("abc")
	fmt.Printf("%s", cipherText)
	if !VerifyPassword(cipherText,"abc"){
		t.Error()
	}
}
