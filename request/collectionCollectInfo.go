package request

import (
	"fmt"
	"io"
)

type CollectionCollectInfo struct {
	CollectionName string `json:"-"`
}

func (c *CollectionCollectInfo) Path() (string, error) {
	return fmt.Sprintf("/collections/%s", c.CollectionName), nil
}

func (c *CollectionCollectInfo) Encode() (io.Reader, error) {
	return nil, nil
}

func (c *CollectionCollectInfo) ContentType() string {
	return ""
}
