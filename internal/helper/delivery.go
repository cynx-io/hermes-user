package helper

import (
	"context"
	"github.com/gorilla/mux"
	"hermes/internal/model/response"
	"hermes/internal/model/response/responsecode"
	"net/http"
)

type HandlerFunc[T any] func(ctx context.Context, req T, res *response.APIResponse)
type HandlerFuncNoReq func(ctx context.Context, res *response.APIResponse)

func InjectGenericRoute[T any](
	r *mux.Router,
	path string,
	handler HandlerFunc[T],
) {
	r.HandleFunc(path, WrapController(createGenericController(handler))).Methods(http.MethodPost)
}

func InjectNoReqRoute(
	r *mux.Router,
	path string,
	handler HandlerFuncNoReq,
) {
	r.HandleFunc(path, WrapController(createNoReqController(handler))).Methods(http.MethodPost)
}

func createGenericController[T any](
	handler HandlerFunc[T],
) func(w http.ResponseWriter, r *http.Request) response.APIResponse {
	return func(w http.ResponseWriter, r *http.Request) response.APIResponse {
		return HandleRequest[T](w, r, handler)
	}
}

func createNoReqController(
	handler HandlerFuncNoReq,
) func(w http.ResponseWriter, r *http.Request) response.APIResponse {
	return func(w http.ResponseWriter, r *http.Request) response.APIResponse {
		return HandleEmptyRequest(w, r, handler)
	}
}

// Generic request handler executor
func HandleRequest[T any](
	w http.ResponseWriter,
	r *http.Request,
	handler HandlerFunc[T],
) response.APIResponse {
	var req T
	var apiResponse response.APIResponse

	ctx := r.Context()

	// Decode and validate
	if err := DecodeAndValidateRequest(r, &req); err != nil {
		apiResponse.Code = responsecode.CodeValidationError
		apiResponse.Error = err.Error()
		return apiResponse
	}

	// Call the actual handler
	handler(ctx, req, &apiResponse)

	return apiResponse
}

// Generic request handler executor
func HandleEmptyRequest(
	w http.ResponseWriter,
	r *http.Request,
	handler HandlerFuncNoReq,
) response.APIResponse {
	var apiResponse response.APIResponse

	ctx := r.Context()

	// Call the actual handler
	handler(ctx, &apiResponse)

	return apiResponse
}
