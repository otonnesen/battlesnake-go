package main

import (
	"encoding/json"
	"net/http"
)

// StartRequest is specified by the
// 2019 battlesnake API
type StartRequest struct {
	GameID int `json:"game_id"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

// StartResponse is specified by the
// 2019 battlesnake API
type StartResponse struct {
	Color          string `json:"color"`
	SecondaryColor string `json:"secondary_color"`
	HeadURL        string `json:"head_url"`
	Taunt          string `json:"taunt"`
	HeadType       string `json:"head_type"`
	TailType       string `json:"tail_type"`
}

// MoveRequest is specified by the
// 2019 battlesnake API
type MoveRequest struct {
	Board board `json:"board"`
	Game  game  `json:"game"`
	Turn  int   `json:"turn"`
	You   Snake `json:"you"`
}

// MoveResponse is specified by the
// 2019 battlesnake API
type MoveResponse struct {
	Move string `json:"move"`
}

type board struct {
	Food   []Point `json:"food"`
	Height int     `json:"height"`
	Width  int     `json:"width"`
	Snakes []Snake `json:"snakes"`
}

type game struct {
	ID string `json:"id"`
}

// NewStartRequest unmarshals JSON from an http.Request into a StartRequest
func NewStartRequest(req *http.Request) (*StartRequest, error) {
	d := StartRequest{}
	err := json.NewDecoder(req.Body).Decode(&d)
	return &d, err
}

// NewMoveRequest unmarshals JSON from an http.Request into a MoveRequest
func NewMoveRequest(req *http.Request) (*MoveRequest, error) {
	d := MoveRequest{}
	err := json.NewDecoder(req.Body).Decode(&d)
	return &d, err
}
