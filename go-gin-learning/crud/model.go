package main
type Book struct{
	Id int64 `json:"id" db:"id"`
	Title string `db:"title" json:"title"`
	Price int64 `json:"price" db:"price"`
}
