package cue

import (
	"bufio"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"unicode/utf8"
)

type Writer interface {
	WriteString(s string) (int, error)

	Flush() error
}

type writer struct {
	writer  *bufio.Writer
	decoder *encoding.Decoder
}

func NewWriter(w io.Writer) *writer {
	var nWriter *bufio.Writer
	if nw, ok := w.(*bufio.Writer); ok {
		nWriter = nw
	} else {
		nWriter = bufio.NewWriter(w)
	}

	return &writer{writer: nWriter, decoder: simplifiedchinese.GBK.NewDecoder()}
}

func (w *writer) WriteString(s string) (int, error) {
	if utf8.ValidString(s) {
		return w.writer.WriteString(s)
	}
	var ns, n, err = transform.String(w.decoder, s)
	if err != nil {
		return n, err
	}
	return w.writer.WriteString(ns)
}

func (w *writer) Flush() error {
	return w.writer.Flush()
}
