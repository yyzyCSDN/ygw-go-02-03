package stock

import (
	"testing"
)

func TestWindowSmoke(t *testing.T) {
	got, err := Window([]int{1, 2, 3}, 0, 0)
	if err != nil || len(got) != 0 {
		t.Fatalf("Window = %v, %v", got, err)
	}
	if _, err := Window([]int{1}, -1, 1); err == nil {
		t.Fatal("expected error")
	}
}
