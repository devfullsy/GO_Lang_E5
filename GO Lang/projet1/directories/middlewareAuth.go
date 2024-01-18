package dictionary

import (
	"net/http"
)

const authToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjMiLCJuYW1lIjoidG9rZW4ifQ.65HrM5nU_61p3vUQGEHr5cP4jjC0KH_u_Hdv0i4tOq0"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != authToken {
			http.Error(w, "Non autorisé. Jeton d'authentification invalide.", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
