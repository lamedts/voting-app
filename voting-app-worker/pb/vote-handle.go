package pb

import (
	"github.com/golang/glog"
	context "golang.org/x/net/context"
)

type EchoServer struct{}

func (s *EchoServer) Echo(ctx context.Context, msg *EchoMessage) (*EchoMessage, error) {
	glog.Info(msg)
	return &EchoMessage{Greeting: "bar"}, nil
}
