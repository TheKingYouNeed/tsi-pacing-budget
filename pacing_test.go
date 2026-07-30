package pacing

import "testing"

func TestBuild(t *testing.T) {
	got := Build(24, 36, 3)
	want := []Checkpoint{{8, 12}, {16, 24}, {24, 36}}
	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("checkpoint %d = %#v, want %#v", i, got[i], want[i])
		}
	}
}
