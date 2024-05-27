create table rooms (
  id int primary key,
  name varchar(255) not null,
  capacity int not null,
  cnt_in_use int not null default 0,
  CONSTRAINT capacity_check CHECK (cnt_in_use <= capacity)
);

create table room_sessions (
  id int primary key auto_increment,
  req_id varchar(32) not null,
  user_id int not null,
  room_id int not null,
  cnt int not null
);

insert into room (id, name, capacity) values 
(1, 'Room 1', 2),
(2, 'Room 2', 2);


#transacion 1
START TRANSACTION;
update rooms set cnt_in_use = cnt_in_use + 2 where id = 1;
update rooms set cnt_in_use = cnt_in_use + 1 where id = 2;
insert into room_sessions (req_id, user_id, room_id, cnt) values 
('req1', 101, 1, 2),
('req1', 101, 2, 1);
COMMIT;

#transacion 2
START TRANSACTION;
update rooms set cnt_in_use = cnt_in_use + 2 where id = 2;
update rooms set cnt_in_use = cnt_in_use + 1 where id = 1;
insert into room_sessions (req_id, user_id, room_id, cnt) values 
('req2', 102, 2, 2),
('req2', 102, 1, 1);
COMMIT;

# error [2024-05-26 22:09:15] [40001][1213] Deadlock found when trying to get lock; try restarting transaction

