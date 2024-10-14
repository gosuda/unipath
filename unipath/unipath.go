package unipath

import "net/url"

type UniPath struct {
	Protocol Protocol
	Host     string
	Path     string
	User     string
	Password string
}

func (u *UniPath) Url() *url.URL {
	var user *url.Userinfo
	if u.User != "" {
		user = url.User(u.User)
	}
	if u.Password != "" {
		user = url.UserPassword(u.User, u.Password)
	}

	return &url.URL{
		Scheme: u.Protocol.String(),
		Host:   u.Host,
		User:   user,
		Path:   u.Path,
	}
}

func (uniPath *UniPath) String() string {
	return uniPath.Url().String()
}
