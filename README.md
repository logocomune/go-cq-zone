# Go CQ Zone Info Package

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/logocomune/go-cq-zone)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/logocomune/go-cq-zone/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/logocomune/go-cq-zone)](https://pkg.go.dev/github.com/logocomune/go-cq-zone)
[![codecov](https://codecov.io/gh/logocomune/go-cq-zone/graph/badge.svg?token=Mhxw3pW0z0)](https://codecov.io/gh/logocomune/go-cq-zone)
[![Go Report Card](https://goreportcard.com/badge/github.com/logocomune/go-cq-zone)](https://goreportcard.com/report/github.com/logocomune/go-cq-zone)


This repository contains a Go package to retrieve information about CQ Zones (used in amateur radio) using geographic coordinates (latitude and longitude)
or the CQ Zone number.

## Features

**Search by coordinates**: Get the CQ Zone corresponding to specific geographic coordinates.

**Search by CQ Zone number**: Retrieve details about an CQ Zone by its identifier number.

## Installation

```console
go get  github.com/logocomune/go-cq-zone
```

## Usage

```golang
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

```

## Acknowledgments

Special thanks to @HB9HIL for providing the GeoJSON files of CQ Zones via their repository: hamradio-zones-geojson.

## License

This project is distributed under the MIT License. See the LICENSE file for more details.
