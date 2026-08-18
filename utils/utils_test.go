package utils

import (
	"testing"
	"time"
)

func TestGetProjectName(t *testing.T) {
	if got := GetProjectName("/home/user/myproject"); got != "myproject" {
		t.Errorf("GetProjectName = %q, want %q", got, "myproject")
	}
	if got := GetProjectName("/a/b/c"); got != "c" {
		t.Errorf("GetProjectName(\"/a/b/c\") = %q, want %q", got, "c")
	}
	// filepath.Base preserves the root separator for "/"
	if got := GetProjectName("/"); got != "/" {
		t.Errorf("GetProjectName(\"/\") = %q, want %q", got, "/")
	}
}

func TestParseDate(t *testing.T) {
	// empty string returns nil
	if got := parseDate(""); got != nil {
		t.Errorf("parseDate(\"\") = %v, want nil", got)
	}
	// valid ISO8601 with numeric offset returns a time.Time
	got := parseDate("2023-01-30T12:45:00+00:00")
	if _, ok := got.(time.Time); !ok {
		t.Errorf("parseDate(valid) = %v (%T), want time.Time", got, got)
	}
	// invalid input returns the original string
	got2 := parseDate("not-a-date")
	if s, ok := got2.(string); !ok || s != "not-a-date" {
		t.Errorf("parseDate(invalid) = %v (%T), want %q", got2, got2, "not-a-date")
	}
}
