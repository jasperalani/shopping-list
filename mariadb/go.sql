CREATE DATABASE `shopping-list`;

USE `shopping-list`;

-- auto-generated definition
create table items
(
    id           int auto_increment
        primary key,
    name         varchar(255) null default '',
    url          varchar(255) null default '',
    image_url    varchar(255) null default '',
    person       varchar(255) null default '',
    quantity     int          null default 1,
    created      timestamp         default current_timestamp,
    deleted      tinyint(1)   null default false,
    completed    tinyint(1)   null default false,
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