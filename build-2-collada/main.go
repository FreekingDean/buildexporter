package main

import (
	"fmt"
	"os"

	build "github.com/FreekingDean/buildengine"
)

func main() {
	if len(os.Args) < 3 {
		printUsage()
		os.Exit(1)
	}

	mapFile, err := os.Open(os.Args[1])
	if err != nil {
		printErr(err)
		os.Exit(1)
	}

	buildMap, err := build.DecodeMap(mapFile)
	if err != nil {
		printErr(err)
		os.Exit(1)
	}

	geom := sectorToColladaGeom(0, buildMap.Sectors()[0])
	if err != nil {
		printErr(err)
		os.Exit(1)
	}

	collada := newCollada()
	collada.LibraryGeometries = append(collada.LibraryGeometries, geom)
	collada.Export(os.Args[2])
}

func printErr(err error) {
	fmt.Printf("ERROR: %s\n", err.Error())
}

func printUsage() {
	fmt.Println(`usage: build-2-collada ./INPUT.MAP ./output.dae`)
}
