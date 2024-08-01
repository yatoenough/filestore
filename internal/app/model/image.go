package model

type Image struct {
	ID           int64  `db:"id"`
	OriginalName string `db:"original_name"`
	StoreName    string `db:"store_name"`
	Path         string `db:"path"`
	Extension    string `db:"ext"`
}
