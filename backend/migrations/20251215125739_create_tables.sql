-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Таблица тренировок
CREATE TABLE trainings (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  title VARCHAR(50) NOT NULL, -- Название тренировки
  date DATE NOT NULL, -- Дата проведения тренировки (2025-12-25)
  duration INTEGER NOT NULL, -- Продолжительность тренировки в секундах
  created_at TIMESTAMP DEFAULT now()
);

-- Таблица справочника упражнений
CREATE TABLE exercises (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  title TEXT NOT NULL, -- Название упражнения TODO: изменить потом на varchar
  muscle_group INTEGER NOT NULL CHECK(muscle_group BETWEEN 0 AND 10) -- На какую группу мышц упражнение
);

-- Таблица упражнений в конкретной тренировке
CREATE TABLE training_exercises (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  training_id UUID REFERENCES trainings(id) ON DELETE CASCADE, -- На какую тренировку ссылаемся
  exercise_id UUID REFERENCES exercises(id) -- На какое упражнение ссылаемся
);

-- Таблица подходов в упражнении, один подход - одна запись в таблице
CREATE TABLE exercise_sets (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  training_exercise_id UUID REFERENCES training_exercises(id) ON DELETE CASCADE, -- Упражнение в тренирвоке
  reps INT NOT NULL, -- Количество повторений
  weight NUMERIC(3,2) -- Вес
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE trainings,
DROP TABLE exercises,
DROP TABLE training_exercises,
DROP TABLE exercise_sets
-- +goose StatementEnd
