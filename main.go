package main

import (
	"finder/scan"
	"finder/search"
)

func main() {
	lines := scan.Scan()

	search.Search(lines)
}
