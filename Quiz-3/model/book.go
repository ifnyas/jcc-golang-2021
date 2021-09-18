package model

import (
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	layoutDateTime = "2006-01-02 15:04:05"
)

type Book struct {
	ID                 int       `json:"id"`
	Title              string    `json:"title"`
	Description        string    `json:"description"`
	Image_url          string    `json:"image_url"`
	Release_year       int       `json:"release_year"`
	Price              string    `json:"price"`
	Total_page         string    `json:"total_page"`
	Kategori_ketebalan string    `json:"kategori_ketebalan"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

func IsImageUrlValid(s string) bool {
	return strings.HasPrefix(s, "http")
}

func IsReleaseYearValid(year int) bool {
	isValid := false
	if year > 1979 && year < 2022 {
		isValid = true
	}
	return isValid
}

func GetPriceWithCurrency(price int) string {
	priceStr := strconv.Itoa(price)
	lenPriceStr := len(priceStr)
	sepAmount := lenPriceStr / 3
	sepPos := []int{}

	for i := 0; i < sepAmount; i++ {
		pos := lenPriceStr - (3 * (i + 1))
		if pos > 0 {
			sepPos = append(sepPos, pos)
		}
	}
	sort.Ints(sepPos)

	for j, k := range sepPos {
		priceStr = priceStr[:j+k] + "." + priceStr[j+k:]
	}

	priceStr = "Rp. " + priceStr + ",-"
	return priceStr
}

func GetKategoriKetebalan(n int) string {
	result := ""
	switch {
	case n >= 201:
		result = "Tebal"
	case n >= 101:
		result = "Sedang"
	default:
		result = "Tipis"
	}
	return result
}
