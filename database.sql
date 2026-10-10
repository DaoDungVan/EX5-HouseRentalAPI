CREATE DATABASE IF NOT EXISTS house_rental_db;

USE house_rental_db;

CREATE TABLE IF NOT EXISTS bills (
    id INT AUTO_INCREMENT PRIMARY KEY,
    room_name VARCHAR(100) NOT NULL,
    tenant_name VARCHAR(255) NOT NULL,
    month VARCHAR(20) NOT NULL,
    room_price DECIMAL(15,2) NOT NULL,
    electricity_old INT NOT NULL,
    electricity_new INT NOT NULL,
    electricity_price DECIMAL(15,2) NOT NULL,
    water_old INT NOT NULL,
    water_new INT NOT NULL,
    water_price DECIMAL(15,2) NOT NULL,
    internet_fee DECIMAL(15,2) NOT NULL DEFAULT 0,
    garbage_fee DECIMAL(15,2) NOT NULL DEFAULT 0,
    other_fee DECIMAL(15,2) NOT NULL DEFAULT 0,
    electricity_cost DECIMAL(15,2) NOT NULL,
    water_cost DECIMAL(15,2) NOT NULL,
    total DECIMAL(15,2) NOT NULL
);
