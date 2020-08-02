-- Blood Table
CREATE TABLE IF NOT EXISTS bloods
(
    id          VARCHAR(150) PRIMARY KEY,
    name        VARCHAR(200) NOT NULL,
    rhesus      BOOLEAN      NOT NULL,
    description TEXT,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP
);

INSERT INTO bloods (id, name, rhesus, description)
values ('blood_1', 'A', true, 'A rhesus + Blood'),
       ('blood_2', 'A', false, 'A rhesus - Blood'),
       ('blood_3', 'B', true, 'B rhesus + Blood'),
       ('blood_4', 'B', false, 'B rhesus - Blood'),
       ('blood_5', 'AB', true, 'AB rhesus + Blood'),
       ('blood_6', 'AB', false, 'AB rhesus - Blood'),
       ('blood_7', 'O', true, 'O rhesus + Blood'),
       ('blood_8', 'O', false, 'O rhesus - Blood');