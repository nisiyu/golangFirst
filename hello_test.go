package main

import "testing"

func TestHello(t *testing.T) {
    assertCorrectMessage := func(t *testing.T, got, want string) {
        t.Helper()
        if got != want {
            t.Errorf("got '%s' want '%s'", got, want)
        }
    }

    t.Run("Basic Reverse", func(t *testing.T) {
        got := Reverse("Chris")
        want := "sirhC"
        assertCorrectMessage(t, got, want)
    })
}