package util_test

import (
	"os"
	"testing"

	"github.com/cloudnative-zoo/go-commons/util"
)

func TestGetEnv(t *testing.T) {
	t.Parallel() // Enable parallel execution of the test function itself

	tests := []struct {
		name  string
		keys  []string
		setup func()
		want  string
	}{
		{
			name: "single key exists",
			keys: []string{"EXISTING_KEY"},
			setup: func() {
				_ = os.Setenv("EXISTING_KEY", "value1")
			},
			want: "value1",
		},
		{
			name: "multiple keys, first exists",
			keys: []string{"EXISTING_KEY1", "EXISTING_KEY2"},
			setup: func() {
				_ = os.Setenv("EXISTING_KEY1", "value1")
				_ = os.Setenv("EXISTING_KEY2", "value2")
			},
			want: "value1",
		},
		{
			name: "multiple keys, second exists",
			keys: []string{"NON_EXISTING_KEY", "EXISTING_KEY"},
			setup: func() {
				_ = os.Setenv("EXISTING_KEY", "value2")
			},
			want: "value2",
		},
		{
			name:  "no keys exist",
			keys:  []string{"NON_EXISTING_KEY1", "NON_EXISTING_KEY2"},
			setup: func() {},
			want:  "",
		},
	}

	for _, tt := range tests {
		tt := tt // Capture range variable
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // Enable parallel execution of individual subtests

			tt.setup()
			if got := util.GetEnv(tt.keys...); got != tt.want {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
