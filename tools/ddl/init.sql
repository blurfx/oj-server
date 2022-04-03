DROP DATABASE onlinejudge;
CREATE DATABASE onlinejudge DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_unicode_ci;
CREATE USER 'judge_admin'@'%' IDENTIFIED BY 'judge_pass';
GRANT ALL PRIVILEGES ON onlinejudge.* TO 'judge_admin'@'%';
