# Blog AggreGATOR

This is a guided project from boot.dev where I practiced using database
and client using Golang.

This app fetches news or links at an interval, it can only fetch websites that provides accessible RSS
## Config
    you will need to have a config like this
    ```json
    {"db_url":"postgres://postgres:postgres@localhost:5432/gator?sslmode=disable","current_user_name":"raven"}
    ```
## Requirements
- postgres
- go

you can install using this command 
`go install https://github.com/wagbubu/gator`

## commands:
every command should have a prefix of `go run . `
e.g. `go run . register raven` which registers a new user

- register <username>
    - registers a new user
- agg <request interval>
    - fetch all RSS at an interval
- browse <limit results>
    - browse all saved RSS posts in the database, you can limit the results, default is 2
- feeds
    - shows all feeds
- follow <feed url>
    - follow a feed url 
- following
    - shows all following feeds
- unfollow <feed url>
    - unfollow a feed
- login <username>
    - login or change current user
- users
    - shows all users
- addfeed <feed url>
    - add feed to aggregate
- reset 
    - resets database
