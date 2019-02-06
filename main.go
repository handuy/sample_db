package main

import "sample_db/model"

func main() {
	db := model.ConnectDb()
	defer db.Close()
	model.InitializeDatabase(db)
	model.SaveClubLeagueData(db)
	model.SaveNationCupData(db)
	model.SavePlayerData(db)
}
