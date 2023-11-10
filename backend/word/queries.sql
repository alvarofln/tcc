-- name: FindWordById :one
SELECT *
FROM words
WHERE id = ?
LIMIT 1;
-- name: FindWordByName :one
SELECT *
FROM words
WHERE name = ?
LIMIT 1;
-- name: FindSimilarWordsById :many
SELECT w.id, w.name, sw.similarity
FROM similar_words sw
         JOIN words w on w.id = sw.word_similar_id
WHERE sw.word_id = ?
ORDER BY sw.similarity DESC;
-- name: FindAllWords :many
SELECT *
FROM words