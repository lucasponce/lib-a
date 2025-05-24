package liba

import (
	"fmt"
	"github.com/lucasponce/shared-x/pkg/tracks"
)

var VERSION = "v0.0.1"

func GetVersion() string {
	return VERSION
}

func AddRecord(id string, content string) {
	tracks.AddTrack(id, fmt.Sprintf("lib-a %s: %s", VERSION, content))
}

func DeleteRecord(id string) {
	tracks.RemoveTrack(id)
}

func GetRecords() map[string]string {
	tracks.DumpTracks()
	return tracks.GetTracks()
}
