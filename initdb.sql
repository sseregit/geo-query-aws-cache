CREATE DATABASE IF NOT exists metting;

use metting;

CREATE TABLE IF NOT EXISTS `user` (
    `t_id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `user_name` VARCHAR(100) UNIQUE  NOT NULL,
    `image` JSON NOT NULL,
    `description` VARCHAR(300) DEFAULT '',
    `hobby` JSON NOT NULL,
    `is_valid` TINYINT(1) DEFAULT false
);

CREATE TABLE IF NOT EXISTS `user_location` (
    `user_name` VARCHAR(100) UNIQUE NOT NULL,
    `latitude` DOUBLE NOT NULL COMMENT '위도',
    `hardness` DOUBLE NOT NULL COMMENT '경도',
    `location` POINT NOT NULL COMMENT '위치',

    SPATIAL INDEX `spatial_index` (`location`)
);

CREATE TABLE IF NOT EXISTS `user_like` (
    `form_user` VARCHAR(100) NOT NULL,
    `to_user` VARCHAR(100) NOT NULL,
    `status` ENUM('send', 'accepted', 'refuse') NOT NULL ,
    `created_time` BIGINT NOT NULL,
    `update_time` BIGINT DEFAULT 0
);