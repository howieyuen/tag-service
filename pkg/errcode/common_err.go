package errcode

var (
	Succeed          = NewError(0, "Succeed")
	InternalError    = NewError(10000000, "InternalError")
	InvalidParams    = NewError(10000001, "InvalidParam")
	Unauthorized     = NewError(10000002, "Unauthorized")
	NotFound         = NewError(10000003, "NotFound")
	Unknown          = NewError(10000004, "Unknown")
	DeadlineExceeded = NewError(10000005, "DeadlineExceeded")
	AccessDenied     = NewError(10000006, "AccessDenied")
	LimitExceed      = NewError(10000007, "LimitExceed")
	MethodNotAllowed = NewError(10000008, "MethodNotAllowed")
)
