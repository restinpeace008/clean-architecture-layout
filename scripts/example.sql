
CREATE TABLE IF NOT EXISTS example
(
	id                 serial       NOT NULL,
	name               varchar(255) NOT NULL,
	related			   int 			NOT NULL,
	CONSTRAINT PK_example_id PRIMARY KEY (id),
	FOREIGN KEY (related) REFERENCES another_one(id)
);

INSERT INTO example (name) VALUES ("test");


CREATE TABLE IF NOT EXISTS another_one
(
	id                 serial       NOT NULL,
	type               varchar(255) NOT NULL,
	data 			   varchar(255) NOT NULL,
	CONSTRAINT PK_another_one_id PRIMARY KEY (id)
);

INSERT INTO another_one (name, data) VALUES ("test-type-1", "test-data-1"), ("test-type-2", "test-data-2"), ("test-type-3", "test-data-3");

INSERT INTO example (name, related) VALUES ("test-1", 1), ("test-2", 2), ("test-3", 3), ("test-4", 3);