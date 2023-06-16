package request

import "io"

type CollectionList struct {
}

func (c *CollectionList) Path() (string, error) {
	return "/collections", nil
}

func (c *CollectionList) Encode() (io.Reader, error) {
	return nil, nil
}

func (c *CollectionList) ContentType() string {
	return ""
}
