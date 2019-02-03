package main

import (
	"log"
	"time"

	"math/rand"

	"github.com/Pallinder/go-randomdata"
	"github.com/Pallinder/sillyname-go"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/pariz/gountries"
	xid "github.com/rs/xid"
)

type League struct {
	TableName []byte `sql:"public.league"`
	ID        string `json:"id"`
	Name      string `json:"name"`
}

type Club struct {
	TableName []byte `sql:"public.club"`
	ID        string `json:"id"`
	Name      string `json:"name"`
	Stadium   string `json:"stadium"`
	CoachName string `json:"coach_name"`
}

type ClubLeague struct {
	TableName []byte `sql:"public.club_league"`
	ClubID    string `json:"club_id" sql:",pk"`
	LeagueID  string `json:"league_id" sql:",pk"`
}

type Cup struct {
	TableName []byte `sql:"public.cup"`
	ID        string `json:"id"`
	Name      string `json:"name"`
}

type Nation struct {
	TableName []byte `sql:"public.nation"`
	ID        string `json:"id"`
	Name      string `json:"name"`
	Continent string `json:"continent"`
	Ranking   int    `json:"ranking"`
	CoachName string `json:"coach_name"`
}

type NationCup struct {
	TableName []byte `sql:"public.nation_cup"`
	NationID  string `json:"nation_id" sql:",pk"`
	CupID     string `json:"cup_id" sql:",pk"`
}

// ConnectDb kết nối CSDL qua cấu hình
func ConnectDb() (db *pg.DB) {
	db = pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "123",
		Database: "postgres",
		Addr:     "localhost:5432",
	})
	return db
}

func InitializeDatabase(db *pg.DB) {
	var league League
	var club Club
	var clubLeague ClubLeague
	var cup Cup
	var nation Nation
	var nationCup NationCup

	// Tạo bảng
	for _, model := range []interface{}{&league, &club, &cup, &nation, &clubLeague, &nationCup} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp:          false,
			FKConstraints: true,
			IfNotExists:   true,
		})
		if err != nil {
			log.Println(err)
		}
	}

	_, err := db.Exec(`
		ALTER TABLE club_league 
		ADD CONSTRAINT club_league_club_id_fkey FOREIGN KEY (club_id) REFERENCES club(id)
	`)
	if err != nil {
		log.Println(err)
	}

	_, err = db.Exec(`
		ALTER TABLE club_league 
		ADD CONSTRAINT club_league_league_id_fkey FOREIGN KEY (league_id) REFERENCES league(id)
	`)
	if err != nil {
		log.Println(err)
	}

	_, err = db.Exec(`
		ALTER TABLE nation_cup 
		ADD CONSTRAINT nation_cup_nation_id_fkey FOREIGN KEY (nation_id) REFERENCES nation(id)
	`)
	if err != nil {
		log.Println(err)
	}

	_, err = db.Exec(`
		ALTER TABLE nation_cup 
		ADD CONSTRAINT nation_cup_cup_id_fkey FOREIGN KEY (cup_id) REFERENCES cup(id)
	`)
	if err != nil {
		log.Println(err)
	}
}

func SaveData(db *pg.DB) {
	for i := 1; i <= 1000; i++ {
		var league League
		league.ID = xid.New().String()
		league.Name = sillyname.GenerateStupidName()

		err := db.Insert(&league)
		if err != nil {
			log.Println(err)
		}
	}

	for i := 1; i <= 20000; i++ {
		var club Club
		club.ID = xid.New().String()
		club.Name = sillyname.GenerateStupidName()
		club.Stadium = sillyname.GenerateStupidName()
		club.CoachName = randomdata.FullName(randomdata.Male)

		err := db.Insert(&club)
		if err != nil {
			log.Println(err)
		}
	}

	_, err := db.Exec(`
		INSERT INTO club_league (club_id, league_id)
		SELECT club.id AS club_id, league.id AS league_id
		FROM club, league
		LIMIT 10000000
	`)
	if err != nil {
		log.Println(err)
	}

	for i := 1; i <= 1000; i++ {
		var cup Cup
		cup.ID = xid.New().String()
		cup.Name = sillyname.GenerateStupidName()

		err := db.Insert(&cup)
		if err != nil {
			log.Println(err)
		}
	}

	query := gountries.New()
	allCountries := query.FindAllCountries()
	min := 1
	max := 247
	for _, v := range allCountries {
		var nation Nation

		nation.ID = xid.New().String()
		nation.Name = v.Name.Official
		nation.Continent = v.Continent

		rand.Seed(time.Now().UnixNano())
		nation.Ranking = rand.Intn(max) + min

		nation.CoachName = randomdata.FullName(randomdata.Male)

		err := db.Insert(&nation)
		if err != nil {
			log.Println(err)
		}
	}

	_, err = db.Exec(`
		INSERT INTO nation_cup (nation_id, cup_id)
		SELECT nation.id AS nation_id, cup.id AS cup_id
		FROM nation, cup
		LIMIT 150000
	`)
	if err != nil {
		log.Println(err)
	}
}
