package model

type GetTotalGamesSalesReport struct {
	Name  string  `db:"name"`
	Total float64 `db:"total"`
}

type GetMostPopularGameReport struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Total int    `db:"total"`
}

type GetTotalRevenuePerGameReport struct{
	Name  string  `db:"name"`
	Total float64 `db:"total"`
}