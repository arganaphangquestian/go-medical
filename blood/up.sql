-- Blood Table
CREATE TABLE IF NOT EXISTS bloods
(
    id          INTEGER PRIMARY KEY,
    name        VARCHAR(200) NOT NULL,
    rhesus      BOOLEAN      NOT NULL,
    description TEXT,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP
);

INSERT INTO bloods (name, rhesus, description)
values ('A', true, 'A rhesus + Blood'),
       ('A', false, 'A rhesus - Blood'),
       ('B', true, 'B rhesus + Blood'),
       ('B', false, 'B rhesus - Blood'),
       ('AB', true, 'AB rhesus + Blood'),
       ('AB', false, 'AB rhesus - Blood'),
       ('O', true, 'O rhesus + Blood'),
       ('O', false, 'O rhesus - Blood');