package model

// Partner ...
type Partner struct {
	ID             string
	Capacity       int
	TheatreTariffs []TheatreTariff
}

func (p *Partner) AddTheatreTariff(theatre Theatre, tariff Tariff) {
	if p.TheatreTariffs == nil {
		p.TheatreTariffs = make([]TheatreTariff, 0)
	}

	for i := 0; i < len(p.TheatreTariffs); i++ {
		if p.TheatreTariffs[i].Theatre == theatre {
			p.TheatreTariffs[i].AddTariff(tariff)
			return
		}
	}

	tt := TheatreTariff{
		Theatre: theatre,
		Tariffs: []Tariff{tariff},
	}
	p.TheatreTariffs = append(p.TheatreTariffs, tt)
}
