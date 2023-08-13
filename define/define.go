package define

import "time"

var (
	DefaultPage = "1"
	DefaultSize = "10"
	TokenExp    = time.Now().Add(time.Hour * 24 * 7)
)
