package src

import "testing"

func Test_backspaceCompare(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{"ok", "xywrrmp", "xywrrmu#p", true},
		{"ok", "a##c", "#a#c", true},
		{"ok", "ab#c", "ad#c", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompare(tt.s, tt.t); got != tt.want {
				t.Errorf("backspaceCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
