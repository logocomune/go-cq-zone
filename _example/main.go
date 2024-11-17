package main

import (
	"fmt"
	"github.com/logocomune/go-cq-zone"
)

func main() {
	cqZone, found := cqzone.GetZoneByCoordinate(cqzone.Coordinate{Lat: 43.71, Lng: 11.75})

	if !found {
		fmt.Println("Zone not found")
	}
	fmt.Println("CQ Zone: ", cqZone.Number)
	fmt.Printf("CQ Zone Name: %+v\n", cqZone.Name)
	fmt.Printf("CQ Zone Center: %+v\n", cqZone.Center)
	fmt.Printf("CQ Zone GeoJSON: %+v\n", cqZone.GeoJSON)
}
