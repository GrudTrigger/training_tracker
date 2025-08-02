CREATE TABLE IF NOT EXISTS exercise_list (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR NOT NULL,
    category_muscle INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS exercise (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  training_id UUID REFERENCES training(id) ON DELETE CASCADE,
  exercise_list_id UUID REFERENCES exercise_list(id) ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS approach (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    exercise_id UUID REFERENCES exercise(id) ON DELETE CASCADE,
    repetition INT NOT NULL,
    weight INT NOT NULL
);

