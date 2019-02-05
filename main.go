package main

func main() {
	db := ConnectDb()
	defer db.Close()
	InitializeDatabase(db)
	SaveClubLeagueData(db)
	SaveNationCupData(db)
	SavePlayerData(db)
}
