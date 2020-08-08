CREATE DATABASE COFFEESHOP;

USE COFFEESHOP;
CREATE TABLE Drinks
(
    Id        VARCHAR(100) PRIMARY KEY,
    Name      VARCHAR(100) NOT NULL,
    Type      VARCHAR(100) NOT NULL,
    Origin    VARCHAR(100) NOT NULL,
    Brand     VARCHAR(100) NOT NULL,
    Price     FLOAT(10)    NOT NULL,
    Stock     INT(6)       NOT NULL,
    Timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


SELECT *
from Drinks;

UPDATE Drinks
SET Stock = 100
where Id = '1';

Insert into Drinks (Id, Name, Type, Origin, Brand, Price, Stock)
VALUES ('1', 'Water', 'Soda', 'Tap', 'Earth', 1.0, '10');