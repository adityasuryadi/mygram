CREATE TABLE "comment" (
    id UUID NOT NULL PRIMARY KEY,
    user_id UUID,
    password VARCHAR(255),
    photo_id UUID,
    message TEXT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);