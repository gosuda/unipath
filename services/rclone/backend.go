package rclone

import (
	_ "github.com/rclone/rclone/backend/http"
	_ "github.com/rclone/rclone/backend/local"
	"github.com/rclone/rclone/fs/config/configfile"
)

func init() {
	configfile.Install()
}
