package authorizationserver

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewHttpHandler () http.Handler{
	mux := mux.NewRouter()
	mux.HandleFunc("/", healthCheck).Methods("GET") //
	mux.HandleFunc("/oauth2/ssocheck", ssocheckHandler).Methods("GET") //SSO 상태인지 아닌지를 확인
	mux.HandleFunc("/login", loginHandler).Methods("POST") // 로그인 요청을 처리
	mux.HandleFunc("/signup", signupHandler).Methods("POST") //회원등록
	return mux
}

