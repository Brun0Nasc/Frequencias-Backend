package frequencias

type IFrequencia interface {
	NovaFrequencia() error
	PegarFrequenciaMaisRecente() (*int, error)
}
