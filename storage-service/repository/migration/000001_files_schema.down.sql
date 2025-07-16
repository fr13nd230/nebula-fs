CREATE TABLE files (
  id UUID PRIMARY KEY,
  user_id UUID NOT NULL,
  filename TEXT NOT NULL,
  mime_type TEXT,
  total_size BIGINT DEFAULT 0,
  chunk_count INT,
  created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE file_chunks (
  id SERIAL PRIMARY KEY,
  file_id UUID REFERENCES files(id) ON DELETE CASCADE,
  chunk_index INT NOT NULL,
  cid TEXT NOT NULL,
  chunk_size INT DEFAULT 0,
  uploaded_at TIMESTAMP DEFAULT now(),
  UNIQUE(file_id, chunk_index)
);