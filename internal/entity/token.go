package entity

type JWT struct {
	AccessToken        string `json:"access_token"`
	RefreshToken       string `json:"refresh_token"`
	AccessTokenMaxAge  int    `json:"-"`
	RefreshTokenMaxAge int    `json:"-"`
	Domain             string `json:"-"`
}
