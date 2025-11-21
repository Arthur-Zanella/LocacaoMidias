-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema locacao_midias
-- -----------------------------------------------------
DROP SCHEMA IF EXISTS `locacao_midias` ;

-- -----------------------------------------------------
-- Schema locacao_midias
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `locacao_midias` DEFAULT CHARACTER SET latin1 ;
USE `locacao_midias` ;

-- -----------------------------------------------------
-- Table `locacao_midias`.`estado`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `locacao_midias`.`estado` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `nome` VARCHAR(30) NOT NULL,
  `sigla` VARCHAR(2) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `sigla_UNIQUE` (`sigla` ASC))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `locacao_midias`.`cidade`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `locacao_midias`.`cidade` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `nome` VARCHAR(30) NOT NULL,
  `estado_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_cidade_estado_idx` (`estado_id` ASC),
  CONSTRAINT `fk_cidade_estado`
    FOREIGN KEY (`estado_id`)
    REFERENCES `locacao_midias`.`estado` (`id`)
    ON DELETE RESTRICT
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `locacao_midias`.`cliente`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `locacao_midias`.`cliente` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `nome` VARCHAR(45) NOT NULL,
  `sobrenome` VARCHAR(45) NOT NULL,
  `data_nascimento` DATE NOT NULL,
  `cpf` VARCHAR(14) NOT NULL,
  `email` VARCHAR(60) NOT NULL,
  `logradouro` VARCHAR(50) NOT NULL,
  `numero` VARCHAR(6) NOT NULL,
  `bairro` VARCHAR(30) NOT NULL,
  `cep` VARCHAR(9) NOT NULL,
  `cidade_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `cpf_UNIQUE` (`cpf` ASC),
  INDEX `fk_cliente_cidade_idx` (`cidade_id` ASC),
  CONSTRAINT `fk_cliente_cidade`
    FOREIGN KEY (`cidade_id`)
    REFERENCES `locacao_midias`.`cidade` (`id`)
    ON DELETE RESTRICT
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `locacao_midias`.`ator`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `locacao_midias`.`ator` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `nome` VARCHAR(45) NOT NULL,
  `sobrenome` VARCHAR(45) NOT NULL,
  `data_estreia` DATE NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `locacao_midias`.`genero`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `locacao_midias`.`genero` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `descricao` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `locacao_midias`.`classificacao_etaria`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `locacao_midias`.`classificacao_etaria` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `descricao` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `locacao_midias`.`tipo`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `locacao_midias`.`tipo` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `descricao` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `locacao_midias`.`classificacao_interna`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `locacao_midias`.`classificacao_interna` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `descricao` VARCHAR(45) NOT NULL,
  `valor_aluguel` DECIMAL(15,2) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `locacao_midias`.`midia`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `locacao_midias`.`midia` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `titulo` VARCHAR(100) NOT NULL,
  `ano_lancamento` VARCHAR(100) NOT NULL,
  `codigo_barras` VARCHAR(13) NOT NULL,
  `duracao_em_minutos` INT NOT NULL,
  `ator_principal` INT NOT NULL,
  `ator_coadjuvante` INT NOT NULL,
  `genero_id` INT NOT NULL,
  `classificacao_etaria_id` INT NOT NULL,
  `tipo_id` INT NOT NULL,
  `classificacao_interna_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `codigoBarras_UNIQUE` (`codigo_barras` ASC),
  INDEX `fk_midia_ator_principal_idx` (`ator_principal` ASC),
  INDEX `fk_midia_ator_coadjuvante_idx` (`ator_coadjuvante` ASC),
  INDEX `fk_midia_genero_idx` (`genero_id` ASC),
  INDEX `fk_midia_classificacao_etaria_idx` (`classificacao_etaria_id` ASC),
  INDEX `fk_midia_tipo_idx` (`tipo_id` ASC),
  INDEX `fk_midia_classificacao_interna1_idx` (`classificacao_interna_id` ASC),
  CONSTRAINT `fk_midia_ator_principal`
    FOREIGN KEY (`ator_principal`)
    REFERENCES `locacao_midias`.`ator` (`id`)
    ON DELETE RESTRICT
    ON UPDATE CASCADE,
  CONSTRAINT `fk_midia_ator_coadjuvante`
    FOREIGN KEY (`ator_coadjuvante`)
    REFERENCES `locacao_midias`.`ator` (`id`)
    ON DELETE RESTRICT
    ON UPDATE CASCADE,
  CONSTRAINT `fk_midia_genero`
    FOREIGN KEY (`genero_id`)
    REFERENCES `locacao_midias`.`genero` (`id`)
    ON DELETE RESTRICT
    ON UPDATE CASCADE,
  CONSTRAINT `fk_midia_classificacao_etaria`
    FOREIGN KEY (`classificacao_etaria_id`)
    REFERENCES `locacao_midias`.`classificacao_etaria` (`id`)
    ON DELETE RESTRICT
    ON UPDATE CASCADE,
  CONSTRAINT `fk_midia_tipo`
    FOREIGN KEY (`tipo_id`)
    REFERENCES `locacao_midias`.`tipo` (`id`)
    ON DELETE RESTRICT
    ON UPDATE CASCADE,
  CONSTRAINT `fk_midia_classificacao_interna`
    FOREIGN KEY (`classificacao_interna_id`)
    REFERENCES `locacao_midias`.`classificacao_interna` (`id`)
    ON DELETE RESTRICT
    ON UPDATE CASCADE)
