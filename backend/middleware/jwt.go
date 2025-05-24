package middleware

import (
    "net/http"
    "os"
    "strings"

    "github.com/golang-jwt/jwt/v5"
)

func JwtVerify(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tokenString := r.Header.Get("Authorization")
        if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
            http.Error(w, "Missing or invalid token", http.StatusUnauthorized)
            return
        }

        tokenString = strings.TrimPrefix(tokenString, "Bearer ")

        _, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, http.ErrAbortHandler
            }
            return []byte(os.Getenv("JWT_SECRET")), nil
        })

        if err != nil {
            http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
            return
        }

        next.ServeHTTP(w, r)
    })
}
