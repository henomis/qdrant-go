package main

import (
	"context"
	"fmt"

	qdrantgo "github.com/henomis/qdrant-go"
	"github.com/henomis/qdrant-go/request"
	"github.com/henomis/qdrant-go/response"
)

func main() {

	client := qdrantgo.New("http://localhost:6333", "")

	resp := &response.PointGet{}
	err := client.PointGet(
		context.Background(),
		&request.PointGet{
			CollectionName: "test",
			ID:             "45b07125-f592-414f-a9d0-160c8ecc283a",
		},
		resp,
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("resp: %#v\n", resp)

}
