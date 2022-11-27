package node

import (
	context "context"

	gogogrpc "github.com/cosmos/gogoproto/grpc"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

// RegisterGRPCGatewayRoutes mounts the node gRPC service's GRPC-gateway routes
// on the given mux object.
func RegisterGRPCGatewayRoutes(clientConn gogogrpc.ClientConn, mux *runtime.ServeMux) {
	_ = RegisterServiceHandlerClient(context.Background(), mux, NewServiceClient(clientConn))
}
