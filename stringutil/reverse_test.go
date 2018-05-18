package stringutil

import "testing"

func TestReverse(t *testing.T) {
	t.Run("reverse lib", func (t *testing.T) {
		got := Reverse("Hello")
		expected := "olleH"
		if got != expected {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})
}