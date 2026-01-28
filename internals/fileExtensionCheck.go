package internals

import (
	"path/filepath"
	"strings"
)

func CheckFileExtension(mediaType string, fileName string) bool{
	ext := strings.ToLower(filepath.Ext(fileName))
	switch mediaType {
	case VideoType:
		switch ext {
		case ".mp4", ".mkv", ".mov", ".avi", ".webm":
			return true
		default:
			return false
		}
	case AudioType:
		switch ext {
		case ".mp3", ".wav", ".aac", ".flac", ".ogg", ".m4a", ".opus", ".wma", ".aiff":
			return true
		default:
			return false
		}
	}
	return false;
}