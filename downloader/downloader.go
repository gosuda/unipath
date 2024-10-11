package downloader

type Downloader interface {
	Download(remoteUrl, localPath string) error
}
