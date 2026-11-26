package unipath

import "net/url"

func NewUniPathFromString(rawurl string) (*UniPath, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	return &UniPath{
		Protocol: parseProtocol(u.Scheme),
		Host:     u.Host,
		Path:     u.Path,
	}, nil
}

type UniPath struct {
	Protocol  Protocol
	Authority string
	Host      string
	Path      string
}

func (u *UniPath) Url() *url.URL {
	return &url.URL{
		Scheme: u.Protocol.String(),
		Host:   u.Host,
		Path:   u.Path,
	}
}

func (u *UniPath) String() string {
	return u.Url().String()
}

func (u *UniPath) IsLocal() bool {
	return u.Protocol == Local
}
