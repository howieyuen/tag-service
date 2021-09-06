package errcode

import (
	pb "github.com/howieyuen/tag-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Status struct {
	*status.Status
}

func FromError(err error) *Status {
	s, _ := status.FromError(err)
	return &Status{s}
}

func TogRPCError(err *Error) error {
	s, _ := status.New(ToRPCCode(err.Code()), err.Msg()).
		WithDetails(
			&pb.Error{
				Code: int32(err.Code()),
				Msg:  err.Msg(),
			})
	return s.Err()
}

func ToRPCStatus(code int, msg string) *Status {
	s, _ := status.New(ToRPCCode(code), msg).
		WithDetails(
			&pb.Error{
				Code: int32(code),
				Msg:  msg,
			})
	return &Status{s}
}

func ToRPCCode(code int) codes.Code {
	var statusCode codes.Code
	switch code {
	case InternalError.Code():
		statusCode = codes.Internal
	case InvalidParams.Code():
		statusCode = codes.InvalidArgument
	case Unauthorized.Code():
		statusCode = codes.Unauthenticated
	case AccessDenied.Code():
		statusCode = codes.PermissionDenied
	case DeadlineExceeded.Code():
		statusCode = codes.DeadlineExceeded
	case NotFound.Code():
		statusCode = codes.NotFound
	case LimitExceed.Code():
		statusCode = codes.ResourceExhausted
	case MethodNotAllowed.Code():
		statusCode = codes.Unimplemented
	default:
		statusCode = codes.Unknown
	}
	
	return statusCode
}
