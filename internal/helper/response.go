package helper

import (
	"hermes/internal/model/response"
	"hermes/internal/model/response/responsecode"
	"net/http"
)

type HandlerFuncWithHelper func(w http.ResponseWriter, r *http.Request) response.APIResponse

func WrapController(handler HandlerFuncWithHelper) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiResponse := handler(w, r)

		if apiResponse.Code != "" {
			apiResponse.CodeName = GetResponseCodeName(apiResponse.Code)
		}

		if apiResponse.Code != responsecode.CodeSuccess {
			WriteJSONResponse(w, http.StatusOK, apiResponse)
			return
		}

		WriteJSONResponse(w, http.StatusOK, apiResponse)
	})
}
