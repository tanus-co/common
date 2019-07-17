package jwt

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	h "github.com/tanus-co/common/util/hash"
	"log"
	"strings"
)

const secret = "0933e54e76b24731a2d84b6b463ec04c"

func Encode(header Header, payload Payload) string {
	headerJson, err := json.Marshal(header)
	if err != nil {
		log.Fatal(err)
	}
	headerBase64 := base64.StdEncoding.EncodeToString(headerJson)
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	payloadBase64 := base64.URLEncoding.EncodeToString(payloadJson)

	body := fmt.Sprintf("%s.%s", headerBase64, payloadBase64)

	jwt := ""
	if hashResult := hash(header.Alg, body); hashResult != "" {
		jwt = fmt.Sprintf("%s.%s",
			body,
			base64.URLEncoding.EncodeToString(
				[]byte(hashResult)),
		)
	}
	return jwt
}

func Decode(jwt string) (Header, Payload, error) {
	splits := strings.Split(jwt, ".")
	if len(splits) != 3 {
		return Header{}, Payload{}, errors.New("jwt error")
	}
	headerBase64, payloadBase64, signature := splits[0], splits[1], splits[2]

	headerJson, err := base64.URLEncoding.DecodeString(headerBase64)
	if err != nil {
		log.Fatal(err)
	}
	payloadJson, err := base64.URLEncoding.DecodeString(payloadBase64)
	if err != nil {
		log.Fatal(err)
	}

	header := Header{}
	err = json.Unmarshal(headerJson, &header)
	if err != nil {
		log.Fatal(err)
	}

	payload := Payload{}
	err = json.Unmarshal(payloadJson, &payload)
	if err != nil {
		log.Fatal(err)
	}

	body := fmt.Sprintf("%s.%s", headerBase64, payloadBase64)
	hashResult := hash(header.Alg, body)
	if base64.URLEncoding.EncodeToString([]byte(hashResult)) != signature {
		return Header{}, Payload{}, errors.New("jwt error")
	}

	return header, payload, nil
}

func hash(alg, body string) string {
	hashResult := ""
	switch alg {
	case "SH256":
		hashResult = h.HmacSha256(body, secret)
	}
	return hashResult
}
