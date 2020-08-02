-- Role Table
CREATE TABLE IF NOT EXISTS roles
(
    id          INTEGER PRIMARY KEY,
    name        VARCHAR(200) NOT NULL,
    description TEXT,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP
);

INSERT INTO roles (name, description)
VALUES ('Admin', 'Admin Role'),
       ('Doctor', 'Doctor Role'),
       ('Staff', 'Staff Role'),
       ('User', 'User Role'),
       ('Guest', 'Guest Role');