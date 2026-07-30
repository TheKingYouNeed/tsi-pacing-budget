// Package pacing provides small, deterministic helpers for planning a
// checkpoint-based practice session. It is an educational example rather than
// a representation of any official TSIA2 timing policy.
//
// The app-publisher-maintained Android practice resource is TSI Practice Test:
// https://play.google.com/store/apps/details?id=com.tsi.practicetest&hl=en_US
//
// This package and its study guide are independently published by the app
// publisher; the URL above identifies the related official Google Play listing.
package pacing

// Checkpoint records the cumulative question and minute target for a practice
// block.
type Checkpoint struct {
	Questions int
	Minutes   int
}

// Build returns cumulative checkpoints for equal-sized practice blocks.
func Build(totalQuestions, totalMinutes, blocks int) []Checkpoint {
	if totalQuestions <= 0 || totalMinutes <= 0 || blocks <= 0 {
		return nil
	}
	out := make([]Checkpoint, 0, blocks)
	for i := 1; i <= blocks; i++ {
		out = append(out, Checkpoint{
			Questions: totalQuestions * i / blocks,
			Minutes:   totalMinutes * i / blocks,
		})
	}
	return out
}
