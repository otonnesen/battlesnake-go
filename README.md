[![Build Status](https://travis-ci.org/otonnesen/battlesnake-go.svg?branch=master)](https://travis-ci.org/otonnesen/battlesnake-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/otonnesen/battlesnake-go)](https://goreportcard.com/report/github.com/otonnesen/battlesnake-go)
[![codecov](https://codecov.io/gh/otonnesen/battlesnake-go/branch/master/graph/badge.svg)](https://codecov.io/gh/otonnesen/battlesnake-go)
# battlesnake-go

Snake AI implemented in Golang for the 2019 Battlesnake competition

### Running the Snake

Clone the repository with<br>
`git clone https://github.com/otonnesen/battlesnake-go`

Go into the directory and compile<br>
`cd battlesnake-go`<br>
`go build`

Run the snake with<br>
`./battesnake-go`

The web server runs on whatever port the $PORT environment variable is set to, and defaults to 8080 if unset.

## TODO
- [ ] Write tests
- [x] Finish implementing move filtering
- [ ] Add filters to filtering system
- [ ] Come up with sets of filters for different scenarios
