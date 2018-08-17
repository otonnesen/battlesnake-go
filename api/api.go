package api

import (
	"encoding/json"
	"net/http"

	"github.com/otonnesen/battlesnake-go/board"
)

type StartRequest struct {
	GameID int `json:"game_id"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type StartResponse struct {
	Color          string `json:"color"`
	SecondaryColor string `json:"secondary_color"`
	HeadURL        string `json:"head_url"`
	Taunt          string `json:"taunt"`
	HeadType       string `json:"head_type"`
	TailType       string `json:"tail_type"`
}

type MoveRequest struct {
	Board Board `json:"board"`
	Game  Game  `json:"game"`
	Turn  int   `json:"turn"`
	You   Snake `json:"you"`
}

type MoveResponse struct {
	Move string `json:"move"`
}

type Snake struct {
	Body   []board.Point `json:"body"`
	Health int           `json:"health"`
	ID     string        `json:"id"`
	Length int           `json:"length"`
	Name   string        `json:"name"`
	Taunt  string        `json:"taunt"`
}

type Board struct {
	Food   []board.Point `json:"food"`
	Height int           `json:"height"`
	Width  int           `json:"width"`
	Snakes []Snake       `json:"snakes"`
}

type Game struct {
	ID string `json:"id"`
}

func NewStartRequest(req *http.Request) (*StartRequest, error) {
	d := StartRequest{}
	dec := json.NewDecoder(req.Body)
	err := dec.Decode(&d)
	return &d, err
}

func NewMoveRequest(req *http.Request) (*MoveRequest, error) {
	d := MoveRequest{}
	dec := json.NewDecoder(req.Body)
	err := dec.Decode(&d)
	return &d, err
}
