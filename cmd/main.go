package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/halja7/kv/db"
	"github.com/halja7/kv/kv"
)

func main() {

	db := db.NewDb("db.json")
	store, err := kv.NewStore(db)
	if err != nil {
		fmt.Printf("Error creating store: %s", err)
		os.Exit(1)
	}

	input, err := parseInputs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch input.Op {
	case "set":
		err := store.Set(input.Key, input.Value)
		if err != nil {
			fmt.Println("Something went terribly wrong.")
			os.Exit(1)
		}
	case "get":
		v := store.Get(input.Key)
		if v == "" {
			fmt.Println("(nil)")
		} else {
			fmt.Println(v)
		}
	default:
		fmt.Println("Something went terribly wrong.")
		os.Exit(1)
	}

}

func parseInputs() (kv.KeyValue, error) {
	set := flag.NewFlagSet("set", flag.ExitOnError)
	get := flag.NewFlagSet("get", flag.ExitOnError)

	if len(os.Args) < 2 {
		return kv.KeyValue{}, fmt.Errorf("expected 'set' or 'get' subcommands")
	}

	var key, value, op string

	switch os.Args[1] {
	case "set":
		set.Parse(os.Args[2:])
		setArgs := set.Args()

		if len(setArgs) < 2 {
			return kv.KeyValue{}, fmt.Errorf("expected arguments to 'set': kv set <key> <value>")
		}

		key = setArgs[0]
		value = setArgs[1]
		op = "set"
	case "get":
		get.Parse(os.Args[2:])
		getArgs := get.Args()

		if len(getArgs) < 1 {
			return kv.KeyValue{}, fmt.Errorf("expected arguments to 'get': kv get <key>")
		}

		key = getArgs[0]
		op = "get"
	default:
		return kv.KeyValue{}, fmt.Errorf("expected 'set' or 'get' subcommands")
	}

	return kv.KeyValue{
		Key:   key,
		Value: value,
		Op:    op,
	}, nil

}
