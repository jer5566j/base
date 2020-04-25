package token

type Token interface {
	Verify(token string, secret string) bool
	Expires() int64
}
