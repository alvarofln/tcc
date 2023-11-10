package main

import (
	"context"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	_ "github.com/mattn/go-sqlite3"
	"gotoolapi/lyric"
	"gotoolapi/word"
	"gotoolapi/word/db"
)

func main() {
	ctx := context.Background()
	database, err := sql.Open("sqlite3", "file:word2vec.db?cache=shared")

	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/api/words", func(c *fiber.Ctx) error {
		q := db.New(database)
		s := word.NewService(q)

		w, err := s.FindWordByName(ctx, c.Query("q"))

		if err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(w)
	})

	app.Get("/api/words/:id", func(c *fiber.Ctx) error {
		q := db.New(database)
		s := word.NewService(q)

		id, err := c.ParamsInt("id")

		if err != nil {
			return err
		}

		w, err := s.FindWordById(ctx, int64(id))

		if err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(w)
	})

	app.Get("/api/words/:id/similars", func(c *fiber.Ctx) error {
		q := db.New(database)
		s := word.NewService(q)

		id, err := c.ParamsInt("id")

		if err != nil {
			return err
		}

		sw, err := s.FindSimilarWordsById(ctx, int64(id))

		if err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(sw)
	})

	app.Get("/api/words-rhyming", func(c *fiber.Ctx) error {
		q := db.New(database)
		s := word.NewService(q)

		//get query parameter q:
		qp := c.Query("q")
		if qp == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "missing query parameter q",
			})
		}

		sw, err := s.FindRhymingWords(ctx, qp)

		if err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(sw)
	})

	app.Post("/api/lyrics", func(c *fiber.Ctx) error {

		input := new(lyric.InputLyric)

		if err := c.BodyParser(input); err != nil {
			return err
		}
		q := db.New(database)

		ls := lyric.NewService(word.NewService(q))

		output, err := ls.Process(ctx, input)

		if err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(output)
	})

	err = app.Listen(":3000")

	if err != nil {
		panic("Error starting the server :(")
	}
}
