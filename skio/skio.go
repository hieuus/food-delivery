package skio

import (
	"github.com/hieuus/food-delivery/common"
	"net"
	"net/url"
)

type Conn interface {
	ID() string
	Close() error
	URL() url.URL
	LocalAddr() net.Addr
	RemoteAddr() net.Addr

	Context() interface{}
	SetContext(v interface{})
	Namespace() string
	Emit(msg string, v ...interface{})

	Join(room string)
	Leave(room string)
	LeaveAll()
	Rooms() []string
}

type AppSocket interface {
	Conn
	common.Requester
}

type appSocket struct {
	Conn
	common.Requester
}

func NewAppSocket(conn Conn, requester common.Requester) *appSocket {
	return &appSocket{conn, requester}
}
