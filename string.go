package cue

type String struct {
	key     string
	value   string
	prefix  string
	suffix  string
	newline bool
}

func NewString(key, value, prefix, suffix string, newline bool) String {
	return String{key: key, value: value, prefix: prefix, suffix: suffix, newline: newline}
}

func (s String) writeTo(w Writer) error {
	if _, err := w.WriteString(s.key); err != nil {
		return err
	}
	if _, err := w.WriteString(s.prefix); err != nil {
		return err
	}
	if _, err := w.WriteString(s.value); err != nil {
		return err
	}
	if _, err := w.WriteString(s.suffix); err != nil {
		return err
	}
	if s.newline {
		w.WriteString("\n")
	}
	return nil
}
