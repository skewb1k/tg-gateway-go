package tggateway

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	CODE_VALID                 string = "code_valid"
	CODE_INVALID               string = "code_invalid"
	CODE_MAX_ATTEMPTS_EXCEEDED string = "code_max_attempts_exceeded"
	CODE_EXPIRED               string = "expired"
	MESSAGE_SENT               string = "sent"
	MESSAGE_READ               string = "read"
	MESSAGE_REVOKED            string = "revoked"
)

type requestStatusJSON struct {
	RequestID        string   `json:"request_id"`
	PhoneNumber      string   `json:"phone_number"`
	RequestCost      float64  `json:"request_cost"`
	RemainingBalance *float64 `json:"remaining_balance"`
	DeliveryStatus   *struct {
		Status    string `json:"status"`
		UpdatedAt int    `json:"updated_at"`
	} `json:"delivery_status"`
	VerificationStatus *struct {
		Status      string  `json:"status"`
		UpdatedAt   int     `json:"updated_at"`
		CodeEntered *string `json:"code_entered"`
	} `json:"verification_status"`
	Payload *string `json:"payload"`
}

type deliveryStatus struct {
	status    string
	updatedAt int
}

type verificationStatus struct {
	status      string
	updatedAt   int
	codeEntered *string
}

type RequestStatus struct {
	rawJSON            []byte
	requestID          string
	phoneNumber        string
	requestCost        float64
	remainingBalance   *float64
	deliveryStatus     *deliveryStatus
	verificationStatus *verificationStatus
	payload            *string
}

func (r *RequestStatus) UnmarshalJSON(data []byte) error {
	var req requestStatusJSON
	if err := json.Unmarshal(data, &req); err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}

	r.rawJSON = data

	r.requestID = req.RequestID
	r.phoneNumber = req.PhoneNumber
	r.requestCost = req.RequestCost
	r.remainingBalance = req.RemainingBalance
	if req.DeliveryStatus != nil {
		r.deliveryStatus = &deliveryStatus{
			status:    req.DeliveryStatus.Status,
			updatedAt: req.DeliveryStatus.UpdatedAt,
		}
	}
	if req.VerificationStatus != nil {
		r.verificationStatus = &verificationStatus{
			status:      req.VerificationStatus.Status,
			updatedAt:   req.VerificationStatus.UpdatedAt,
			codeEntered: req.VerificationStatus.CodeEntered,
		}
	}
	r.payload = req.Payload

	return nil
}

// Raw JSON data returned by the API.
func (r *RequestStatus) RawJSON() []byte {
	return r.rawJSON
}

// Unique identifier of the verification request.
func (r *RequestStatus) RequestID() string {
	return r.requestID
}

// The phone number to which the verification code was sent, in the E.164 format.
func (r *RequestStatus) PhoneNumber() string {
	return r.phoneNumber
}

// Total request cost incurred by either checkSendAbility or sendVerificationMessage.
func (r *RequestStatus) RequestCost() float64 {
	return r.requestCost
}

// Remaining balance in credits. Returned only in response to a request that incurs a charge.
func (r *RequestStatus) RemainingBalance() *float64 {
	return r.remainingBalance
}

// The current status of the message.
func (r *RequestStatus) DeliveryStatus() *string {
	if r.deliveryStatus != nil {
		return &r.deliveryStatus.status
	}
	return nil
}

// The timestamp when the status was last updated.
func (r *RequestStatus) DeliveryUpdatedAt() *int {
	if r.deliveryStatus != nil {
		return &r.deliveryStatus.updatedAt
	}
	return nil
}

// Delivery status updated at in UTC format.
func (r *RequestStatus) DeliveryUpdatedAtUTC() *time.Time {
	if r.deliveryStatus != nil {
		t := unixToTime(r.deliveryStatus.updatedAt)
		return &t
	}
	return nil
}

// True if the message has been sent to the recipient's device(s).
func (r *RequestStatus) IsMessageSent() bool {
	return r.deliveryStatus != nil && r.deliveryStatus.status == MESSAGE_SENT
}

// True if the message has been read by the recipient.
func (r *RequestStatus) IsMessageRead() bool {
	return r.deliveryStatus != nil && r.deliveryStatus.status == MESSAGE_READ
}

// True if the message has been revoked.
func (r *RequestStatus) IsMessageRevoked() bool {
	return r.deliveryStatus != nil && r.deliveryStatus.status == MESSAGE_REVOKED
}

// The current status of the verification process.
func (r *RequestStatus) VerificationStatus() *string {
	if r.verificationStatus != nil {
		return &r.verificationStatus.status
	}
	return nil
}

// The timestamp for this particular status. Represents the time when the status was last updated.
func (r *RequestStatus) VerificationUpdatedAt() *int {
	if r.verificationStatus != nil {
		return &r.verificationStatus.updatedAt
	}
	return nil
}

// Verification status updated at in UTC format.
func (r *RequestStatus) VerificationUpdatedAtUTC() *time.Time {
	if r.verificationStatus != nil {
		t := unixToTime(r.verificationStatus.updatedAt)
		return &t
	}
	return nil
}

// The code entered by the user.
func (r *RequestStatus) VerificationCodeEntered() *string {
	if r.verificationStatus != nil {
		return r.verificationStatus.codeEntered
	}
	return nil
}

// True if the code entered by the user is correct.
func (r *RequestStatus) IsCodeValid() bool {
	return r.verificationStatus != nil && r.verificationStatus.status == CODE_VALID
}

// True if the code entered by the user is incorrect.
func (r *RequestStatus) IsCodeInvalid() bool {
	return r.verificationStatus != nil && r.verificationStatus.status == CODE_INVALID
}

// True if the maximum number of attempts to enter the code has been exceeded.
func (r *RequestStatus) IsCodeMaxAttemptsExceeded() bool {
	return r.verificationStatus != nil && r.verificationStatus.status == CODE_MAX_ATTEMPTS_EXCEEDED
}

// True if the code has expired and can no longer be used for verification.
func (r *RequestStatus) IsCodeExpired() bool {
	return r.verificationStatus != nil && r.verificationStatus.status == CODE_EXPIRED
}

// Custom payload if it was provided in the request, 0-256 bytes.
func (r *RequestStatus) Payload() *string {
	return r.payload
}
