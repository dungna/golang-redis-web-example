package models

import (
	"github.com/mediocregopher/radix.v2/pool"
	"log"
	"errors"
	"strconv"
)

var db *pool.Pool

func init()  {
	var err error
	db, err = pool.New("tcp", "localhost:6379", 10)
	if err != nil {
		log.Panic(err)
	}
}

var ErrNoAlbum = errors.New("models: no album found")

type Album struct {
	Title string
	Artist string
	Price  float64
	Likes  int
}

func populateAlbum(reply map[string]string) (*Album, error) {
	var err error
	ab := new(Album)
	ab.Title = reply["title"]
	ab.Artist = reply["artist"]
	ab.Price, err = strconv.ParseFloat(reply["price"], 64)
	if err != nil {
		return nil, err
	}
	ab.Likes, err = strconv.Atoi(reply["likes"])
	if err != nil {
		return nil, err
	}
	return ab, err
}

func FindAlbum(id string) (*Album, error){
	conn, err := db.Get()
	if err != nil {
		return nil, err
	}
	defer db.Put(conn)
	reply, err := conn.Cmd("HGETALL", "album:" + id).Map()
	if err != nil {
		return nil, err
	} else if len(reply) == 0 {
		return nil, ErrNoAlbum
	}

	return populateAlbum(reply)

}
