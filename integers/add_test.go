package integers

import "testing"

func TestAdder(t *testing.T) {
    assertCorrectMessage := func(t *testing.T, got, want string) {
        t.Helper()
        if got != want {
            t.Errorf("got '%s' want '%s'", got, want)
        }
    }

    t.Run("Adder test", func(t *testing.T) {
        sum := Add(2, 2)
    	expected := 4
        assertCorrectMessage(t, sum, expected)
    })
}