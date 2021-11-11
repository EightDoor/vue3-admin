package configure

type Code int

const (
	ApiStartError         Code = 10000
	ApiInnerResponseError Code = 100001
)

const (
	ApiGenSignatureError Code = 5000
)

const (
	RequestError               Code = -1
	RequestSuccess             Code = 0
	RequestOtherError          Code = 4000
	RequestKeyNotFound         Code = 4001
	RequestParameterTypeError  Code = 4002
	RequestAuthorizedFailed    Code = 4003
	RequestNotFound            Code = 4004
	RequestParameterMiss       Code = 4005
	RequestMethodNotAllowed    Code = 4006
	RequestExpired             Code = 4007
	RequestParameterRangeError Code = 4008
	RequestAccessDeny          Code = 4009
	RequestUserPassValidationError Code = 4010
)
