-- Active: 1715157815254@@127.0.0.1@5432@6semsmall

DROP TABLE IF EXISTS "metric" CASCADE;
DROP TABLE IF EXISTS "audio" CASCADE;
DROP TABLE IF EXISTS "genre" CASCADE;
DROP TABLE IF EXISTS "author" CASCADE;
DROP TABLE IF EXISTS "external resource" CASCADE;
DROP TABLE IF EXISTS "snippet" CASCADE;
DROP TABLE IF EXISTS "album" CASCADE;
DROP TABLE IF EXISTS "user" CASCADE;
----------------------------------------------
DROP TABLE IF EXISTS "album_genre" CASCADE;
DROP TABLE IF EXISTS "album_audio" CASCADE;
DROP TABLE IF EXISTS "author_album" CASCADE;
DROP TABLE IF EXISTS "author_audio" CASCADE;
DROP TABLE IF EXISTS "audio_genre" CASCADE; 
DROP TABLE IF EXISTS "user_album" CASCADE;
DROP TABLE IF EXISTS "user_audio" CASCADE;

CREATE TABLE "metric" (
    id SERIAL PRIMARY KEY,
    "views" INTEGER NOT NULL,
    "likes" INTEGER NOT NULL,
    "reposts" INTEGER NOT NULL,  
    "retention" NUMERIC NOT NULL,
    "downloads" INTEGER NOT NULL,
    "year-popularity" NUMERIC NOT NULL 
);

-- гавнище сначала метрику, потом аудио
CREATE TABLE "audio" (
    id SERIAL PRIMARY KEY,
    "appellation" VARCHAR(1000) NOT NULL,
    "lyric" VARCHAR(10000),
    "release data" INTEGER NOT NULL,
    "metric id" INTEGER NOT NULL,
    FOREIGN KEY ("metric id") REFERENCES "metric"(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "genre" (
    id SERIAL PRIMARY KEY,
    "appellation" VARCHAR(100) UNIQUE NOT NULL,
    "description" VARCHAR(1000),
    "metric id" INTEGER NOT NULL,
    FOREIGN KEY ("metric id") REFERENCES "metric"(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "author" (
    id SERIAL PRIMARY KEY,
    "appellation" VARCHAR(50) UNIQUE NOT NULL,
    "description" VARCHAR(1000),
    "metric id" INTEGER NOT NULL,
    FOREIGN KEY ("metric id") REFERENCES "metric"(id) ON DELETE CASCADE ON UPDATE CASCADE 
);

CREATE TABLE "external resource" (
    id SERIAL PRIMARY KEY,
    "link" VARCHAR(100) NOT NULL,
    "type" VARCHAR(100) NOT NULL,
    "author id" INTEGER NOT NULL,
    FOREIGN KEY ("author id") REFERENCES "author"(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "snippet" (
    id SERIAL PRIMARY KEY,
    "start" INTEGER NOT NULL,
    "end" INTEGER NOT NULL,
    "audio id" INTEGER NOT NULL,
    FOREIGN KEY ("audio id") REFERENCES "audio"(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "album" (
    id SERIAL PRIMARY KEY,
    "appellation" VARCHAR(100) NOT NULL,
    "release data" INTEGER NOT NULL,
    "metric id" INTEGER NOT NULL,
    FOREIGN KEY ("metric id") REFERENCES "metric"(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "user" (
    id SERIAL PRIMARY KEY,
    "login" VARCHAR(100) UNIQUE NOT NULL,
    "email" VARCHAR(50) UNIQUE NOT NULL,
    "password" VARCHAR(50) NOT NULL,
    "role" INT NOT NULL,
    "author id" INTEGER,
    FOREIGN KEY ("author id") REFERENCES "author"(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- многие к многим баля
-- альбом: аудио, жанр

CREATE TABLE "album_genre" (
    id SERIAL PRIMARY KEY,
    "album id" INTEGER NOT NULL,
    "genre id" INTEGER NOT NULL,
    FOREIGN KEY ("album id") REFERENCES "album"(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY ("genre id") REFERENCES "genre"(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "album_audio" (
    id SERIAL PRIMARY KEY,
    "album id" INTEGER NOT NULL,
    "audio id" INTEGER NOT NULL,
    FOREIGN KEY ("album id") REFERENCES "album"(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY ("audio id") REFERENCES "audio"(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- автор: aльбом, аудио

CREATE TABLE "author_album" (
    id SERIAL PRIMARY KEY,
    "author id" INTEGER NOT NULL,
    "album id" INTEGER NOT NULL,
    FOREIGN KEY ("author id") REFERENCES "author"(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY ("album id") REFERENCES "album"(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "author_audio" (
    id SERIAL PRIMARY KEY,
    "author id" INTEGER NOT NULL,
    "audio id" INTEGER NOT NULL,
    FOREIGN KEY ("author id") REFERENCES "author"(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY ("audio id") REFERENCES "audio"(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- аудио: жанр

CREATE TABLE "audio_genre" (
    id SERIAL PRIMARY KEY,
    "audio id" INTEGER NOT NULL,
    "genre id" INTEGER NOT NULL,
    FOREIGN KEY ("audio id") REFERENCES "audio"(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY ("genre id") REFERENCES "genre"(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- пользователь: альбом, аудио

CREATE TABLE "user_album" (
    id SERIAL PRIMARY KEY,
    "user id" INTEGER NOT NULL,
    "album id" INTEGER NOT NULL,
    FOREIGN KEY ("user id") REFERENCES "user"(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY ("album id") REFERENCES "album"(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "user_audio" (
    id SERIAL PRIMARY KEY,
    "user id" INTEGER NOT NULL,
    "audio id" INTEGER NOT NULL,
    FOREIGN KEY ("user id") REFERENCES "user"(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY ("audio id") REFERENCES "audio"(id) ON DELETE CASCADE ON UPDATE CASCADE
); 