package handlers

import (
	"book-api/internal/models/entities"
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

func (h *authHandlers) Signup(c *fiber.Ctx) error {
	var req struct {
		FullName        string `json:"full_name" validate:"required"`
		Username        string `json:"username" validate:"required"`
		Password        string `json:"password" validate:"required gte=6 lte=20"`
		ConfirmPassword string `json:"confirm_password" validate:"required gte=6 lte=20"`
	}

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

	userRes, token, err := h.authService.Signup(c.Context(), user)
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusInternalServerError, utils.UserCreatedFailed, err.Error())
	}

	return utils.RespondWithSuccess(c, fiber.StatusCreated, utils.UserCreatedSuccessfully, fiber.Map{
		"user":  userRes,
		"token": token,
	})
}

func (h *authHandlers) Login(c *fiber.Ctx) error {
	var req struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required gte=6 lte=20"`
	}

	if err := c.BodyParser(&req); err != nil {
		return utils.RespondWithError(c, fiber.StatusBadRequest, utils.InvalidRequestBody, err.Error())
	}

	if err := h.validate.Struct(req); err != nil {
		return utils.RespondWithError(c, fiber.StatusBadRequest, utils.ValidationFailed, err.Error())
	}

	userRes, token, err := h.authService.Login(c.Context(), req.Username, req.Password)
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusUnauthorized, utils.UserCreatedFailed, err.Error())
	}

	return utils.RespondWithSuccess(c, fiber.StatusOK, utils.UserLoggedIn, fiber.Map{
		"user":  userRes,
		"token": token,
	})
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
