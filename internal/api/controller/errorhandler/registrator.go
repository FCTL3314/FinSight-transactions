package errorhandler

func RegisterAllErrorHandlers(errorHandler *ErrorHandler) {
	errorHandler.RegisterHandler(handleObjectNotFound)
	errorHandler.RegisterHandler(handleAccessDenied)
	errorHandler.RegisterHandler(handleInvalidParam)
	errorHandler.RegisterHandler(handlePaginationLimitExceeded)
	errorHandler.RegisterHandler(handleUniqueConstraint)
}
