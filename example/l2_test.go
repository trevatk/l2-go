package example

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/trevatk/l2-go"
	"github.com/trevatk/l2-go/credential"
	"github.com/trevatk/l2-go/model"
)

// TestMakeBlock
func TestMakeBlock(t *testing.T) {

	ctx := context.TODO()

	key := "8d1414ec-fa88-11ec-afff-637eb61dd5bc"
	secret := "8d1415fa-fa88-11ec-b000-af4e0525c5e6"

	creds := credential.NewFromConfig(&credential.Config{
		Key:    key,
		Secret: secret,
	})

	l2 := l2.NewL2(creds)

	data := []byte("hello world")
	m := &model.Mock{
		Stake:      0.1,
		Difficulty: 1,
		Data:       data,
	}

	block, err := l2.MakeBlock(ctx, m)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("block hash: ", block.Hash)
}
