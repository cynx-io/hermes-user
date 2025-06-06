package responsecode

type ResponseCode string

const (

	// Expected Error
	CodeSuccess             ResponseCode = "00"
	CodeValidationError     ResponseCode = "VE"
	CodeAuthenticationError ResponseCode = "AU"
	CodeNotAllowed          ResponseCode = "NA"
	CodeNotFound            ResponseCode = "NF"
	CodeInvalidCredentials  ResponseCode = "IC"

	// Internal
	CodeJwtError      ResponseCode = "JWTERR"
	CodeInternalError ResponseCode = "I-IE"

	// DB Error
	CodeTblUserError ResponseCode = "TBLUSR"
)

var ResponseCodeNames = map[ResponseCode]string{
	CodeSuccess:             "Success",
	CodeValidationError:     "Validation Error",
	CodeAuthenticationError: "Authentication Error",
	CodeInternalError:       "Internal Error",
	CodeNotAllowed:          "Not Allowed",
	CodeNotFound:            "Not Found",
	CodeInvalidCredentials:  "Invalid Credentials",

	CodeTblUserError: "TblUser Error",
}
