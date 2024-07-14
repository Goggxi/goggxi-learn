package handlers

import (
	"book-api/internal/models/entities"
	"book-api/internal/models/requests"
	"book-api/internal/services"
	"book-api/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type AuthHandlers interface {
	Signup(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
	RefreshToken(c *fiber.Ctx) error
	GetCurrentUser(c *fiber.Ctx) error
}

type authHandlers struct {
	authService services.AuthService
	validate    *validator.Validate
}

// Signup godoc
// @Summary Register a new user
// @Description Register a new user with the provided information
// @Tags auth-public
// @Accept json
// @Produce json
// @Param user body requests.SignupRequest true "User SignUp"
// @Success 201 {object} responses.UserTokenResponse
// @Failure 400 {object} utils.SwaggerErrorResponse
// @Failure 500 {object} utils.SwaggerErrorResponse
// @Router /auth/signup [post]
func (h *authHandlers) Signup(c *fiber.Ctx) error {
	var req = requests.SignupRequest{}
	if err := c.BodyParser(&req); err != nil {
		return utils.RespondWithError(c, fiber.StatusBadRequest, utils.InvalidRequestBody, err.Error())
	}

	if err := h.validate.Struct(req); err != nil {
		return utils.RespondWithError(c, fiber.StatusBadRequest, utils.ValidationFailed, err.Error())
	}

	if req.Password != req.ConfirmPassword {
		return utils.RespondWithError(c, fiber.StatusBadRequest, utils.PasswordMismatch, nil)
	}

	user := &entities.User{
		FullName: req.FullName,
		Username: req.Username,
		Password: req.Password,
	}

	userRes, err := h.authService.Signup(c.Context(), user)
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusInternalServerError, utils.UserCreatedFailed, err.Error())
	}

	return utils.RespondWithSuccess(c, fiber.StatusCreated, utils.UserCreatedSuccessfully, userRes)
}

// Login godoc
// @Summary Login a user
// @Description Login a user with the provided credentials
// @Tags auth-public
// @Accept json
// @Produce json
// @Param user body requests.LoginRequest true "User Login"
// @Success 200 {object} responses.UserTokenResponse
// @Failure 400 {object} utils.SwaggerErrorResponse
// @Failure 401 {object} utils.SwaggerErrorResponse
// @Router /auth/login [post]
func (h *authHandlers) Login(c *fiber.Ctx) error {
	var req = requests.LoginRequest{}

	if err := c.BodyParser(&req); err != nil {
		return utils.RespondWithError(c, fiber.StatusBadRequest, utils.InvalidRequestBody, err.Error())
	}

	if err := h.validate.Struct(req); err != nil {
		return utils.RespondWithError(c, fiber.StatusBadRequest, utils.ValidationFailed, err.Error())
	}

	userRes, err := h.authService.Login(c.Context(), req.Username, req.Password)
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusUnauthorized, utils.UserAuthenticationFailed, err.Error())
	}

	return utils.RespondWithSuccess(c, fiber.StatusOK, utils.UserLoggedIn, userRes)
}

func (h *authHandlers) Logout(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (h *authHandlers) RefreshToken(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return utils.RespondWithError(c, fiber.StatusUnauthorized, utils.MissingAuthorizationHeader, nil)
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
		return utils.RespondWithError(c, fiber.StatusBadRequest, utils.InvalidAuthorizationFormat, nil)
	}

	oldToken := tokenParts[1]

	newToken, err := h.authService.RefreshToken(oldToken)
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusInternalServerError, utils.InvalidToken, err.Error())
	}

	return utils.RespondWithSuccess(c, fiber.StatusOK, utils.TokenRefreshedSuccessfully, fiber.Map{
		"token": newToken,
	})
}

func (h *authHandlers) GetCurrentUser(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	userRes, err := h.authService.GetCurrentUser(c.Context(), userID)
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusNotFound, utils.UserNotFound, err.Error())
	}

	return utils.RespondWithSuccess(c, fiber.StatusOK, utils.UserFound, fiber.Map{
		"user": userRes,
	})
}

func NewAuthHandlers(authService services.AuthService) AuthHandlers {
	return &authHandlers{
		authService: authService,
		validate:    validator.New(),
	}
}
