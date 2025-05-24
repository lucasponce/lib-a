package liba

import "testing"

func TestRecords(t *testing.T) {

	AddRecord("Monday", "Beatles")
	AddRecord("Tuesday", "Rollings")
	AddRecord("Wednesday", "The Doors")

	DeleteRecord("Wednesday")

	version := GetVersion()
	if version != "v0.0.1" {
		t.Errorf("Version %s, want %s", version, "v0.0.1")
	}

	records := GetRecords()
	if len(records) != 2 {
		t.Errorf("GetRecords returned %d tracks, want 2", len(records))
	}
}
