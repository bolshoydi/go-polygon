package main

import (
	"github.com/labstack/gommon/log"
	"github.com/paulsmith/gogeos/geos"
)

type TrashWKT struct {
	id       string
	wkt_geom string
	name     string
	lat      string
	lon      string
}
type TrashWKTList []TrashWKT

type WKT struct {
	wkt_geom string
}

type WKTList []WKT

type GeosGeomList []*geos.Geometry

func (wl WKTList) GetGeosGeom() (GeosGeomList, error) {
	var (geosList GeosGeomList
	err error = nil)

	for _, trashWKT := range wl {
		geom, err := geos.FromWKT(trashWKT.wkt_geom)
		if err != nil {
			log.Error(err)
			return geosList, err
		}
		geosList = append(geosList, geom)
		//fmt.Println(geom.Type())
	}
	return geosList, err
}

func (gl GeosGeomList) IsIntersect(point *geos.Geometry) (bool, error) {
	for _, geom := range gl {
		intersect, err := point.Intersects(geom)
		if err != nil {
			return false, err
		} else if intersect {
			return true, nil
		}
	}
	return false, nil
}
