package main

import (
	"context"
	"fmt"

	qdrantgo "github.com/henomis/qdrant-go"
	collreq "github.com/henomis/qdrant-go/request/collection"
	collresp "github.com/henomis/qdrant-go/response/collection"
)

func main() {

	client := qdrantgo.New("http://localhost:6333", "lingoose")

	resp := &collresp.List{}
	client.CollectionList(context.Background(), &collreq.List{}, resp)

	fmt.Printf("resp: %#v\n", resp)

}
