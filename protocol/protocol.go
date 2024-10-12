package protocol

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
