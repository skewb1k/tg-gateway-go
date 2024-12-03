# Golang Telegram Gateway API

Zero-dependencies Go package for [Telegram Gateway API](https://core.telegram.org/gateway).

**Key features:**

- Full support of Telegram Gateway API.
- 100% documented.
- Context support.
- Type-safe error handling.
- Immutable response structs.
- No external dependencies.

## Installation

```bash
go get github.com/skewb1k/tg-gateway-go
```

## Usage

Basic example:

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	tggateway "github.com/skewb1k/tg-gateway-go"
)

func main() {
	token := "XXXXXXXX-XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"

	client := tggateway.NewClient(token)

	ctx := context.Background()

	sendVerificationMessageResp, err := client.SendVerificationMessage(ctx, &tggateway.SendVerificationMessageParams{
		PhoneNumber: "11111111111",
		CodeLength:  4,
		// or set your generated code.
		// Code: "1234",
	})
	if err != nil {
		// check error type.
		if errors.Is(err, tggateway.ErrPhoneNumberInvalid) {
			// ...handle error
			log.Fatal("phone number is invalid")
		}

		log.Fatalf("send verification message error: %s", err.Error())
	}

	// store recieved request id.
	requestId := sendVerificationMessageResp.RequestId()

	// get code from some user input.
	enteredCode := userInput.Code

	checkVerificationStatusResp, err := client.CheckVerificationStatus(ctx, &tggateway.CheckVerificationStatusParams{
		RequestId: requestId,
		Code:      enteredCode,
	})
	if err != nil {
		log.Fatalf("check verification status error: %s", err.Error())
	}

	switch *checkVerificationStatusResp.VerificationStatus() {
	case tggateway.CODE_VALID:
		fmt.Println("code is valid")
	case tggateway.CODE_INVALID:
		fmt.Println("code is invalid")
	case tggateway.CODE_MAX_ATTEMPTS_EXCEEDED:
		fmt.Println("code max attempts exceeded")
	case tggateway.CODE_EXPIRED:
		fmt.Println("code is expired")
	}
}
```

## Todo list

- [ ] Add example for every method.
- [ ] Add support for custom loggers.
- [ ] Cover with tests.

## Contributing

No contributing guidelines yet.
