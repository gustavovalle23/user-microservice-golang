package tests

import "testing"

func Add(x, y int) (res int) {
	return x + y
}

func TestAdd(t *testing.T) {
	got := Add(2, 2)
	want := 4
	t.Log("Result: ")
	t.Log(got)

	if got != want {
		t.Errorf("Got %q, wanted %q", got, want)
	}
}
