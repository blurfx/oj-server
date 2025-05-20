-- +goose Up
CREATE TABLE IF NOT EXISTS "user"
(
    id                  BIGSERIAL PRIMARY KEY,
    username            VARCHAR(20) NOT NULL,
    password            VARCHAR(255) NOT NULL,
    last_login_dt       TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    create_dt           TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    delete_dt           TIMESTAMP,

    CONSTRAINT user_username_uniq UNIQUE (username)
);

CREATE INDEX user_username ON "user" (username);

CREATE TABLE IF NOT EXISTS problem
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

CREATE TABLE IF NOT EXISTS problem_testcase
(
    id                BIGSERIAL PRIMARY KEY,
    problem_id        BIGINT NOT NULL,
    input_filepath    VARCHAR(255) NOT NULL,
    output_filepath   VARCHAR(255) NOT NULL,

    CONSTRAINT problem_testcase_problem_id_fk FOREIGN KEY (problem_id) REFERENCES problem (id)
);

CREATE TABLE IF NOT EXISTS language
(
    id            BIGSERIAL PRIMARY KEY,
    name          VARCHAR(20) NOT NULL,
    create_dt     TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    delete_dt     TIMESTAMP
);

CREATE TABLE IF NOT EXISTS submission
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

    CONSTRAINT submission_language_id_fk FOREIGN KEY (language_id) REFERENCES language (id),
    CONSTRAINT submission_user_id_fk FOREIGN KEY (user_id) REFERENCES "user" (id),
    CONSTRAINT submission_problem_id_fk FOREIGN KEY (problem_id) REFERENCES problem (id)
);

CREATE INDEX submission_user_id ON submission (user_id);
CREATE INDEX submission_problem_id ON submission (problem_id);

-- +goose Down
DROP TABLE IF EXISTS submission, language, problem_testcase, problem, "user";
