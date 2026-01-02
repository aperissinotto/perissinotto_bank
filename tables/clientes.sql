CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE clientes (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,

    nome_completo VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    data_nascimento DATE NOT NULL,

    cpf CHAR(11) NOT NULL UNIQUE,
    rg VARCHAR(20) NOT NULL,

    cep CHAR(8) NOT NULL,
    endereco VARCHAR(100) NOT NULL,
    bairro VARCHAR(60) NOT NULL,
    cidade VARCHAR(60) NOT NULL,
    estado CHAR(2) NOT NULL,

    renda_mensal NUMERIC(12,2)
);
