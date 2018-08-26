package cake

import "testing"

func Test_baker(t *testing.T) {
	cake := <-iced
	if cake.state != "iced" {
		t.Errorf("expected iced cake got %v\n", cake.state)
	}
}
