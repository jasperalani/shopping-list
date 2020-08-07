CREATE DATABASE `shopping-list`;

USE `shopping-list`;

-- auto-generated definition
create table items
(
    id        int auto_increment
        primary key,
    name      varchar(255) default ''                  null,
    url       varchar(255) default ''                  null,
    image_id  varchar(255) default '0'                 null,
    person    varchar(255) default ''                  null,
    quantity  int          default 1                   null,
    created   timestamp    default current_timestamp() not null,
    deleted   tinyint(1)   default 0                   null,
    completed tinyint(1)   default 0                   null,
    completed_on timestamp         default '0000-00-00 00:00:00'
);

-- auto-generated definition
create table errors
(
    err_httpnotfound   varchar(255) null,
    err_idnotfound     varchar(255) null,
    err_nodataprovided varchar(255) null,
    err_noitems        varchar(255) null
);


INSERT INTO `shopping-list`.errors (err_httpnotfound, err_idnotfound, err_nodataprovided, err_noitems)
VALUES ('HTTP Not Found Error', 'ID Not Found', 'Error No Data Provided', 'Error No Items');