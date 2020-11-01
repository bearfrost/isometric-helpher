package main

import (
	"flag"
	"fmt"
	"github.com/photostorm/tiled"
)

var (
	mapFile = flag.String("f", "testIMap.tmx", "Tile map (.tmx)")
	mode    = flag.String("m", "basic", "control what information is showed")
)

func main() {
	flag.Parse()

	var tileMap *tiled.Map
	var err error

	tileMap, err = tiled.LoadFromFile(*mapFile)
	if err != nil {
		println(err.Error())
		return
	}

	switch *mode {
	case "basic":
		fmt.Printf("Render Order: %s", tileMap.RenderOrder)
		fmt.Printf("Map Height: %d", tileMap.Height)
		fmt.Printf("Map Width: %d", tileMap.Width)
		fmt.Printf("Tile Height: %d", tileMap.TileHeight)
		fmt.Printf("Tile Width: %d", tileMap.TileWidth)
		return
	}
}
