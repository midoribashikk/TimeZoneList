package main

import (
	"fmt"
	"time"
)

const oneYear = 60 * 60 * 24 * 365

var aYearAgo int64

func main() {
	aYearAgo = time.Now().Unix() - oneYear

	list, err := loadLocation()
	if err == nil {
		printAll(list)
	}
}

func printAll(list []Location) {
	for i := range list {
		loc := list[i]

		for j := range loc.zone {
			neg := false
			offset := loc.zone[j].offset
			offset /= 60
			if offset < 0 {
				offset *= -1
				neg = true
			}
			h := offset / 60
			m := offset % 60
			if neg {
				h *= -1
			}

			if judgeRecent(loc, uint8(j)) {
				fmt.Printf("%s,%s,%+d:%02d\n", loc.name, loc.zone[j].name, h, m)
			}
		}
	}
}

func judgeRecent(loc Location, idx uint8) bool {
	return true
	/*
		for i := range loc.tx {
			if loc.tx[i].when > aYearAgo && loc.tx[i].index == idx {
				return true
			} else if &loc.zone[idx] == loc.cacheZone {
				return true
			}
		}

		return false
	*/
}
