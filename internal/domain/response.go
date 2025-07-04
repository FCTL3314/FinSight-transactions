package domain

type PaginatedResponse[T any] struct {
	Count   int64 `json:"count"`
	Limit   int   `json:"limit"`
	Offset  int   `json:"offset"`
	Results []T   `json:"results"`
}

type ErrorResponse struct {
	Detail   string `json:"detail"`
	Codename string `json:"codename"`
}

func NewErrorResponse(detail string, codename string) *ErrorResponse {
	return &ErrorResponse{detail, codename}
}

func NewPaginationErrorResponse(detail string) *ErrorResponse {
	return NewErrorResponse(
		detail,
		"pagination_error",
	)
}

func NewValidationErrorResponse(detail string) *ErrorResponse {
	return NewErrorResponse(
		detail,
		"validation_error",
	)
}

func NewUniqueConstraintErrorResponse(detail string) *ErrorResponse {
	return NewErrorResponse(
		detail,
		"unique_constraint_error",
	)
}

func NewInvalidURLParamResponse(detail string) *ErrorResponse {
	return NewErrorResponse(
		detail,
		"invalid_url_param",
	)
}

func NewUnauthorizedErrorResponse() *ErrorResponse {
	return NewErrorResponse(
		"Authentication required - X-User-ID header is missing",
		"unauthorized",
	)
}

var (
	InternalServerErrorResponse = NewErrorResponse(
		"Internal Server Error",
		"internal_server_error",
	)
	NotFoundResponse = NewErrorResponse(
		"Not found",
		"not_found",
	)
	ForbiddenResponse = NewErrorResponse(
		"Forbidden",
		"forbidden",
	)
)
