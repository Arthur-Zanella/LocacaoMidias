package models

import "time"

// Estado
type Estado struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Sigla string `json:"sigla"`
}

// Cidade
type Cidade struct {
	ID       int    `json:"id"`
	Nome     string `json:"nome"`
	EstadoID int    `json:"estado_id"`
}

// Cliente
type Cliente struct {
	ID             int       `json:"id"`
	Nome           string    `json:"nome"`
	Sobrenome      string    `json:"sobrenome"`
	DataNascimento time.Time `json:"data_nascimento"`
	CPF            string    `json:"cpf"`
	Email          string    `json:"email"`
	Logradouro     string    `json:"logradouro"`
	Numero         string    `json:"numero"`
	Bairro         string    `json:"bairro"`
	CEP            string    `json:"cep"`
	CidadeID       int       `json:"cidade_id"`
}

// Ator

type Ator struct {
	ID          int       `json:"id"`
	Nome        string    `json:"nome"`
	Sobrenome   string    `json:"sobrenome"`
	DataEstreia time.Time `json:"data_estreia"`
}

// Genero

type Genero struct {
	ID        int    `json:"id"`
	Descricao string `json:"descricao"`
}

// ClassificacaoEtaria

type ClassificacaoEtaria struct {
	ID        int    `json:"id"`
	Descricao string `json:"descricao"`
}

// Tipo

type Tipo struct {
	ID        int    `json:"id"`
	Descricao string `json:"descricao"`
}

// ClassificacaoInterna

type ClassificacaoInterna struct {
	ID           int     `json:"id"`
	Descricao    string  `json:"descricao"`
	ValorAluguel float64 `json:"valor_aluguel"`
}

// Midia

type Midia struct {
	ID                     int    `json:"id"`
	Titulo                 string `json:"titulo"`
	AnoLancamento          string `json:"ano_lancamento"`
	CodigoBarras           string `json:"codigo_barras"`
	DuracaoEmMinutos       int    `json:"duracao_em_minutos"`
	AtorPrincipalID        int    `json:"ator_principal"`
	AtorCoadjuvanteID      int    `json:"ator_coadjuvante"`
	GeneroID               int    `json:"genero_id"`
	ClassificacaoEtariaID  int    `json:"classificacao_etaria_id"`
	TipoID                 int    `json:"tipo_id"`
	ClassificacaoInternaID int    `json:"classificacao_interna_id"`
}

// Locacao

type Locacao struct {
	ID         int       `json:"id"`
	DataInicio time.Time `json:"data_inicio"`
	DataFim    time.Time `json:"data_fim"`
	Cancelada  bool      `json:"cancelada"`
	ClienteID  int       `json:"cliente_id"`
}

// Exemplar

type Exemplar struct {
	CodigoInterno int  `json:"codigo_interno"`
	Disponivel    bool `json:"disponivel"`
	MidiaID       int  `json:"midia_id"`
}

// ItemLocacao

type ItemLocacao struct {
	LocacaoID      int     `json:"locacao_id"`
	ExemplarCodigo int     `json:"exemplar_codigo_interno"`
	Valor          float64 `json:"valor"`
}
