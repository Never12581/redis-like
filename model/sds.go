package model

const miB int64 = 1024 * 1024

type Sds struct {
	buf []byte
}

func NewSds(bs []byte) *Sds {
	return &Sds{buf: bs}
}

func (s *Sds) Size() int64 {
	return int64(len(s.buf))
}

func (s *Sds) Append(appendBs []byte) {
	s.buf = append(s.buf, appendBs...)
}

func (s *Sds) Value() []byte {
	return s.buf
}
