package soal1

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

func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() int {
	return 2 * (p.Panjang + p.Lebar)
}

func (b Balok) Volume() float64 {
	return float64(b.Panjang * b.Lebar * b.Tinggi)
}

func (b Balok) LuasPermukaan() float64 {
	return float64((2 * (b.Panjang + b.Lebar)) + (2 * (b.Panjang + b.Tinggi)) + (2 * (b.Tinggi + b.Lebar)))
}

func (t Tabung) Volume() float64 {
	if int(t.JariJari)%7 == 0 {
		return float64(22 / 7 * t.JariJari * t.JariJari * t.Tinggi)
	} else {
		return float64(3.14 * t.JariJari * t.JariJari * t.Tinggi)
	}
}

func (t Tabung) LuasPermukaan() float64 {
	if int(t.JariJari)%7 == 0 {
		return float64(2 * 22 / 7 * t.JariJari * (t.JariJari + t.Tinggi))
	} else {
		return float64(2 * 3.14 * t.JariJari * (t.JariJari + t.Tinggi))
	}
}
