package model

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

type Player struct {
	TableName []byte `sql:"public.player"`
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Position  string `json:"position"`
	Height    int    `json:"height"`
	Weight    int    `json:"weight"`
	ClubID    string `json:"club_id"`
	NationID  string `json:"nation_id"`
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
	var player Player

	// Tạo bảng
	for _, model := range []interface{}{&league, &club, &cup, &nation, &clubLeague, &nationCup, &player} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp:          false,
			FKConstraints: true,
			IfNotExists:   true,
		})
		if err != nil {
			log.Println(err)
		}
	}

	// Thêm FK constraints cho bảng club_league
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


	// Thêm FK constraints cho bảng nation_cup
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


	// Thêm FK constraints cho bảng player
	_, err = db.Exec(`
		ALTER TABLE player 
		ADD CONSTRAINT player_club_id_fkey FOREIGN KEY (club_id) REFERENCES club(id)
	`)
	if err != nil {
		log.Println(err)
	}
	_, err = db.Exec(`
		ALTER TABLE player 
		ADD CONSTRAINT player_nation_id_fkey FOREIGN KEY (nation_id) REFERENCES nation(id)
	`)
	if err != nil {
		log.Println(err)
	}
}

func SaveClubLeagueData(db *pg.DB) {
	log.Println("Begin save club data")

	var leagueList []League
	for i := 1; i <= 1000; i++ {
		var league League
		league.ID = xid.New().String()
		league.Name = sillyname.GenerateStupidName()

		leagueList = append(leagueList, league)
	}

	var clubList []Club
	for i := 1; i <= 20000; i++ {
		var club Club
		club.ID = xid.New().String()
		club.Name = sillyname.GenerateStupidName()
		club.Stadium = sillyname.GenerateStupidName()
		club.CoachName = randomdata.FullName(randomdata.Male)

		clubList = append(clubList, club)
	}

	err := db.Insert(&leagueList)
	if err != nil {
		log.Println(err)
	}

	err = db.Insert(&clubList)
	if err != nil {
		log.Println(err)
	}

	_, err = db.Exec(`
		INSERT INTO club_league (club_id, league_id)
		SELECT club.id AS club_id, league.id AS league_id
		FROM club, league
		LIMIT 1000000
	`)
	if err != nil {
		log.Println(err)
	}

	log.Println("Finished save club data")
}

func SaveNationCupData(db *pg.DB) {
	log.Println("Begin save nation data")

	var cupList []Cup
	for i := 1; i <= 1000; i++ {
		var cup Cup
		cup.ID = xid.New().String()
		cup.Name = sillyname.GenerateStupidName()

		cupList = append(cupList, cup)
	}

	query := gountries.New()
	allCountries := query.FindAllCountries()
	min := 1
	max := 247
	var nationList []Nation
	for _, v := range allCountries {
		var nation Nation
		nation.ID = xid.New().String()
		nation.Name = v.Name.Official
		nation.Continent = v.Continent
		rand.Seed(time.Now().UnixNano())
		nation.Ranking = rand.Intn(max) + min
		nation.CoachName = randomdata.FullName(randomdata.Male)

		nationList = append(nationList, nation)
	}

	err := db.Insert(&cupList)
	if err != nil {
		log.Println(err)
	}

	err = db.Insert(&nationList)
	if err != nil {
		log.Println(err)
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

	log.Println("Finished save nation data")
}

func SavePlayerData(db *pg.DB) {
	log.Println("Begin save player data")

	var clubIds []string
	_, err := db.Query(&clubIds, `SELECT id AS club_ids FROM club`)
	if err != nil {
		log.Println(err)
	}
	minClub := 0
	maxClub := len(clubIds) - 1

	var nationIds []string
	_, err = db.Query(&nationIds, `SELECT id AS nation_ids FROM nation`)
	if err != nil {
		log.Println(err)
	} 
	minNation := 0
	maxNation := len(nationIds) - 1

	minAge := 18
	maxAge := 22 

	positions := []string{"Goalkeeper", "Defender", "Midfielder", "Striker"}
	minPos := 0
	maxPos := len(positions) - 1

	minHeight := 160
	maxHeight := 200

	minWeight := 65
	maxWeight := 90

	var playerList []Player

	for i := 1; i <= 1000000; i++ {
		var player Player
		rand.Seed(time.Now().UnixNano())
		player.ID = xid.New().String()
		player.FirstName = randomdata.FirstName(randomdata.Male)
		player.LastName = randomdata.LastName()
		player.Age = rand.Intn(maxAge) + minAge
		player.Position = positions[ rand.Intn(maxPos) + minPos ]
		player.Height = rand.Intn(maxHeight) + minHeight
		player.Weight = rand.Intn(maxWeight) + minWeight
		player.ClubID = clubIds[ rand.Intn(maxClub) + minClub ]
		player.NationID = nationIds[ rand.Intn(maxNation) + minNation ]

		playerList = append(playerList, player)
	}

	err = db.Insert(&playerList)
	if err != nil {
		log.Println(err)
	}

	log.Println("Finished save player data")
}