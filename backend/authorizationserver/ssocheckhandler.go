package authorizationserver

import (
	"net/http"
)

func ssocheckHandler (rw http.ResponseWriter, req *http.Request){
	//기존 요청에 세션 쿠키가 있는지 없는지만 확인하고 없을 경우 로그인 화면을 보여줄 거임. 

	// ssoCookie, err := req.Cookie("sso")

	// fmt.Println(ssoCookie,"ssoCookie")

	// if err != nil {
	// 	//세션쿠키가 없다. 로그인 페이지로 사용자를 리디렉션시킨다.
	// 	http.Redirect(rw, req,"http://localhost:3000",http.StatusSeeOther)
	// 	return 
	// }

	// //이미 로그인을 한 전적이 있다 == 세션 쿠키가 존재한다
	// //새로운 서비스로 접속하는 경우 액세스토큰 발급(?) ..
	// //그후 사용자가 쓰려던 서비스 화면으로 그대로 리디렉션
	// http.Redirect(rw, req, req.Host, http.StatusSeeOther)//********************************
}

	// cookie := &http.Cookie{
	// 	Name: "sso",
	// 	Value: "hello~",
	// 	MaxAge: 300,
	// }	
	
	// http.SetCookie(rw, cookie)