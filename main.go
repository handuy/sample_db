package main

import "log"

func main() {
	db := ConnectDb()
	defer db.Close()
	InitializeDatabase(db)
	// SaveClubLeagueData(db)
	// SaveNationCupData(db)
	log.Println("begin")
	SavePlayerData(db)
	log.Println("finished")
}
