package utils

const (
	InvalidRequestBody = "invalid request body"
	ValidationFailed   = "validation failed"
	PasswordMismatch   = "password and confirm password do not match"

	MissingAuthorizationHeader = "missing authorization header"
	InvalidAuthorizationFormat = "invalid authorization header format"
	InvalidToken               = "invalid token"
	TokenRefreshedSuccessfully = "token refreshed successfully"

	UserCreatedSuccessfully  = "User created successfully"
	UserCreatedFailed        = "Failed to create user"
	UserUpdatedSuccessfully  = "User updated successfully"
	UserUpdatedFailed        = "Failed to update user"
	UserDeletedSuccessfully  = "User deleted successfully"
	UserDeletedFailed        = "Failed to delete user"
	UserFound                = "User found"
	UserNotFound             = "User not found"
	UserLoggedIn             = "User logged in successfully"
	UserLoggedOut            = "User logged out successfully"
	UserAuthenticationFailed = "User authentication failed"
)
