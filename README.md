# Geos in go 
_______________________ 

Small project that shows how to work with [goes](https://github.com/libgeos/geos) library in golang.

## Build

```bash
$ go build
```

## Main

* Main file loads 2 csv files (Polygon and Points) in corresponding structures.

* Then check if point intersects polygon and log the result.

## Utils

* Some utils to work with csv files that have WKT geometry field.

## Types
* Structs for loading csv files

* To run tests on types:

```bash
$ go test
```