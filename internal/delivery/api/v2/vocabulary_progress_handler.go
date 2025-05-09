package v2

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nguyenminhhoang/JapaneseCourses/internal/domain"
	"github.com/nguyenminhhoang/JapaneseCourses/internal/infrastructure/auth"
)

type VocabularyProgressHandler struct {
	vocabularyUseCase domain.VocabularyUseCase
	jwtService        *auth.JWTService
}

// NewVocabularyProgressHandler creates a new vocabulary progress handler
func NewVocabularyProgressHandler(vocabularyUseCase domain.VocabularyUseCase, jwtService *auth.JWTService) *VocabularyProgressHandler {
	return &VocabularyProgressHandler{
		vocabularyUseCase: vocabularyUseCase,
		jwtService:        jwtService,
	}
}

// Register registers the vocabulary progress routes
func (h *VocabularyProgressHandler) Register(g *echo.Group) {
	vocab := g.Group("/vocabulary-progress")
	{
		vocab.GET("/learned", h.GetLearnedVocabulary)
	}
}

// GetLearnedVocabulary returns all vocabulary items that the user has learned
func (h *VocabularyProgressHandler) GetLearnedVocabulary(c echo.Context) error {
	// Get user ID from JWT token
	tokenString := c.Request().Header.Get("Authorization")
	if tokenString == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Missing authorization token",
		})
	}

	// Remove "Bearer " prefix if present
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	claims, err := h.jwtService.ValidateToken(tokenString)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Invalid token",
		})
	}

	// Get learned vocabulary for the user
	vocabularies, err := h.vocabularyUseCase.GetLearnedVocabulary(claims.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, vocabularies)
}
