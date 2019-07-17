package jwt

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestDef(t *testing.T) {
	payload := &Payload{
		Iss: "iss_ooo",
		Exp: (time.Now().UnixNano() / 1e6) + (15 * 24 * 60 * 60 * 1000),
		Nbf: time.Now().UnixNano() / 1e6,
		Iat: 0,
	}

	b, _ := json.Marshal(payload)
	fmt.Println(string(b))

	unix := time.Unix(0, time.Now().UnixNano())
	fmt.Println(unix)
}

//1562118265000
//1562118296242
//1562118414295
//1562118256456585
