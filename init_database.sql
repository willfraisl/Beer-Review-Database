DROP TABLE IF EXISTS brewery;
DROP TABLE IF EXISTS rater;
DROP TABLE IF EXISTS vendor;
DROP TABLE IF EXISTS beer;
DROP TABLE IF EXISTS rating;


CREATE TABLE brewery(
    name VARCHAR(50),
    password VARCHAR(50),
    location VARCHAR(50),
    PRIMARY KEY (name)
);

CREATE TABLE rater(
    username VARCHAR(50),
    password VARCHAR(50),
    PRIMARY KEY (username)
);

CREATE TABLE vendor(
    name VARCHAR(50),
    password VARCHAR(50),
    location VARCHAR(50),
    PRIMARY KEY (name)
);

CREATE TABLE beer(
    name VARCHAR(50),
    abv DECIMAL(4,3),
    ibu INT,
    PRIMARY KEY (name)
);

CREATE TABLE rating(
    id INT,
    stars INT,
    description VARCHAR(150),
    date DATE
);