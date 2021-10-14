package services

import (
	"testing"
)

func TestHomeSvc(t *testing.T) {
	hs := NewHomeSvc()
	expected := "At this stage, available endpoints are /read and /fetch."

	t.Run("home message", func(t *testing.T) {
		msg := hs.Inform()
		if msg != expected {
			t.Fatalf("message is not as expected:\ngot:%v\nwant:%v\n", msg, expected)
		}
	})
}
