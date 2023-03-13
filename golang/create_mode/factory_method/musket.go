package factory_method

type Musket struct {
	Gun
}

// TODO 报错
func NewMusket() Gun {
	return &Musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}
