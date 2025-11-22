package models

import "time"

//estado
type Estado struct {
	ID    int    `json:"id" gorm:"primaryKey;column:id"`
	Nome  string `json:"nome" gorm:"column:nome"`
	Sigla string `json:"sigla" gorm:"column:sigla"`
}

func (Estado) TableName() string {
	return "estado"
}

//cidade
type Cidade struct {
	ID       int    `json:"id" gorm:"primaryKey;column:id"`
	Nome     string `json:"nome" gorm:"column:nome"`
	EstadoID int    `json:"estado_id" gorm:"column:estado_id"`
	Estado   Estado `json:"estado,omitempty" gorm:"foreignKey:EstadoID"`
}

func (Cidade) TableName() string {
	return "cidade"
}

//cliente
type Cliente struct {
	ID             int       `json:"id" gorm:"primaryKey;column:id"`
	Nome           string    `json:"nome" gorm:"column:nome"`
	Sobrenome      string    `json:"sobrenome" gorm:"column:sobrenome"`
	DataNascimento time.Time `json:"data_nascimento" gorm:"column:data_nascimento"`
	CPF            string    `json:"cpf" gorm:"column:cpf"`
	Email          string    `json:"email" gorm:"column:email"`
	Logradouro     string    `json:"logradouro" gorm:"column:logradouro"`
	Numero         string    `json:"numero" gorm:"column:numero"`
	Bairro         string    `json:"bairro" gorm:"column:bairro"`
	CEP            string    `json:"cep" gorm:"column:cep"`
	CidadeID       int       `json:"cidade_id" gorm:"column:cidade_id"`
	Cidade         Cidade    `json:"cidade,omitempty" gorm:"foreignKey:CidadeID"`
}

func (Cliente) TableName() string {
	return "cliente"
}

//ator
type Ator struct {
	ID          int       `json:"id" gorm:"primaryKey;column:id"`
	Nome        string    `json:"nome" gorm:"column:nome"`
	Sobrenome   string    `json:"sobrenome" gorm:"column:sobrenome"`
	DataEstreia time.Time `json:"data_estreia" gorm:"column:data_estreia"`
}

func (Ator) TableName() string {
	return "ator"
}

//genero
type Genero struct {
	ID        int    `json:"id" gorm:"primaryKey;column:id"`
	Descricao string `json:"descricao" gorm:"column:descricao"`
}

func (Genero) TableName() string {
	return "genero"
}

//classificacao etaria
type ClassificacaoEtaria struct {
	ID        int    `json:"id" gorm:"primaryKey;column:id"`
	Descricao string `json:"descricao" gorm:"column:descricao"`
}

func (ClassificacaoEtaria) TableName() string {
	return "classificacao_etaria"
}

//tipo
type Tipo struct {
	ID        int    `json:"id" gorm:"primaryKey;column:id"`
	Descricao string `json:"descricao" gorm:"column:descricao"`
}

func (Tipo) TableName() string {
	return "tipo"
}

//classificacao interna
type ClassificacaoInterna struct {
	ID           int     `json:"id" gorm:"primaryKey;column:id"`
	Descricao    string  `json:"descricao" gorm:"column:descricao"`
	ValorAluguel float64 `json:"valor_aluguel" gorm:"column:valor_aluguel"`
}

func (ClassificacaoInterna) TableName() string {
	return "classificacao_interna"
}

//midia
type Midia struct {
	ID                     int    `json:"id" gorm:"primaryKey;column:id"`
	Titulo                 string `json:"titulo" gorm:"column:titulo"`
	AnoLancamento          string `json:"ano_lancamento" gorm:"column:ano_lancamento"`
	CodigoBarras           string `json:"codigo_barras" gorm:"column:codigo_barras"`
	DuracaoEmMinutos       int    `json:"duracao_em_minutos" gorm:"column:duracao_em_minutos"`
	AtorPrincipalID        int    `json:"ator_principal" gorm:"column:ator_principal"`
	AtorCoadjuvanteID      int    `json:"ator_coadjuvante" gorm:"column:ator_coadjuvante"`
	GeneroID               int    `json:"genero_id" gorm:"column:genero_id"`
	ClassificacaoEtariaID  int    `json:"classificacao_etaria_id" gorm:"column:classificacao_etaria_id"`
	TipoID                 int    `json:"tipo_id" gorm:"column:tipo_id"`
	ClassificacaoInternaID int    `json:"classificacao_interna_id" gorm:"column:classificacao_interna_id"`

	//relacionamento
	Genero               Genero               `json:"genero,omitempty" gorm:"foreignKey:GeneroID"`
	ClassificacaoEtaria  ClassificacaoEtaria  `json:"classificacao_etaria,omitempty" gorm:"foreignKey:ClassificacaoEtariaID"`
	Tipo                 Tipo                 `json:"tipo,omitempty" gorm:"foreignKey:TipoID"`
	ClassificacaoInterna ClassificacaoInterna `json:"classificacao_interna,omitempty" gorm:"foreignKey:ClassificacaoInternaID"`
	AtorPrincipal        Ator                 `json:"ator_principal_obj,omitempty" gorm:"foreignKey:AtorPrincipalID"`
	AtorCoadjuvante      Ator                 `json:"ator_coadjuvante_obj,omitempty" gorm:"foreignKey:AtorCoadjuvanteID"`
}

func (Midia) TableName() string {
	return "midia"
}

//exemplar
type Exemplar struct {
	CodigoInterno int   `json:"codigo_interno" gorm:"primaryKey;column:codigo_interno"`
	Disponivel    bool  `json:"disponivel" gorm:"column:disponivel"`
	MidiaID       int   `json:"midia_id" gorm:"column:midia_id"`
	Midia         Midia `json:"midia,omitempty" gorm:"foreignKey:MidiaID"`
}

func (Exemplar) TableName() string {
	return "exemplar"
}

//locacao
type Locacao struct {
	ID         int           `json:"id" gorm:"primaryKey;column:id"`
	DataInicio time.Time     `json:"data_inicio" gorm:"column:data_inicio"`
	DataFim    time.Time     `json:"data_fim" gorm:"column:data_fim"`
	Cancelada  bool          `json:"cancelada" gorm:"column:cancelada"`
	ClienteID  int           `json:"cliente_id" gorm:"column:cliente_id"`
	Cliente    Cliente       `json:"cliente,omitempty" gorm:"foreignKey:ClienteID"`
	Itens      []ItemLocacao `json:"itens,omitempty" gorm:"foreignKey:LocacaoID"`
}

func (Locacao) TableName() string {
	return "locacao"
}

//itemlocacao
type ItemLocacao struct {
	LocacaoID      int      `json:"locacao_id" gorm:"primaryKey;column:locacao_id"`
	ExemplarCodigo int      `json:"exemplar_codigo_interno" gorm:"primaryKey;column:exemplar_codigo_interno"`
	Valor          float64  `json:"valor" gorm:"column:valor"`
	Locacao        Locacao  `json:"locacao,omitempty" gorm:"foreignKey:LocacaoID"`
	Exemplar       Exemplar `json:"exemplar,omitempty" gorm:"foreignKey:ExemplarCodigo;references:CodigoInterno"`
}

func (ItemLocacao) TableName() string {
	return "item_locacao"
}
