-- Insert sample users
INSERT INTO users (username, password, email, full_name, role) VALUES
('admin', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'admin@example.com', 'Admin User', 'admin'),
('student1', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'student1@example.com', 'Student One', 'user'),
('student2', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'student2@example.com', 'Student Two', 'user');

-- Insert topics
INSERT INTO topics (name, description, level) VALUES
('Basic Greetings', 'Essential Japanese greetings and introductions', 'N5'),
('Daily Conversations', 'Common phrases used in everyday situations', 'N5'),
('Family Members', 'Vocabulary related to family relationships', 'N5'),
('Numbers and Counting', 'Japanese number system and counters', 'N5'),
('Basic Kanji', 'Fundamental kanji characters for beginners', 'N5');

-- Insert lessons
INSERT INTO lessons (topic_id, title, content, order_number) VALUES
(1, 'Common Greetings', 'Learn the most common Japanese greetings used throughout the day.', 1),
(1, 'Self Introduction', 'Learn how to introduce yourself in Japanese.', 2),
(2, 'Ordering at Restaurant', 'Essential phrases for ordering food and drinks.', 1),
(2, 'Shopping Phrases', 'Useful expressions for shopping in Japan.', 2),
(3, 'Family Vocabulary', 'Basic vocabulary for family members.', 1);

-- Insert vocabulary
INSERT INTO vocabulary (word, meaning, kanji, example_sentence, lesson_id, topic_id) VALUES
('おはようございます', 'Good morning (polite)', NULL, 'おはようございます、先生。', 1, 1),
('こんにちは', 'Good afternoon', NULL, 'こんにちは、元気ですか。', 1, 1),
('さようなら', 'Goodbye', NULL, 'さようなら、また会いましょう。', 1, 1),
('はじめまして', 'Nice to meet you', NULL, 'はじめまして、田中です。', 2, 1),
('お願いします', 'Please', NULL, 'コーヒーを一つお願いします。', 3, 2),
('いくらですか', 'How much is it?', NULL, 'このりんごはいくらですか。', 4, 2),
('父', 'Father', '父', '父は医者です。', 5, 3),
('母', 'Mother', '母', '母は先生です。', 5, 3);

-- Insert grammar points
INSERT INTO grammar (title, explanation, examples, lesson_id, topic_id) VALUES
('です/だ', 'Basic copula for describing things', ARRAY['これは本です。', 'それは私の本だ。'], 2, 1),
('ます form', 'Polite form of verbs', ARRAY['食べます', '飲みます', '行きます'], 1, 1),
('の particle', 'Possessive particle', ARRAY['私の本', '友達の車'], 2, 1),
('を particle', 'Object marker', ARRAY['本を読みます', '水を飲みます'], 3, 2);

-- Insert tests
INSERT INTO tests (title, description, lesson_id, topic_id, time_limit, passing_score) VALUES
('Greetings Test', 'Test your knowledge of Japanese greetings', 1, 1, 15, 70),
('Self Introduction Test', 'Practice introducing yourself', 2, 1, 20, 70),
('Restaurant Phrases', 'Test your restaurant vocabulary', 3, 2, 15, 70);

-- Insert test questions
INSERT INTO test_questions (test_id, question_text, question_type, options, correct_answer, points) VALUES
(1, 'How do you say "Good morning" in polite Japanese?', 'multiple_choice', 
    '{"a": "おはようございます", "b": "こんにちは", "c": "こんばんは", "d": "さようなら"}', 
    'a', 1),
(1, 'What is the appropriate greeting for afternoon?', 'multiple_choice', 
    '{"a": "おはようございます", "b": "こんにちは", "c": "こんばんは", "d": "さようなら"}', 
    'b', 1),
(2, 'Complete the sentence: "わたしの なまえは ___ です"', 'fill_blank', 
    NULL, 
    '[your name]', 1),
(3, 'How do you say "Please give me one coffee"?', 'multiple_choice', 
    '{"a": "コーヒーを一つお願いします", "b": "コーヒーをください", "c": "コーヒーが好きです", "d": "コーヒーを飲みます"}', 
    'a', 1);

-- Insert some user progress
INSERT INTO user_progress (user_id, lesson_id, completed, score) VALUES
(2, 1, true, 90),
(2, 2, true, 85),
(2, 3, false, NULL),
(3, 1, true, 95);

-- Insert vocabulary progress
INSERT INTO vocabulary_progress (user_id, vocabulary_id, mastery_level, last_reviewed_at, next_review_at) VALUES
(2, 1, 3, CURRENT_TIMESTAMP - INTERVAL '1 day', CURRENT_TIMESTAMP + INTERVAL '3 days'),
(2, 2, 2, CURRENT_TIMESTAMP - INTERVAL '2 days', CURRENT_TIMESTAMP + INTERVAL '2 days'),
(3, 1, 4, CURRENT_TIMESTAMP - INTERVAL '1 day', CURRENT_TIMESTAMP + INTERVAL '4 days');

-- Insert test results
INSERT INTO user_test_results (user_id, test_id, score, completed_at) VALUES
(2, 1, 90, CURRENT_TIMESTAMP - INTERVAL '5 days'),
(2, 2, 85, CURRENT_TIMESTAMP - INTERVAL '3 days'),
(3, 1, 95, CURRENT_TIMESTAMP - INTERVAL '4 days'); 