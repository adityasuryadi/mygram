CREATE TABLE "socialmedia" (
    id UUID NOT NULL PRIMARY KEY,
    name VARCHAR (255),
    social_media_url VARCHAR(255),
    user_id UUID,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);