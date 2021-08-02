package main

import (
	"github.com/smartwalle/cue"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	filepath.Walk("/Volumes/SmartWalle/音乐/未命名文件夹/齐秦", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if strings.HasPrefix(info.Name(), ".") {
			os.Remove(path)
			return nil
		}

		var ext = strings.ToLower(filepath.Ext(info.Name()))
		switch ext {
		case ".cue":
			cue.GBKFileToUTF8File(path)
			cue.Clear(path, TrimRemComment, TrimRemGenre, FixFileWave)
			if info.Name() != "CDImage.cue" {
				os.Rename(path, filepath.Join(filepath.Dir(path), "CDImage.cue"))
			}
		case ".wav":
			if info.Name() != "CDImage.wav" {
				os.Rename(path, filepath.Join(filepath.Dir(path), "CDImage.wav"))
			}
		case ".flac":
			if info.Name() != "CDImage.flac" {
				os.Rename(path, filepath.Join(filepath.Dir(path), "CDImage.flac"))
			}
		case ".ape":
			if info.Name() != "CDImage.ape" {
				os.Rename(path, filepath.Join(filepath.Dir(path), "CDImage.ape"))
			}
		case ".log":
			os.Remove(path)
		case ".url":
			os.Remove(path)
		case ".lnk":
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
				if strings.Contains(path, "下载") {
					os.Remove(path)
				} else {
					cue.GBKFileToUTF8File(path)
				}
			}
		}
		return nil
	})
}

// TrimRemComment 去除 REM COMMENT
func TrimRemComment(s string) string {
	if strings.HasPrefix(strings.TrimSpace(s), "REM COMMENT") {
		return ""
	}
	return s
}

func TrimRemGenre(s string) string {
	if strings.HasPrefix(strings.TrimSpace(s), "REM GENRE") {
		return ""
	}
	return s
}

var FileWave = regexp.MustCompile(`FILE "([^\n]+).(wav|flac|ape)" WAVE`)

func FixFileWave(s string) string {
	if strings.HasPrefix(s, "FILE \"") {
		var matches = FileWave.FindStringSubmatch(s)
		if len(matches) == 3 {
			s = `FILE "CDImage.` + matches[2] + `" WAVE`
			return s
		}
		return s
	}
	return s
}
