package unipath

import "strings"

type Protocol uint16

const (
	Unknown Protocol = iota
	Local
	Http
	Https
	Ftp
	Sftp
	Nfs
	Torrent
	Magnet
	Ipfs
	S3
	GoogleCloudStorage
	Dropbox
	GooglePhotos
	Onedrive
	OracleObjectStorage
)

func parseProtocol(scheme string) Protocol {
	switch strings.ToLower(scheme) {
	case "file", "":
		return Local
	case "http":
		return Http
	case "https":
		return Https
	case "ftp":
		return Ftp
	case "sftp":
		return Sftp
	case "nfs":
		return Nfs
	case "torrent":
		return Torrent
	case "magnet":
		return Magnet
	case "ipfs":
		return Ipfs
	case "s3":
		return S3
	case "gs":
		return GoogleCloudStorage
	case "dropbox":
		return Dropbox
	case "gphotos":
		return GooglePhotos
	case "onedrive":
		return Onedrive
	case "oracle":
		return OracleObjectStorage
	default:
		return Unknown
	}
}

func (p Protocol) String() string {
	switch p {
	case Local:
		return "file://"
	case Http:
		return "http://"
	case Https:
		return "https://"
	case Ftp:
		return "ftp://"
	case Sftp:
		return "sftp://"
	case Nfs:
		return "nfs://"
	case Torrent:
		return "torrent://"
	case Magnet:
		return "magnet:?"
	case Ipfs:
		return "ipfs://"
	case Dropbox:
		return "dropbox://"
	case GoogleCloudStorage:
		return "gs://"
	case GooglePhotos:
		return "gphotos://"
	case Onedrive:
		return "onedrive://"
	case OracleObjectStorage:
		return "oracle://"
	case S3:
		return "s3://"
	default:
		return "unknown://"
	}
}
