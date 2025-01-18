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