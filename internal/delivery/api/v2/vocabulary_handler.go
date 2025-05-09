package v2

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nguyenminhhoang/JapaneseCourses/internal/domain"
)

type VocabularyHandler struct {
	vocabularyUseCase domain.VocabularyUseCase
}

// NewVocabularyHandler creates a new vocabulary handler
func NewVocabularyHandler(vocabularyUseCase domain.VocabularyUseCase) *VocabularyHandler {
	return &VocabularyHandler{
		vocabularyUseCase: vocabularyUseCase,
	}
}

// Register registers the vocabulary routes
func (h *VocabularyHandler) Register(g *echo.Group) {
	vocabulary := g.Group("/vocabulary")
	{
		vocabulary.GET("", h.GetAll)
		vocabulary.GET("/:id", h.GetByID)
		vocabulary.GET("/lesson/:lesson_id", h.GetByLessonID)
		vocabulary.GET("/topic/:topic_id", h.GetByTopicID)
	}
}

func (h *VocabularyHandler) GetAll(c echo.Context) error {
	vocabularies, err := h.vocabularyUseCase.GetAllVocabulary()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, vocabularies)
}

func (h *VocabularyHandler) GetByID(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid ID format",
		})
	}

	vocabulary, err := h.vocabularyUseCase.GetVocabularyByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	if vocabulary == nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Vocabulary not found",
		})
	}

	return c.JSON(http.StatusOK, vocabulary)
}

func (h *VocabularyHandler) GetByLessonID(c echo.Context) error {
	lessonID, err := strconv.ParseInt(c.Param("lesson_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid lesson ID format",
		})
	}

	vocabularies, err := h.vocabularyUseCase.GetVocabularyByLessonID(lessonID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, vocabularies)
}

func (h *VocabularyHandler) GetByTopicID(c echo.Context) error {
	topicID, err := strconv.ParseInt(c.Param("topic_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid topic ID format",
		})
	}

	vocabularies, err := h.vocabularyUseCase.GetVocabularyByTopicID(topicID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, vocabularies)
}
