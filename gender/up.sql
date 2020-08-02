-- Gender Table
CREATE TABLE IF NOT EXISTS genders
(
    id          VARCHAR(150) PRIMARY KEY,
    name        VARCHAR(200) NOT NULL,
    description TEXT,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP
);

INSERT INTO genders (id, name, description)
values ('gender_Male', 'Male', 'Gender Male'),
       ('gender_Female', 'Female', 'Gender Female');