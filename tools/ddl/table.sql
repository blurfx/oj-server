CREATE TABLE IF NOT EXISTS `user`
(
    `id`                  BIGINT    AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `username`            VARCHAR(20) NOT NULL,
    `password`            VARCHAR(255) NOT NULL,
    `last_login_dt`       DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `create_dt`           DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `delete_dt`           DATETIME,

    CONSTRAINT `user_username_uniq` UNIQUE (`username`),

    INDEX `user_username` (`username`)
)
ENGINE = InnoDB
DEFAULT CHARSET = utf8mb4;

CREATE TABLE IF NOT EXISTS `problem`
(
    `id`                BIGINT NOT NULL PRIMARY KEY,
    `title`             VARCHAR(255) NOT NULL,
    `description`       TEXT NOT NULL,
    `spoiler`           TEXT NOT NULL,
    `time_limit`        BIGINT NOT NULL,
    `memory_limit`      BIGINT NOT NULL,
    `create_dt`         DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `delete_dt`         DATETIME
)
ENGINE = InnoDB
DEFAULT CHARSET = utf8mb4;

CREATE TABLE IF NOT EXISTS `problem_testcase`
(
    `id`                BIGINT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `problem_id`        BIGINT NOT NULL,
    `input_filepath`    VARCHAR(255) NOT NULL,
    `output_filepath`   VARCHAR(255) NOT NULL,

    CONSTRAINT `problem_testcase_problem_id_fk` FOREIGN KEY (`problem_id`) REFERENCES `problem` (`id`)
)
ENGINE = InnoDB
DEFAULT CHARSET = utf8mb4;

CREATE TABLE IF NOT EXISTS `language`
(
    `id`            BIGINT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `name`          VARCHAR(20) NOT NULL,
    `create_dt`     DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `delete_dt`     DATETIME
)
ENGINE = InnoDB
DEFAULT CHARSET = utf8mb4;

CREATE TABLE IF NOT EXISTS `submission`
(
    `id`            BIGINT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `user_id`       BIGINT NOT NULL,
    `problem_id`    BIGINT NOT NULL,
    `language_id`   BIGINT NOT NULL,
    `code`          TEXT NOT NULL,
    `time_limit`    INTEGER,
    `memory_limit`  INTEGER,
    `status`        VARCHAR(20) NOT NULL,
    `visibility`    VARCHAR(10) NOT NULL,
    `create_dt`     DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT `submission_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
    CONSTRAINT `submission_problem_id_fk` FOREIGN KEY (`problem_id`) REFERENCES `problem` (`id`),
    CONSTRAINT `submission_language_id_fk` FOREIGN KEY (`language_id`) REFERENCES `language` (`id`),

    INDEX `submission_user_id` (`user_id`),
    INDEX `submission_problem_id` (`problem_id`)
)
ENGINE = InnoDB
DEFAULT CHARSET = utf8mb4;
