package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"snippetdemo/internal/snippetdemo/helpers"
	"strings"
)

func (mgr MiddleWareManager) Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]

			contextInfo, err := helpers.ValidateToken(mgr.SecretKey, jwtToken)

			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			} else {
				b := context.WithValue(r.Context(), "userid", contextInfo.UserID)
				b = context.WithValue(b, "username", contextInfo.Username)
				r = r.WithContext(b)
				next(w, r)
			}
		}
	})
}
