CREATE TABLE IF NOT EXISTS cliente_has_tag (
  tag_id UUID,
  cliente_cpf VARCHAR(11),
  FOREIGN KEY(tag_id) REFERENCES tag(id) ON DELETE CASCADE,
  FOREIGN KEY(cliente_cpf) REFERENCES cliente(cpf) ON DELETE CASCADE,
  CONSTRAINT pk_client_tag PRIMARY KEY(tag_id, cliente_cpf)
);
