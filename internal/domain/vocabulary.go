package domain

import "github.com/nguyenminhhoang/JapaneseCourses/internal/models"

// VocabularyRepository defines the interface for vocabulary data operations
type VocabularyRepository interface {
	GetAll() ([]*models.Vocabulary, error)
	GetByID(id int64) (*models.Vocabulary, error)
	GetByLessonID(lessonID int64) ([]*models.Vocabulary, error)
	GetByTopicID(topicID int64) ([]*models.Vocabulary, error)
	GetLearnedByUserID(userID int64) ([]*models.Vocabulary, error)
	Create(vocabulary *models.Vocabulary) error
	Update(vocabulary *models.Vocabulary) error
	Delete(id int64) error
}

// VocabularyUseCase defines the interface for vocabulary business logic
type VocabularyUseCase interface {
	GetAllVocabulary() ([]*models.Vocabulary, error)
	GetVocabularyByID(id int64) (*models.Vocabulary, error)
	GetVocabularyByLessonID(lessonID int64) ([]*models.Vocabulary, error)
	GetVocabularyByTopicID(topicID int64) ([]*models.Vocabulary, error)
	GetLearnedVocabulary(userID int64) ([]*models.Vocabulary, error)
	CreateVocabulary(vocabulary *models.Vocabulary) error
	UpdateVocabulary(vocabulary *models.Vocabulary) error
	DeleteVocabulary(id int64) error
}
