-- 创建账户表：
create table account(
	id int primary key auto_increment,
	uname varchar(10) not null,
	balance double
);

-- 查看账户表：
select * from account;
-- 在表中插入数据：
insert into account values (null,'丽丽',2000),(null,'小刚',2000);

-- 丽丽给小刚 转200元：
update account set balance = balance - 200 where id = 1;
update account set balance = balance + 200 where id = 2;
-- 默认一个DML语句是一个事务，所以上面的操作执行了2个事务。

update account set balance = balance - 200 where id = 1;
update account set balance = balance2 + 200 where id = 2;

-- 必须让上面的两个操作控制在一个事务中：

-- 手动开启事务：
start transaction;

update account set balance = balance - 200 where id = 1;
update account set balance = balance + 200 where id = 2;

-- 手动回滚：刚才执行的操作全部取消：
rollback;
-- 手动提交：
commit;

-- 在回滚和提交之前，数据库中的数据都是操作的缓存中的数据，而不是数据库的真实数据




-- 查看默认的事务隔离级别  MySQL默认的是repeatable read  
select @@transaction_isolation;  
-- 设置事务的隔离级别   （设置当前会话的隔离级别）
set session transaction isolation level read uncommitted;  
set session transaction isolation level read committed;  
set session transaction isolation level repeatable read;  
set session transaction isolation level serializable;  

start transaction ;
select * from account where id = 1;














