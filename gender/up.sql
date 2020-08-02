-- Gender Table
CREATE TABLE IF NOT EXISTS genders
(
    id          INTEGER PRIMARY KEY,
    name        VARCHAR(200) NOT NULL,
    description TEXT,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP
);

INSERT INTO genders (name, description)
values ('Male', 'Gender Male'),
       ('Female', 'Gender Female');