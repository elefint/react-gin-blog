package database

import (
	"github.com/go-pg/pg/v10"
)

func NewDBOptions() *pg.Options {
	return &pg.Options{
		Addr:     "ec2-176-34-211-0.eu-west-1.compute.amazonaws.com",
		Database: "d4rcgn11tuuc8q",
		User:     "jvziibuspteepf",
		Password: "1742f8ef327cd1de02ec31bb93b41c58cebf8409ee5bf7285acde247e0fd37bd",
	}
}
