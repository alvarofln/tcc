// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import ()

type SimilarWord struct {
	WordID        int64
	WordSimilarID int64
	Similarity    float64
}

type Word struct {
	ID   int64
	Name string
}
