package sprites

import (
	"fmt"
	"testing"
)

// TODO Fix TestDecodeSpecs testing function
func TestDecodeSpecs(t *testing.T) {
	results, err := DecodeSpecs("../../assets/horse/horse.json")
	if err != nil {
		fmt.Println("error:\t", err)
	}

	for _, image := range results.Images {
		fmt.Println(image.Name)
	}
}
