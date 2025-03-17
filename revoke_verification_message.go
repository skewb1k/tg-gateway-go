package tggateway

import (
	"context"
	"encoding/json"
	"fmt"
)

type RevokeVerificationMessageParams struct {
	// The unique identifier of the request whose verification message you want to revoke.
	RequestID string `json:"request_id"`
}

// Use this method to revoke a verification message that was sent previously. Returns True if the revocation request was
// received. However, this does not guarantee that the message will be deleted. For example, if the message has already
// been delivered or read, it will not be removed.
func (c Client) RevokeVerificationMessage(ctx context.Context, params *RevokeVerificationMessageParams) (bool, error) {
	if params == nil {
		panic("param must be not nil")
	}

	var result bool
	resultBytes, err := c.makeAPIRequest(ctx, "revokeVerificationMessage", params, &result)
	if err != nil {
		return false, err
	}

	if err := json.Unmarshal(resultBytes, &result); err != nil {
		return false, fmt.Errorf("failed to unmarshal result: %w", err)
	}

	return result, nil
}
