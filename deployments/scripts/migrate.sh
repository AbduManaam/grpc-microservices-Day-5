#!/bin/sh
echo "Running DB migrations..."

psql postgres://user:pass@postgres:5432/app <<EOF
CREATE TABLE IF NOT EXISTS users (
    id       SERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password TEXT        NOT NULL
);
EOF

echo "Migration complete"