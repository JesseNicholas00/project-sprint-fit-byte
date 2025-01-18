CREATE TABLE activities (
  activity_id UUID PRIMARY KEY,
  user_id UUID,
  activity_type VARCHAR(255) NOT NULL,
  done_at TIMESTAMP NOT NULL,
  calories_burned INTEGER NOT NULL,
  duration_in_minutes INTEGER NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS activities_activity_type_idx ON activities USING HASH (activity_type);
CREATE INDEX IF NOT EXISTS activities_done_at_idx ON activities(done_at);
CREATE INDEX IF NOT EXISTS activities_calories_burned_idx ON activities(calories_burned);