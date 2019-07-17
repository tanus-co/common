package jwt

type Header struct {
	Alg string `json:"alg,omitempty"`
	Typ string `json:"typ,omitempty"`
}

type Payload struct {
	Iss     string `json:"iss,omitempty"`
	Exp     int64  `json:"exp,omitempty"`
	Sub     string `json:"sub,omitempty"`
	Aud     string `json:"aud,omitempty"`
	Nbf     int64  `json:"nbf,omitempty"`
	Iat     int64  `json:"iat,omitempty"`
	Jti     string `json:"jti,omitempty"`
	Tenant  int64  `json:"tenant,string,omitempty"`
	UserId  int64  `json:"user_id,string,omitempty"`
	GrantId int64  `json:"grant_id,string,omitempty"`
}
