package usecase

import (
	"github.com/nguyenminhhoang/JapaneseCourses/internal/domain"
	"github.com/nguyenminhhoang/JapaneseCourses/internal/models"
)

type vocabularyUseCase struct {
	vocabularyRepo domain.VocabularyRepository
}

// NewVocabularyUseCase creates a new instance of vocabularyUseCase
func NewVocabularyUseCase(vocabularyRepo domain.VocabularyRepository) domain.VocabularyUseCase {
	return &vocabularyUseCase{
		vocabularyRepo: vocabularyRepo,
	}
}

func (uc *vocabularyUseCase) GetAllVocabulary() ([]*models.Vocabulary, error) {
	return uc.vocabularyRepo.GetAll()
}

func (uc *vocabularyUseCase) GetVocabularyByID(id int64) (*models.Vocabulary, error) {
	return uc.vocabularyRepo.GetByID(id)
}

func (uc *vocabularyUseCase) GetVocabularyByLessonID(lessonID int64) ([]*models.Vocabulary, error) {
	return uc.vocabularyRepo.GetByLessonID(lessonID)
}

func (uc *vocabularyUseCase) GetVocabularyByTopicID(topicID int64) ([]*models.Vocabulary, error) {
	return uc.vocabularyRepo.GetByTopicID(topicID)
}

func (uc *vocabularyUseCase) GetLearnedVocabulary(userID int64) ([]*models.Vocabulary, error) {
	return uc.vocabularyRepo.GetLearnedByUserID(userID)
}

func (uc *vocabularyUseCase) CreateVocabulary(vocabulary *models.Vocabulary) error {
	return uc.vocabularyRepo.Create(vocabulary)
}

func (uc *vocabularyUseCase) UpdateVocabulary(vocabulary *models.Vocabulary) error {
	return uc.vocabularyRepo.Update(vocabulary)
}

func (uc *vocabularyUseCase) DeleteVocabulary(id int64) error {
	return uc.vocabularyRepo.Delete(id)
}
