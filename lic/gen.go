package lic

import (
	"encoding/json"
	"fmt"
)

func GenKeyPair() ([]byte, []byte) {
	return nil, nil
}

func ToXml(m map[string]any, publicKey []byte) (string, error) {
	marshal, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	fmt.Println(string(marshal))
	return "", nil
}

func FromXml(xml string, privateKey []byte) {

}
