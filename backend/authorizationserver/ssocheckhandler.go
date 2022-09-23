package authorizationserver

import (
	"fmt"
	"log"
	"net/http"
)

func ssocheckHandler (rw http.ResponseWriter, req *http.Request){
	// This context will be passed to all methods.
	ctx := req.Context()

	// Let's create an AuthorizeRequest object!
	// It will analyze the request and extract important information like scopes, response type and others.
	ar, err := oauth2.NewAuthorizeRequest(ctx, req) //서버에 등록된 클라이언트 정보와 입력값을 비교
	if err != nil {
		//이 코드 블록은 콜백페이지 로딩에 필요한 요소가 올바른지 확인하는 곳이다.
		// 요청 객체의 쿼리 파라미터에는 redirect_uri, client_id가 들어가 있을 것이다.
		// 이 두 요소가 올바르게 제공되지 않았을 경우에는 에러 메세지를 제공한다. 
		log.Printf("Error occurred in NewAuthorizeRequest: %+v", err)
		oauth2.WriteAuthorizeError(rw, ar, err)
		return
	}
	//ar는 초기에 인증서버 생성시 각 클라이언트에 대해 설정해준 값을 담고 있음. 
	//이런식으로 ... 일단 지금은 안 쓸거임. 
	//기존 요청에 세션 쿠키가 있는지 없는지만 확인하고 없을 경우 로그인 화면을 보여줄 거임. 

	ssoCookie, err := req.Cookie("sso")

	if err != nil {
		//세션쿠키가 없다. 로그인 페이지로 사용자를 리디렉션시킨다.
		http.Redirect(rw, req,"http://localhost:3000",301)
		return 
	}

	//이미 로그인을 한 전적이 있다 == 세션 쿠키가 존재한다
	//ID 토큰과 액세스(?) 토큰 발급
	//사용자가 쓰려던 서비스 화면으로 그대로 리디렉션
	fmt.Println(ssoCookie,"ssoCookie")
	fmt.Println(ar.GetRedirectURI(),"서비스 url")
	http.Redirect(rw, req, ar.GetRedirectURI().Scheme + "://" + ar.GetRedirectURI().Host + ar.GetRedirectURI().Path, 301)
}