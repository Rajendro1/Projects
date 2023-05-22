package pgdatabase

var (
	CreatePlatformDatabaseQuery = `
    CREATE DATABASE platform;
    `
)

var CreateTableQuery = `

CREATE TABLE IF NOT EXISTS public.users
(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name text DEFAULT NULL,
    email text DEFAULT NULL UNIQUE,
    updated_by UUID DEFAULT NULL,
    created_at timestamp DEFAULT NULL,
    updated_at timestamp DEFAULT NULL
);
CREATE TABLE IF NOT EXISTS public.notes
(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id UUID DEFAULT NULL,
    note text DEFAULT NULL,
	created_by UUID DEFAULT NULL,
    updated_by UUID DEFAULT NULL,
    created_at timestamp DEFAULT NULL,
    updated_at timestamp DEFAULT NULL
);
`
