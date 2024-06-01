CREATE TABLE movies (
    id              SERIAL PRIMARY KEY,
    title           TEXT NOT NULL,
    year            NUMERIC NOT NULL,
    description     TEXT,
    runtime         TEXT,
    posterURL       TEXT,
    rating          FLOAT,
    votes           INTEGER,
    lastUpdated     TIMESTAMP
);

CREATE TABLE movieToPerson (
    movieId         INTEGER NOT NULL,
    personId        INTEGER NOT NULL,
    roleId          INTEGER NOT NULL
);

CREATE TABLE role (
    roleId          SERIAL PRIMARY KEY,
    roleName        TEXT NOT NULL
);

CREATE TABLE person (
    personId        SERIAL PRIMARY KEY,
    firstName       TEXT NOT NULL,
    lastName        TEXT NOT NULL,
    description     TEXT,
    photoURL        TEXT
);

CREATE TABLE genre (
    genreId         SERIAL PRIMARY KEY,
    genreName       TEXT NOT NULL
);

CREATE TABLE movieToGenre (
    movieId         INTEGER NOT NULL,
    genreId         INTEGER NOT NULL
);

CREATE TABLE appUser (
    userId          TEXT NOT NULL PRIMARY KEY,
    nickname        TEXT NOT NULL,
    photoURL        TEXT
);

CREATE TABLE favoritesList (
    favoritesId     SERIAL PRIMARY KEY,
    userId          TEXT NOT NULL
);

CREATE TABLE favoritesListToMovie (
    favoritesId     INTEGER NOT NULL,
    movieId         INTEGER NOT NULL
);

CREATE TABLE movieComment (
    commentId       SERIAL PRIMARY KEY,
    movieId         INTEGER NOT NULL,
    userId          TEXT NOT NULL,
    commentText     TEXT NOT NULL,
    replyCommentId  INTEGER,
    commentTime     TIMESTAMP NOT NULL
);