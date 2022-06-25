CREATE TABLE `User` (
    `id`       bigint unsigned NOT NULL AUTO_INCREMENT,
    `name`     varchar(255) NOT NULL,
    `created`  datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modified` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
)
ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `User` (`id`,`name`) VALUES
(1, 'foo'), (2, 'bar'), (3, 'hoge');


CREATE TABLE `Post` (
    `id`       bigint unsigned NOT NULL AUTO_INCREMENT,
    `authorId` bigint unsigned NOT NULL,
    `title`    varchar(255) NOT NULL,
    `posted`   datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `created`  datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modified` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY `author` (`authorId`) REFERENCES `User` (`id`)
)
ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `Post` (`authorId`,`title`,`posted`) VALUES
(1, 'foo1', '2022-01-02 15:00'),
(1, 'foo2', '2022-01-02 18:00'),
(2, 'bar1', '2022-01-03 15:00');
