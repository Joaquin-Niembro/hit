package hit

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Send(client *http.Client, req *http.Request) Result {
	started := time.Now()
	var (
		bytes int64
		code  int
	)
	resp, err := client.Do(req)
	if err == nil { // no error
		defer resp.Body.Close()
		code = resp.StatusCode
		bytes, err = io.Copy(io.Discard, resp.Body)
	}
	return Result{
		Duration: time.Since(started),
		Bytes:    bytes,
		Status:   code,
		Error:    err,
	}
}

func SendN(ctx context.Context, n int, req *http.Request, opts Options) (Results, error) {
	ctx, cancel := context.WithCancel(ctx)
	opts = withDefaults(opts)
	if n <= 0 {
		return nil, fmt.Errorf("n must be positive: got %d", n)
	}

	results := runPipeline(ctx, n, req, opts)

	return func(yield func(Result) bool) {
		defer cancel()
		for result := range results {
			if !yield(result) {
				return
			}
		}
	}, nil
}
