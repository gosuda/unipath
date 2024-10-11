package downloader

import "net/url"

type Protocol uint16

const (
	Local Protocol = iota
	Http
	Ftp
	Sftp
	Nfs
	Torrent
	Ipfs

	Dropbox
	Googlecloudstorage
	Googlephotos
	Onedrive
	Oracleobjectstorage
	S3
)

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
		Scheme: string(u.Protocol),
		Host:   u.Host,
		User:   user,
		Path:   u.Path,
	}
}

func (uniPath *UniPath) String() string {
	return uniPath.Url().String()
}
