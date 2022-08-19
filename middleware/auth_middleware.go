package middleware

import (
	"go-salaries-app/helper"
	"go-salaries-app/model/web"
	"net/http"
)

type AuthMiddleWare struct {
	Handler http.Handler
}

func (middleware *AuthMiddleWare) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "rahasia" == request.Header.Get("x-api-key") {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		resp := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}
		helper.WriteToResponseBody(writer, &resp)
	}
}

func NewAuthMiddleWare(handler http.Handler) *AuthMiddleWare {
	return &AuthMiddleWare{Handler: handler}
}
