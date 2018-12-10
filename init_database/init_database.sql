DROP TABLE IF EXISTS rating;
DROP TABLE IF EXISTS rater;
DROP TABLE IF EXISTS beer;
DROP TABLE IF EXISTS brewery;
DROP TABLE IF EXISTS vendor;


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
    FOREIGN KEY (brewery) REFERENCES brewery(name) ON DELETE CASCADE
);

CREATE TABLE rating(
    id INT AUTO_INCREMENT,
    beer VARCHAR(50),
    brewery VARCHAR(50),
    stars INT,
    description VARCHAR(120),
    date DATE,
    PRIMARY KEY (id),
    FOREIGN KEY (beer) REFERENCES beer(name) ON DELETE CASCADE,
    FOREIGN KEY (brewery) REFERENCES brewery(name) ON DELETE CASCADE
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