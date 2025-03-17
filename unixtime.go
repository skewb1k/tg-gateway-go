package tggateway

import (
	"encoding/json"
	"fmt"
	"time"
)

type UnixTime struct {
	time.Time
}

// Custom unmarshal function to convert Unix timestamp to time.Time
func (ut *UnixTime) UnmarshalJSON(data []byte) error {
	// The timestamp in the data will be a number (Unix timestamp in seconds)
	var unixTimestamp int64
	if err := json.Unmarshal(data, &unixTimestamp); err != nil {
		return fmt.Errorf("failed to unmarshal unix time string: %w", err)
	}

	ut.Time = time.Unix(unixTimestamp, 0)
	return nil
}
