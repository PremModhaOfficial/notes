\c postgres
DROP DATABASE IF EXISTS nats_db;
CREATE DATABASE nats_db;

CREATE TABLE events (
	id serial PRIMARY KEY,
	subject varchar(255),
	payload varchar(255)
);

