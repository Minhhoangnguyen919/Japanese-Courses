package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/nguyenminhhoang/JapaneseCourses/internal/domain"
	"github.com/nguyenminhhoang/JapaneseCourses/internal/models"
)

type vocabularyRepository struct {
	db *pgx.Conn
}

// NewVocabularyRepository creates a new instance of vocabularyRepository
func NewVocabularyRepository(db *pgx.Conn) domain.VocabularyRepository {
	return &vocabularyRepository{db: db}
}

func (r *vocabularyRepository) GetAll() ([]*models.Vocabulary, error) {
	query := `
		SELECT id, word, meaning, kanji, example_sentence, lesson_id, topic_id, created_at, updated_at
		FROM vocabulary
		ORDER BY id
	`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vocabularies []*models.Vocabulary
	for rows.Next() {
		var v models.Vocabulary
		err := rows.Scan(
			&v.ID,
			&v.Word,
			&v.Meaning,
			&v.Kanji,
			&v.ExampleSentence,
			&v.LessonID,
			&v.TopicID,
			&v.CreatedAt,
			&v.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		vocabularies = append(vocabularies, &v)
	}

	return vocabularies, nil
}

func (r *vocabularyRepository) GetByID(id int64) (*models.Vocabulary, error) {
	query := `
		SELECT id, word, meaning, kanji, example_sentence, lesson_id, topic_id, created_at, updated_at
		FROM vocabulary
		WHERE id = $1
	`

	var v models.Vocabulary
	err := r.db.QueryRow(context.Background(), query, id).Scan(
		&v.ID,
		&v.Word,
		&v.Meaning,
		&v.Kanji,
		&v.ExampleSentence,
		&v.LessonID,
		&v.TopicID,
		&v.CreatedAt,
		&v.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &v, nil
}

func (r *vocabularyRepository) GetByLessonID(lessonID int64) ([]*models.Vocabulary, error) {
	query := `
		SELECT id, word, meaning, kanji, example_sentence, lesson_id, topic_id, created_at, updated_at
		FROM vocabulary
		WHERE lesson_id = $1
		ORDER BY id
	`

	rows, err := r.db.Query(context.Background(), query, lessonID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vocabularies []*models.Vocabulary
	for rows.Next() {
		var v models.Vocabulary
		err := rows.Scan(
			&v.ID,
			&v.Word,
			&v.Meaning,
			&v.Kanji,
			&v.ExampleSentence,
			&v.LessonID,
			&v.TopicID,
			&v.CreatedAt,
			&v.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		vocabularies = append(vocabularies, &v)
	}

	return vocabularies, nil
}

func (r *vocabularyRepository) GetByTopicID(topicID int64) ([]*models.Vocabulary, error) {
	query := `
		SELECT id, word, meaning, kanji, example_sentence, lesson_id, topic_id, created_at, updated_at
		FROM vocabulary
		WHERE topic_id = $1
		ORDER BY id
	`

	rows, err := r.db.Query(context.Background(), query, topicID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vocabularies []*models.Vocabulary
	for rows.Next() {
		var v models.Vocabulary
		err := rows.Scan(
			&v.ID,
			&v.Word,
			&v.Meaning,
			&v.Kanji,
			&v.ExampleSentence,
			&v.LessonID,
			&v.TopicID,
			&v.CreatedAt,
			&v.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		vocabularies = append(vocabularies, &v)
	}

	return vocabularies, nil
}

func (r *vocabularyRepository) Create(vocabulary *models.Vocabulary) error {
	query := `
		INSERT INTO vocabulary (word, meaning, kanji, example_sentence, lesson_id, topic_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`

	now := time.Now()
	vocabulary.CreatedAt = now
	vocabulary.UpdatedAt = now

	return r.db.QueryRow(
		context.Background(),
		query,
		vocabulary.Word,
		vocabulary.Meaning,
		vocabulary.Kanji,
		vocabulary.ExampleSentence,
		vocabulary.LessonID,
		vocabulary.TopicID,
		vocabulary.CreatedAt,
		vocabulary.UpdatedAt,
	).Scan(&vocabulary.ID)
}

func (r *vocabularyRepository) Update(vocabulary *models.Vocabulary) error {
	query := `
		UPDATE vocabulary 
		SET word = $1, meaning = $2, kanji = $3, example_sentence = $4, 
			lesson_id = $5, topic_id = $6, updated_at = $7
		WHERE id = $8
	`

	vocabulary.UpdatedAt = time.Now()

	_, err := r.db.Exec(
		context.Background(),
		query,
		vocabulary.Word,
		vocabulary.Meaning,
		vocabulary.Kanji,
		vocabulary.ExampleSentence,
		vocabulary.LessonID,
		vocabulary.TopicID,
		vocabulary.UpdatedAt,
		vocabulary.ID,
	)

	return err
}

func (r *vocabularyRepository) Delete(id int64) error {
	query := `DELETE FROM vocabulary WHERE id = $1`

	_, err := r.db.Exec(context.Background(), query, id)
	return err
}

func (r *vocabularyRepository) GetLearnedByUserID(userID int64) ([]*models.Vocabulary, error) {
	query := `
		SELECT v.id, v.word, v.meaning, v.kanji, v.example_sentence, v.lesson_id, v.topic_id, v.created_at, v.updated_at
		FROM vocabulary v
		INNER JOIN vocabulary_progress vp ON v.id = vp.vocabulary_id
		WHERE vp.user_id = $1
		ORDER BY vp.mastery_level DESC, v.id
	`

	rows, err := r.db.Query(context.Background(), query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vocabularies []*models.Vocabulary
	for rows.Next() {
		var v models.Vocabulary
		err := rows.Scan(
			&v.ID,
			&v.Word,
			&v.Meaning,
			&v.Kanji,
			&v.ExampleSentence,
			&v.LessonID,
			&v.TopicID,
			&v.CreatedAt,
			&v.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		vocabularies = append(vocabularies, &v)
	}

	return vocabularies, nil
}
