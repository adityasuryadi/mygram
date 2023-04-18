CREATE TABLE "user_role" (
    id SERIAL PRIMARY KEY,
    user_id UUID,
    role_id INTEGER
);