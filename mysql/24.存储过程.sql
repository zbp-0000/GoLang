-- 定义一个没有返回值 存储过程
-- 实现：模糊查询操作：
select * from emp where ename like '%A%';

create procedure mypro01(name varchar(10))
begin
	if name is null or name = "" then
		select * from emp;
	else
    select * from emp where ename like concat('%',name,'%');
	end if;	
end;

-- 删除存储过程：
drop procedure mypro01;

-- 调用存储过程：
call mypro01(null);
call mypro01('R');


-- 定义一个  有返回值的存储过程：
-- 实现：模糊查询操作：
-- in 参数前面的in可以省略不写
-- found_rows()mysql中定义的一个函数，作用返回查询结果的条数
create procedure mypro02(in name varchar(10),out num int(3))
begin
	if name is null or name = "" then
		select * from emp;
	else
    select * from emp where ename like concat('%',name,'%');
	end if;	
	select found_rows() into num;
end;

-- -- 调用存储过程：
call mypro02(null,@num);
select @num;

call mypro02('R',@aaa);
select @aaa;