package downloader

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

func (u *UniPath) Url() string {
	return u.Host + ":" + u.Path
}
