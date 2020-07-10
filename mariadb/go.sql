-- auto-generated definition
create table items
(
    ID       int auto_increment
        primary key,
    Name     varchar(255) null,
    URL      varchar(255) null,
    ImageURL varchar(255) null,
    Person   varchar(255) null,
    Quantity int          null,
    Deleted  tinyint(1)   null default false
);

-- auto-generated definition
create table errors
(
    HTTPNotFound   varchar(50) null,
    IDNotFound     varchar(50) null,
    NoDataProvided varchar(50) null,
    NoItems        varchar(50) null
);

INSERT INTO `shopping-list`.errors (HTTPNotFound, IDNotFound, NoDataProvided, NoItems)
VALUES ('HTTP Not Found Error', 'ID Not Found', 'Error No Data Provided', 'Error No Items');