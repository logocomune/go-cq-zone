package cqzone

type CqZone struct {
	Number  int
	Name    string
	Center  Coordinate
	GeoJSON GeoJSON
}

func GetZoneByCoordinate(c Coordinate) (CqZone, bool) {
	if c.Lat < -90 || c.Lat > 90 || c.Lng < -180 || c.Lng > 180 {
		return CqZone{}, false
	}

	for i := range zones {
		if zones[i].polygon.isPointInZone(c) {
			return zoneToCqZone(zones[i]), true
		}
	}

	return CqZone{}, false
}

func GetZoneByNumber(id int) (CqZone, bool) {
	if id < 1 || id > len(zones) {
		return CqZone{}, false
	}

	return zoneToCqZone(zones[id]), true

}

func GetGoeJson() FeatureCollection {
	return zonesToGeoJSON(zones)
}

func zoneToCqZone(z zone) CqZone {
	return CqZone{
		Number:  z.number,
		Name:    z.name,
		Center:  z.center,
		GeoJSON: zoneToGeoJSON(z),
	}
}
