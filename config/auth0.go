package config

import (
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"time"
)

type Auth0ConfigType struct {
	Domain             string
	Audience           []string
	Issuer             string
	SignatureAlgorithm validator.SignatureAlgorithm
	CacheDuration      time.Duration
}

var Auth0Config = Auth0ConfigType{
	Domain:             "dev-h77prh7uq1xin4u0.us.auth0.com",
	Audience:           []string{"https://api.neplode.com/"},
	Issuer:             "https://" + "dev-h77prh7uq1xin4u0.us.auth0.com" + "/",
	SignatureAlgorithm: validator.RS256,
	CacheDuration:      15 * time.Minute,
}
