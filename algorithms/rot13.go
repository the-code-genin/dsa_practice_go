package algorithms

import "strings"

type Rot13 struct {
	alphabets []string
}

func (r *Rot13) getAlphabetIndex(letter string) int8 {
	for i := 0; i < len(r.alphabets); i++ {
		if r.alphabets[i] == letter {
			return int8(i)
		}
	}

	return -1
}

func (r *Rot13) displace(data string) string {
	buffer := strings.Split(data, "")

	for i := 0; i < len(buffer); i++ {
		letter := buffer[i]
		if index := r.getAlphabetIndex(letter); index != -1 {
			displacedIndex := (index + 13) % 26
			buffer[i] = r.alphabets[displacedIndex]
		}
	}

	return strings.Join(buffer, "")
}

func (r *Rot13) Encrypt(data string) string {
	return r.displace(data)
}

func (r *Rot13) Decrypt(data string) string {
	return r.displace(data)
}

func NewRot13() *Rot13 {
	return &Rot13{
		[]string{
			"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
			"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		},
	}
}
