package collection

import "io"

type List struct {
}

func (r *List) Path() (string, error) {
	return "/collections", nil
}

func (l *List) Encode() (io.Reader, error) {
	return nil, nil
}

func (l *List) ContentType() string {
	return ""
}
