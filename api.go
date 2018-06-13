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
	Food   PointList `json:"food"`
	Height int       `json:"height"`
	ID     int       `json:"id"`
	Snakes SnakeList `json:"snakes"`
	Turn   int       `json:"turn"`
	Width  int       `json:"width"`
	You    Snake     `json:"you"`
}

type MoveResponse struct {
	Move string `json:"move"`
}

type EndRequest struct {
	GameID     int           `json:"game_id"`
	Winners    []string      `json:"winners"`
	DeadSnakes DeadSnakeList `json:"dead_snakes"`
}

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type PointList []Point

type Snake struct {
	Body   PointList `json:"body"`
	Health int       `json:"health"`
	ID     string    `json:"id"`
	Length int       `json:"length"`
	Name   string    `json:"name"`
	Taunt  string    `json:"taunt"`
}

type SnakeList []Snake

type DeathRecap struct {
	Turn   int      `json:"turn"`
	Causes []string `json:"causes"`
}

type DeadSnake struct {
	ID     string     `json:"id"`
	Length int        `json:"length"`
	Death  DeathRecap `json:"death"`
}

type DeadSnakeList []DeadSnake

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

func NewEndRequest(req *http.Request) (*EndRequest, error) {
	d := EndRequest{}
	dec := json.NewDecoder(req.Body)
	err := dec.Decode(&d)
	return &d, err
}

func (l *PointList) UnmarshalJSON(b []byte) error {
	var list struct {
		Data []Point `json:"data"`
	}

	err := json.Unmarshal(b, &list)
	if err != nil {
		return err
	}
	*l = list.Data
	return nil
}

func (l *SnakeList) UnmarshalJSON(b []byte) error {
	var list struct {
		Data []Snake `json:"data"`
	}

	err := json.Unmarshal(b, &list)
	if err != nil {
		return err
	}
	*l = list.Data
	return nil
}

func (l *DeadSnakeList) UnmarshalJSON(b []byte) error {
	var list struct {
		Data []DeadSnake `json:"data`
	}

	err := json.Unmarshal(b, &list)
	if err != nil {
		return err
	}
	*l = list.Data
	return nil
}
