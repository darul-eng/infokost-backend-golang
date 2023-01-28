CREATE TABLE boarding(
    id serial not null primary key ,
    name varchar(255) not null ,
    description varchar(255) not null ,
    address varchar(255) not null ,
    contact varchar(255) not null ,
    price int not null ,
    longlat int not null
)