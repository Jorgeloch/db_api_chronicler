CREATE TABLE IF NOT EXISTS tag (
  id UUID UNIQUE NOT NULL,
  nome VARCHAR(45),
  cor VARCHAR(7),
  PRIMARY KEY(id)
);
