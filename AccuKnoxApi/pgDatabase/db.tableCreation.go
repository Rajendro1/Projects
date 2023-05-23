package pgdatabase

var (
	CreatePlatformDatabaseQuery = `
    CREATE DATABASE accuknox;
    `
)

var CreateTableQuery = `

CREATE TABLE IF NOT EXISTS public.users
(
    id SERIAL PRIMARY KEY,
    name text DEFAULT NULL,
    email text DEFAULT NULL UNIQUE,
    password text DEFAULT NULL,
    session_id text DEFAULT NULL
);
CREATE TABLE IF NOT EXISTS public.notes
(
    id SERIAL PRIMARY KEY,
    session_id text DEFAULT NULL,
    user_id int DEFAULT NULL,
    note text DEFAULT NULL,
);
`
