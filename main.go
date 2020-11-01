package main

import (
	"fmt"
	"github.com/photostorm/tiled"
)

func main() {
	var tileMap *tiled.Map
	var err error

	tileMap, err = tiled.LoadFromFile("testIMap.tmx")
	if err != nil {
		println(err.Error())
		return
	}

	// TODO Layer sorter

	fmt.Printf("Map Height: %d", tileMap.Height)
	fmt.Printf("Map Width: %d", tileMap.Width)
	fmt.Printf("Tile Height: %d", tileMap.TileHeight)
	fmt.Printf("Tile Width: %d", tileMap.TileWidth)
}
