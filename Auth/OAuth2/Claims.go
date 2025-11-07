package OAuth2

import "time"

type Claims struct {
	Name       string
	Issuer     string
	Subject    int32
	Audience   int32
	Expiration time.Time
	NotBefore  time.Time
	IssuedAt   time.Time
}
