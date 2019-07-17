package jwt

import (
	"testing"
	"time"
)

func TestJwt(t *testing.T) {
	p := &Payload{
		Iss: "iss_ooo",
		Exp: (time.Now().UnixNano() / 1e6) + (15 * 24 * 60 * 60 * 1000),
		Nbf: time.Now().UnixNano() / 1e6,
		Iat: 0,
	}
	h := &Header{
		Alg: "SH256",
		Typ: "",
	}

	encode := Encode(*h, *p)
	t.Log(encode)

	header, payload, err := Decode(encode)
	if err != nil {
		t.Error(err)
	}

	t.Log(header.Alg)
	t.Log(p.Iss, payload.Iss)
	t.Log(p.Exp, payload.Exp)
}
