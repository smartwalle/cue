package cue

import (
	"bufio"
	"os"
	"strings"
)

func Decode(filename string) (*Sheet, error) {
	var file, err = os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var sheet = NewSheet()

	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			return nil, err
		}

		var line = strings.TrimSpace(scanner.Text())
		var kEnd = strings.Index(line, " ")

		var key = strings.ToUpper(line[:kEnd])
		var value = line[kEnd+1:]

		switch key {
		case "TITLE":
			var title = strings.TrimRight(strings.TrimLeft(value, "\""), "\"")
			sheet.setTitle(title)
		case "PERFORMER":
			var performer = strings.TrimRight(strings.TrimLeft(value, "\""), "\"")
			sheet.setPerformer(performer)
		case "SONGWRITER":
			var writer = strings.TrimRight(strings.TrimLeft(value, "\""), "\"")
			sheet.setSongWriter(writer)
		case "CATALOG":
			var catalog = strings.TrimRight(strings.TrimLeft(value, "\""), "\"")
			sheet.setCatalog(catalog)
		case "CDTEXTFILE":
			var cdFile = strings.TrimRight(strings.TrimLeft(value, "\""), "\"")
			sheet.setCDTextFile(cdFile)
		case "FILE":
			var sep = strings.LastIndex(value, " ")
			sheet.setFile(value[1:sep-1], value[sep+1:])
		case "REM":
			decodeComment(sheet, value)
		case "TRACK":
			var values = strings.Split(value, " ")
			sheet.AddTrack(values[0], values[1])
		case "ISRC":
			var isrc = strings.TrimRight(strings.TrimLeft(value, "\""), "\"")
			sheet.setISRC(isrc)
		case "FLAGS":
			sheet.setFlags(value)
		case "INDEX":
			var values = strings.Split(value, " ")
			sheet.setIndex(values[0], values[1])
		}
	}
	return sheet, nil
}

func decodeComment(sheet *Sheet, comment string) {
	if comment[0] == '"' {
		sheet.current.setComment("", comment)
	} else {
		var kEnd = strings.Index(comment, " ")
		if kEnd > 0 {
			var key = comment[:kEnd]
			var value = comment[kEnd+1:]
			sheet.current.setComment(key, value)
		} else {
			var key = strings.ToUpper(comment)
			switch key {
			case "GENRE":
				sheet.current.setComment(key, "")
			case "DISCID":
				sheet.current.setComment(key, "")
			case "DATE":
				sheet.current.setComment(key, "")
			case "COMMENT":
				sheet.current.setComment(key, "")
			default:
				sheet.current.setComment("", comment)
			}
		}
	}
}
