package passgen

import (
	"fmt"
)

//Algorithm defines behaviour of an algorithm that can be used in PassGen.
type Algorithm interface {
	Encrypt(data []byte, key []byte) ([]byte, error)
	Decrypt(encrypted []byte, key []byte) ([]byte, error)
}

func NewAlgorithm(algo string) Algorithm {
	switch algo {
	case "aes":
		return new(AES)
	default:
		return new(AES)
	}
}

//PassGen ...
type PassGen struct {
	Master    []byte
	Algorithm Algorithm
}

func NewPassGen(master []byte, algo Algorithm) *PassGen {
	return &PassGen{Master: master, Algorithm: algo}
}

//GenFor generates a new password based on given uniqueName, and configured salt and master.
func (p *PassGen) GenFor(uniqueName string) ([]byte, error) {
	data := []byte(fmt.Sprintf("%s-%s", uniqueName, string(p.Master)))
	return p.Algorithm.Encrypt(data, p.Master)
}
