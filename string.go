package cue

import "strings"

type String []string

func NewString(vs ...string) String {
	return vs
}

func (s String) writeTo(w Writer) error {
	if _, err := w.WriteString(strings.Join(s, "")); err != nil {
		return err
	}
	return nil
}
