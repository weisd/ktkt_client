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

type User struct {
	// base
	Id       int64
	NickName string
	UserName string
	Email    string
	Phone    string
	Avatar   string
	Hash     string
	Password string
	Status   string

	// hexun
	HexunId          int64
	HexunUserName    string
	Hexun2026Endtime time.Time `json:hexun_2026_endtime`
	Hexun2027Endtime time.Time `json:hexun_2027_endtime`

	// kingstrader
	SingleScreenQuoteExpired time.Time `json:single_screen_quote_expired`
	MultiScreenQuoteExpired  time.Time `json:multi_screen_quote_expired`
}

type UserService struct {
	UserGetUidByToken  func(token string, origin TokenType) (int64, error)
	UserGetInfoById    func(uid int64) (*User, error)
	UserGetInfoByToken func(token string, origin TokenType) (*User, error)
}

var UserClient *UserService
