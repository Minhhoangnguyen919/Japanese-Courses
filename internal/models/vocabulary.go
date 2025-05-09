package models

import "time"

// Vocabulary represents a Japanese vocabulary item
type Vocabulary struct {
	ID              int64     `json:"id"`
	Word            string    `json:"word"`
	Meaning         string    `json:"meaning"`
	Kanji           string    `json:"kanji,omitempty"`
	ExampleSentence string    `json:"example_sentence,omitempty"`
	LessonID        int64     `json:"lesson_id,omitempty"`
	TopicID         int64     `json:"topic_id,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
