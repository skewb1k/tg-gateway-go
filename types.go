package tggateway

type DeliveryStatusEnum string

const (
	DeliveryStatusSent    DeliveryStatusEnum = "sent"
	DeliveryStatusRead    DeliveryStatusEnum = "read"
	DeliveryStatusRevoked DeliveryStatusEnum = "revoked"
)

type DeliveryStatus struct {
	// The current status of the message. One of the following:
	//
	// - sent – the message has been sent to the recipient's device(s),
	//
	// - delivered – the message has been delivered to the recipient's device(s),
	//
	// - read – the message has been read by the recipient,
	//
	// - expired – the message has expired without being delivered or read,
	//
	// - revoked – the message has been revoked.
	Status DeliveryStatusEnum `json:"status"`
	// The timestamp when the status was last updated.
	UpdatedAt UnixTime `json:"updated_at"`
}

type VerificationStatusEnum string

func (vs VerificationStatusEnum) IsValid() bool {
	return vs == VerificationStatusValid
}

const (
	VerificationStatusValid               VerificationStatusEnum = "code_valid"
	VerificationStatusInvalid             VerificationStatusEnum = "code_invalid"
	VerificationStatusMaxAttemptsExceeded VerificationStatusEnum = "code_max_attempts_exceeded"
	VerificationStatusExpired             VerificationStatusEnum = "expired"
)

type VerificationStatus struct {
	// The current status of the verification process. One of the following:
	//
	// - code_valid – the code entered by the user is correct,
	//
	// - code_invalid – the code entered by the user is incorrect,
	//
	// - code_max_attempts_exceeded – the maximum number of attempts to enter the code has been exceeded,
	//
	// - expired – the code has expired and can no longer be used for verification.
	Status VerificationStatusEnum `json:"status"`
	// The timestamp for this particular status. Represents the time when the status was last updated.
	UpdatedAt UnixTime `json:"updated_at"`
	// The code entered by the user.
	CodeEntered *string `json:"code_entered"`
}

type RequestStatus struct {
	// Unique identifier of the verification request.
	RequestID string `json:"request_id"`
	// The phone number to which the verification code was sent, in the E.164 format.
	PhoneNumber string `json:"phone_number"`
	// Total request cost incurred by either checkSendAbility or sendVerificationMessage.
	RequestCost float64 `json:"request_cost"`
	// If True, the request fee was refunded.
	IsRefunded *bool `json:"is_refunded"`
	// Remaining balance in credits. Returned only in response to a request that incurs a charge.
	RemainingBalance *float64 `json:"remaining_balance"`
	// The current message delivery status. Returned only if a verification message was sent to the user.
	DeliveryStatus *DeliveryStatus `json:"delivery_status"`
	// The current status of the verification process.
	VerificationStatus *VerificationStatus `json:"verification_status"`
	// Custom payload if it was provided in the request, 0-256 bytes.
	Payload *string `json:"payload"`
}
