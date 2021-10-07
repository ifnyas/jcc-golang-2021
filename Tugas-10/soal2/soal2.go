package soal2

import "strconv"

type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

type Description interface {
	ShowDescription() string
}

func (p Phone) ShowDescription() (desc string) {
	var colors string

	for index, item := range p.Colors {
		if index == 0 {
			colors += item
		} else {
			colors += ", " + item
		}
	}

	desc = "name : " + p.Name + "\n" +
		"brand : " + p.Brand + "\n" +
		"year : " + strconv.Itoa(p.Year) + "\n" +
		"color: " + colors
	return
}
