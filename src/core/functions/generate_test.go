package functions

import "testing"

func TestGenerateArguments(t *testing.T) {
	t.SkipNow()
	t.Parallel()
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			GenerateArguments()
		})
	}
}
