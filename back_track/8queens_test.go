package back_track

import "testing"

func TestQueensSolutions(t *testing.T) {
	var (
		isPrintIncomplete = false
	)
	QueensSolutions(8, isPrintIncomplete)
}
