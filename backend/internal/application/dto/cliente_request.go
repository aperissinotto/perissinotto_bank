package dto

type CriarClienteRequest struct {
	NomeCompleto   string  `json:"nomeCompleto"`
	Email          string  `json:"email"`
	DataNascimento string  `json:"dataNascimento"`
	CPF            string  `json:"cpf"`
	RG             string  `json:"rg"`
	CEP            string  `json:"cep"`
	Endereco       string  `json:"endereco"`
	Bairro         string  `json:"bairro"`
	Cidade         string  `json:"cidade"`
	Estado         string  `json:"estado"`
	RendaMensal    float64 `json:"rendaMensal"`
	Senha          string  `json:"senha"`
}
