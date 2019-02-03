package main

func main() {
	db := ConnectDb()
	defer db.Close()
	InitializeDatabase(db)
	SaveData(db)
}