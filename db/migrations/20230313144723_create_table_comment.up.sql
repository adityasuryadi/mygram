CREATE TABLE "comment" (
    id UUID NOT NULL PRIMARY KEY,
    user_id UUID,
    photo_id UUID,
    message TEXT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);