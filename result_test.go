package hit

import (
	"slices"
	"testing"
	"time"
)

func TestSummarizeFastestResult(t *testing.T) {
	results := []Result{
		{Duration: 2 * time.Second},
		{Duration: 5 * time.Second},
	}
	type sliceOfResultsType []Result
	sum := Summarize(Results(slices.Values[sliceOfResultsType](results)))
	if sum.Fastest != 2*time.Second {
		t.Errorf("Fastest=%v; want 2s", sum.Fastest)
	}
}

func TestSummarizeNilResults(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Errorf("should not panic: %v", err)
		}
	}()
	_ = Summarize(nil)
}
