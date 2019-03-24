package main


import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"encoding/json"
    "log"
	"net/http"
	"strings"
)

type banners struct{
	/* category_id ,category_name , banner_link*/
	CategoryID string
	CategoryName string
	BannerLink string
}

type nations struct{
	/* id, name, image, link*/
	ID string 
	Name string 
	Image string
}

type leagues struct{
	/* id, nationID, leagueName, league_image, league_link */
	ID string
	NationID string 
	LeagueName string 
	LeagueImage string 
}

type teams struct{
	ID string
	LeagueID string 
	TeamName string 
	TeamImage string
}

type players struct{
	/* id, teamID, name, image, age, position, value, exprValue, exprDate, rating, potential*/
	ID string 
	TeamID string 
	Name string
	Image string 
	Age string 
	Position string 
	Value string 
	ExprValue string 
	ExprDate string 
	Rating string 
	Potential string 
}


var db *sql.DB

const (
	db_host		= "localhost"
	db_user     = "postgres"
	db_password = "~"
	db_name     = "~"
	db_port		= "~"
)

var nationList []nations
var leagueList []leagues
var teamList 	[]teams
var playerList []players
var bannerList []banners


func main(){
	// connect to postgres database.
	initDB()
	defer db.Close()
	/* Writter method for API Response. */
	http.HandleFunc("/nations", GetNations)

	http.HandleFunc("/leagues", Getleagues)

	http.HandleFunc("/teams",GetTeams)

	http.HandleFunc("/players",GetPlayers)

	http.HandleFunc("/banners",GetBanner)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func initDB(){
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s "+
	"password=%s dbname=%s sslmode=disable",db_host,db_port,db_user,db_password,db_name)
	var err error

	db,err = sql.Open("postgres",dbinfo)

	if err != nil {
        panic(err)
    }
    err = db.Ping()
    if err != nil {
        panic(err)
    }
    fmt.Println("Successfully connected!")

	//check error.
	checkErr(err)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}


/* Display all banner datas and query datas.*/
func GetBanner(w http.ResponseWriter, req *http.Request) {

	err  := queryBanner()

	if err != nil {
		fmt.Print("err boş değil.")
		http.Error(w, err.Error(), 500)
		return
	}

    out, err := json.Marshal(bannerList)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, string(out))
}

func queryBanner() error{
	
	query 	  	:= "select category.category_id,category.category_name,banner.banner_link from category inner join banner on (category.category_id=banner.category_id)"

	rows, err 	:= db.Query(query)

	defer rows.Close()

	bannerList=nil
	for rows.Next() {
		banner		 := banners{}
		err 		 = rows.Scan(&banner.CategoryID,&banner.CategoryName,&banner.BannerLink)
	
		bannerList  = append(bannerList,banner)
	}
	err = rows.Err()


	if err != nil {
		return err
	}
	return nil

}


/* Display all nations datas and query datas.*/
func GetNations(w http.ResponseWriter, req *http.Request) {

	id   := req.URL.Query().Get("id")
	name := req.URL.Query().Get("name")

	err  := queryNations(id,name)

	if err != nil {
		fmt.Print("err boş değil.")
		http.Error(w, err.Error(), 500)
		return
	}

    out, err := json.Marshal(nationList)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, string(out))
}

func queryNations(id string, name string) error{
	
	var query string

	if id=="" && name=="" {
		query	    = fmt.Sprintf(`SELECT id,name,image FROM nations`)
	}else{
		query 	  	= fmt.Sprintf("SELECT id,name,image FROM nations WHERE id ='%s' OR name='%s' ", id,name)
	}

	rows, err 	:= db.Query(query)

	defer rows.Close()

	nationList=nil
	for rows.Next() {
		nation		 := nations{}
		err 		 = rows.Scan(&nation.ID,&nation.Name,&nation.Image)
	
		nationList  = append(nationList,nation)
	}
	err = rows.Err()


	if err != nil {
		return err
	}
	return nil

}

/* Display all league datas and query data.*/
func Getleagues(w http.ResponseWriter, req *http.Request) {

	id        	:= req.URL.Query().Get("id")
	nationID 	:= req.URL.Query().Get("nationID")
	leagueName := req.URL.Query().Get("leagueName")

	err := queryleagues(id,nationID,leagueName)

	if err != nil {
		fmt.Print("err boş değil.")
		http.Error(w, err.Error(), 500)
		return
	}

    out, err := json.Marshal(leagueList)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, string(out))
}

func queryleagues(id string,nationID string,leagueName string) error{
	//select * from leagues where id='' or nationID='' or leagueName=''

	var query string

	if id=="" && nationID=="" && leagueName=="" {
		query	    = fmt.Sprintf(`SELECT * FROM leagues`)
	}else{
		query 	  	= fmt.Sprintf("SELECT * FROM leagues WHERE id='%s' OR nation_id='%s' OR league_name='%s' ", id,nationID,leagueName)
	}

	rows, err := db.Query(query)

	if err != nil {
		return err
	}

	defer rows.Close()

	leagueList=nil
	for rows.Next() {
		league		 := leagues{}
		err 		  = rows.Scan(&league.ID,&league.NationID,&league.LeagueName,&league.LeagueImage)
		leagueList    = append(leagueList,league)
	}
	err = rows.Err()


	if err != nil {
		return err
	}
	return nil
}

