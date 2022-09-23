package authorizationserver

import (
	// "log"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sso/ldap"
)

type loginForm struct {
	OrgId              string `json:"orgid"`
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
	//기관번호, 아이디, 비밀번호 입력 받아서 LDAP 서버에 저장된 정보와 비교하고, 
	//올바른 정보를 제공하였을 시 authorzation code 또는  ID 토큰. 액세스 토큰, 리프레시 토큰을 제공함.
	//1. 등록안된 기관번호를 입력한 경우
	//2. 기관은 등록되어 있으나, 그 기관에 등록안된 아이디를 입력한 경우
	//3. 기관과 아이디 모두 맞게 입력했는데 비밀번호가 틀린 경우
	//4. 기관번호, 아이디, 비밀번호 모두 올바르게 입력한 경우: authorizationcode 발급

	//각 경우에 대해서 ldap 서버에 검색 요청을 보내야 한다. (LDAP은 디비다.)

	ldapconn := ldap.DialAndBind()


}