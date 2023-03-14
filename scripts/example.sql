
CREATE TABLE IF NOT EXISTS example
(
	id                 serial       NOT NULL,
	name               varchar(255)  NOT NULL,
	CONSTRAINT PK_example_id PRIMARY KEY (id)
);

INSERT INTO example (name) VALUES ("test");