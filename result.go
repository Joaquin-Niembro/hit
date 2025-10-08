package hit

import (
	"iter"
	"net/http"
	"time"
)

type Result struct {
	Status   int
	Bytes    int64
	Duration time.Duration
	Error    error
}

type Results iter.Seq[Result] // same as: type Results func(yield func(Result) bool)

type Summary struct {
	Requests int           // Requests is the total number of requests made
	Errors   int           // Errors is the total number of failed requests
	Bytes    int64         // Bytes is the total number of bytes transferred
	RPS      float64       // RPS is the number of requests sent per second
	Duration time.Duration // Duration is the total time taken by requests
	Fastest  time.Duration // Fastest is the fastest request duration
	Slowest  time.Duration // Slowest is the slowest request duration
	Success  float64       // Success is the ratio of successful requests
}

func Summarize(results Results) Summary {
	var s Summary
	if results == nil {
		return s
	}
	started := time.Now()
	for r := range results {
		s.Requests++
		s.Bytes += r.Bytes

		if r.Error != nil || r.Status != http.StatusOK {
			s.Errors++
		}
		if s.Fastest == 0 {
			s.Fastest = r.Duration
		}
		if r.Duration < s.Fastest {
			s.Fastest = r.Duration
		}
		if r.Duration > s.Slowest {
			s.Slowest = r.Duration
		}
	}

	if s.Requests > 0 {
		s.Success = (float64(s.Requests-s.Errors) / float64(s.Requests)) * 100
	}
	s.Duration = time.Since(started)
	s.RPS = float64(s.Requests) / s.Duration.Seconds()

	return s
}
