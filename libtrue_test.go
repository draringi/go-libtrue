package truth

import "testing"

func TestGetTrue(t *testing.T) {
	val := GetTrue()
	expected := true
	if val != expected {
		t.Fatalf("Expected '%v', got '%v'", expected, val)
	}
}
