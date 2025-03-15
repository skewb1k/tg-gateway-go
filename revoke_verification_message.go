package tggateway

import (
	"context"
	"fmt"
)

type RevokeVerificationMessageParams struct {
	// The unique identifier of the request whose verification message you want to revoke.
	RequestID string `json:"request_id"`
}

// Use this method to revoke a verification message that was sent previously. Returns True if the revocation request was
// received. However, this does not guarantee that the message will be deleted. For example, if the message has already
// been delivered or read, it will not be removed.
func (c Client) RevokeVerificationMessage(ctx context.Context, params *RevokeVerificationMessageParams) (*bool, error) {
	var result struct {
		Ok    bool    `json:"ok"`
		Error *string `json:"error"`
		// todo: wrap result in struct with IsOk() method.
		Result *bool `json:"result"`
	}

	err := c.makeAPIRequest(ctx, "revokeVerificationMessage", params, &result)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, fmt.Errorf("revoke verication message failed: %w", c.mapErr(*result.Error))
	}

	return result.Result, nil
}
