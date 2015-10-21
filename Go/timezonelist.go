package main

import "fmt"

func main() {
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
			fmt.Printf("%s,%s,%+d:%02d\n", loc.name, loc.zone[j].name, h, m)
		}
	}
}
