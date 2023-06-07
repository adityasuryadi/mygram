CREATE TABLE "user_token" (
    id UUID NOT NULL PRIMARY KEY,
    user_id UUID,
    token text,
    expired_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

