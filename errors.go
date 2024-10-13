package tggateway

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrCodeInvalid             = errors.New("CODE_INVALID")
	ErrCodeExpired             = errors.New("CODE_EXPIRED")
	ErrCodeLengthRequired      = errors.New("CODE_LENGTH_REQUIRED")
	ErrCodeLengthInvalid       = errors.New("CODE_LENGTH_INVALID")
	ErrCodeMaxAttemptsExceeded = errors.New("CODE_MAX_ATTEMPTS_EXCEEDED")
	ErrPhoneNumberInvalid      = errors.New("PHONE_NUMBER_INVALID")
	ErrPhoneNumberNotFound     = errors.New("PHONE_NUMBER_NOT_FOUND")
	ErrPhoneNumberMismatch     = errors.New("PHONE_NUMBER_MISMATCH")
	ErrRequestIdInvalid        = errors.New("REQUEST_ID_INVALID")
	ErrRequestIdRequired       = errors.New("REQUEST_ID_REQUIRED")
	ErrPayloadInvalid          = errors.New("PAYLOAD_INVALID")
	ErrSenderUsernameInvalid   = errors.New("SENDER_USERNAME_INVALID")
	ErrSenderNotVerified       = errors.New("SENDER_NOT_VERIFIED")
	ErrSenderNotOwned          = errors.New("SENDER_NOT_OWNED")
	ErrCallbackUrlInvalid      = errors.New("CALLBACK_URL_INVALID")
	ErrTtlInvalid              = errors.New("TTL_INVALID")
	ErrAccessTokenInvalid      = errors.New("ACCESS_TOKEN_INVALID")
	ErrAccessTokenRequired     = errors.New("ACCESS_TOKEN_REQUIRED")
	ErrMessageAlreadySent      = errors.New("MESSAGE_ALREADY_SENT")
	ErrBalanceNotEnough        = errors.New("BALANCE_NOT_ENOUGH")
	ErrFloodWait               = errors.New("FLOOD_WAIT")
	ErrUnknownMethod           = errors.New("UNKNOWN_METHOD")
	ErrUnknown                 = errors.New("UNKNOWN_ERROR")
)

var errorsMap = map[string]error{
	"CODE_INVALID":               ErrCodeInvalid,
	"CODE_EXPIRED":               ErrCodeExpired,
	"CODE_LENGTH_REQUIRED":       ErrCodeLengthRequired,
	"CODE_LENGTH_INVALID":        ErrCodeLengthInvalid,
	"CODE_MAX_ATTEMPTS_EXCEEDED": ErrCodeMaxAttemptsExceeded,
	"PHONE_NUMBER_INVALID":       ErrPhoneNumberInvalid,
	"PHONE_NUMBER_NOT_FOUND":     ErrPhoneNumberNotFound,
	"PHONE_NUMBER_MISMATCH":      ErrPhoneNumberMismatch,
	"REQUEST_ID_INVALID":         ErrRequestIdInvalid,
	"REQUEST_ID_REQUIRED":        ErrRequestIdRequired,
	"PAYLOAD_INVALID":            ErrPayloadInvalid,
	"SENDER_USERNAME_INVALID":    ErrSenderUsernameInvalid,
	"SENDER_NOT_VERIFIED":        ErrSenderNotVerified,
	"SENDER_NOT_OWNED":           ErrSenderNotOwned,
	"CALLBACK_URL_INVALID":       ErrCallbackUrlInvalid,
	"TTL_INVALID":                ErrTtlInvalid,
	"ACCESS_TOKEN_INVALID":       ErrAccessTokenInvalid,
	"ACCESS_TOKEN_REQUIRED":      ErrAccessTokenRequired,
	"MESSAGE_ALREADY_SENT":       ErrMessageAlreadySent,
	"BALANCE_NOT_ENOUGH":         ErrBalanceNotEnough,
	"UNKNOWN_METHOD":             ErrUnknownMethod,
}

func (c Client) mapErr(errVal string) error {
	if strings.Contains(errVal, "FLOOD_WAIT") {
		return ErrFloodWait
	}

	if err, found := errorsMap[errVal]; found {
		return err
	}

	fmt.Printf("unknown error:%s\n", errVal)
	return ErrUnknown
}
