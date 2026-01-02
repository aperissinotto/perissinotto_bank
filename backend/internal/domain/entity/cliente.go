package entity

type Cliente struct {
	ID             string
	NomeCompleto   string
	Email          string
	DataNascimento string
	CPF            string
	RG             string
	CEP            string
	Endereco       string
	Bairro         string
	Cidade         string
	Estado         string
	RendaMensal    float64
}
