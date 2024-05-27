-- create table
create table room_trans (
    id bigint primary key auto_increment,
    room_name varchar(32),
    ref_id varchar(16) comment 'room_id + yymmdd',
    order_id varchar(32) comment 'order_id of trans, if empty room is available' default '',
    key `idx_ref_id`(ref_id)
);

-- insert data test
insert into room_trans(room_name, ref_id) values
('room1', '1-240527'),
('room1', '1-240528'),
('room1', '1-240529'),
('room1', '1-240530'),
('room2', '2-240527'),
('room2', '2-240528'),
('room2', '2-240529'),
('room2', '2-240530');

-- transaction 1
start transaction;
select * from room_trans where order_id = '' and room_name='room1' limit 2
for update skip locked ;
-- we will get 2 rows with id 1 and 2
-- update order
update room_trans set order_id='order1' where id in (1,2);
commit;


-- transaction 2
start transaction;
select * from room_trans where order_id = '' and room_name='room1' limit 2
for update skip locked ;
-- we will get 2 rows with id 3 and 4
-- update order

-- check if len(rows) < 2 rollback
update room_trans set order_id='order2' where id in (3,4);
commit;

