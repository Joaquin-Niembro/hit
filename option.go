package hit

import (
	"net/http"
	"time"
)

type SendFunc func(*http.Request) Result

type Options struct {
	Concurrency int
	RPS         int
	Send        SendFunc
}

func Defaults() Options {
	return withDefaults(Options{})
}

func withDefaults(o Options) Options {
	if o.Concurrency == 0 {
		o.Concurrency = 1
	}
	if o.Send == nil {
		client := &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost: o.Concurrency,
			},
			CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Timeout: 30 * time.Second,
		}
		o.Send = func(r *http.Request) Result {
			return Send(client, r)
		}
	}

	return o
}
