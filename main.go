package main

import "ewc-backend-go/database"

func main() {
	// Initialize Database
	database.Connect("root@tcp(127.0.0.1:3306)/crud_demo?parseTime=true")
	database.Migrate()
}
