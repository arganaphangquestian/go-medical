-- Role Table
CREATE TABLE IF NOT EXISTS roles
(
    id          VARCHAR(150) PRIMARY KEY,
    name        VARCHAR(200) NOT NULL,
    description TEXT,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP
);

INSERT INTO roles (id, name, description)
VALUES ('role_Admin', 'Admin', 'Admin Role'),
       ('role_Doctor', 'Doctor', 'Doctor Role'),
       ('role_Staff', 'Staff', 'Staff Role'),
       ('role_User', 'User', 'User Role'),
       ('role_Guest', 'Guest', 'Guest Role');