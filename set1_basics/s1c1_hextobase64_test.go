package set1_basics

import "testing"

func TestHexToBase64(t *testing.T) {
	hexVal := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
    
	actual := HexToBase64(hexVal)
    expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

    if actual != expected {
        t.Errorf("actual %q, expected %q", actual, expected)
    }
}