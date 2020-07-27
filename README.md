GoBot
===

This project aims at providing an administration bot for [Urban Terror](https://www.urbanterror.info/home/).  I am aware the game is not currently very active, but the idea is to have a proper code base for the upcoming Urban Terror. The architecture should also make GoBot a viable option for other FPS games.

It is still highly WIP, you can find behind the things I have done/plan to do. It is built using Goland and MySQL.

## How to install

### Preliminary steps

- Install [golang](https://golang.org/doc/install) (it is super fast to do)
- Install MariaDB and setup a database (you can use create_database.sh for that).

### The bot

1. Configure the bot (tbd)
2. Launch using `go run .`


## Checklist

### Now: before an alpha

- [x] Core architecture: a parsing module passes parsed Events to the main scheduler which calls the plugins.
- [x] Server status. The server maintains a list of connected players and cvars.
- [x] RCON interface.
- [x] Logging module
- [ ] Proper config module
- [ ] Database logging of events
- [ ] Plugin: admin. !iamgod, !kick is done. Bans/warns to be implemented.
- [ ] Plugin: stats. SQL to do
- [ ] Proper linting/Fix last TODOs
- [ ] All fancy plugins
- [ ] Debugging

### Future

- [ ] Proper testing
- [ ] Web back-end for administration/stats
- [ ] Web front-end for administration/stats