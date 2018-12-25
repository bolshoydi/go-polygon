package main

import (
	"bufio"
	"encoding/csv"
	"github.com/labstack/gommon/log"
	"io"
	"os"
	"path/filepath"
)

func ReadCsvToWKT(input string, hasHeader bool) TrashWKTList {
	var res TrashWKTList

	in, _ := os.Open(input)
	reader := csv.NewReader(bufio.NewReader(in))
	reader.Comma = ';'
	// read header
	if hasHeader {
		_, err := reader.Read()
		if err != nil {
			log.Error(err)
		}
	}

	for {
		rec, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Error(err)
		}
		res = append(res, TrashWKT{
			id:       rec[0],
			wkt_geom: rec[1],
			name:     rec[2],
			lat:      rec[3],
			lon:      rec[4],
		})
	}
	return res
}

func ReadCsvOnlyWKT(input string, wkt_idx int, hasHeader bool) WKTList {
	absPath,_ := filepath.Abs(input)
	in, _ := os.Open(absPath)
	reader := csv.NewReader(bufio.NewReader(in))
	reader.Comma = ';'
	// read header
	if hasHeader {
		_, err := reader.Read()
		if err != nil {
			log.Error(err)
		}
	}

	var res WKTList
	for {
		rec, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Error(err)
		}
		res = append(res, WKT{
			wkt_geom: rec[wkt_idx],
		})
	}
	return res
}
