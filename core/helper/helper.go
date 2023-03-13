package helper

import (
	"cloud-disk/core/define"
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-wright/email"
	"github.com/satori/go.uuid"
	"math/rand"
	"net/smtp"
	"time"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func GenerateToken(id int, identity string, name string) (string, error) {
	//ID Identity Name
	//定义对象
	uc := new(define.UserClaim)
	uc.Name = name
	uc.Id = id
	uc.Identity = identity
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))

	if err != nil {
		return "token SignedString failed,err : ", err
	}

	//ucc :=define.UserClaim{
	//	Id: id,
	//	Name: name,
	//	Identity: identity,
	//}

	return tokenString, nil
}

func MailSendCode(mail, code string) error {
	e := email.NewEmail()
	e.From = "Get <yeppydeyouxiang@163.com>"
	e.To = []string{mail}
	e.Subject = "验证码发送测试"
	//e.Text = []byte("你的验证码为：<h1>" + + "<h1>")
	e.HTML = []byte("你的验证码为：<h1>" + code + "<h1>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "yeppydeyouxiang@163.com", define.Password, "smtp.163.com"), &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         "smtp.163.com",
	})
	if err != nil {
		return err
	}
	return nil
}

func RandCode() string {
	s := "1234567890"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < define.CodeLength; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}

func GenerateUuid() string {
	return uuid.NewV4().String()
}
