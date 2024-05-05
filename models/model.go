package models

type ToDoList struct {
	Id         int64    `json: "id"`
	Title      string   `json: "title"`
	List       []string `json: "list"`
	StatusList Status   `json: "status"`
}

type Status int64

const (
	EmAndamento = iota
	Concluido
	Pendente
)

func (s Status) String() string {
	switch s {
	case EmAndamento:
		return "Em Andamento"
	case Concluido:
		return "Concluido"
	case Pendente:
		return "Pendente"
	}

	return "Status inválido"
}
