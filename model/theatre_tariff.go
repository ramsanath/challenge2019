package model

type TheatreTariff struct {
	Theatre Theatre
	Tariffs []Tariff
}

func (tt *TheatreTariff) AddTariff(t Tariff) {
	tt.Tariffs = append(tt.Tariffs, t)
}