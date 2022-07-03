package l2

import (
	"context"
	"fmt"
	"time"

	"github.com/trevatk/l2-go/model"
	pb "github.com/trevatk/l2-go/proto/l2_v1"
	"google.golang.org/grpc"
)

// MakeBlock
func (l2 *L2) MakeBlock(ctx context.Context, input *model.Mock) (*model.Block, error) {

	timeout, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	conn, client, err := l2.getClient(timeout)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	msg := &pb.Mock{
		ApplicationKey:  l2.credentials.Key(),
		OrganizationKey: l2.credentials.Secret(),
		Stake:           input.Stake,
		Difficulty:      input.Difficulty,
		Contents:        input.Data,
	}

	fmt.Println("msg: ", msg)

	opt := []grpc.CallOption{}
	b, err := client.MakeBlock(timeout, msg, opt...)
	if err != nil {
		return nil, fmt.Errorf("client.MakeBlock: %v", err)
	}

	return &model.Block{Hash: b.Hash}, nil
}
