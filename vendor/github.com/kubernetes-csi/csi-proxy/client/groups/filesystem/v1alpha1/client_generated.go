// Code generated by csi-proxy-api-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"net"

	"github.com/Microsoft/go-winio"
	"github.com/kubernetes-csi/csi-proxy/client"
	"github.com/kubernetes-csi/csi-proxy/client/api/filesystem/v1alpha1"
	"github.com/kubernetes-csi/csi-proxy/client/apiversion"
	"google.golang.org/grpc"
)

const groupName = "filesystem"

var version = apiversion.NewVersionOrPanic("v1alpha1")

type Client struct {
	client     v1alpha1.FilesystemClient
	connection *grpc.ClientConn
}

// NewClient returns a client to make calls to the filesystem API group version v1alpha1.
// It's the caller's responsibility to Close the client when done.
func NewClient() (*Client, error) {
	pipePath := client.PipePath(groupName, version)

	connection, err := grpc.Dial(pipePath,
		grpc.WithContextDialer(func(context context.Context, s string) (net.Conn, error) {
			return winio.DialPipeContext(context, s)
		}),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := v1alpha1.NewFilesystemClient(connection)
	return &Client{
		client:     client,
		connection: connection,
	}, nil
}

// Close closes the client. It must be called before the client gets GC-ed.
func (w *Client) Close() error {
	return w.connection.Close()
}

// ensures we implement all the required methods
var _ v1alpha1.FilesystemClient = &Client{}

func (w *Client) IsMountPoint(context context.Context, request *v1alpha1.IsMountPointRequest, opts ...grpc.CallOption) (*v1alpha1.IsMountPointResponse, error) {
	return w.client.IsMountPoint(context, request, opts...)
}

func (w *Client) LinkPath(context context.Context, request *v1alpha1.LinkPathRequest, opts ...grpc.CallOption) (*v1alpha1.LinkPathResponse, error) {
	return w.client.LinkPath(context, request, opts...)
}

func (w *Client) Mkdir(context context.Context, request *v1alpha1.MkdirRequest, opts ...grpc.CallOption) (*v1alpha1.MkdirResponse, error) {
	return w.client.Mkdir(context, request, opts...)
}

func (w *Client) PathExists(context context.Context, request *v1alpha1.PathExistsRequest, opts ...grpc.CallOption) (*v1alpha1.PathExistsResponse, error) {
	return w.client.PathExists(context, request, opts...)
}

func (w *Client) Rmdir(context context.Context, request *v1alpha1.RmdirRequest, opts ...grpc.CallOption) (*v1alpha1.RmdirResponse, error) {
	return w.client.Rmdir(context, request, opts...)
}
