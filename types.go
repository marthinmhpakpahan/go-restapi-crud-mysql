package main

type VideoGame struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Genre string `json:"genre"`
	Year int64 `json:"year"`
}