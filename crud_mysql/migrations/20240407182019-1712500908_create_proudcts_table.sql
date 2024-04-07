-- +migrate Up
create table `products`
(
    `id`          INT AUTO_INCREMENT PRIMARY KEY ,
    `name`        varchar(50)  not null,
    `price`       float        not null,
    `quantity`    int          not null,
    `description` varchar(255) not null
);
-- +migrate Down
drop table `products`;
