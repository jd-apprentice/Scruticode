package functions

import "testing"

func TestReadConfigFile(t *testing.T) {
	t.SkipNow()
	t.Parallel()
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := ReadConfigFile(); got != tt.want {
				t.Errorf("ReadConfigFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