ENGINE = InnoDB
COMMENT = '		';


-- -----------------------------------------------------
-- Table `locacao_midias`.`locacao`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `locacao_midias`.`locacao` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `data_inicio` DATE NOT NULL,
  `data_fim` DATE NOT NULL,
  `cancelada` TINYINT(1) NOT NULL,
  `cliente_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_venda_cliente_idx` (`cliente_id` ASC),
  CONSTRAINT `fk_venda_cliente`
    FOREIGN KEY (`cliente_id`)
    REFERENCES `locacao_midias`.`cliente` (`id`)
    ON DELETE RESTRICT
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `locacao_midias`.`exemplar`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `locacao_midias`.`exemplar` (
  `codigo_interno` INT NOT NULL AUTO_INCREMENT,
  `disponivel` TINYINT(1) NOT NULL,
  `midia_id` INT NOT NULL,
  PRIMARY KEY (`codigo_interno`),
  INDEX `fk_exemplar_midia_idx` (`midia_id` ASC),
  CONSTRAINT `fk_exemplar_midia`
    FOREIGN KEY (`midia_id`)
    REFERENCES `locacao_midias`.`midia` (`id`)
    ON DELETE RESTRICT
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `locacao_midias`.`item_locacao`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `locacao_midias`.`item_locacao` (
  `locacao_id` INT NOT NULL,
  `exemplar_codigo_interno` INT NOT NULL,
  `valor` DECIMAL(15,2) NOT NULL,
  PRIMARY KEY (`locacao_id`, `exemplar_codigo_interno`),
  INDEX `fk_item_locacao_exemplar_idx` (`exemplar_codigo_interno` ASC),
  INDEX `fk_item_locacao_locacao_idx` (`locacao_id` ASC),
  CONSTRAINT `fk_item_locacao_locacao`
    FOREIGN KEY (`locacao_id`)
    REFERENCES `locacao_midias`.`locacao` (`id`)
    ON DELETE RESTRICT
    ON UPDATE CASCADE,
  CONSTRAINT `fk_item_locacao_exemplar`
    FOREIGN KEY (`exemplar_codigo_interno`)
    REFERENCES `locacao_midias`.`exemplar` (`codigo_interno`)
    ON DELETE RESTRICT
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Data for table `locacao_midias`.`estado`
-- -----------------------------------------------------
START TRANSACTION;
USE `locacao_midias`;
INSERT INTO `locacao_midias`.`estado` (`id`, `nome`, `sigla`) VALUES (1, 'São Paulo', 'SP');
INSERT INTO `locacao_midias`.`estado` (`id`, `nome`, `sigla`) VALUES (2, 'Minas Gerais', 'MG');

COMMIT;


-- -----------------------------------------------------
-- Data for table `locacao_midias`.`cidade`
-- -----------------------------------------------------
START TRANSACTION;
USE `locacao_midias`;
INSERT INTO `locacao_midias`.`cidade` (`id`, `nome`, `estado_id`) VALUES (1, 'Vargem Grande do Sul', 1);
INSERT INTO `locacao_midias`.`cidade` (`id`, `nome`, `estado_id`) VALUES (2, 'São João da Boa Vista', 1);
INSERT INTO `locacao_midias`.`cidade` (`id`, `nome`, `estado_id`) VALUES (3, 'Poços de Caldas', 2);

COMMIT;


-- -----------------------------------------------------
-- Data for table `locacao_midias`.`cliente`
-- -----------------------------------------------------
START TRANSACTION;
USE `locacao_midias`;
INSERT INTO `locacao_midias`.`cliente` (`id`, `nome`, `sobrenome`, `data_nascimento`, `cpf`, `email`, `logradouro`, `numero`, `bairro`, `cep`, `cidade_id`) VALUES (1, 'João', 'da Silva', '1999-10-10', '111.111.111-11', 'joao@joao.com.br', 'Rua Hermenegildo Cossi', '750', 'Jardim Fortaleza', '13880-000', 1);
INSERT INTO `locacao_midias`.`cliente` (`id`, `nome`, `sobrenome`, `data_nascimento`, `cpf`, `email`, `logradouro`, `numero`, `bairro`, `cep`, `cidade_id`) VALUES (2, 'Maria', 'Rodrigues', '1974-05-16', '222.222.222-22', 'mariarod@gmail.com', 'Rua Patrocínio Rodrigues', '120', 'Centro', '13880-000', 1);
INSERT INTO `locacao_midias`.`cliente` (`id`, `nome`, `sobrenome`, `data_nascimento`, `cpf`, `email`, `logradouro`, `numero`, `bairro`, `cep`, `cidade_id`) VALUES (3, 'Marcela', 'dos Santos', '1985-09-25', '333.333.333-33', 'mdossantos@uol.com.br', 'Rua Primeirod de Maio', '219', 'Centro', '13880-000', 1);

COMMIT;


-- -----------------------------------------------------
-- Data for table `locacao_midias`.`ator`
-- -----------------------------------------------------
START TRANSACTION;
USE `locacao_midias`;
INSERT INTO `locacao_midias`.`ator` (`id`, `nome`, `sobrenome`, `data_estreia`) VALUES (1, 'Keanu', 'Reeves', '1980-10-10');
INSERT INTO `locacao_midias`.`ator` (`id`, `nome`, `sobrenome`, `data_estreia`) VALUES (2, 'Laurence', 'Fishburne', '1970-10-10');
INSERT INTO `locacao_midias`.`ator` (`id`, `nome`, `sobrenome`, `data_estreia`) VALUES (3, 'Daniel', 'Radcliffe', '1999-10-10');
INSERT INTO `locacao_midias`.`ator` (`id`, `nome`, `sobrenome`, `data_estreia`) VALUES (4, 'Emma', 'Watson', '2000-10-10');

COMMIT;


-- -----------------------------------------------------
-- Data for table `locacao_midias`.`genero`
-- -----------------------------------------------------
START TRANSACTION;
USE `locacao_midias`;
INSERT INTO `locacao_midias`.`genero` (`id`, `descricao`) VALUES (1, 'Comédia');
INSERT INTO `locacao_midias`.`genero` (`id`, `descricao`) VALUES (2, 'Aventura');
INSERT INTO `locacao_midias`.`genero` (`id`, `descricao`) VALUES (3, 'Ação');

COMMIT;


-- -----------------------------------------------------
-- Data for table `locacao_midias`.`classificacao_etaria`
-- -----------------------------------------------------
START TRANSACTION;
USE `locacao_midias`;
INSERT INTO `locacao_midias`.`classificacao_etaria` (`id`, `descricao`) VALUES (1, 'Infantil');
INSERT INTO `locacao_midias`.`classificacao_etaria` (`id`, `descricao`) VALUES (2, 'Maior que 18');

COMMIT;


-- -----------------------------------------------------
-- Data for table `locacao_midias`.`tipo`
-- -----------------------------------------------------
START TRANSACTION;
USE `locacao_midias`;
INSERT INTO `locacao_midias`.`tipo` (`id`, `descricao`) VALUES (1, 'BluRay');
INSERT INTO `locacao_midias`.`tipo` (`id`, `descricao`) VALUES (2, 'DVD');

COMMIT;


-- -----------------------------------------------------
-- Data for table `locacao_midias`.`classificacao_interna`
-- -----------------------------------------------------
START TRANSACTION;
USE `locacao_midias`;
INSERT INTO `locacao_midias`.`classificacao_interna` (`id`, `descricao`, `valor_aluguel`) VALUES (1, 'Lançamento', 10.00);
INSERT INTO `locacao_midias`.`classificacao_interna` (`id`, `descricao`, `valor_aluguel`) VALUES (2, 'Padrão', 5.50);
INSERT INTO `locacao_midias`.`classificacao_interna` (`id`, `descricao`, `valor_aluguel`) VALUES (3, 'Promoção', 4.00);

COMMIT;


-- -----------------------------------------------------
-- Data for table `locacao_midias`.`midia`
-- -----------------------------------------------------
START TRANSACTION;
USE `locacao_midias`;
INSERT INTO `locacao_midias`.`midia` (`id`, `titulo`, `ano_lancamento`, `codigo_barras`, `duracao_em_minutos`, `ator_principal`, `ator_coadjuvante`, `genero_id`, `classificacao_etaria_id`, `tipo_id`, `classificacao_interna_id`) VALUES (1, 'Harry Potter e a Pedra Filosofal', '2001', '1231231231231', 125, 3, 4, 2, 1, 1, 2);
INSERT INTO `locacao_midias`.`midia` (`id`, `titulo`, `ano_lancamento`, `codigo_barras`, `duracao_em_minutos`, `ator_principal`, `ator_coadjuvante`, `genero_id`, `classificacao_etaria_id`, `tipo_id`, `classificacao_interna_id`) VALUES (2, 'Matrix', '1999', '4564564564564', 136, 1, 2, 3, 2, 2, 3);

COMMIT;


-- -----------------------------------------------------
-- Data for table `locacao_midias`.`exemplar`
-- -----------------------------------------------------
START TRANSACTION;
USE `locacao_midias`;
INSERT INTO `locacao_midias`.`exemplar` (`codigo_interno`, `disponivel`, `midia_id`) VALUES (1, 1, 1);
INSERT INTO `locacao_midias`.`exemplar` (`codigo_interno`, `disponivel`, `midia_id`) VALUES (2, 1, 1);
INSERT INTO `locacao_midias`.`exemplar` (`codigo_interno`, `disponivel`, `midia_id`) VALUES (3, 1, 1);
INSERT INTO `locacao_midias`.`exemplar` (`codigo_interno`, `disponivel`, `midia_id`) VALUES (4, 1, 1);
INSERT INTO `locacao_midias`.`exemplar` (`codigo_interno`, `disponivel`, `midia_id`) VALUES (5, 1, 1);
INSERT INTO `locacao_midias`.`exemplar` (`codigo_interno`, `disponivel`, `midia_id`) VALUES (6, 1, 1);
INSERT INTO `locacao_midias`.`exemplar` (`codigo_interno`, `disponivel`, `midia_id`) VALUES (7, 1, 2);
INSERT INTO `locacao_midias`.`exemplar` (`codigo_interno`, `disponivel`, `midia_id`) VALUES (8, 1, 2);
INSERT INTO `locacao_midias`.`exemplar` (`codigo_interno`, `disponivel`, `midia_id`) VALUES (9, 1, 2);
INSERT INTO `locacao_midias`.`exemplar` (`codigo_interno`, `disponivel`, `midia_id`) VALUES (10, 1, 2);

COMMIT;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
