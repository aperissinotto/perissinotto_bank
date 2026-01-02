CREATE TABLE contas (
    agencia NUMERIC(4,0) NOT NULL,
    conta NUMERIC(10,0) NOT NULL,

    cliente_id UUID NOT NULL,
    senha VARCHAR(255) NOT NULL,

    CONSTRAINT pk_contas PRIMARY KEY (agencia, conta),

    CONSTRAINT fk_contas_cliente
        FOREIGN KEY (cliente_id)
        REFERENCES clientes(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_contas_cliente ON contas(cliente_id);