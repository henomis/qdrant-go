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

	resp := &response.CollectionList{}
	err := client.CollectionList(context.Background(), &request.CollectionList{}, resp)
	if err != nil {
		panic(err)
	}

	fmt.Printf("resp: %#v\n", resp)

}
