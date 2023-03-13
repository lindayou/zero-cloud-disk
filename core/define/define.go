package define

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

//jwtKey
var JwtKey = "cloud-disk-key"

// 连接邮箱授权码
var Password = "VFKEMTNHAWUFWQNZ"

//验证码长度
var CodeLength = 6

// 验证码过期时间
var TimeExpired = 60 * time.Second
