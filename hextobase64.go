/* Set 1 Challenge 1 - Convert hex to base64 */

package main

import (
	"encoding/hex"
	"encoding/base64"
)

func hexToBase64(hexVal string) string {
	byteEquivalent := hexToByte(hexVal)
	base64Equivalent := byteToBase64(byteEquivalent)

	return base64Equivalent
}

func hexToByte(hexVal string) []byte {
	byteEquivalent, err := hex.DecodeString(hexVal)
	if err != nil {
		panic(err)
	}

	return byteEquivalent
}

func byteToBase64(byteVal []byte) string {
	base64Equivalent := base64.StdEncoding.EncodeToString(byteVal)

	return base64Equivalent
}