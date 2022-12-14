package set1_basics

import (
	"testing"
	utils "wryhder/cryptopals-crypto-challenges/utilities"
)

func TestEncryptAES128_ECB(t *testing.T) {
	plaintext := "The afternoon is as hot as a fireplace. Are these walls purple??"
    
	actual := EncryptAES128_ECB(plaintext, "YELLOW SUBMARINE")
    expected := "3kHmcp52LR447kBgC/tWQK2bO78YBSfA2M62iJd/B8gmBcXdax38DuUaC3bgpyuejFtWXieUHN7QUm3JA00c2g=="

    if actual != expected {
        t.Errorf("actual %q, expected %q", actual, expected)
    }
}

func TestDecryptAES128_ECB(t *testing.T) {
	// Test 1
	AESinECBFileContent := utils.ReadTextFile("../data/s1c7_encodedAESinECBmodesample.txt")

	actual_1 := DecryptAES128_ECB(AESinECBFileContent, "YELLOW SUBMARINE")
    expected_1 := `I'm back and I'm ringin' the bell 
A rockin' on the mike while the fly girls yell 
In ecstasy in the back of me 
Well that's my DJ Deshay cuttin' all them Z's 
Hittin' hard and the girlies goin' crazy 
Vanilla's on the mike, man I'm not lazy. 

I'm lettin' my drug kick in 
It controls my mouth and I begin 
To just let it flow, let my concepts go 
My posse's to the side yellin', Go Vanilla Go! 

Smooth 'cause that's the way I will be 
And if you don't give a damn, then 
Why you starin' at me 
So get off 'cause I control the stage 
There's no dissin' allowed 
I'm in my own phase 
The girlies sa y they love me and that is ok 
And I can dance better than any kid n' play 

Stage 2 -- Yea the one ya' wanna listen to 
It's off my head so let the beat play through 
So I can funk it up and make it sound good 
1-2-3 Yo -- Knock on some wood 
For good luck, I like my rhymes atrocious 
Supercalafragilisticexpialidocious 
I'm an effect and that you can bet 
I can take a fly girl and make her wet. 

I'm like Samson -- Samson to Delilah 
There's no denyin', You can try to hang 
But you'll keep tryin' to get my style 
Over and over, practice makes perfect 
But not if you're a loafer. 

You'll get nowhere, no place, no time, no girls 
Soon -- Oh my God, homebody, you probably eat 
Spaghetti with a spoon! Come on and say it! 

VIP. Vanilla Ice yep, yep, I'm comin' hard like a rhino 
Intoxicating so you stagger like a wino 
So punks stop trying and girl stop cryin' 
Vanilla Ice is sellin' and you people are buyin' 
'Cause why the freaks are jockin' like Crazy Glue 
Movin' and groovin' trying to sing along 
All through the ghetto groovin' this here song 
Now you're amazed by the VIP posse. 

Steppin' so hard like a German Nazi 
Startled by the bases hittin' ground 
There's no trippin' on mine, I'm just gettin' down 
Sparkamatic, I'm hangin' tight like a fanatic 
You trapped me once and I thought that 
You might have it 
So step down and lend me your ear 
'89 in my time! You, '90 is my year. 

You're weakenin' fast, YO! and I can tell it 
Your body's gettin' hot, so, so I can smell it 
So don't be mad and don't be sad 
'Cause the lyrics belong to ICE, You can call me Dad 
You're pitchin' a fit, so step back and endure 
Let the witch doctor, Ice, do the dance to cure 
So come up close and don't be square 
You wanna battle me -- Anytime, anywhere 

You thought that I was weak, Boy, you're dead wrong 
So come on, everybody and sing this song 

Say -- Play that funky music Say, go white boy, go white boy go 
play that funky music Go white boy, go white boy, go 
Lay down and boogie and play that funky music till you die. 

Play that funky music Come on, Come on, let me hear 
Play that funky music white boy you say it, say it 
Play that funky music A little louder now 
Play that funky music, white boy Come on, Come on, Come on 
Play that funky music 
` + "\x04\x04\x04\x04"

	if actual_1 != expected_1 {
        t.Errorf("actual: %q\n\nexpected: %q\n", actual_1, expected_1)
    }

	// Test 2
	ciphertext := "3kHmcp52LR447kBgC/tWQK2bO78YBSfA2M62iJd/B8gmBcXdax38DuUaC3bgpyuejFtWXieUHN7QUm3JA00c2g=="

	actual_2 := DecryptAES128_ECB(ciphertext, "YELLOW SUBMARINE")
    expected_2 := "The afternoon is as hot as a fireplace. Are these walls purple??"
    if actual_2 != expected_2 {
        t.Errorf("actual %q, expected %q", actual_2, expected_2)
    }
}