package tggateway

import (
	"encoding/json"
	"time"
)

const (
	CODE_VALID                 string = "code_valid"
	CODE_INVALID                      = "code_invalid"
	CODE_MAX_ATTEMPTS_EXCEEDED        = "code_max_attempts_exceeded"
	CODE_EXPIRED                      = "expired"
	MESSAGE_SENT                      = "sent"
	MESSAGE_READ                      = "read"
	MESSAGE_REVOKED                   = "revoked"
)

type requestStatusJson struct {
	RequestId        string   `json:"request_id"`
	PhoneNumber      string   `json:"phone_number"`
	RequestCost      float32  `json:"request_cost"`
	RemainingBalance *float32 `json:"remaining_balance"`
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
	rawJson            []byte
	requestId          string
	phoneNumber        string
	requestCost        float32
	remainingBalance   *float32
	deliveryStatus     *deliveryStatus
	verificationStatus *verificationStatus
	payload            *string
}

func (r *RequestStatus) UnmarshalJSON(data []byte) error {
	var requestStatusJson requestStatusJson
	if err := json.Unmarshal(data, &requestStatusJson); err != nil {
		return err
	}

	r.rawJson = data

	r.requestId = requestStatusJson.RequestId
	r.phoneNumber = requestStatusJson.PhoneNumber
	r.requestCost = requestStatusJson.RequestCost
	r.remainingBalance = requestStatusJson.RemainingBalance
	if requestStatusJson.DeliveryStatus != nil {
		r.deliveryStatus = &deliveryStatus{
			status:    requestStatusJson.DeliveryStatus.Status,
			updatedAt: requestStatusJson.DeliveryStatus.UpdatedAt,
		}
	}
	if requestStatusJson.VerificationStatus != nil {
		r.verificationStatus = &verificationStatus{
			status:      requestStatusJson.VerificationStatus.Status,
			updatedAt:   requestStatusJson.VerificationStatus.UpdatedAt,
			codeEntered: requestStatusJson.VerificationStatus.CodeEntered,
		}
	}
	r.payload = requestStatusJson.Payload

	return nil
}

// Raw JSON data returned by the API.
func (r RequestStatus) RawJson() []byte {
	return r.rawJson
}

// Unique identifier of the verification request.
func (r RequestStatus) RequestId() string {
	return r.requestId
}

// The phone number to which the verification code was sent, in the E.164 format.
func (r RequestStatus) PhoneNumber() string {
	return r.phoneNumber
}

// Total request cost incurred by either checkSendAbility or sendVerificationMessage.
func (r RequestStatus) RequestCost() float32 {
	return r.requestCost
}

// Remaining balance in credits. Returned only in response to a request that incurs a charge.
func (r RequestStatus) RemainingBalance() *float32 {
	return r.remainingBalance
}

// The current status of the message.
func (r RequestStatus) DeliveryStatus() *string {
	if r.deliveryStatus != nil {
		return &r.deliveryStatus.status
	}
	return nil
}

// The timestamp when the status was last updated.
func (r RequestStatus) DeliveryUpdatedAt() *int {
	if r.deliveryStatus != nil {
		return &r.deliveryStatus.updatedAt
	}
	return nil
}

// Delivery status updated at in UTC format.
func (r RequestStatus) DeliveryUpdatedAtUTC() *time.Time {
	if r.deliveryStatus != nil {
		t := unixToTime(r.deliveryStatus.updatedAt)
		return &t
	}
	return nil
}

// True if the message has been sent to the recipient's device(s).
func (r RequestStatus) IsMessageSent() bool {
	return r.deliveryStatus != nil && r.deliveryStatus.status == MESSAGE_SENT
}

// True if the message has been read by the recipient.
func (r RequestStatus) IsMessageRead() bool {
	return r.deliveryStatus != nil && r.deliveryStatus.status == MESSAGE_READ
}

// True if the message has been revoked.
func (r RequestStatus) IsMessageRevoked() bool {
	return r.deliveryStatus != nil && r.deliveryStatus.status == MESSAGE_REVOKED
}

// The current status of the verification process.
func (r RequestStatus) VerificationStatus() *string {
	if r.verificationStatus != nil {
		return &r.verificationStatus.status
	}
	return nil
}

// The timestamp for this particular status. Represents the time when the status was last updated.
func (r RequestStatus) VerificationUpdatedAt() *int {
	if r.verificationStatus != nil {
		return &r.verificationStatus.updatedAt
	}
	return nil
}

// Verification status updated at in UTC format.
func (r RequestStatus) VerificationUpdatedAtUTC() *time.Time {
	if r.verificationStatus != nil {
		t := unixToTime(r.verificationStatus.updatedAt)
		return &t
	}
	return nil
}

// The code entered by the user.
func (r RequestStatus) VerificationCodeEntered() *string {
	if r.verificationStatus != nil {
		return r.verificationStatus.codeEntered
	}
	return nil
}

// True if the code entered by the user is correct.
func (r RequestStatus) IsCodeValid() bool {
	return r.verificationStatus != nil && r.verificationStatus.status == CODE_VALID
}

// True if the code entered by the user is incorrect.
func (r RequestStatus) IsCodeInvalid() bool {
	return r.verificationStatus != nil && r.verificationStatus.status == CODE_INVALID
}

// True if the maximum number of attempts to enter the code has been exceeded.
func (r RequestStatus) IsCodeMaxAttemptsExceeded() bool {
	return r.verificationStatus != nil && r.verificationStatus.status == CODE_MAX_ATTEMPTS_EXCEEDED
}

// True if the code has expired and can no longer be used for verification.
func (r RequestStatus) IsCodeExpired() bool {
	return r.verificationStatus != nil && r.verificationStatus.status == CODE_EXPIRED
}

// Custom payload if it was provided in the request, 0-256 bytes.
func (r RequestStatus) Payload() *string {
	return r.payload
}
