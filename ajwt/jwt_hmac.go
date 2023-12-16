package ajwt

import (
	"gitee.com/asktop_golib/util/acast"
	"github.com/dgrijalva/jwt-go"
)

// 生成Jwt
// exp 签名过期时间戳，为0不过期
// keyStr 秘钥
func CreateJwtHmac(params map[string]interface{}, exp int64, keyStr string) (jwtStr string, err error) {
	// 赋值
	mapClaims := make(jwt.MapClaims)
	if exp >= 0 {
		// 超时时间
		mapClaims["exp"] = exp
	}
	for k, v := range params {
		mapClaims[k] = v
	}

	// 创建 hs256 类型的 token 对象
	tk := jwt.New(jwt.SigningMethodHS256)
	tk.Claims = mapClaims
	jwtStr, err = tk.SignedString([]byte(keyStr))
	return
}

// 解析Jwt
// keyStr 秘钥
func ParseJwtHmac(jwtStr string, keyStr string) (data map[string]interface{}, err error) {
	// 将token字符串解析成token对象，会自动校验有效性，超时会报错
	tk, err := jwt.Parse(jwtStr, func(*jwt.Token) (interface{}, error) {
		return []byte(keyStr), nil
	})
	if err != nil {
		return
	}

	return tk.Claims.(jwt.MapClaims), nil
}

// 解析Jwt为指定结构
// keyStr 密钥
func ParseJwtHmacObj(jwtStr string, obj interface{}, keyStr string) error {
	data, err := ParseJwtHmac(jwtStr, keyStr)
	if err != nil {
		return err
	} else {
		return acast.MapToStruct(data, obj)
	}
}

// 解析Jwt，不校验过期
// keyStr 密钥
func ParseJwtHmacUnValid(jwtStr string, keyStr string) (data map[string]interface{}, err error) {
	// 解析jwt字符串，会自动校验有效性，超时会报错
	tk, err := jwt.ParseWithClaims(jwtStr, MapClaims{}, func(*jwt.Token) (interface{}, error) {
		return keyStr, nil
	})
	if err != nil {
		return
	}

	return tk.Claims.(MapClaims), nil
}

// 解析Jwt为指定结构，不校验过期
// keyStr 密钥
func ParseJwtHmacObjUnValid(token string, info interface{}, keyStr string) error {
	data, err := ParseJwtHmacUnValid(token, keyStr)
	if err != nil {
		return err
	} else {
		return acast.MapToStruct(data, info)
	}
}