/* Display all team datas and query data*/
func GetTeams(w http.ResponseWriter, req *http.Request) {

	id        	:= req.URL.Query().Get("id")
	leagueID 	:= req.URL.Query().Get("leagueID")
	teamName 	:= req.URL.Query().Get("teamName")

	err := queryTeams(id,leagueID,teamName)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Print(teamList)
	out, err := json.Marshal(teamList)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, string(out))
}


func queryTeams(id string,leagueID string,teamName string) error{

	var query string

	if id=="" && leagueID=="" && teamName=="" {
		query	    = fmt.Sprintf(`SELECT * FROM teams LIMIT 50`)
	}else{
		query 	  	= fmt.Sprintf("SELECT * FROM nations WHERE id='%s' or league_id='%s' or team_name='%s' ", id,leagueID,teamName)
	}

	rows, err 	:= db.Query(query)

	if err != nil {
		return err
	}

	defer rows.Close()

	teamList= nil 
	for rows.Next() {
		team		 		:= teams{}
		err 		 		= rows.Scan(&team.ID,&team.LeagueID,&team.TeamName,&team.TeamImage)
		teamList			= append(teamList,team)
	}

	//fmt.Println(teamList)
	err = rows.Err()

	if err != nil {
		return err
	}
	return nil
}

/* Display all team datas and query data*/
func GetPlayers(w http.ResponseWriter, req *http.Request) {

	id   	 	:= req.URL.Query().Get("id")
	teamID  	:= req.URL.Query().Get("teamID")
	name 	 	:= req.URL.Query().Get("name")
	age 	 	:= strings.Split(req.URL.Query().Get("age"),"-")
	position  	:= req.URL.Query().Get("position")
	value  	   	:= strings.Split(req.URL.Query().Get("value"),"-")
	exprValue 	:= strings.Split(req.URL.Query().Get("exprValue"),"-")
	exprDate 	:= strings.Split(req.URL.Query().Get("exprDate"),"-")
	rating 		:= strings.Split(req.URL.Query().Get("rating"),"-")
	potential 	:= strings.Split(req.URL.Query().Get("potential"),"-")

	fmt.Print("1")
    err := queryPlayers(id,teamID,name,age,position,value,exprValue,exprDate,rating,potential)

	if err != nil {
		fmt.Print("err boş değil.")
		http.Error(w, err.Error(), 500)
		return
	}

    out, err := json.Marshal(playerList)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, string(out))
}

func queryPlayers(id string,teamID string,name string,age [] string,position string,value [] string,exprValue [] string,exprDate [] string,rating [] string,potential [] string) error{
	
	fmt.Print("2")

	var query string

	query	    	= fmt.Sprintf(`SELECT * FROM players LIMIT 250`)

	/*if age!=nil && value!=nil && exprDate!=nil && exprValue !=nil && rating != nil && potential !=nil {
		if(len(age)==2 && len(value)==2 && le)
		if id=="" && teamID=="" && name=="" && age[0]=="" && age[1]=="" && position=="" && value[0]=="" && value[1]=="" && exprValue[0]=="" && exprValue[1]=="" && exprDate[0]=="" && exprDate[1]=="" && rating[0]=="" && rating[1]=="" && potential[0]=="" && potential[1]=="" {
			query	    = fmt.Sprintf(`SELECT * FROM players LIMIT 250`)
		}else{
			query 	  	= fmt.Sprintf("SELECT id,team_id,name,age,position,value,expr_value,expr_date,rating,potentialfrom playerswhere id='%s' or team_id='%s' or name='%s' or age between '%s' and '%s' or position like '% %s %' or value between '%s' and '%s' or expr_value between '%s' and '%s' or expr_date between '%s' and '%s' or rating between '%s' and '%s' or potential between '%s' and '%s' ",id,teamID,name,age[0],age[1],position,value[0],value[1],exprValue[0],exprValue[1],exprDate[0],exprDate[1],rating[0],rating[1],potential[0],potential[1])
		}
	}else{
	}*/
	query	    = fmt.Sprintf(`SELECT * FROM players LIMIT 250`)


	rows, err 	:= db.Query(query)

	if err != nil {
		return err
	}

	defer rows.Close()

	playerList=nil
	for rows.Next() {
		player		 := players{}
		err 	 	 = rows.Scan(&player.ID,&player.TeamID,&player.Name,&player.Image,&player.Age,&player.Position,&player.Value,&player.ExprValue,&player.ExprDate,&player.Rating,&player.Potential)
		playerList  = append(playerList,player)
	}
	err = rows.Err()


	if err != nil {
		return err
	}
	return nil
}
