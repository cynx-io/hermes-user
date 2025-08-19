package response

import "github.com/cynx-io/cynx-core/src/response"

const (
	// Expected Error
	codeSuccess            response.Code = "00"
	codeValidationError    response.Code = "VE"
	codeUnauthorized       response.Code = "UA"
	codeNotAllowed         response.Code = "NA"
	codeNotFound           response.Code = "NF"
	codeInvalidCredentials response.Code = "IC"

	// Internal
	codeInternalError response.Code = "I-IE"

	// External Errors
	// Database Errors
	codeDbUserError      response.Code = "DB-USR"
	codeDbUserGuestError response.Code = "DB-USG"
)

var responseCodeNames = map[response.Code]string{
	// Expected Error
	codeSuccess:            "Success",
	codeValidationError:    "Validation Error",
	codeUnauthorized:       "Not Authorized",
	codeNotAllowed:         "Not Allowed",
	codeNotFound:           "Not Found",
	codeInvalidCredentials: "Invalid Credentials",

	// Internal
	codeInternalError: "Internal Error",

	// Database Errors
	codeDbUserError:      "Database User Error",
	codeDbUserGuestError: "Database User Guest Error",
}
