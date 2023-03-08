CREATE TABLE IF NOT EXISTS clientes (
  cpf VARCHAR(11) NOT NULL PRIMARY KEY,
  private INT NOT NULL,
  incompleto INT NOT NULL,
  ultimacompra DATE NOT NULL,
  ticketmedio FLOAT NOT NULL,
  ticketultcompra FLOAT NOT NULL,
  lojafrequente VARCHAR(100) NOT NULL,
  lojaultcompra VARCHAR(100) NOT NULL
);
