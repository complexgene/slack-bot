package repo

import (
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

const User = "qvttnfslgwtvmm"
const Db = "ddus9c4pca9cj7"
const Host = "ec2-54-155-61-133.eu-west-1.compute.amazonaws.com"

//"postgres://user:pass@localhost:5432/db_name"
const DBStringProd = "postgres://" + User + ":58886887368b5edc4def71e815ae49c3610bdfa750c5eb1ab99c221d8be5366f@" + Host + ":5432/" + Db

func GetDBConn() *bun.DB {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(DBStringProd)))
	db := bun.NewDB(sqldb, pgdialect.New())
	return db
}
