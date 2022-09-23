package authorizationserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	ldapserver "sso/ldap"

	"github.com/go-ldap/ldap/v3"
	"golang.org/x/crypto/bcrypt"
)

type signupForm struct {
	OrgId              string `json:"hospital_code"`
	UserId             string `json:"userid"`
	Password           string `json:"password"`
	Sn                 string `json:"surname"`
	Cn             	   string `json:"commonname"`
	EmployeeNumber 	   string `json:"employee_number"`
	Mobile             string `json:"mobile"`
	//그 외 소속정보는 추가 예정!
}

func signupHandler(rw http.ResponseWriter, req *http.Request) {
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	defer req.Body.Close()

	var userinfo signupForm
	json.Unmarshal(data, &userinfo)

	ldapconn, err := ldapserver.Connect()

	if err != nil {
		log.Fatal(err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userinfo.Password), 10)

	if err != nil {
		fmt.Fprint(rw, "회원등록 중 문제가 발생했습니다. 관리자에게 연락바랍니다")
	}

	dn := "uid="+userinfo.UserId+",ou="+userinfo.OrgId+",ou=hospitals,dc=int,dc=trustnhope,dc=com"
	fmt.Println(dn,"*")

	newuser := ldap.NewAddRequest(dn,nil)

	newuser.Attribute("sn", []string{userinfo.Sn})
	newuser.Attribute("cn", []string{userinfo.Cn})
	newuser.Attribute("objectClass", []string{"top", "inetOrgPerson", "organizationalPerson", "person"})
	newuser.Attribute("userPassword", []string{string(hashedPassword)})
	newuser.Attribute("uid", []string{userinfo.UserId})
	newuser.Attribute("mobile", []string{userinfo.Mobile})
	newuser.Attribute("employeeNumber", []string{userinfo.EmployeeNumber})

	_, adderr := ldapserver.Add(ldapconn, newuser)

	if adderr != nil {
		fmt.Fprint(rw, "신규회원 등록에 실패했습니다")
		return
	}
	//신규회원 가입 성공
	fmt.Println("가입성공")
}