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

	onDisk := true
	resp := &response.CollectionUpdate{}
	err := client.CollectionUpdate(
		context.Background(),
		&request.CollectionUpdate{
			CollectionName: "test",
			Vectors: &request.UpdateVectorsParams{
				OnDisk: &onDisk,
			},
		},
		resp,
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("resp: %#v\n", resp)

}
