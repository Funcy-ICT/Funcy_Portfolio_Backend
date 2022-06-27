-- SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
-- SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
-- SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';
SET CHARSET UTF8;
DROP SCHEMA IF EXISTS `funcy`;
CREATE SCHEMA IF NOT EXISTS `funcy` DEFAULT CHARACTER SET utf8;
USE `funcy`;


CREATE TABLE IF NOT EXISTS `users` (
    `id` VARCHAR(64) NOT NULL COMMENT 'ユーザID',
    `display_name` VARCHAR(32) NOT NULL COMMENT '表示名',
    `icon` VARCHAR(128) NOT NULL COMMENT 'プロフィール画像',
    `family_name` VARCHAR(32) NOT NULL COMMENT '性',
    `first_name` VARCHAR(32) NOT NULL COMMENT '名',
    `mail` VARCHAR(32) NOT NULL COMMENT 'メールアドレス',
    `password` VARCHAR(128) NOT NULL COMMENT 'パスワード',
    `grade` VARCHAR(32) NOT NULL COMMENT '学年',
    `course` VARCHAR(32) NOT NULL COMMENT 'コース名',
    `token` VARCHAR(128) NOT NULL COMMENT '認証用Token',
    PRIMARY KEY (`id`),
    INDEX `idx_auth_token` (`id` ASC)
    );

CREATE TABLE IF NOT EXISTS `funcy`.`works` (
    `id` VARCHAR(64) NOT NULL COMMENT 'ID',
    `work_id` VARCHAR(64) NOT NULL COMMENT '作品ID',
    `title` VARCHAR(64) NOT NULL COMMENT 'タイトル',
    `description` VARCHAR(512) NOT NULL COMMENT '作品概要',
    `url` VARCHAR(128) NOT NULL COMMENT '成果物url(github)',
    `movie_url` VARCHAR(128) NOT NULL COMMENT '成果物url(youtube)',
    `security` int NOT NULL COMMENT '公開設定',
    `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    PRIMARY KEY (`work_id`)
    );

CREATE TABLE IF NOT EXISTS `funcy`.`work_images` (
    `id` VARCHAR(64) NOT NULL COMMENT '作品ID',
    `image_id` VARCHAR(64) NOT NULL COMMENT '画像ID',
    `image_url` VARCHAR(128) NOT NULL COMMENT '画像URL',
    PRIMARY KEY (`image_id`)
    );

CREATE TABLE IF NOT EXISTS `funcy`.`work_tags` (
    `id` VARCHAR(64) NOT NULL COMMENT '作品ID',
    `tag_id` VARCHAR(64) NOT NULL COMMENT 'タグID',
    `tag` VARCHAR(128) NOT NULL COMMENT 'タグ',
    PRIMARY KEY (`tag_id`)
    );

CREATE TABLE IF NOT EXISTS `funcy`.`tags` (
    `id` VARCHAR(64) NOT NULL COMMENT 'タグID',
    `name` VARCHAR(64) NOT NULL COMMENT 'タグ名',
    `is_skill` tinyint(1) DEFAULT NULL COMMENT 'スキルタグ判定',
    PRIMARY KEY (`id`)
    );

-- user
INSERT INTO `users` VALUES ("1","山本",".com","yamamoto","yuhei","yamamoto@fun.ac.jp","pass","修士1年","情報アーキテクチャ領域","Token1");
INSERT INTO `users` VALUES ("2","まっすー",".com","まっすー","だよ","増田@fun.ac.jp","pass","修士1年","情報アーキテクチャ領域","Token2");

-- -- work
-- INSERT INTO `works` VALUES ("1","1","初めての投稿だよ💓","もうやめましょうよ！！！","https://github.com/Funcy-ICT/Funcy_Portfolio_Android","https://www.youtube.com/watch?v=ViOzYSYWCMM&list=RDViOzYSYWCMM&start_radio=1",1);
-- -- work images
-- INSERT INTO `work_images` VALUES ("1","1","https://avatars.githubusercontent.com/u/40165303?v=4");
-- INSERT INTO `work_images` VALUES ("1","2","https://avatars.githubusercontent.com/u/40165303?v=4");
-- INSERT INTO `work_images` VALUES ("1","3","https://avatars.githubusercontent.com/u/40165303?v=4");
-- -- work tags
-- INSERT INTO `work_tags` VALUES ("1","1","android");
-- INSERT INTO `work_tags` VALUES ("1","2","go");
-- INSERT INTO `work_tags` VALUES ("1","3","天下統一");