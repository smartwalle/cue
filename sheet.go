package cue

type property interface {
	setTitle(title string)
	setPerformer(performer string)
	setSongWriter(writer string)
	setISRC(isrc string)
	setCatalog(catalog string)
	setCDTextFile(filename string)
	setFile(filename, fileType string)
	setComment(key, value string)
	setIndex(index, beginTime string)
}

type Sheet struct {
	Header  *Header
	Tracks  []*Track
	current property
}

func NewSheet() *Sheet {
	var s = &Sheet{}
	s.Header = &Header{}
	s.current = s.Header
	return s
}

func (s *Sheet) WriteTo(w Writer) error {
	if err := s.Header.writeTo(w); err != nil {
		return err
	}
	for _, t := range s.Tracks {
		if err := t.writeTo(w); err != nil {
			return err
		}
	}
	return w.Flush()
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

func (s *Sheet) setCDTextFile(filename string) {
	s.current.setCDTextFile(filename)
}

func (s *Sheet) setFile(filename, fileType string) {
	s.current.setFile(filename, fileType)
}

func (s *Sheet) setComment(key, value string) {
	s.current.setComment(key, value)
}

func (s *Sheet) setIndex(index, beginTime string) {
	s.current.setIndex(index, beginTime)
}

func (s *Sheet) AddTrack(id, trackType string) *Track {
	var t = NewTrack(id, trackType)
	s.Tracks = append(s.Tracks, t)
	s.current = t
	return t
}

type Header struct {
	Title      string
	Performer  string
	SongWriter string
	Catalog    string
	CDTextFile string
	Comments   []Comment
	File       File
}

func (h *Header) writeTo(w Writer) error {
	if err := NewString("TITLE \"", h.Title, "\"\n").writeTo(w); err != nil {
		return err
	}
	if err := NewString("PERFORMER \"", h.Performer, "\"\n").writeTo(w); err != nil {
		return err
	}
	if len(h.SongWriter) > 0 {
		if err := NewString("SONGWRITER \"", h.SongWriter, "\"\n").writeTo(w); err != nil {
			return err
		}
	}
	if len(h.Catalog) > 0 {
		if err := NewString("CATALOG \"", h.Catalog, "\"\n").writeTo(w); err != nil {
			return err
		}
	}
	if len(h.CDTextFile) > 0 {
		if err := NewString("CDTEXTFILE \"", h.CDTextFile, "\"\n").writeTo(w); err != nil {
			return err
		}
	}
	for _, comment := range h.Comments {
		if len(comment.Key) > 0 {
			if err := NewString("REM ", comment.Key).writeTo(w); err != nil {
				return err
			}
			if len(comment.Value) > 0 {
				if err := NewString(" ", comment.Value).writeTo(w); err != nil {
					return err
				}
			}
			w.WriteString("\n")
		} else {
			if err := NewString("REM ", comment.Value, "\n").writeTo(w); err != nil {
				return err
			}
		}
	}
	if err := NewString("FILE \"", h.File.Filename, "\" ", h.File.FileType, "\n").writeTo(w); err != nil {
		return err
	}
	return nil
}

func (h *Header) setTitle(title string) {
	h.Title = title
}

func (h *Header) setPerformer(performer string) {
	h.Performer = performer
}

func (h *Header) setSongWriter(writer string) {
	h.SongWriter = writer
}

func (h *Header) setISRC(isrc string) {
	panic("Header not implemented method setISRC")
}

func (h *Header) setCatalog(catalog string) {
	h.Catalog = catalog
}

func (h *Header) setCDTextFile(filename string) {
	h.CDTextFile = filename
}

func (h *Header) setFile(filename, fileType string) {
	h.File = File{Filename: filename, FileType: fileType}
}

func (h *Header) setComment(key, value string) {
	h.Comments = append(h.Comments, Comment{Key: key, Value: value})
}

func (h *Header) setIndex(index, beginTime string) {
	panic("Header not implemented method setIndex")
}

func (h *Header) AddComment(key, value string) {
	h.setComment(key, value)
}

type Comment struct {
	Key   string
	Value string
}

type File struct {
	Filename string
	FileType string
}

type Track struct {
	Id         string
	TrackType  string
	Title      string
	Performer  string
	SongWriter string
	Catalog    string
	ISRC       string
	Comments   []Comment
	Indexes    []Index
}

func NewTrack(id, trackType string) *Track {
	var t = &Track{}
	t.Id = id
	t.TrackType = trackType
	return t
}

func (t *Track) writeTo(w Writer) error {
	if err := NewString("  TRACK ", t.Id, " ", t.TrackType, "\n").writeTo(w); err != nil {
		return err
	}
	if err := NewString("    TITLE \"", t.Title, "\"\n").writeTo(w); err != nil {
		return err
	}
	if len(t.Performer) > 0 {
		if err := NewString("    PERFORMER \"", t.Performer, "\"\n").writeTo(w); err != nil {
			return err
		}
	}
	if len(t.SongWriter) > 0 {
		if err := NewString("    SONGWRITER \"", t.SongWriter, "\"\n").writeTo(w); err != nil {
			return err
		}
	}
	if len(t.Catalog) > 0 {
		if err := NewString("    CATALOG \"", t.Catalog, "\"\n").writeTo(w); err != nil {
			return err
		}
	}
	if len(t.ISRC) > 0 {
		if err := NewString("    ISRC \"", t.ISRC, "\"\n").writeTo(w); err != nil {
			return err
		}
	}
	for _, comment := range t.Comments {
		if len(comment.Key) > 0 {
			if err := NewString("REM ", comment.Key).writeTo(w); err != nil {
				return err
			}
			if len(comment.Value) > 0 {
				if err := NewString(" ", comment.Value).writeTo(w); err != nil {
					return err
				}
			}
			w.WriteString("\n")
		} else {
			if err := NewString("REM ", comment.Value, "\n").writeTo(w); err != nil {
				return err
			}
		}
	}
	for _, index := range t.Indexes {
		if err := NewString("    INDEX ", index.Index, " ", index.BeginTime, "\n").writeTo(w); err != nil {
			return err
		}
	}
	return nil
}

func (t *Track) setTitle(title string) {
	t.Title = title
}

func (t *Track) setPerformer(performer string) {
	t.Performer = performer
}

func (t *Track) setSongWriter(writer string) {
	t.SongWriter = writer
}

func (t *Track) setISRC(isrc string) {
	t.ISRC = isrc
}

func (t *Track) setCatalog(catalog string) {
	t.Catalog = catalog
}

func (t *Track) setCDTextFile(filename string) {
	panic("Track not implemented method setCDTextFile")
}

func (t *Track) setFile(filename, fileType string) {
	panic("Track not implemented method setFile")
}

func (t *Track) setComment(key, value string) {
	t.Comments = append(t.Comments, Comment{Key: key, Value: value})
}

func (t *Track) setIndex(index, beginTime string) {
	t.Indexes = append(t.Indexes, Index{Index: index, BeginTime: beginTime})
}

func (t *Track) AddComment(key, value string) {
	t.setComment(key, value)
}

func (t *Track) AddIndex(index, beginTime string) {
	t.setIndex(index, beginTime)
}

type Index struct {
	Index     string
	BeginTime string
}
