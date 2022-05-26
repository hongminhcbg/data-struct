create table `users` (
    id int(20) primary key AUTO_INCREMENT,
    balance int(20) default 0,
    name VARCHAR(255)
);

insert into users(name, balance) values 
    ('minh', 1000000),
    ('nguyen', 100000),
    ('hong', 300000);
