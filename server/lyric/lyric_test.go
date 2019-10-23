package lyric

import (
	"testing"
	"time"
)

func TestSubStrTime(t *testing.T) {
	type args struct {
		t1 string
		t2 string
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{"test-float1-sec", args{"01:02.1","01:02.3"}, 0.2*1e9},
		{"test-float2-sec", args{"01:02.11","01:02.22"}, 0.11*1e9},
		{"test-sec", args{"01:02","01:03"}, 1*1e9},
		{"test-min", args{"01:02","02:02"}, 60*1e9},
		{"test-hour", args{"01:02:02","02:02:02"}, 60*60*1e9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubStrTime(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("SubStrTime() = %v, want %v", got, tt.want)
			}
		})
	}
}