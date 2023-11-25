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
	resp := &response.VectorsDelete{}
	err := client.VectorsDelete(
		context.Background(),
		&request.VectorsDelete{
			CollectionName: "test",
			Wait:           &wait,
			Vector:         []string{"vec1"},
		},
		resp,
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("resp: %#v\n", resp)

}
