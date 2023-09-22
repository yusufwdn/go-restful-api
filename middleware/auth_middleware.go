package middleware

import (
	"net/http"
	"yusufwdn/golang-restful-api/helper"
	"yusufwdn/golang-restful-api/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	apiKey := request.Header.Get("X-API-Key")
	if apiKey == "RAHASIA" {
		// ok
		// jika tidak ada error, lanjutkan ke handler selanjutnya
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		// error
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
	}

}
