package main

import (
	"github.com/smartwalle/cue"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var tagInject = regexp.MustCompile("\\[.+\\]$")

func main() {

	filepath.Walk("/Volumes/Data/Download/D/动力火车", func(path string, info fs.FileInfo, err error) error {
		if info == nil {
			return nil
		}
		if info.IsDir() {
			if info.Name() == "[原始文件]" {
				os.RemoveAll(path)
				return nil
			}

			// 替换最后一级目录的字符
			var name = strings.ReplaceAll(info.Name(), "-", " ")
			// 替换最后一级目录的歌手名称
			name = strings.ReplaceAll(name, "动力火车", "")
			// 替换最后一级目录中 [] 内的内容
			name = tagInject.ReplaceAllString(name, "")

			os.Rename(path, filepath.Join(filepath.Dir(path), name))

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
		case ".htm":
			os.Remove(path)
		case ".html":
			os.Remove(path)
		case ".db":
			switch info.Name() {
			case "Thumbs.db":
				os.Remove(path)
			}
		case ".torrent":
			os.Remove(path)
		case ".jpg":
			switch info.Name() {
			case "光盘光谱图.jpg":
				os.Remove(path)
			case "光盘检测图.jpg":
				os.Remove(path)
			}
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
