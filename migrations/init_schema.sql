CREATE TYPE pr_status AS ENUM ('OPEN', 'MERGED');

CREATE TABLE IF NOT EXISTS teams (
     name text PRIMARY KEY,
     created_at  timestamp DEFAULT current_timestamp,
     updated_at timestamp DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS users (
    id text PRIMARY KEY,
    username text NOT NULL,
    team_name text NOT NULL REFERENCES teams(name) ON DELETE RESTRICT,
    is_active boolean NOT NULL DEFAULT true,
    created_at  timestamp DEFAULT current_timestamp,
    updated_at timestamp DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS pull_requests (
    id text PRIMARY KEY,
    name text NOT NULL,
    author_id text REFERENCES users(id) ON DELETE SET NULL,
    status pr_status NOT NULL,
    reviewer_ids text[] DEFAULT NULL,
    created_at  timestamp DEFAULT current_timestamp,
    updated_at timestamp DEFAULT NULL,
    merged_at  timestamp DEFAULT NULL
);

CREATE INDEX pull_requests_reviewer_index ON pull_requests USING GIN (reviewer_ids);
