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

func (p Protocol) String() string {
	switch p {
	case Local:
		return "file://"
	case Http:
		return "http://"
	case Ftp:
		return "ftp://"
	case Sftp:
		return "sftp://"
	case Nfs:
		return "nfs://"
	case Torrent:
		return "torrent://"
	case Ipfs:
		return "ipfs://"
	case Dropbox:
		return "dropbox://"
	case Googlecloudstorage:
		return "gs://"
	case Googlephotos:
		return "gphotos://"
	case Onedrive:
		return "onedrive://"
	case Oracleobjectstorage:
		return "oracle://"
	case S3:
		return "s3://"
	default:
		return "unknown://"
	}
}

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
