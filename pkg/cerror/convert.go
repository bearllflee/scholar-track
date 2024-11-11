package cerror

import "google.golang.org/grpc/status"

func FromError(err error) bool {
	_, ok := status.FromError(err)

	return ok
}
