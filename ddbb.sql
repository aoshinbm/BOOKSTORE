use bookstore;

show tables;
select * from userregistration;
desc userregistration;

select * from bookrecord;
desc bookrecord;

select * from cartdetails;
desc cartdetails;
delete from cartdetails;

ALTER TABLE cartdetails
DROP COLUMN status;

select * from orderdetails;
desc orderdetails;

create table login(
	Username varchar(100) not null primary key,
	Password varchar(300) not null 
);

desc login;

SELECT * FROM BookRecord ;

update BookRecord set Title="Flash" ,Author="DC" ,Year="2015" where BOOKID=6;

SELECT BOOKID ,Title ,Author ,Year FROM BookRecord where Title="golang";

