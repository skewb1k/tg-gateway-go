package tggateway

import "time"

func unixToTime(unix int) time.Time {
	return time.Unix(int64(unix), 0)
}
