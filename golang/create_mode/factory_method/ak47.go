package factory_method

type Ak47 struct {
	Gun
}

// TODO 报错
func NewAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}
