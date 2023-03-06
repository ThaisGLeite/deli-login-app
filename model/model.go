package model

type User struct {
	Nome           string `json:"nome"`
	CPF            string `json:"cpf"`
	NomeCompleto   string `json:"nome-completo"`
	DataNascimento string `json:"data-nascimento"`
	Senha          string `json:"senha"`
	DataCriacao    string `json:"data-criacao"`
}
