/* Set 2 Challenge 10 - Implement CBC mode */

package set2_blockcrypto

import (
	"bytes"
	"crypto/aes"
	set1 "wryhder/cryptopals-crypto-challenges/set1_basics"
	utils "wryhder/cryptopals-crypto-challenges/utilities"
)

/*
Approach (encryption):
- Pad plaintext
- Chunkify plaintext
- XOR each block with previous (Fixed XOR), then encrypt (AES in ECB mode)
*/

// XOR each block with previous, then encrypt
func EncryptAES128_CBC(plaintext, key string, IV []byte) string {
	blockSize := aes.BlockSize
	paddedPlaintext := PKCS7padding([]byte(plaintext), blockSize)
	plainTextBlocks := utils.ChunkifyText(paddedPlaintext, blockSize)

	var cipherTextBlocks [][]byte
	var XORed []byte
	var encrypted string
	for index, block := range plainTextBlocks {

		if index == 0 {
			XORed = set1.FixedXOR(IV, block)
		} else {
			previousCiphertextBlock := cipherTextBlocks[index - 1]
			XORed = set1.FixedXOR(previousCiphertextBlock, block)
		}

		encrypted = set1.EncryptAES128_ECB(string(XORed), key)
		cipherTextBlocks = append(cipherTextBlocks, []byte(utils.DecodeBase64(encrypted)))
	}

	return set1.ByteToBase64(bytes.Join(cipherTextBlocks, []byte("")))
}

// Decrypt each block before XORing with previous block to recover plaintext
func DecryptAES128_CBC(ciphertext, key string, IV []byte) string {
	decodedCiphertext := utils.DecodeBase64(ciphertext)
	cipherTextBlocks := utils.ChunkifyText([]byte(decodedCiphertext), aes.BlockSize)

	var plainTextBlocks [][]byte
	var XORed []byte
	var decrypted string
	for index, block := range cipherTextBlocks {
		decrypted = set1.DecryptAES128_ECB(set1.ByteToBase64(block), key)

		if index == 0 {
			XORed = set1.FixedXOR(IV, []byte(decrypted))
		} else {
			previousCiphertextBlock := cipherTextBlocks[index - 1]
			XORed = set1.FixedXOR(previousCiphertextBlock, []byte(decrypted))
			
		}
		
		plainTextBlocks = append(plainTextBlocks, XORed)
	}

	return string(bytes.Join(plainTextBlocks, []byte("")))
}
