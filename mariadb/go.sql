-- auto-generated definition
create table items
(
    id        int auto_increment
        primary key,
    name      varchar(255) null,
    url       varchar(255) null,
    image_url varchar(255) null,
    person    varchar(255) null,
    quantity  int          null,
    deleted   tinyint(1)   null default false
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