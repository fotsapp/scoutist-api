# Scoutist API
Written in Golang for fmscout app's backend.


# Includes

- Banners
- Nations
- League
- Teams 
- Players 


## Items

### Banners

Field | Description
------|------------
**banner_id** | The item's unique id.
category_name | Banner name get value.
banner_link | The URL of the banner image.

### Nation

Field | Description
------|------------
**id** | The item's unique id.
name | Nation name get value.
image | The URL of the nation image.

### League

Field | Description
------|------------
**id** | The item's unique id.
nation_id | Get nation id value.
league_name | The name of the league name.
league_image | The URL of the league image.

### Teams

Field | Description
------|------------
**id** | The item's unique id.
league_id | Get league id value.
team_name | The name of the team name.
team_image | The URL of the team image.

### Players

Field | Description
------|------------
**id** | The item's unique id.
team_id | Get team id value.
name | The player of the name.
image | The player of the image.
age | The player of the age.
position | The player of the position.
value | The player of the value.
expr_value | The player of the exprience vale.
expr_date | Provides information on the end date of the contract
rating | Returns the current talent value of the player.
potential | Returns the potential value of the player. 


# Execute Links

- Banners : http://localhost:8080/banners
- Nations : http://localhost:8080/teams
- Leagues : http://localhost:8080/leagues
- Teams   : http://localhost:8080/teams
- Players : http://localhost:8080/players 

# Output

- Banner 

```javascript
[
    {
    BannerID: "1",
    CategoryName: "Goalkeepers of Future",
    BannerLink: "http://allpicts.in/download/7122/2016/02/Gianluigi_Donnarumma_AC_Milan_best_goalkeepers_3-1600x900.jpg"
    }
    ...
]
```


- Nation 

```javascript
[
    {
        "ID": "106",
        "Name": "Afghanistan",
        "Image": "https://sortitoutsi.net/uploads/flags/106.png"
    },
    {
        "ID": "752",
        "Name": "Albania",
        "Image": "https://sortitoutsi.net/uploads/flags/752.png"
    }
  ...
]
```

- Leagues

```javascript
[
    {
        "id": "23106020",
        "nationID": "106",
        "leagueName": "Afghan Premier League",
        "league_image": "https://sortitoutsi.net/uploads/comp_sm/23106020.png"
    },
    {
        "id": "23151830",
        "nationID": "106",
        "leagueName": "Afghan Lower Division",
        "league_image": "https://sortitoutsi.net/uploads/comp_sm/23151830.png"
    }
  ...
]
```

- Teams 

```javascript
[
    {
        "ID": "23184348",
        "LeagueID": "23106020",
        "TeamName": "De Abasin Sape",
        "TeamImage": "https://sortitoutsi.net/uploads/team_sm/23184348.png"
    },
    {
        "ID": "23184346",
        "LeagueID": "23106020",
        "TeamName": "De Maiwand Atalan",
        "TeamImage": "https://sortitoutsi.net/uploads/team_sm/23184346.png"
    }
    ...
]
```

- Players

```javascript
[
    {
        "id": "83110775",
        "teamID": "83111761",
        "Name": "Rxfaben Castro",
        "image": "https://sortitoutsi.net/uploads/iconface/83110775.png",
        "age": "23",
        "position": "AM LC, F C",
        "value": "30",
        "expr_value": "30",
        "exprDate": "30-06-2019",
        "rating": "44.50",
        "potential": "49.80"
    },
    {
        "id": "83115497",
        "teamID": "83111761",
        "Name": "Sxe9rgio Oliveira",
        "image": "https://sortitoutsi.net/uploads/iconface/83115497.png",
        "age": "22",
        "position": "D C",
        "value": "30",
        "expr_value": "30",
        "exprDate": "30-06-2019",
        "rating": "42.80",
        "potential": "-4"
    }
    ...
]
```

# Usage 

1. Git clone this project.

*git@github.com:scoutist/scoutist-api.git*

2. Run project.

*go run scoutist.go*

3. Execute query.
- http://localhost:8080/banners
- http://localhost:8080/nations
- http://localhost:8080/leagues
- http://localhost:8080/teams
- http://localhost:8080/players


> Domains will change.







The MIT License- see [`LICENSE`](LICENSE) for more details
