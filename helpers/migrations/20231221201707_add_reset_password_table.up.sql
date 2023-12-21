CREATE TABLE IF NOT EXISTS reset_pass (
    "id" SERIAL PRIMARY KEY,
    "reset_code" TEXT NOT NULL UNIQUE,
    "user_id" INTEGER REFERENCES users (id),
    "used" BOOLEAN DEFAULT FALSE
)
