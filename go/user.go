package client

import (
	"time"
)

type TokenType string

const (
	TOKEN_TYPE_KINGTRADER TokenType = "kt"
	TOKEN_TYPE_KTKT       TokenType = "ktkt"
	TOKEN_TYPE_LIANGTOU   TokenType = "pc"
	TOKEN_TYPE_MOBILE     TokenType = "mobile"
)

var ErrNotSuportLoginType = errors.New("请使用手机或邮箱账号登陆")

type User struct {
	// base
	Id       int64  `json:"id"`
	NickName string `json:"nick_name"`
	UserName string `json:"user_name"`
	Email    string `json:email`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
	Hash     string `json:"-"`
	Password string `json:"-"`
	Status   string `json:"status"`

	// hexun
	HexunId          int64     `json:"hexun_id"`
	HexunUserName    string    `json:"hexun_user_name"`
	Hexun2026Endtime time.Time `json:"hexun_2026_endtime"`
	Hexun2027Endtime time.Time `json:"hexun_2027_endtime"`

	// kingstrader
	SingleScreenQuoteExpired time.Time `json:"single_screen_quote_expired"`
	MultiScreenQuoteExpired  time.Time `json:"multi_screen_quote_expired"`
}

type UserService struct {
	UserGetUidByToken  func(token string, origin TokenType) (int64, error)
	UserGetInfoById    func(uid int64) (*User, error)
	UserGetInfoByToken func(token string, origin TokenType) (*User, error)
	UserGetTokenByUid  func(uid int64, origin TokenType) (string, error)

	// 通过用户名查
	UserGetInfoByName func(name string) (*User, error)

	// 通过手机查用户
	UserGetInfoByPhone func(phone string) (*User, error)
	// 通过email查用户
	UserGetInfoByEmail func(email string) (*User, error)

	// 验证用户密码
	UserCheckPassword func(uid int64, pass string) (bool, error)

	// 加密
	UserHashPassword func(pass, hash string) string

	// 验证用户账号
	UserAuth func(user, pass string) (*User, error)

	// 更新用户token
	UserTokenUpdate func(uid int64, origin string) (string, error)
}

var UserClient *UserService

func IsTokenType(t TokenType) bool {
	switch t {
	case TOKEN_TYPE_KINGTRADER:
		return true
	case TOKEN_TYPE_KTKT:
		return true
	case TOKEN_TYPE_LIANGTOU:
		return true
	case TOKEN_TYPE_MOBILE:
		return true
	default:
		return false
	}
}
