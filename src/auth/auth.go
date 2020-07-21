package auth

import (
	"net/http"
)

type Token interface {
	IsValid() bool
}

func valid(t Token) bool {
	return t.IsValid()
}

func IsAuth(r *http.Request) bool {
	t := MockToken{token: r.Header.Get("Authorization")}
	return valid(&t)
}