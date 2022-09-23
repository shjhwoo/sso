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
		//authorizationcode를 발급한다. (어디에 저장하지??: 디비에 저장해둔다. 그리고 일정 시간이 지나면 삭제해버렷)
	}
}