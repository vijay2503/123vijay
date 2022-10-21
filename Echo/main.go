package main

import (
	conn "echo-task/dbconnection"
	rt "echo-task/router"
)

func main() {
	conn.DbConn()
	e := rt.ApiRouting()
	e.Logger.Fatal(e.Start(":8000"))
}
