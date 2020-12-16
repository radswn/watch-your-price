CREATE TABLE Users
(
    user_id  int          NOT NULL AUTO_INCREMENT,
    email    varchar(100) NOT NULL,
    password varchar(100) NOT NULL,
    PRIMARY KEY (user_id)
);

CREATE TABLE Products
(
    product_id   int          NOT NULL AUTO_INCREMENT,
    name         varchar(100) NOT NULL,
    lowest_price int          NOT NULL,
    PRIMARY KEY (product_id)
);

CREATE TABLE Prices
(
    price_id      int          NOT NULL AUTO_INCREMENT,
    product_id    int          NOT NULL,
    date_of_check DATE         NOT NULL,
    value         int          NOT NULL,
    shop_name     varchar(100) NOT NULL,
    PRIMARY KEY (price_id),
    FOREIGN KEY (product_id) REFERENCES Products (product_id)
);

CREATE TABLE Pages
(
    page_id       int          NOT NULL AUTO_INCREMENT,
    product_id    int          NOT NULL,
    shop_name     varchar(100) NOT NULL,
    link          varchar(255) NOT NULL,
    current_price int          NOT NULL,
    PRIMARY KEY (page_id),
    FOREIGN KEY (product_id) REFERENCES Products (product_id)
);

CREATE TABLE Watched
(
    user_id    int NOT NULL,
    product_id int NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users (user_id),
    FOREIGN KEY (product_id) REFERENCES Products (product_id)
);