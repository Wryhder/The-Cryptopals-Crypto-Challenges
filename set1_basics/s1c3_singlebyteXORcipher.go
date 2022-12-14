/* Set 1 Challenge 3 - Single-byte XOR cipher */

/*

Approach (based on https://www.codementor.io/@arpitbhayani/deciphering-single-byte-xor-ciphertext-17mtwlzh30)

1. Convert hex-encoded ciphertext to byte
2. Brute-force decryption: XOR ciphertext with all possible key values (1 - 255)
3. Score results based on letter frequency (uses this table: http://mathcenter.oxford.emory.edu/site/math125/englishLetterFreqScores/) to find most probable answer. Formula for scoring text is from article linked above (Fitting Quotient)

*/

package set1_basics

import (
	"fmt"
	"math"
	"strings"
)

func countChar(s string) map[string]int {
	var charCount = make(map[string]int)

	for _, char := range s {
		charCount[strings.ToLower(string(char))] += 1
	}

	return charCount
}

var englishLetterFreqScores = map[string]float64{
	" ": 0.15,
	"a": 0.08167,
	"b": 0.01492,
	"c": 0.02782,
	"d": 0.04253,
	"e": 0.12702,
	"f": 0.02228,
	"g": 0.02015,
	"h": 0.06094,
	"i": 0.06966,
	"j": 0.00153,
	"k": 0.00772,
	"l": 0.04025,
	"m": 0.02406,
	"n": 0.06749,
	"o": 0.07507,
	"p": 0.01929,
	"q": 0.00095,
	"r": 0.05987,
	"s": 0.06327,
	"t": 0.09056,
	"u": 0.02758,
	"v": 0.00978,
	"w": 0.02360,
	"x": 0.00150,
	"y": 0.01974,
	"z": 0.00074,
}

/*
This function scores strings based on letter frequency
(when compared with the letter frequency scores of standard English strings),
to determine how likely a string is to being correct English.
*/
func TextScorer(text string) float64 {
	lengthOfText := len(text)

	letterFreqScoresInText := func() map[string]float64 {
		var freqMap = make(map[string]float64)
		charCount := countChar(text)

		for char, _ := range englishLetterFreqScores {
			charCode := []rune(char)[0]

			if (charCode == ' ') || (charCode >= 'a' && charCode <= 'z')  {
				freqMap[char] = float64(charCount[char]) * 100 / float64(lengthOfText)
			} else {
				continue
			}	
		}
		return freqMap
	}()

	// Calculate the sum of the absolute differences between the frequencies
	// of each letter in the text and the corresponding frequency of the letter
	// in standard English usage
	sumOfAbsDiff := func() float64 {
		total := 0.0

		for textChar, _ := range letterFreqScoresInText {
			total += math.Abs(englishLetterFreqScores[textChar] - letterFreqScoresInText[textChar])
		}

		return total
	}()

	// Calculate the average of the absolute differences
	textScore := sumOfAbsDiff / float64(len(letterFreqScoresInText))

	return textScore
}

func SingleByteXORCipher(text []byte) (string, string) {
	lengthOfText := len(text)
	XORCombination := make([]byte, lengthOfText)

	/*
		allResults maps each score to associated decrypted text and key used for decryption

		Example:

		{
			"[score]": {
				"key": "[key]",
				"decryptedText":[decryptedText],
			},
		}
	*/
	var allResults = make(map[string]map[string]string)
	highestTextScore := 0.0

	// Brute-force decryption: XOR ciphertext with all possible key values (1 - 255)
	for key := 0; key <= 255; key++ {

		for j := 0; j < lengthOfText; j++ {
			XORCombination[j] = text[j] ^ uint8(key)
		}

		// Score results based on letter frequency
		currentScore := TextScorer(string(XORCombination))

		allResults[fmt.Sprintf("%.15f", currentScore)] = map[string]string{
			"key":           string(uint8(key)),
			"decryptedText": string(XORCombination),
		}

		if currentScore < highestTextScore {
			continue
		} else {
			highestTextScore = currentScore
		}
	}

	encryptionKey := allResults[fmt.Sprintf("%.15f", highestTextScore)]["key"]
	decryptedText := allResults[fmt.Sprintf("%.15f", highestTextScore)]["decryptedText"]

	return encryptionKey, decryptedText
}
