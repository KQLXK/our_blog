# CREATE DATABASE IF NOT EXISTS `ourblog` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;
# USE `ourblog`;
#
# DROP TABLE IF EXISTS `articles`;
# CREATE TABLE `articles`
# (
#     `article_id`  int          NOT NULL AUTO_INCREMENT COMMENT '文章ID',
#     `user_id`     int          NOT NULL COMMENT '作者ID',
#     `title`       varchar(255)  NOT NULL COMMENT '文章标题',
#     `excerpt`     text          COMMENT '文章摘要',
#     `category`    varchar(100)  COMMENT '文章分类',
#     `content`     text          NOT NULL COMMENT '文章内容',
#     `status`      varchar(50)   COMMENT '文章状态',
#     `create_time` timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
#     `update_time` timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
#     PRIMARY KEY (`article_id`)
# ) ENGINE = InnoDB
#   DEFAULT CHARSET=utf8mb4;
#
# DROP TABLE  IF EXISTS `user`;
# CREATE TABLE `user` (
#     `user_id` int NOT NULL AUTO_INCREMENT COMMENT '用户ID',
#     `username` varchar(255) NOT NULL COMMENT '用户名',
#     `password` varchar(255) NOT NULL COMMENT '用户密码',
#     `email` varchar(255) NOT NULL COMMENT '电子邮件',
#     `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
#     PRIMARY KEY (`user_id`)
# ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户信息表';

CREATE TABLE `likes` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `article_id` int(11) NOT NULL COMMENT '文章ID',
    `user_id` int(11) NOT NULL COMMENT '用户ID',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT '软删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_article_id` (`article_id`),
    KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章点赞表';