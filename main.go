package main

import (
	"blog-aggregator/internal/config"
	"blog-aggregator/internal/database"
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)



func main () {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	/*
	* Open Database
	*/
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error oppening the database %v", err)
	}

	dbQueries := database.New(db)
	/*
	* Initialize state
	*/
	appState := &State{ cfg: &cfg, db: dbQueries }
	/*
	* Initialize commands list 
	*/
	commands := Commands{ handlers: make(map[string]func(*State, Command) error) }
	/*
	* Register commands
	*/
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
  commands.register("users", handlerUsers)
	commands.register("agg", handlerAgg)
	commands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	commands.register("feeds", handlerFeeds)
	commands.register("follow", middlewareLoggedIn(handlerFollow))
	commands.register("following", handlerFollowing)
	commands.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	commands.register("browse", middlewareLoggedIn(handlerBrowse))

	userInput := os.Args
	/*
	* [0]first argument is file directory, [1]second is the command name, [2]third and do on must be the command args
	*/
	if len(userInput) < 2 {
		log.Fatal("Too few arguments")
		os.Exit(1)
	}

	commandName := userInput[1]
	arguments := userInput[2:]
  userCommand := Command{Name: commandName, Args: arguments}
  
	err = commands.run(appState, userCommand)
	if err != nil {
		log.Fatal(err)
	}
}	