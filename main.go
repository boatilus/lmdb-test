package main

import (
	"fmt"
	"log"

	"github.com/bmatsuo/lmdb-go/lmdb"
)

// open a database that can be used as long as the enviroment is mapped.
var dbi lmdb.DBI

var env *lmdb.Env

func init() {
	// create an environment and make sure it is eventually closed.
	var err error

	env, err = lmdb.NewEnv()
	if err != nil {
		log.Fatal(err)
	}

	if err := env.SetMaxDBs(1); err != nil {
		log.Fatal(err)
	}

	if err := env.SetMapSize(1 << 30); err != nil {
		log.Fatal(err)
	}

	if err := env.Open("./", 0, 0644); err != nil {
		log.Fatal(err)
	}
}

func main() {
	err := env.Update(func(txn *lmdb.Txn) (err error) {
		dbi, err = txn.OpenDBI("dbnew", lmdb.Create)
		return err
	})
	if err != nil {
		log.Fatal(err)
	}

	// Run only once!
	// err = env.Update(func(txn *lmdb.Txn) (err error) {
	// 	err = txn.Put(dbi, []byte("hello"), []byte("world!"), 0)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	return nil
	// })

	// the database is now ready for use.  read the value for a key and print
	// it to standard output.
	err = env.View(func(txn *lmdb.Txn) (err error) {
		v, err := txn.Get(dbi, []byte("hello"))
		if err != nil {
			return err
		}

		fmt.Println(string(v))

		return nil
	})
	if err != nil {
		log.Print(err)
	}
}
