package common

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Error struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func NewError(statusCode int, message string) *Error {
	return &Error{
		StatusCode: statusCode,
		Message:    message,
	}
}

func (e Error) Error() string {
	return fmt.Sprintf("Status Code: %d, Message: %s", e.StatusCode, e.Message)
}

var (
	ErrInvalidParams        = NewError(fiber.StatusBadRequest, "invalid parameters")
	ErrMissingFields        = NewError(fiber.StatusBadRequest, "missing required fields")
	ErrInvalidRequestFormat = NewError(fiber.StatusBadRequest, "invalid request format")
	ErrInvalidJSON          = NewError(fiber.StatusBadRequest, "invalid JSON payload")
	ErrValidationFailed     = NewError(fiber.StatusUnprocessableEntity, "validation failed")
	ErrNotFound             = NewError(fiber.StatusNotFound, "resource not found")
	ErrUnauthorized         = NewError(fiber.StatusUnauthorized, "unauthorized access")
	ErrForbidden            = NewError(fiber.StatusForbidden, "access forbidden")
	ErrConflict             = NewError(fiber.StatusConflict, "conflict detected")
	ErrAlreadyExists        = NewError(fiber.StatusConflict, "resource already exists")
	ErrRateLimitExceeded    = NewError(fiber.StatusTooManyRequests, "rate limit exceeded")
	ErrRequestTimeout       = NewError(fiber.StatusRequestTimeout, "request timeout")
	ErrInternalServerError  = NewError(fiber.StatusInternalServerError, "internal server error")
	ErrServiceUnavailable   = NewError(fiber.StatusServiceUnavailable, "service temporarily unavailable")
	ErrGatewayTimeout       = NewError(fiber.StatusGatewayTimeout, "gateway timeout")
	ErrDatabaseError        = NewError(fiber.StatusInternalServerError, "database operation failed")
	ErrCacheMiss            = NewError(fiber.StatusNotFound, "cache miss")
	ErrDependencyFailed     = NewError(fiber.StatusFailedDependency, "failed dependency")
	ErrSessionExpired       = NewError(fiber.StatusUnauthorized, "session expired")
	ErrTokenInvalid         = NewError(fiber.StatusUnauthorized, "invalid token")
	ErrTokenExpired         = NewError(fiber.StatusUnauthorized, "token expired")
	ErrTokenRevoked         = NewError(fiber.StatusUnauthorized, "token has been revoked")
	ErrCSRFTokenMismatch    = NewError(fiber.StatusForbidden, "CSRF token mismatch")
	ErrMethodNotAllowed     = NewError(fiber.StatusMethodNotAllowed, "method not allowed")
	ErrPayloadTooLarge      = NewError(fiber.StatusRequestEntityTooLarge, "payload too large")
	ErrUnsupportedMediaType = NewError(fiber.StatusUnsupportedMediaType, "unsupported media type")
	ErrUnprocessableEntity  = NewError(fiber.StatusUnprocessableEntity, "unprocessable entity")
	ErrLocked               = NewError(fiber.StatusLocked, "resource is locked")
	ErrInsufficientFunds    = NewError(fiber.StatusPaymentRequired, "insufficient funds")
	ErrNotImplemented       = NewError(fiber.StatusNotImplemented, "not implemented")
)
