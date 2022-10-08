package algorithms

import "testing"

func FuzzRot13(f *testing.F) {
	rot13 := NewRot13()

	f.Add("hello world")
	f.Add("foo bar")
	f.Add("Mohammed Adekunle")

	f.Fuzz(func(t *testing.T, input string) {
		encrypted := rot13.Encrypt(input)
		decrypted := rot13.Decrypt(encrypted)

		if decrypted != input {
			t.Errorf("expected %v to match %v", decrypted, input)
		}
	})
}