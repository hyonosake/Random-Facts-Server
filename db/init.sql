create extension if not exists pg_stat_statements;

-- Creation of facts table
CREATE TABLE IF NOT EXISTS facts (
    id BIGSERIAL UNIQUE NOT NULL PRIMARY KEY ,
    title text NOT NULL,
    description text NOT NULL,
    links text ARRAY
);

