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

	wait := true
	resp := &response.IndexCreate{}
	err := client.IndexCreate(
		context.Background(),
		&request.IndexCreate{
			CollectionName: "test",
			Wait:           &wait,
			FieldName:      "test_field",
			FieldSchema:    "keyword",
		},
		resp,
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("resp: %#v\n", resp)

}
