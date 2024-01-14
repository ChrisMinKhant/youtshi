create database yout_shi_v_1_0_0;

show databases;

use yout_shi_v_1_0_0;

create table if not exists buses (
    id int not null,
    latest_location varchar(255),
    primary key(id)
);