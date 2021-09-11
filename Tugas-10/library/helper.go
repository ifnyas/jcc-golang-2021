package library

import (
	"math"
	"strconv"
	"strings"
)

// soal 1
type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

type Tabung struct {
	JariJari, Tinggi float64
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

func (s SegitigaSamaSisi) Luas() int {
	return s.Alas * s.Tinggi / 2
}

func (s SegitigaSamaSisi) Keliling() int {
	return s.Alas * 3
}

func (s PersegiPanjang) Luas() int {
	return s.Panjang * s.Lebar
}

func (s PersegiPanjang) Keliling() int {
	return 2 * (s.Panjang + s.Lebar)
}

func (s Balok) Volume() float64 {
	return float64(s.Lebar) * float64(s.Panjang) * float64(s.Tinggi)
}

func (s Balok) LuasPermukaan() float64 {
	return (2 * (float64(s.Lebar) + float64(s.Panjang))) +
		(2 * (float64(s.Lebar) + float64(s.Tinggi))) +
		(2 * (float64(s.Tinggi) + float64(s.Panjang)))
}

func (s Tabung) Volume() float64 {
	return math.Pi * math.Pow(s.JariJari, 2) * s.Tinggi
}

func (s Tabung) LuasPermukaan() float64 {
	return 2 * math.Pi * s.JariJari * (s.JariJari + s.Tinggi)
}

type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

var IFace interface {
	GetData() string
	PrintData()
}

func (p Phone) GetData() string {
	var result string
	result += "name : " + p.Name
	result += "\nbrand : " + p.Brand
	result += "\nyear : " + strconv.Itoa(p.Year)
	result += "\ncolors : " + strings.Join(p.Colors, ", ")
	return result
}
