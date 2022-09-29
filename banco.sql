CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;

$$ LANGUAGE plpgsql;

CREATE TABLE usuarios(
	id serial primary key not null,
	tipo int not null,
	nome varchar(100) not null ,
	email varchar(200) not null unique,
	senha text not null,
	created_at timestamp default (now()),
	updated_at timestamp default (now()),
	removed_at timestamp
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON "usuarios"
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE  frequencias (
	id serial primary key not null,
	data_atual timestamp default (now())
);

CREATE TABLE frequencia_usuario (
	id serial primary key not null,
	usuario_id int not null, 
	frequencia_id int not null,
	entrada timestamp default (now()),
	saida timestamp,
	constraint fk_usuario foreign key(usuario_id) references usuarios(id) on delete no action,
	constraint fk_frequencia foreign key(frequencia_id) references frequencias(id) on delete no action
)
