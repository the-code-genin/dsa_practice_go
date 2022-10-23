package algorithms

import (
	"fmt"
	"testing"
)

func TestGenerateRandomNumber(t *testing.T) {
	for i := int64(1); i <= 100; i++ {
		v, err := GenerateRandomNumber(i * 10)
		if err != nil {
			t.Error(err)
		}
	
		fmt.Printf("Rand Int: %v\n", v)
	}
}
