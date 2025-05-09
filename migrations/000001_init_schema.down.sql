-- Drop indexes
DROP INDEX IF EXISTS idx_vocabulary_lesson_id;
DROP INDEX IF EXISTS idx_vocabulary_topic_id;
DROP INDEX IF EXISTS idx_grammar_lesson_id;
DROP INDEX IF EXISTS idx_grammar_topic_id;
DROP INDEX IF EXISTS idx_lessons_topic_id;
DROP INDEX IF EXISTS idx_user_progress_user_id;
DROP INDEX IF EXISTS idx_user_progress_lesson_id;
DROP INDEX IF EXISTS idx_vocabulary_progress_user_id;
DROP INDEX IF EXISTS idx_vocabulary_progress_vocabulary_id;
DROP INDEX IF EXISTS idx_tests_lesson_id;
DROP INDEX IF EXISTS idx_tests_topic_id;
DROP INDEX IF EXISTS idx_test_questions_test_id;
DROP INDEX IF EXISTS idx_user_test_results_user_id;
DROP INDEX IF EXISTS idx_user_test_results_test_id;

-- Drop tables in reverse order of creation
DROP TABLE IF EXISTS user_test_results;
DROP TABLE IF EXISTS test_questions;
DROP TABLE IF EXISTS tests;
DROP TABLE IF EXISTS vocabulary_progress;
DROP TABLE IF EXISTS user_progress;
DROP TABLE IF EXISTS grammar;
DROP TABLE IF EXISTS vocabulary;
DROP TABLE IF EXISTS lessons;
DROP TABLE IF EXISTS topics;
DROP TABLE IF EXISTS users;

-- Drop extensions
DROP EXTENSION IF EXISTS "uuid-ossp"; 