package authorizationserver

import "github.com/ory/fosite"

//서비스 서버 목록. (DB에서 불러와? 아니면 LDAP 서버에서 불러와?)
//DB든 LDAP 서버든 어디에다가 어떻게든 저장해서 이 인증서버랑 연결하고 데이터를 가져와서 할당해주면 돼.
var clientlist = map[string]fosite.Client{
	"service1": &fosite.DefaultClient{
		ID:             "my-client",
		Secret:         []byte(`$2a$10$IxMdI6d.LIRZPpSfEwNoeu4rY3FhDREsxFJXikcgdRRAStxUlsuEO`),            // = "foobar"
		RotatedSecrets: nil, //이 값 사용할 경우 서비스 서버도 매번 비번을 바꿔야함ㅍ, [][]byte{[]byte(`$2y$10$X51gLxUQJ.hGw1epgHTE5u0bt64xM0COU7K9iAp.OFg8p2pUd.1zC `)}  = "foobaz", 
		RedirectURIs:   []string{"http://vegas-solution.com/callback"},
		ResponseTypes:  []string{"code","id_token token"}, //둘 중 하나 택해야 함: 토큰을 직접 줄건지, 코드 줘서 교환하게 할건지
		GrantTypes:     []string{"authorization_code"},
		Scopes:         []string{"openid", "offline"},
	},
	"service2": &fosite.DefaultClient{
		ID:             "encoded:client",
		Secret:         []byte(`$2a$10$A7M8b65dSSKGHF0H2sNkn.9Z0hT8U1Nv6OWPV3teUUaczXkVkxuDS`), // = "encoded&password"
		RotatedSecrets: nil,
		RedirectURIs:   []string{"http://hanchartcloud.com/callback"},
		ResponseTypes:  []string{"code","id_token token"},
		GrantTypes:     []string{"authorization_code"},
		Scopes:         []string{"openid", "offline"},
	},
}