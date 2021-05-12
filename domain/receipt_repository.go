package domain

// ReciptRepository Reciptモデルのリポジトリ
type ReciptRepository interface {
	AddRecipt(newRecipt *ReciptModel) (*ReciptModel, error)
	UpdateRecipt(newRecipt *ReciptModel) error
	GetRecipt(name string) (*ReciptModel, error)
	DeleteRecipt(name string) error
}
