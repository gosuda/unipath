package downloader

import (
	_ "github.com/rclone/rclone/backend/local"
	"github.com/rclone/rclone/fs/config/configfile"
	_ "github.com/rclone/rclone/lib/plugin"
)

func init() {
	configfile.Install()
}
