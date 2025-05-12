package exception

import (
	"github.com/go-playground/validator/v10"
	"github.com/helmigandi/go-restapi-1/helper"
	"github.com/helmigandi/go-restapi-1/model/web"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if notFoundErrorHandler(w, r, err) {
		return
	}

	if validationErrorHandler(w, r, err) {
		return
	}

	internalServerErrorHandler(w, r, err)
}

func notFoundErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(*NotFoundError)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		response := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(w, response)
		return true
	}

	return false
}

func validationErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(w, response)
		return true
	}

	return false
}

func internalServerErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	response := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	helper.WriteToResponseBody(w, response)
}
