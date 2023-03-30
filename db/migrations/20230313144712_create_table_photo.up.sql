CREATE TABLE "photo" (
    id UUID NOT NULL PRIMARY KEY,
    title VARCHAR (255),
    caption VARCHAR(255),
    photo_url VARCHAR (255) NOT NULL,
    user_id UUID,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);