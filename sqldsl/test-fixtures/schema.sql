package test_fixtures

CREATE TABLE language (
    id              INT     NOT NULL PRIMARY KEY,
    cd              CHAR(2)       NOT NULL,
    description     VARCHAR(50)
);

CREATE TABLE author (
    id              INT     NOT NULL PRIMARY KEY,
    first_name      VARCHAR(50),
    last_name       VARCHAR(50)  NOT NULL,
    date_of_birth   DATETIME,
    year_of_birth   INT,
    distinguished   TINYINT
);

CREATE TABLE book (
    id              INT     NOT NULL PRIMARY KEY,
    author_id       INT     NOT NULL,
    title           VARCHAR(400) NOT NULL,
    published_in    INT     NOT NULL,
    language_id     INT     NOT NULL,
);

CREATE TABLE book_store (
    name            VARCHAR(400) NOT NULL UNIQUE
);

CREATE TABLE book_to_book_store (
    name            VARCHAR(400) NOT NULL,
    book_id         INTEGER       NOT NULL,
    stock           INTEGER,

    PRIMARY KEY(name, book_id),
);