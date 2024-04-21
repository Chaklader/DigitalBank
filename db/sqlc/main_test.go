package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const dbSource = "postgresql://root:secret@127.0.0.1:5432/simple_bank?sslmode=disable"

var testStore Store

func TestMain(m *testing.M) {
	//config, err := util.LoadConfig("../..")
	//if err != nil {
	//	log.Fatal("cannot load config:", err)
	//}

	connPool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testStore = NewStore(connPool)
	os.Exit(m.Run())
}
