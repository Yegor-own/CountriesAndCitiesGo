DROP TABLE IF EXISTS countries;
DROP TABLE IF EXISTS cities;

CREATE TABLE countries (
    id INT NOT NULL PRIMARY KEY,
    country VARCHAR(255)
);

CREATE TABLE cities (
    id INT NOT NULL PRIMARY KEY,
    country_id INT,
    city VARCHAR(255)
);