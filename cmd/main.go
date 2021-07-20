package main

import (
	"github.com/smartwalle/cue"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	filepath.Walk("./", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		var ext = filepath.Ext(info.Name())
		switch ext {
		case ".cue":
			cue.GBKFileToUTF8File(path)
			cue.Clear(path, TrimRemComment)
		case ".log":
			os.Remove(path)
		case ".url":
			os.Remove(path)
		case ".db":
			switch info.Name() {
			case "Thumbs.db":
				os.Remove(path)
			}
		case ".torrent":
			os.Remove(path)
		case ".txt":
			switch info.Name() {
			case "免责声明.txt":
				os.Remove(path)
			case "说明.txt":
				os.Remove(path)
			default:
				cue.GBKFileToUTF8File(path)
			}
		}
		return nil
	})
}

// TrimRemComment 去除 REM COMMENT
func TrimRemComment(s string) string {
	if strings.HasPrefix(s, "REM COMMENT") {
		return ""
	}
	return s
}
