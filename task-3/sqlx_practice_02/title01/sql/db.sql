create database sqlx_study;

use sqlx_study;

drop table if exists `employees`;
create table employees(
                          `id` int primary key auto_increment,
                          `name` varchar(20),
                          `department` varchar(20),
                          `salary` decimal(10,2)
)charset ='utf8mb4';

insert into employees(`name`,`department`,`salary`)
values ('张三','人事部',3000),
       ('李四','技术部',8000),
       ('王五','财务部',4300),
       ('赵六','人事部',2400),
       ('小明','技术部',7200),
       ('小红','技术部',9000),
       ('小黑','技术部',10000),
       ('小兰','财务部',3000),
       ('小紫','技术部',8500),
       ('小黄','技术部',13000),
       ('小白','技术部',7500);


select id,name,department,salary from employees order by salary desc limit 1