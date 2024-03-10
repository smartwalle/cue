package cue

import (
	"bufio"
	"os"
	"strings"
)

func Decode(cue string) (*Sheet, error) {
	var file, err = os.Open(cue)
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
		var kIndex = strings.Index(line, " ")

		var key = strings.ToUpper(line[:kIndex])
		var value = line[kIndex+1:]

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
			sheet.setComment(value)
		case "TRACK":
			var values = strings.Split(value, " ")
			sheet.AddTrack(values[0], values[1])
		case "ISRC":
			var isrc = strings.TrimRight(strings.TrimLeft(value, "\""), "\"")
			sheet.setISRC(isrc)
		case "INDEX":
			var values = strings.Split(value, " ")
			sheet.setIndex(values[0], values[1])
		}
	}

	return sheet, nil
}
