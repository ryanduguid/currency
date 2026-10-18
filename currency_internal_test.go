// SPDX-License-Identifier: MIT

package currency

import "testing"

func TestContainsLocale(t *testing.T) {
	tests := []struct {
		name     string
		list     string
		localeID string
		want     bool
	}{
		{"empty list", "", "en", false},
		{"empty locale", "en fr", "", false},
		{"single entry", "en", "en", true},
		{"first entry", "en fr", "en", true},
		{"middle entry", "en fr de", "fr", true},
		{"last entry", "en fr", "fr", true},
		{"prefix only", "en-US", "en", false},
		{"suffix only", "en-US", "US", false},
		{"later exact match", "en-US en", "en", true},
		{"absent entry", "en fr", "de", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsLocale(tt.list, tt.localeID); got != tt.want {
				t.Errorf("containsLocale(%q, %q) = %v, want %v", tt.list, tt.localeID, got, tt.want)
			}
		})
	}
}
