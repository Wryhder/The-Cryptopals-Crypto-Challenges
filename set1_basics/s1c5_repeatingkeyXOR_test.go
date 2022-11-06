package set1_basics

import "testing"

func TestRepeatingKeyXOR(t *testing.T) {
	var openingStanza string = `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`

	actual := RepeatingKeyXOR(openingStanza, "ICE")
    expected := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a2622632427276" +
	"5272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

    if actual != expected {
        t.Errorf("actual %q, expected %q", actual, expected)
    }
}