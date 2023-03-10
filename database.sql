CREATE TABLE  (
    cpf VARCHAR(20) PRIMARY KEY,
    private BOOLEAN,
    incompleto BOOLEAN,
    ultima_compra DATE,
    ticket_medio NUMERIC(13,3),
    ticket_ult_compra NUMERIC(13,3),
    loja_frequente VARCHAR(100),
    loja_ult_compra VARCHAR(100)
);