package cue

import (
	"bufio"
	"strings"
)

type element interface {
	setTitle(title string)
	setPerformer(performer string)
	setSongWriter(writer string)
	setISRC(isrc string)
	setCatalog(catalog string)
	setCDTextFile(file string)
	setFile(name, fType string)
	setComment(key, value string)
	setIndex(index, beginTime string)
}

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

func (s String) WriteTo(w *bufio.Writer) error {
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

type Sheet struct {
	Header  *Header
	Tracks  []*Track
	current element
}

func NewSheet() *Sheet {
	var s = &Sheet{}
	s.Header = &Header{}
	s.current = s.Header
	return s
}

func (s *Sheet) WriteTo(w *bufio.Writer) error {
	if err := s.Header.WriteTo(w); err != nil {
		return err
	}
	for _, t := range s.Tracks {
		if err := t.WriteTo(w); err != nil {
			return err
		}
	}
	w.Flush()
	return nil
}

func (s *Sheet) setTitle(title string) {
	s.current.setTitle(title)
}

func (s *Sheet) setPerformer(performer string) {
	s.current.setPerformer(performer)
}

func (s *Sheet) setSongWriter(writer string) {
	s.current.setSongWriter(writer)
}

func (s *Sheet) setISRC(isrc string) {
	s.current.setISRC(isrc)
}

func (s *Sheet) setCatalog(catalog string) {
	s.current.setCatalog(catalog)
}

func (s *Sheet) setCDTextFile(file string) {
	s.current.setCDTextFile(file)
}

func (s *Sheet) setFile(name, fType string) {
	s.current.setFile(name, fType)
}

func (s *Sheet) setComment(comment string) {
	if comment[0] == '"' {
		s.current.setComment("", comment)
	} else {
		var kIndex = strings.Index(comment, " ")
		if kIndex > 0 {
			var key = comment[:kIndex]
			var value = comment[kIndex+1:]
			s.current.setComment(key, value)
		} else {
			var key = strings.ToUpper(comment)
			switch key {
			case "GENRE":
				s.current.setComment(key, "")
			case "DISCID":
				s.current.setComment(key, "")
			case "DATE":
				s.current.setComment(key, "")
			case "COMMENT":
				s.current.setComment(key, "")
			default:
				s.current.setComment("", comment)
			}
		}
	}
}

func (s *Sheet) setIndex(index, beginTime string) {
	s.current.setIndex(index, beginTime)
}

func (s *Sheet) AddTrack(id, tType string) *Track {
	var t = NewTrack(id, tType)
	s.Tracks = append(s.Tracks, t)
	s.current = t
	return t
}

type Header struct {
	title      string
	performer  string
	songWriter string
	catalog    string
	cdTextFile string
	comments   []Comment
	file       File
}

func (h *Header) WriteTo(w *bufio.Writer) error {
	if err := NewString("TITLE ", h.title, "\"", "\"", true).WriteTo(w); err != nil {
		return err
	}
	if err := NewString("PERFORMER ", h.performer, "\"", "\"", true).WriteTo(w); err != nil {
		return err
	}
	if len(h.songWriter) > 0 {
		if err := NewString("SONGWRITER ", h.songWriter, "\"", "\"", true).WriteTo(w); err != nil {
			return err
		}
	}
	if len(h.catalog) > 0 {
		if err := NewString("CATALOG ", h.catalog, "\"", "\"", true).WriteTo(w); err != nil {
			return err
		}
	}
	if len(h.cdTextFile) > 0 {
		if err := NewString("CDTEXTFILE ", h.cdTextFile, "\"", "\"", true).WriteTo(w); err != nil {
			return err
		}
	}
	for _, comment := range h.comments {
		if len(comment.key) > 0 {
			if err := NewString("REM ", comment.key, "", "", len(comment.value) == 0).WriteTo(w); err != nil {
				return err
			}
			if len(comment.value) > 0 {

				if err := NewString(" ", comment.value, "", "", true).WriteTo(w); err != nil {
					return err
				}
			}
		} else {
			if err := NewString("REM ", comment.value, "", "", true).WriteTo(w); err != nil {
				return err
			}
		}
	}
	if err := h.file.WriteTo(w); err != nil {
		return err
	}
	return nil
}

func (h *Header) setTitle(title string) {
	h.title = title
}

func (h *Header) SetTitle(title string) {
	h.setTitle(title)
}

func (h *Header) setPerformer(performer string) {
	h.performer = performer
}

func (h *Header) SetPerformer(performer string) {
	h.setPerformer(performer)
}

func (h *Header) setSongWriter(writer string) {
	h.songWriter = writer
}

func (h *Header) SetSongWriter(writer string) {
	h.setSongWriter(writer)
}

func (h *Header) setISRC(isrc string) {
	panic("Header not implemented method setISRC")
}

func (h *Header) setCatalog(catalog string) {
	h.catalog = catalog
}

func (h *Header) SetCatalog(catalog string) {
	h.setCatalog(catalog)
}

func (h *Header) setCDTextFile(file string) {
	h.cdTextFile = file
}

func (h *Header) SetCDTextFile(file string) {
	h.setCDTextFile(file)
}

func (h *Header) setFile(name, fType string) {
	h.file = File{name: name, fileType: fType}
}

func (h *Header) SetFile(name, fType string) {
	h.setFile(name, fType)
}

func (h *Header) setComment(key, value string) {
	h.comments = append(h.comments, Comment{key: key, value: value})
}

func (h *Header) SetComment(key, value string) {
	h.setComment(key, value)
}

func (h *Header) setIndex(index, beginTime string) {
	panic("Header not implemented method setIndex")
}

type Comment struct {
	key   string
	value string
}

type File struct {
	name     string
	fileType string
}

func (f *File) WriteTo(w *bufio.Writer) error {
	if err := NewString("FILE ", f.name, "\"", "\"", false).WriteTo(w); err != nil {
		return err
	}
	if err := NewString(" ", f.fileType, "", "", true).WriteTo(w); err != nil {
		return err
	}
	return nil
}

type Track struct {
	id         string
	trackType  string
	title      string
	performer  string
	songWriter string
	catalog    string
	isrc       string
	comments   []Comment
	indexes    []Index
}

func NewTrack(id, tType string) *Track {
	var t = &Track{}
	t.id = id
	t.trackType = tType
	return t
}

func (t *Track) WriteTo(w *bufio.Writer) error {
	if err := NewString("  TRACK ", t.id, "", "", false).WriteTo(w); err != nil {
		return err
	}
	if err := NewString(" ", t.trackType, "", "", true).WriteTo(w); err != nil {
		return err
	}
	if err := NewString("    TITLE ", t.title, "\"", "\"", true).WriteTo(w); err != nil {
		return err
	}
	if err := NewString("    PERFORMER ", t.performer, "\"", "\"", true).WriteTo(w); err != nil {
		return err
	}
	if len(t.songWriter) > 0 {
		if err := NewString("    SONGWRITER ", t.songWriter, "\"", "\"", true).WriteTo(w); err != nil {
			return err
		}
	}
	if len(t.catalog) > 0 {
		if err := NewString("    CATALOG ", t.catalog, "\"", "\"", true).WriteTo(w); err != nil {
			return err
		}
	}
	if len(t.isrc) > 0 {
		if err := NewString("    ISRC ", t.isrc, "\"", "\"", true).WriteTo(w); err != nil {
			return err
		}
	}
	for _, comment := range t.comments {
		if len(comment.key) > 0 && len(comment.value) > 0 {
			if err := NewString("    REM ", comment.key, "", "", len(comment.value) == 0).WriteTo(w); err != nil {
				return err
			}
			if len(comment.value) > 0 {
				if err := NewString(" ", comment.value, "", "", true).WriteTo(w); err != nil {
					return err
				}
			}
		} else {
			if err := NewString("    REM ", comment.value, "", "", true).WriteTo(w); err != nil {
				return err
			}
		}
	}
	for _, index := range t.indexes {
		if err := NewString("    INDEX ", index.index, "", "", false).WriteTo(w); err != nil {
			return err
		}
		if err := NewString(" ", index.beginTime, "", "", true).WriteTo(w); err != nil {
			return err
		}
	}
	return nil
}

func (t *Track) setTitle(title string) {
	t.title = title
}

func (t *Track) SetTitle(title string) {
	t.setTitle(title)
}

func (t *Track) setPerformer(performer string) {
	t.performer = performer
}

func (t *Track) SetPerformer(performer string) {
	t.setPerformer(performer)
}

func (t *Track) setSongWriter(writer string) {
	t.songWriter = writer
}

func (t *Track) SetSongWriter(writer string) {
	t.setSongWriter(writer)
}

func (t *Track) setISRC(isrc string) {
	t.isrc = isrc
}

func (t *Track) SetISRC(isrc string) {
	t.setISRC(isrc)
}

func (t *Track) setCatalog(catalog string) {
	t.catalog = catalog
}

func (t *Track) SetCatalog(catalog string) {
	t.setCatalog(catalog)
}

func (t *Track) setCDTextFile(file string) {
	panic("Track not implemented method setCDTextFile")
}

func (t *Track) setFile(name, fType string) {
	panic("Track not implemented method setFile")
}

func (t *Track) setComment(key, value string) {
	t.comments = append(t.comments, Comment{key: key, value: value})
}

func (t *Track) SetComment(key, value string) {
	t.setComment(key, value)
}

func (t *Track) setIndex(index, beginTime string) {
	t.indexes = append(t.indexes, Index{index: index, beginTime: beginTime})
}

func (t *Track) SetIndex(index, beginTime string) {
	t.setIndex(index, beginTime)
}

type Index struct {
	index     string
	beginTime string
}
