CREATE TABLE IF NOT EXISTS exercise (
   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
   training_id UUID,
   title VARCHAR NOT NULL,
   muscle_group INT NOT NULL,
   approach_count INT NOT NULL,
   weight INT NOT NULL,
   created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),

   FOREIGN KEY(training_id) REFERENCES training (id) ON DELETE CASCADE
)