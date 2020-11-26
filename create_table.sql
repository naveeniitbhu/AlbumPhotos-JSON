-- create database `ngnaven@gmail.com`;
-- USE ngnaven@gmail.com
CREATE TABLE album(
    id INT,
    userId INT,
    title VARCHAR(100),
    PRIMARY KEY(id)
);

CREATE TABLE photo(
    id INT,
    albumId INT,
    photoId INT AUTO_INCREMENT,
    title VARCHAR(100),
    url VARCHAR(255),
    thumbnailUrl VARCHAR(255),
    PRIMARY KEY(photoId),
    FOREIGN KEY(albumId) REFERENCES album(id)
);