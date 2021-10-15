package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var _ = godotenv.Load(".env")
var (
	ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("user"),
		os.Getenv("pass"),
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("db_name"),
	)
)

const AllowedCORSDomain = "http://localhost"