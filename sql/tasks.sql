-- Active: 1715273381048@@127.0.0.1@5432@6sem
-- ●	Выбрать все песни в библиотеке пользователя с указанием автора и жанра
WITH 
choosenUser as (
SELECT "user".id as id FROM "user"
JOIN user_audio
ON "user".id = user_audio."user id"
GROUP BY "user".id
ORDER BY count(*) DESC
LIMIT 1
),
choosenUserAudio as (
    SELECT audio.id as id, audio.appellation as appellation FROM user_audio
    JOIN choosenUser
    ON user_audio."user id" = choosenUser.id
    JOIN audio
    ON audio.id = user_audio."audio id"
)
SELECT author.appellation as "author", choosenUserAudio.appellation as "audio", genre.appellation as "genre" FROM choosenUserAudio
JOIN audio_genre
ON audio_genre."audio id" = choosenUserAudio.id
JOIN genre
ON genre.id = audio_genre."genre id"
JOIN author_audio
ON author_audio."audio id" = choosenUserAudio.id
JOIN author
ON author.id = author_audio."author id"

-- ●	Выбрать для пользователя наиболее подходящий по жанру 
-- плейлист от других пользователей

WITH 
choosenUser as (
SELECT "user".id as id FROM "user"
JOIN user_audio
ON "user".id = user_audio."user id"
GROUP BY "user".id
ORDER BY count(*) DESC
LIMIT 1
),
choosenUserTopGenre as (
    SELECT audio_genre."genre id" as "genre id" FROM "user"
    JOIN user_audio
    ON user_audio."user id" = "user".id
    JOIN audio_genre
    ON audio_genre."audio id" = user_audio."audio id"
    JOIN choosenUser
    ON choosenUser.id = "user".id
    GROUP BY audio_genre."genre id"
    ORDER BY count(*) DESC 
    LIMIT 1
),
mixedAlbumsGenre as (
    SELECT mixedAlbums.id as id, mixedAlbums.appellation as appellation, album_genre."genre id" as "genre id" FROM
    (
        SELECT album.id as id, album.appellation as appellation FROM album
        WHERE album.appellation LIKE 'mixed%'
    ) as mixedAlbums
    JOIN album_genre
    ON album_genre."album id" = mixedAlbums.id
)
SELECT mixedAlbumsGenre.appellation FROM mixedAlbumsGenre
JOIN choosenUserTopGenre
ON choosenUserTopGenre."genre id" = mixedAlbumsGenre."genre id"


-- ●	Выбрать топ 10 самых популярных авторов за последний год

SELECT author.appellation FROM author
JOIN metric
ON author."metric id" = metric.id
ORDER BY metric."year-popularity"
LIMIT 10

-- ●	Среди 10 самых популярных альбомов за последний год, 
-- выбрать самую залайканую песню, которая жанрово совпадает с метриками пользователя 
-- и у нее есть сниппет

WITH 
choosenUser as (
SELECT "user".id as id FROM "user"
JOIN user_audio
ON "user".id = user_audio."user id"
GROUP BY "user".id
ORDER BY count(*) DESC
LIMIT 1
),
choosenUserTopGenre as (
    SELECT audio_genre."genre id" as "genre id" FROM "user"
    JOIN user_audio
    ON user_audio."user id" = "user".id
    JOIN audio_genre
    ON audio_genre."audio id" = user_audio."audio id"
    JOIN choosenUser
    ON choosenUser.id = "user".id
    GROUP BY audio_genre."genre id"
    ORDER BY count(*) DESC 
    LIMIT 1
),
audioInTop10Albums as (
    SELECT audio.id as id, audio.appellation as appellation, audio."metric id" as "metric id" FROM
    (
        SELECT album.id as id FROM album
        JOIN metric
        ON album."metric id" = metric.id
        ORDER BY metric."year-popularity"
        LIMIT 10
    ) as top10AlbumsID
    JOIN album_audio
    ON album_audio."album id" = top10AlbumsID.id
    JOIN audio
    ON audio.id = album_audio."audio id"
)
SELECT audioInTop10Albums.id, audioInTop10Albums.appellation FROM audioInTop10Albums
JOIN audio_genre
ON audio_genre."audio id" = audioInTop10Albums.id
JOIN snippet
ON snippet."audio id" = audioInTop10Albums.id
JOIN metric
ON audioInTop10Albums."metric id" = metric.id
JOIN choosenUserTopGenre
ON audio_genre."genre id" = choosenUserTopGenre."genre id"
ORDER BY metric."likes" DESC
LIMIT 1


-- ●	Выбрать самый просматриваемыей альбом, который в свою библиотеку добавили авторы, 
-- чей основной жанр совпадает с любимым жанром пользователя
WITH 
choosenUser as (
SELECT "user".id as id FROM "user"
JOIN user_audio
ON "user".id = user_audio."user id"
GROUP BY "user".id
ORDER BY count(*) DESC
LIMIT 1
),
choosenUserTopGenre as (
    SELECT audio_genre."genre id" as "genre id" FROM "user"
    JOIN user_audio
    ON user_audio."user id" = "user".id
    JOIN audio_genre
    ON audio_genre."audio id" = user_audio."audio id"
    JOIN choosenUser
    ON choosenUser.id = "user".id
    GROUP BY audio_genre."genre id"
    ORDER BY count(*) DESC 
    LIMIT 1
),
authorsAddedAlbums as (
    SELECT user_album."album id" as id FROM
    (
        SELECT "user".id as id FROM "user"
        WHERE "user".role = 1
    ) as authorUsersIDs
    JOIN user_album
    ON user_album."user id" = authorUsersIDs.id
),
authorsAddedAlbumsGenre as (
    SELECT authorsAddedAlbums.id as id, album_genre."genre id" as "genre id" FROM authorsAddedAlbums
    JOIN album_genre
    ON album_genre."album id" = authorsAddedAlbums.ID
)
SELECT album.appellation FROM authorsAddedAlbumsGenre
JOIN choosenUserTopGenre
ON authorsAddedAlbumsGenre."genre id" = choosenUserTopGenre."genre id"
JOIN album
ON album.id = authorsAddedAlbumsGenre."id"
JOIN metric
ON album."metric id" = metric.id
ORDER BY metric.views DESC
LIMIT 1


SELECT genre.appellation, count(*) FROM user_audio
JOIN audio_genre
ON user_audio."audio id" = audio_genre."audio id"
JOIN genre
ON genre.id = audio_genre."genre id"
GROUP BY genre.appellation

SELECT * FROM audio

SELECT audio."release data" as "year", SUM(metric.likes) as likes FROM audio
JOIN metric 
ON audio."metric id" = audio.id
GROUP BY audio."release data"
