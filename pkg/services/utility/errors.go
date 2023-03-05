// FIXME: golangci-lint
// nolint:revive
package utility

// PayloadTypeAssertionError indicates the image name is not defined
type PayloadTypeAssertionError struct{}

func (e *PayloadTypeAssertionError) Error() string {
	return "Payload type assertion failed"
}
