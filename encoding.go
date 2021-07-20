package cue

import (
	"bufio"
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os"
	"unicode/utf8"
)

func GBKToUTF8(src []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(src), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func GBKFileToUTF8File(p string) error {
	var file, err = os.OpenFile(p, os.O_RDWR|os.O_SYNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	rData, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	if utf8.Valid(rData) {
		return nil
	}

	var reader = transform.NewReader(bytes.NewReader(rData), simplifiedchinese.GBK.NewDecoder())
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	file.Seek(0, 0)

	var writer = bufio.NewWriter(file)
	if _, err = writer.Write(data); err != nil {
		return err
	}
	return writer.Flush()
}
