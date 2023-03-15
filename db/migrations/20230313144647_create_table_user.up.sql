CREATE TABLE "user" (
    id UUID NOT NULL PRIMARY KEY,
    username VARCHAR (255),
    password VARCHAR(255),
    email VARCHAR (255) UNIQUE NOT NULL,
    name VARCHAR (255),
    age INT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);