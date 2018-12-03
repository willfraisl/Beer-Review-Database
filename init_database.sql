DROP TABLE IF EXISTS beer;
DROP TABLE IF EXISTS brewery;
DROP TABLE IF EXISTS rater;
DROP TABLE IF EXISTS vendor;
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
    brewery VARCHAR(50),
    abv DECIMAL(4,3),
    ibu INT,
    PRIMARY KEY (name),
    FOREIGN KEY (brewery) REFERENCES brewery(name)
);

CREATE TABLE rating(
    id INT,
    stars INT,
    description VARCHAR(150),
    date DATE
);

INSERT INTO brewery VALUES
("No-Li", "password", "Spokane, WA");

INSERT INTO beer VALUES
("Amber", "No-Li", 0.05, 10),
("Red, White & No-Li Pale Ale", "No-Li", 0.061, 35),
("Born & Raised", "No-Li", 0.07, 85),
("Big Juicy", "No-Li", 0.061, 55),
("Wrecking Ball", "No-Li", 0.095, 100),
("Corner Coast", "No-Li", 0.048, 20),
("Falls Porter", "No-Li", 0.061, 39);