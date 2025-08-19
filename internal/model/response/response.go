package response

import (
	"github.com/cynx-io/cynx-core/src/response"
)

func setResponse[Resp response.Generic](resp Resp, code response.Code) {
	resp.GetBase().Code = code.String()
	resp.GetBase().Desc = responseCodeNames[code]
}

func Success[Resp response.Generic](resp Resp) {
	setResponse(resp, codeSuccess)
}

func ErrorValidation[Resp response.Generic](resp Resp) {
	setResponse(resp, codeValidationError)
}

func ErrorUnauthorized[Resp response.Generic](resp Resp) {
	setResponse(resp, codeUnauthorized)
}

func ErrorNotAllowed[Resp response.Generic](resp Resp) {
	setResponse(resp, codeNotAllowed)
}

func ErrorNotFound[Resp response.Generic](resp Resp) {
	setResponse(resp, codeNotFound)
}

func ErrorInvalidCredentials[Resp response.Generic](resp Resp) {
	setResponse(resp, codeInvalidCredentials)
}

func ErrorInternal[Resp response.Generic](resp Resp) {
	setResponse(resp, codeInternalError)
}

func ErrorDbUser[Resp response.Generic](resp Resp) {
	setResponse(resp, codeDbUserError)
}

func ErrorDbUserGuest[Resp response.Generic](resp Resp) {
	setResponse(resp, codeDbUserGuestError)
}
