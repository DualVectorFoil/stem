package pb

import "google.golang.org/grpc"

func NewAccountClientWrapper(cc *grpc.ClientConn) interface{} {
	return NewAccountServiceClient(cc)
}
