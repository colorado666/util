package ajwt

import (
	"gitee.com/asktop_golib/util/acast"
	"github.com/dgrijalva/jwt-go"
)

func IsExpired(err error) bool {
	if err == nil {
		return false
	} else {
		return err.Error() == "Token is expired"
	}
}

//JWT : Json Web Token
//规则：36位base64( header的json串 )+"."+base64( claims的json串 )+"."+43位加密算法秘钥加密生成的signature签名( 36位base64( header的json串 )+"."+base64( claims的json串 ) )
//示例：eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjAwODg4MjEsInVzZXJpZCI6MTIzfQ.N5HT1gpwA2tXip9V9-47iwd9fWwHAY5waUZVKleMIkQ

//Header 组成部分：头
//type Header struct {
//	typ string //JWT规范：JWT
//	alg string //签名加密算法：HS256
//}

//Claims 组成部分：有效载荷
//type Claims struct {
//	Id        string `json:"jti,omitempty"` //id
//	Subject   string `json:"sub,omitempty"` //主题
//	Audience  string `json:"aud,omitempty"` //用户
//	Issuer    string `json:"iss,omitempty"` //发行者
//	IssuedAt  int64  `json:"iat,omitempty"` //发行时间
//	NotBefore int64  `json:"nbf,omitempty"` //生效时间
//	ExpiresAt int64  `json:"exp,omitempty"` //过期时间
//}

//Jwt加密生成token
//exp 签名过期时间戳，为0不过期
//secretKey 签名加密秘钥
func Encrypt(info map[string]interface{}, exp int64, secretKey string) (token string, err error) {
	//赋值
	mapClaims := make(jwt.MapClaims)
	if exp >= 0 {
		//超时时间
		mapClaims["exp"] = exp
	}
	for k, v := range info {
		mapClaims[k] = v
	}
	//创建 hs256 类型的 token 对象
	tk := jwt.New(jwt.SigningMethodHS256)
	tk.Claims = mapClaims
	token, err = tk.SignedString([]byte(secretKey))
	return
}

//Jwt解密token
//secretKey 签名加密秘钥
func Decrypt(token string, secretKey string) (info map[string]interface{}, err error) {
	//将token字符串解析成token对象，会自动校验有效性，超时会报错
	tk, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return
	}
	return tk.Claims.(jwt.MapClaims), nil
}

//Jwt解密token
//secretKey 签名加密秘钥
func DecryptObj(token string, info interface{}, secretKey string) error {
	data, err := Decrypt(token, secretKey)
	if err != nil {
		return err
	} else {
		return acast.MapToStruct(data, info)
	}
}

//Jwt解密token（不校验过期等）
//secretKey 签名加密秘钥
func DecryptUnValid(token string, secretKey string) (info map[string]interface{}, err error) {
	//将token字符串解析成token对象，会自动校验有效性，超时会报错
	tk, err := jwt.ParseWithClaims(token, MapClaims{}, func(*jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return
	}
	return tk.Claims.(MapClaims), nil
}

//Jwt解密token（不校验过期等）
//secretKey 签名加密秘钥
func DecryptObjUnValid(token string, info interface{}, secretKey string) error {
	data, err := DecryptUnValid(token, secretKey)
	if err != nil {
		return err
	} else {
		return acast.MapToStruct(data, info)
	}
}
