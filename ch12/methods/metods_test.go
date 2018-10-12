package methods

import (
	"strings"
	"testing"
	"time"
)

func TestMethods(t *testing.T) {
	Print(time.Hour)
	Print(new(strings.Replacer))
}
