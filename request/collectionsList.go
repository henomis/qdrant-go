package request

import "io"

type CollectionList struct {
}

func (r *CollectionList) Path() (string, error) {
	return "/collections", nil
}

func (l *CollectionList) Encode() (io.Reader, error) {
	return nil, nil
}

func (l *CollectionList) ContentType() string {
	return ""
}
