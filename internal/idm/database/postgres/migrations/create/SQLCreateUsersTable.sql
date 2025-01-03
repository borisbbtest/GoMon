CREATE SCHEMA IF NOT EXISTS idm;
CREATE TABLE IF NOT EXISTS idm.users (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    login text NOT NULL,
    firstname text,
    lastname text,
    password text NOT NULL,
    source text,
    created_at timestamptz,
    CONSTRAINT users_pkey PRIMARY KEY (login)
);
CREATE TABLE IF NOT EXISTS idm.sessions (
    id uuid NOT NULL,
    config json,
    login text NOT NULL,
    created timestamptz NOT NULL,
    duration timestamptz,
    CONSTRAINT session_pkey PRIMARY KEY (id, login)
);