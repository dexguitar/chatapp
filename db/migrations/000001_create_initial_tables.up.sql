CREATE TABLE IF NOT EXISTS "users" (
    "id" BIGSERIAL PRIMARY KEY,
    "username" VARCHAR NOT NULL,
    "email" VARCHAR NOT NULL,
    "password" VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS "messages" (
    "id" BIGSERIAL PRIMARY KEY,
    "timestamp" TIMESTAMP NOT NULL,
    "sender_id" VARCHAR NOT NULL,
    "receiver_id" VARCHAR NOT NULL,
    "content" TEXT NOT NULL
);