CREATE TABLE review_api(
user_id SERIAL PRIMARY KEY,
user_name TEXT NOT NULL UNIQUE,
email TEXT NOT NULL,
messages TEXT NOT NULL,
create_at TIMESTAMP DEFAULT now(),
updated_at TIMESTAMP DEFAULT now(),
delete_at TIMESTAMP DEFAULT now()
);