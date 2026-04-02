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
	ErrInvalidParams        = NewError(fiber.StatusBadRequest, "invalid parameters") // 400
	ErrMissingFields        = NewError(fiber.StatusBadRequest, "missing required fields") // 400
	ErrInvalidRequestFormat = NewError(fiber.StatusBadRequest, "invalid request format") // 400
	ErrInvalidJSON          = NewError(fiber.StatusBadRequest, "invalid JSON payload") // 400
	ErrValidationFailed     = NewError(fiber.StatusUnprocessableEntity, "validation failed") // 422
	ErrNotFound             = NewError(fiber.StatusNotFound, "resource not found") // 404
	ErrUnauthorized         = NewError(fiber.StatusUnauthorized, "unauthorized access") // 401
	ErrForbidden            = NewError(fiber.StatusForbidden, "access forbidden") // 403
	ErrConflict             = NewError(fiber.StatusConflict, "conflict detected") // 409
	ErrAlreadyExists        = NewError(fiber.StatusConflict, "resource already exists") // 409
	ErrRateLimitExceeded    = NewError(fiber.StatusTooManyRequests, "rate limit exceeded") // 429
	ErrRequestTimeout       = NewError(fiber.StatusRequestTimeout, "request timeout") // 408
	ErrInternalServerError  = NewError(fiber.StatusInternalServerError, "internal server error") // 500
	ErrServiceUnavailable   = NewError(fiber.StatusServiceUnavailable, "service temporarily unavailable") // 503
	ErrGatewayTimeout       = NewError(fiber.StatusGatewayTimeout, "gateway timeout") // 504
	ErrDatabaseError        = NewError(fiber.StatusInternalServerError, "database operation failed") // 500
	ErrCacheMiss            = NewError(fiber.StatusNotFound, "cache miss") // 404
	ErrDependencyFailed     = NewError(fiber.StatusFailedDependency, "failed dependency") // 424
	ErrSessionExpired       = NewError(fiber.StatusUnauthorized, "session expired") // 401
	ErrTokenInvalid         = NewError(fiber.StatusUnauthorized, "invalid token") // 401
	ErrTokenExpired         = NewError(fiber.StatusUnauthorized, "token expired") // 401
	ErrTokenRevoked         = NewError(fiber.StatusUnauthorized, "token has been revoked") // 401
	ErrCSRFTokenMismatch    = NewError(fiber.StatusForbidden, "CSRF token mismatch") // 403
	ErrMethodNotAllowed     = NewError(fiber.StatusMethodNotAllowed, "method not allowed") // 405
	ErrPayloadTooLarge      = NewError(fiber.StatusRequestEntityTooLarge, "payload too large") // 413
	ErrUnsupportedMediaType = NewError(fiber.StatusUnsupportedMediaType, "unsupported media type") // 415
	ErrUnprocessableEntity  = NewError(fiber.StatusUnprocessableEntity, "unprocessable entity") // 422
	ErrLocked               = NewError(fiber.StatusLocked, "resource is locked") // 423
	ErrInsufficientFunds    = NewError(fiber.StatusPaymentRequired, "insufficient funds") // 402
	ErrNotImplemented       = NewError(fiber.StatusNotImplemented, "not implemented") // 501
)
