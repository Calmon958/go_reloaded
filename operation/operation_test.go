package reloaded

import (
	"fmt"
	"testing"
)

func TestHex(t *testing.T) {
	input := "1E (hex) files were added"
	expected := "30 files were added"
	output := Operation(input)

	if expected != output {
		fmt.Printf("Hex failed: Got: %s\n Expected: %s", output, expected)
		return
	}
}

func TestBin(t *testing.T) {
	input := "It has been 10 (bin) years"
	expected := "It has been 2 years"
	output := Operation(input)

	if expected != output {
		fmt.Printf("Bin failed: Got: %s\n Expected: %s", output, expected)
		return
	}
}

func TestUp(t *testing.T) {
	input := "Ready, set, go (up) !"
	expected := "Ready, set, GO!"
	output := Operation(input)

	if expected != output {
		fmt.Printf("toUpper failed: Got: %s\n Expected: %s", output, expected)
		return
	}
}

func TestLow(t *testing.T) {
	input := "I should stop SHOUTING (low)"
	expected := "I should stop shouting"
	output := Operation(input)

	if expected != output {
		fmt.Printf("toLower failed: Got: %s\n Expected: %s", output, expected)
		return
	}
}

func TestPunctuation(t *testing.T) {
	input := "I was sitting over there ,and then BAMM !!"
	expected := "I was sitting over there, and then BAMM!!"
	output := Operation(input)

	if expected != output {
		fmt.Printf("Punctution failed: Got: %s\n Expected: %s", output, expected)
		return
	}
}
