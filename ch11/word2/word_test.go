package word

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
	"unicode"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"empty string", "", true},
		{"a", "a", true},
		{"aa", "aa", true},
		{"ab", "ab", false},
		{"kayak", "kayak", true},
		{"detartrated", "detartrated", true},
		{"phrase", "A man, a plan, a canal: Panama", true},
		{"phrase", "Evil I did dwell; lewd did I live.", true},
		{"phrase", "Able was I ere I saw Elba", true},
		{"french", "été", true},
		{"french2", "Et se resservir, ivresse reste.", true},
		{"non-palindrome", "palindrome", false}, // non-palindrome
		{"semi-palindrome", "desserts", false},  // semi-palindrome
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.input); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func addPunctionToPalindrome(runes []rune, rng *rand.Rand) []rune {
	n := len(runes)
	if n == 0 {
		return runes
	}
	var result []rune
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			result = append(result, rune(' '))
		} else {
			result = append(result, rune(','))
		}
		result = append(result, runes[i])
	}
	return result
}

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[n-i-1] = r
	}
	return string(addPunctionToPalindrome(runes, rng))
}

func randomNonPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) + 2
	start := rng.Intn(0x100)
	step := 1
	runes := make([]rune, n)
	value := start
	for i := 0; i < n; {
		r := rune(value)
		if unicode.IsLetter(r) {
			runes[i] = rune(value)
			i++
		}
		value += step
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		t.Run("random palindrome", func(t *testing.T) {
			if !IsPalindrome(p) {
				t.Errorf("IsPalindrome(%q) = false", p)
			}
		})
	}
}
func TestRandomNonPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomNonPalindrome(rng)
		t.Run("random non palindrome", func(t *testing.T) {
			if IsPalindrome(p) {
				t.Errorf("IsPalindrome(%q) = true", p)
			}
		})
	}
}

func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(IsPalindrome("Palindrome"))
	// Output:
	// true
	// false
}
