# Telegram Gateway API Client for Go

This Go package provides a full-featured client for interacting with the [Telegram Gateway API](https://core.telegram.org/gateway).

**Key features:**

- Full support of Telegram Gateway API
- Simple API
- 100% documented
- Context package support
- Type-safe error handling
- No external dependencies


## Installation

```bash
go get github.com/skewb1k/tg-gateway-go
```

## Usage

Basic example of sending and verifying code:

```go
package main

import (
	"context"
	"errors"
	"log"
	"os"

	tggateway "github.com/skewb1k/tg-gateway-go"
)

func main() {
	token := os.Getenv("TGGW_API_TOKEN")

	client := tggateway.NewClient(token)

	ctx := context.Background()

	sendVerificationMessageResp, err := client.SendVerificationMessage(ctx, &tggateway.SendVerificationMessageParams{
		PhoneNumber: "11111111111",
		CodeLength:  4,
		// or set your generated code
		// Code: "1234",
	})
	if err != nil {
		// check error type if necessary
		if errors.Is(err, tggateway.ErrPhoneNumberInvalid) {
			// ...handle error
		}

		log.Fatalf("failed to send verification message: %s", err.Error())
	}

	// store recieved request id
	requestId := sendVerificationMessageResp.RequestID

	// get code from some user input.
	enteredCode := userInput.Code

	checkVerificationStatusResp, err := client.CheckVerificationStatus(ctx, &tggateway.CheckVerificationStatusParams{
		RequestID: requestId,
		Code:      enteredCode,
	})
	if err != nil {
		log.Fatalf("failed to check verification status error: %s", err.Error())
	}

	if checkVerificationStatusResp.VerificationStatus.Status.IsValid() {
		// ...grant user accesses
	}
}
```


## Docs

For detailed documentation, method signatures, and examples, visit the package page on [pkg.go.dev](https://pkg.go.dev/github.com/skewb1k/tg-gateway-go).


## Supported go versions

This library requires Go 1.18+.


## Contributing

Feel free to submit issues, fork the repository and send pull requests!


## License

This project is licensed under the MIT License.
