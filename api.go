package main

import (
	"encoding/json"
	"net/http"
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

type Board struct {
	Food   []Point `json:"food"`
	Height int     `json:"height"`
	Width  int     `json:"width"`
	Snakes []Snake `json:"snakes"`
}

type Game struct {
	ID string `json:"id"`
}

func NewStartRequest(req *http.Request) (*StartRequest, error) {
	d := StartRequest{}
	err := json.NewDecoder(req.Body).Decode(&d)
	return &d, err
}

func NewMoveRequest(req *http.Request) (*MoveRequest, error) {
	d := MoveRequest{}
	err := json.NewDecoder(req.Body).Decode(&d)
	return &d, err
}
