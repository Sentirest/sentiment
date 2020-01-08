package sentirest

import "testing"

func Test_Unauthorised(t *testing.T) {
  Key = ""
  want := "{\"message\":\"Unauthorized\"}"
  got := GetSentiment("Test")
  if got != want {
    t.Errorf("GetSentiment() = %q, want %q", got, want)
  }
}

func Test_BadKey(t *testing.T) {
  Key = "bad key"
  want := "{\"message\":null}"
  got := GetSentiment("Test")
  if got != want {
    t.Errorf("GetSentiment() = %q, want %q", got, want)
  }
}
