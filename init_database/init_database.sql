DROP TABLE IF EXISTS rating;
DROP TABLE IF EXISTS rater;
DROP TABLE IF EXISTS inventory;
DROP TABLE IF EXISTS beer;
DROP TABLE IF EXISTS brewery;
DROP TABLE IF EXISTS vendor;

CREATE TABLE brewery(
    name        VARCHAR(50),
    password    CHAR(64),
    location    VARCHAR(50),
    PRIMARY KEY (name)
);

CREATE TABLE rater(
    name        VARCHAR(50),
    password    CHAR(64),
    location    VARCHAR(50),
    PRIMARY KEY (name)
);

CREATE TABLE vendor(
    vid         INT AUTO_INCREMENT,
    name        VARCHAR(50),
    password    CHAR(64),
    location    VARCHAR(50),
    PRIMARY KEY (vid)
);

CREATE TABLE beer(
    name    VARCHAR(50),
    brewery     VARCHAR(50),
    abv         DECIMAL(2,1),
    ibu         INT,
    PRIMARY KEY (name, brewery),
    FOREIGN KEY (brewery) REFERENCES brewery(name) ON DELETE CASCADE
);

CREATE TABLE rating(
    rid         INT AUTO_INCREMENT,
    beername    VARCHAR(50),
    brewery     VARCHAR(50),
    stars       INT,
    description VARCHAR(120),
    date        DATE,
    PRIMARY KEY (rid),
    FOREIGN KEY (beername) REFERENCES beer(name) ON DELETE CASCADE,
    FOREIGN KEY (brewery) REFERENCES brewery(name) ON DELETE CASCADE
);

CREATE TABLE inventory(
    vid         INT,
    beername    VARCHAR(50),
    brewery     VARCHAR(50),
    quantity    INT,
    PRIMARY KEY (vid, beername),
    FOREIGN KEY (beername) REFERENCES beer(name) ON DELETE CASCADE,
    FOREIGN KEY (brewery) REFERENCES brewery(name) ON DELETE CASCADE,
    FOREIGN KEY (vid) REFERENCES vendor(vid) ON DELETE CASCADE
);

INSERT INTO brewery VALUES
-- Password = password
-- Password = Spacedust
("No-Li", "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8", "Spokane, WA"),
("Elysian", "e46a0cb9f37b861b63ac1b1192d683341abaabb58494b61ff34fd248ca45b695", "Seattle, WA");

INSERT INTO beer VALUES
("Amber", "No-Li", 5.0, 10),
("Red, White & No-Li Pale Ale", "No-Li", 6.1, 35),
("Born & Raised", "No-Li", 7.0, 85),
("Big Juicy", "No-Li", 6.1, 55),
("Wrecking Ball", "No-Li", 9.5, 100),
("Corner Coast", "No-Li", 4.8, 20),
("Falls Porter", "No-Li", 6.1, 39),
("Amber", "Elysian", 5.5, 15),
("Space Dust IPA", "Elysian", 8.2, 73),
("Dayglow IPA", "Elysian", 7.4, 65),
("Immortal", "Elysian", 6.3, 62),
("Dragonstooth Stout", "Elysian", 8.1, 56),
("Mens Room Red", "Elysian", 5.6, 51),
("Avatar", "Elysian", 6.3, 43);

INSERT INTO vendor (name, password, location) VALUES
("Winco", "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8", "Spokane, WA"),
("Yokes", "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8", "Spokane, WA"),
("Haagens", "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8", "Seattle, WA"),
("Costco", "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8", "Seattle, WA");

INSERT INTO inventory VALUES
(1, "Amber", "No-Li", 34),
(2, "Amber", "No-Li", 21),
(1, "Red, White & No-Li Pale Ale", "No-Li", 17),
(3, "Red, White & No-Li Pale Ale", "No-Li", 13),
(1, "Born & Raised", "No-Li", 18),
(2, "Born & Raised", "No-Li", 7),
(3, "Born & Raised", "No-Li", 3),
(2, "Big Juicy", "No-Li", 8),
(2, "Wrecking Ball", "No-Li", 5),
(3, "Wrecking Ball", "No-Li", 10),
(2, "Space Dust IPA", "Elysian", 10),
(3, "Space Dust IPA", "Elysian", 4),
(4, "Space Dust IPA", "Elysian", 50),
(3, "Dayglow IPA", "Elysian", 6),
(4, "Immortal", "Elysian", 26),
(3, "Mens Room Red", "Elysian", 18),
(4, "Mens Room Red", "Elysian", 15),
(2, "Avatar", "Elysian", 18);

INSERT INTO rating (beername, brewery, stars, description, date) VALUES
("Corner Coast", "No-Li", 5, "My favorite beer ever!!", "2018-12-12"),
("Amber", "No-Li", 4, "Oh so good", "2018-12-10"),
("Corner Coast", "No-Li", 5, "mmmmmmm", "2018-12-11"),
("Corner Coast", "No-Li", 4, "Great beer", "2018-12-09");