package l2

import (
	"context"
	"fmt"
	"net"

	"github.com/trevatk/l2-go/credential"
	pb "github.com/trevatk/l2-go/proto/l2_v1"
	"google.golang.org/grpc"
)

const (
	_hostname = ""
	_port     = ""
)

// L2
type L2 struct {
	credentials credential.ICredentials
}

// NewL2
func NewL2(creds credential.ICredentials) *L2 {
	return &L2{credentials: creds}
}

func (l2 *L2) getClient(ctx context.Context) (*grpc.ClientConn, pb.L2V1Client, error) {

	opt := []grpc.DialOption{}
	conn, err := grpc.Dial(net.JoinHostPort(_hostname, _port), opt...)
	if err != nil {
		return nil, nil, fmt.Errorf("grpc.Dial: %v", err)
	}

	return conn, pb.NewL2V1Client(conn), nil
}
