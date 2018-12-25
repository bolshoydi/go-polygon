package main

import (
	"github.com/labstack/gommon/log"
	"github.com/paulsmith/gogeos/geos"
)

const CSV_TRASH_PATH = `trash_polygons_MO.csv`
const CSV_TEST_POINT = `trash_test_points.csv`
const WKT_IDX = 1

func main() {

	trashWKTList := ReadCsvOnlyWKT(CSV_TRASH_PATH, WKT_IDX, true)
	geosTrashList, err := trashWKTList.GetGeosGeom()
	if err != nil {
		log.Error(err)
	}

	PointsWKTList := ReadCsvOnlyWKT(CSV_TEST_POINT, WKT_IDX, true)
	geosPointList, err := PointsWKTList.GetGeosGeom()
	if err != nil {
		log.Error(err)
	}
	for _, point := range geosPointList {
		log.Info(point)
		da, err := geosTrashList.IsIntersect(point)
		if err != nil {
			log.Error(err)
		} else {
			log.Info(da)
		}
	}

	s, err := geos.NewPoint(geos.Coord{X: 37.49841795020783763, Y: 55.90410272123053659})
	if err != nil {
		log.Error(err)
	} else {
		log.Info(s)
	}

}
