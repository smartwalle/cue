package cue

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
)

type Middleware func(string) string

func Clear(p string, mids ...Middleware) error {
	var rFile, err = os.OpenFile(p, os.O_RDWR|os.O_SYNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer rFile.Close()

	var reader = bufio.NewReader(rFile)

	var buffer = &bytes.Buffer{}
	var writer = bufio.NewWriter(buffer)

	var line []byte

	for {
		if line, _, err = reader.ReadLine(); err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		var nLine = string(line)

		if nLine == "" {
			continue
		}

		for _, mid := range mids {
			nLine = mid(nLine)
			if nLine == "" {
				break
			}
		}

		if strings.TrimSpace(nLine) == "" {
			continue
		}

		writer.WriteString(nLine)
		writer.WriteString("\n")
	}
	writer.Flush()

	// 写入文件
	rFile.Truncate(0)
	rFile.Seek(0, 0)
	rFile.Write(buffer.Bytes())

	return rFile.Sync()
}
