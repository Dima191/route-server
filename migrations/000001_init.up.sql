create table route
(
    domain varchar(70) not null unique primary key,
    host   varchar(15) not null,
    port   varchar(6)  not null
);