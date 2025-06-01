-- +goose Up
CREATE TABLE IF NOT EXISTS users
(
    id                  BIGSERIAL PRIMARY KEY,
    username            VARCHAR(20) NOT NULL,
    password            VARCHAR(255) NOT NULL,
    last_login_dt       TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    create_dt           TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    delete_dt           TIMESTAMP,

    CONSTRAINT user_username_uniq UNIQUE (username)
);

CREATE INDEX users_username ON users (username);

CREATE TABLE IF NOT EXISTS problems
(
    id                BIGINT PRIMARY KEY,
    title             VARCHAR(255) NOT NULL,
    description       TEXT NOT NULL,
    spoiler           TEXT,
    time_limit        BIGINT NOT NULL,
    memory_limit      BIGINT NOT NULL,
    create_dt         TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    delete_dt         TIMESTAMP
);

CREATE TABLE IF NOT EXISTS problem_testcases
(
    id                BIGSERIAL PRIMARY KEY,
    problem_id        BIGINT NOT NULL,
    input_filepath    VARCHAR(255) NOT NULL,
    output_filepath   VARCHAR(255) NOT NULL,

    CONSTRAINT problem_testcases_problem_id_fk FOREIGN KEY (problem_id) REFERENCES problems (id)
);

CREATE TABLE IF NOT EXISTS languages
(
    id            BIGSERIAL PRIMARY KEY,
    name          VARCHAR(20) NOT NULL,
    create_dt     TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    delete_dt     TIMESTAMP
);

CREATE TABLE IF NOT EXISTS submissions
(
    id            BIGSERIAL PRIMARY KEY,
    user_id       BIGINT NOT NULL,
    problem_id    BIGINT NOT NULL,
    language_id   BIGINT NOT NULL,
    code          TEXT NOT NULL,
    time_limit    INTEGER,
    memory_limit  INTEGER,
    status        VARCHAR(20) NOT NULL,
    visibility    VARCHAR(10) NOT NULL,
    create_dt     TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT submissions_languages_id_fk FOREIGN KEY (language_id) REFERENCES languages (id),
    CONSTRAINT submissions_users_id_fk FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT submissions_problem_id_fk FOREIGN KEY (problem_id) REFERENCES problems (id)
);

CREATE INDEX submissions_users_id ON submissions (user_id);
CREATE INDEX submissions_problems_id ON submissions (problem_id);

-- +goose Down
DROP TABLE IF EXISTS submissions, languages, problem_testcases, problems, users;
