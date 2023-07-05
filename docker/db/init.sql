CREATE TABLE users (
    user_id serial PRIMARY KEY,
	nickname VARCHAR ( 255 ) UNIQUE NOT NULL,
	email VARCHAR ( 255 ) UNIQUE NOT NULL,
    password VARCHAR ( 255 ) NOT NULL
);

CREATE EXTENSION IF NOT EXISTS pgcrypto;

INSERT INTO users(nickname, email, password)
VALUES ('test', 'test@gmail.com', crypt('secretpassword', gen_salt('bf')));

GRANT ALL PRIVILEGES ON DATABASE quizzes TO aizek;