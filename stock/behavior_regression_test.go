package stock

import (
	"testing"
)

func TestRegressionBehavior(t *testing.T) {
	got, err := Window([]string{"a", "b", "c", "d"}, 1, 2)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || got[0] != "b" || got[1] != "c" {
		t.Fatalf("Window = %v; want [b c]", got)
	}
}
