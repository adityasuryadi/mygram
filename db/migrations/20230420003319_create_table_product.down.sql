CREATE TABLE "product" (
    id UUID NOT NULL PRIMARY KEY,
    name VARCHAR (255),
    stock INTEGER,
    user_id UUID,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);