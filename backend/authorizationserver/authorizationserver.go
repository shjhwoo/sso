package authorizationserver

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"

	"github.com/ory/fosite"
	//"net/http"
	"time"

	"github.com/ory/fosite/compose"
	"github.com/ory/fosite/handler/openid"
	"github.com/ory/fosite/token/jwt"

	//"github.com/ory/fosite/handler/openid"
	"github.com/ory/fosite/storage"
	//"github.com/ory/fosite/token/jwt"
)

//oauth+oidc 인증 서버 설정하고 빌드하는 코드입니다.

var(
	config = &compose.Config{
		AccessTokenLifespan: time.Minute * 10, //10분
		RefreshTokenLifespan: time.Minute * 60 * 24 * 14, //2주
		AuthorizeCodeLifespan: time.Minute * 5, //5분
		IDTokenLifespan: time.Minute * 10, //10분
		IDTokenIssuer: "trustnhope",
		DisableRefreshTokenValidation: false, //리프레시 토큰 검증 비활성화 취소
		EnforcePKCE: true, //PKCE를 통한 authcode 인터셉트 공격방지
		EnforcePKCEForPublicClients: true,
		//RedirectSecureChecker:
		UseLegacyErrorFormat: true, //로그인 에러 보여주기
	}

	//여기가 문제임ㅠ
	//클라이언트 아이디와 시크릿, 사용자 계정 정보 저장:: 커스텀하게 어떻게 등록해ㅠㅠ??
	TNHSSOMemoryStore = &storage.MemoryStore{
		IDSessions: make(map[string]fosite.Requester),
		Clients: clientlist,
		// Users: map[string]storage.MemoryUserRelation{
		// 	"suhyun": {
		// 		// This store simply checks for equality, a real storage implementation would obviously use
		// 		// a hashing algorithm for encrypting the user password.
		// 		Username: "suhyun",
		// 		Password: "secret",
		// 	},
		// }, 사용자 정보는 LDAP 서버에서 들고올거임. 어짜피 auth code로 인증할거라 이 부분은 여기서 굳이 안써도 된다.
		AuthorizeCodes:         map[string]storage.StoreAuthorizeCode{},
		AccessTokens:           map[string]fosite.Requester{},
		RefreshTokens:          map[string]storage.StoreRefreshToken{},
		PKCES:                  map[string]fosite.Requester{},
		AccessTokenRequestIDs:  map[string]string{},
		RefreshTokenRequestIDs: map[string]string{},
		IssuerPublicKeys:       map[string]storage.IssuerPublicKeys{},
	}
	
	store = storage.NewMemoryStore()

	
	//액세스 및 리프레시 토큰 및 auth code를 생성하기 위해 사용되는 시크릿 키.
	//반드시 글자수는 공백 포함 32자여야한다.
	secret = []byte("but i`m to late to make you mine")
	
	//jwt토큰에 서명할 때 사용되는 비밀키 기본값으로 RS256(RSA Signature with SHA-256)
	privateKey, _ = rsa.GenerateKey(rand.Reader, 2048)
)

func printsomething () {
	fmt.Println(store,"store 정보")
}



var oauth2 = compose.ComposeAllEnabled(config, store, secret, privateKey)// 두번째 인자에는 무슨 타입을 넣어도 상관이없다! 그치만 규칙은 지켜야제.

//이 아래에는 세션 생성 코드를 작성한다.
//추측건대 인증서버와 사용자 간의 세션을 만들어주는 역할을 하는 것 같다.
//일단 얘는 제쳐두고 핵심에 집중했다가 다시 볼 예정.
func newSession(user string) *openid.DefaultSession {
	printsomething()
	fmt.Println("세션이 생성됩니다:: 로그인 페이지에서 사용자가 로그인 성공 후 callback 페이지에서 auth code 발급시 세션이 만들어짐")
	return &openid.DefaultSession{
		Claims: &jwt.IDTokenClaims{
			Issuer:      "http://sso.trustnhope.com",//무슨 url? idp인증서버 url인가?
			Subject:     user,
			Audience:    []string{"http://vegas-solution.com", "http://hanchartcloud.com"},
			ExpiresAt:   time.Now().Add(time.Hour * 6),
			IssuedAt:    time.Now(),
			RequestedAt: time.Now(),
			AuthTime:    time.Now(),
		},
		Headers: &jwt.Headers{
			Extra: make(map[string]interface{}),
		},
	}
}
