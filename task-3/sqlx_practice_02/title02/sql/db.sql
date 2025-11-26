USE sqlx_study;

CREATE TABLE books
(
    id     INT AUTO_INCREMENT PRIMARY KEY,
    title  VARCHAR(200)   NOT NULL,
    author VARCHAR(100)   NOT NULL,
    price  DECIMAL(10, 2) NOT NULL
);

INSERT INTO books (title, author, price)
VALUES ('Go语言编程', '张三', 45.00),
       ('深入理解计算机系统', '李四', 89.00),
       ('算法导论', '王五', 128.00),
       ('设计模式', '赵六', 55.50),
       ('数据库系统概念', '钱七', 75.00),
       ('Python基础教程', '孙八', 39.90);


select * from books where price > 50;