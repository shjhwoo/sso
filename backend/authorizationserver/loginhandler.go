package authorizationserver

import (
	// "log"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	ldapserver "sso/ldap"

	"golang.org/x/crypto/bcrypt"
)

type loginForm struct {
	OrgId              string `json:"hospital_code"`
	UserId             string `json:"userid"`
	Password           string `json:"password"`
}

//로그인 버튼 눌렀을 때 실행됨.
func loginHandler (rw http.ResponseWriter, req *http.Request) {
	// 기관번호, 아이디, 비밀번호가 없을 경우 에러 메세지를 던진다. 
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	defer req.Body.Close()

	var userinfo loginForm
	json.Unmarshal(data, &userinfo)

	ldapconn, err := ldapserver.Connect()

	if err != nil {
		log.Fatal(err)
		return
	}
	//병원코드
	orgcode, err := ldapserver.Search(ldapconn,"ou=hospitals,dc=int,dc=trustnhope,dc=com","(&(objectClass=organizationalUnit)(ou="+ userinfo.OrgId +"))")
	if err != nil {
		fmt.Fprint(rw, "존재하지 않는 기관번호입니다")
		return
	}
	fmt.Println(orgcode,"기관번호 확인")
	//직원아이디
	user, err := ldapserver.Search(ldapconn, "ou="+ userinfo.OrgId +",ou=hospitals,dc=int,dc=trustnhope,dc=com","(&(objectClass=inetOrgPerson)(uid="+ userinfo.UserId +"))" )

	if err != nil {
		fmt.Fprint(rw, "존재하지 않는 아이디입니다")
		return
	}else{
		ldappw := user.Entries[0].GetAttributeValue("userPassword")
		err := bcrypt.CompareHashAndPassword([]byte(ldappw),[]byte(userinfo.Password))
		if err != nil {
			fmt.Fprint(rw, "비밀번호가 올바르지 않습니다")
			return
		}
		//로그인 성공. 
		fmt.Println("로그인 성공했어용~")

		// This context will be passed to all methods.
		ctx := req.Context()

		//서버에 등록된 클라이언트 정보와 입력값을 비교. ar는 초기에 인증서버 생성시 각 클라이언트에 대해 설정해준 값을 담고 있음. 
		ar, err := oauth2.NewAuthorizeRequest(ctx, req) 
		if err != nil {
			//이 코드 블록은 콜백페이지 로딩에 필요한 요소가 올바른지 확인하는 곳이다.
			// 요청 객체의 쿼리 파라미터에는 redirect_uri, client_id가 들어가 있을 것이다.
			// 이 두 요소가 올바르게 제공되지 않았을 경우에는 에러 메세지를 제공한다. 
			log.Printf("Error occurred in NewAuthorizeRequest: %+v", err)
			oauth2.WriteAuthorizeError(rw, ar, err)
			return
		}

		ssoSessionData := newSession(userinfo.UserId)

		response,err := oauth2.NewAuthorizeResponse(ctx, ar, ssoSessionData) 

		if err != nil {
			log.Printf("Error occurred in NewAuthorizeResponse: %+v", err)
			oauth2.WriteAuthorizeError(rw, ar, err)
			return
		}

		oauth2.WriteAuthorizeResponse(rw, ar, response)
	}
}