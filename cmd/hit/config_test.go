package main

import (
	"io"
	"testing"
)

type parseArgsTest struct {
	name string
	args []string
	want config
}

func TestParseArgsValidInput(t *testing.T) {
	t.Parallel()
	tests := []parseArgsTest{
		{
			name: "all_flags",
			args: []string{"-n=10", "-c=5", "-rps=5", "http://test"},
			want: config{n: 10, c: 5, rps: 5, url: "http://test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got config
			if err := parseArgs(&got, tt.args, io.Discard); err != nil {
				t.Fatalf("parseArgs() error = %v, want no error", err)
			}
			if got != tt.want {
				t.Errorf("parseArgs() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseArgsInvalidInput(t *testing.T) {
	t.Parallel()
	tests := []parseArgsTest{
		{
			name: "all_flags",
			args: []string{"-n=10", "-c=5", "-rps=5", "invalid"},
		},
		{name: "n_syntax", args: []string{"-n=ONE", "http://test"}},
		{name: "n_zero", args: []string{"-n=0", "http://test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got config
			if err := parseArgs(&got, tt.args, io.Discard); err == nil {
				t.Fatalf("parseArgs() error = %v, want error", err)
			}
		})
	}
}
