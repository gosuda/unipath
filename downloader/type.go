package downloader

type UrlType uint8

const (
	File UrlType = iota
	Http
	Ftp
	Sftp
	Torrent
	Ipfs

	Dropbox
	Googlecloudstorage
	Googlephotos
	Onedrive
	Oracleobjectstorage
	S3
)
