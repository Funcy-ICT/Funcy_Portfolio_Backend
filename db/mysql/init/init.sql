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
    `id` VARCHAR(64) NOT NULL COMMENT 'ユーザID',
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

-- work
-- INSERT INTO `works` VALUES ("1","w1","初めての投稿だよ💓","もうやめましょうよ！！！","https://github.com/Funcy-ICT/Funcy_Portfolio_Android","https://www.youtube.com/watch?v=ViOzYSYWCMM&list=RDViOzYSYWCMM&start_radio=1",1);
INSERT INTO works (id,work_id,title,description,url,movie_url,security) VALUES ("1","w1","title","des","http","httpv",1);
INSERT INTO works (id,work_id,title,description,url,movie_url,security) VALUES ("1","w2","初めての投稿だよ💓","もうやめましょうよ！！！","https://github.com/Funcy-ICT/Funcy_Portfolio_Android","https://www.youtube.com/watch?v=ViOzYSYWCMM&list=RDViOzYSYWCMM&start_radio=1",1);
INSERT INTO works (id,work_id,title,description,url,movie_url,security) VALUES ("1","w3","初めての投稿だよ2💓","もうやめましょうよ！！！","https://github.com/Funcy-ICT/Funcy_Portfolio_Android","https://www.youtube.com/watch?v=ViOzYSYWCMM&list=RDViOzYSYWCMM&start_radio=1",1);
INSERT INTO works (id,work_id,title,description,url,movie_url,security) VALUES ("1","w4","初めての投稿だよ3💓","もうやめましょうよ！！！","https://github.com/Funcy-ICT/Funcy_Portfolio_Android","https://www.youtube.com/watch?v=ViOzYSYWCMM&list=RDViOzYSYWCMM&start_radio=1",1);
INSERT INTO works (id,work_id,title,description,url,movie_url,security) VALUES ("1","w5","初めての投稿だよ4💓","もうやめましょうよ！！！","https://github.com/Funcy-ICT/Funcy_Portfolio_Android","https://www.youtube.com/watch?v=ViOzYSYWCMM&list=RDViOzYSYWCMM&start_radio=1",1);
INSERT INTO works (id,work_id,title,description,url,movie_url,security) VALUES ("1","w6","初めての投稿だよ5💓","もうやめましょうよ！！！","https://github.com/Funcy-ICT/Funcy_Portfolio_Android","https://www.youtube.com/watch?v=ViOzYSYWCMM&list=RDViOzYSYWCMM&start_radio=1",1);
INSERT INTO works (id,work_id,title,description,url,movie_url,security) VALUES ("1","w7","初めての投稿だよ6💓","もうやめましょうよ！！！","https://github.com/Funcy-ICT/Funcy_Portfolio_Android","https://www.youtube.com/watch?v=ViOzYSYWCMM&list=RDViOzYSYWCMM&start_radio=1",1);
INSERT INTO works (id,work_id,title,description,url,movie_url,security) VALUES ("1","w8","初めての投稿だよ7💓","もうやめましょうよ！！！","https://github.com/Funcy-ICT/Funcy_Portfolio_Android","https://www.youtube.com/watch?v=ViOzYSYWCMM&list=RDViOzYSYWCMM&start_radio=1",1);
INSERT INTO works (id,work_id,title,description,url,movie_url,security) VALUES ("1","w9","初めての投稿だよ8💓","もうやめましょうよ！！！","https://github.com/Funcy-ICT/Funcy_Portfolio_Android","https://www.youtube.com/watch?v=ViOzYSYWCMM&list=RDViOzYSYWCMM&start_radio=1",1);
INSERT INTO works (id,work_id,title,description,url,movie_url,security) VALUES ("1","w10","初めての投稿だよ9💓","もうやめましょうよ！！！","https://github.com/Funcy-ICT/Funcy_Portfolio_Android","https://www.youtube.com/watch?v=ViOzYSYWCMM&list=RDViOzYSYWCMM&start_radio=1",1);
INSERT INTO works (id,work_id,title,description,url,movie_url,security) VALUES ("1","w11","初めての投稿だよ11💓","もうやめましょうよ！！！","https://github.com/Funcy-ICT/Funcy_Portfolio_Android","https://www.youtube.com/watch?v=ViOzYSYWCMM&list=RDViOzYSYWCMM&start_radio=1",1);
INSERT INTO works (id,work_id,title,description,url,movie_url,security) VALUES ("1","w12","初めての投稿だよ12💓","もうやめましょうよ！！！","https://github.com/Funcy-ICT/Funcy_Portfolio_Android","https://www.youtube.com/watch?v=ViOzYSYWCMM&list=RDViOzYSYWCMM&start_radio=1",1);
INSERT INTO works (id,work_id,title,description,url,movie_url,security) VALUES ("1","w13","初めての投稿だよ13💓","もうやめましょうよ！！！","https://github.com/Funcy-ICT/Funcy_Portfolio_Android","https://www.youtube.com/watch?v=ViOzYSYWCMM&list=RDViOzYSYWCMM&start_radio=1",1);
INSERT INTO works (id,work_id,title,description,url,movie_url,security) VALUES ("1","w14","初めての投稿だよ14💓","もうやめましょうよ！！！","https://github.com/Funcy-ICT/Funcy_Portfolio_Android","https://www.youtube.com/watch?v=ViOzYSYWCMM&list=RDViOzYSYWCMM&start_radio=1",1);
INSERT INTO works (id,work_id,title,description,url,movie_url,security) VALUES ("1","w15","初めての投稿だよ15💓","もうやめましょうよ！！！","https://github.com/Funcy-ICT/Funcy_Portfolio_Android","https://www.youtube.com/watch?v=ViOzYSYWCMM&list=RDViOzYSYWCMM&start_radio=1",1);
INSERT INTO works (id,work_id,title,description,url,movie_url,security) VALUES ("1","w16","初めての投稿だよ16💓","もうやめましょうよ！！！","https://github.com/Funcy-ICT/Funcy_Portfolio_Android","https://www.youtube.com/watch?v=ViOzYSYWCMM&list=RDViOzYSYWCMM&start_radio=1",1);
INSERT INTO works (id,work_id,title,description,url,movie_url,security) VALUES ("1","w17","初めての投稿だよ17💓","もうやめましょうよ！！！","https://github.com/Funcy-ICT/Funcy_Portfolio_Android","https://www.youtube.com/watch?v=ViOzYSYWCMM&list=RDViOzYSYWCMM&start_radio=1",1);
INSERT INTO works (id,work_id,title,description,url,movie_url,security) VALUES ("1","w18","初めての投稿だよ18💓","もうやめましょうよ！！！","https://github.com/Funcy-ICT/Funcy_Portfolio_Android","https://www.youtube.com/watch?v=ViOzYSYWCMM&list=RDViOzYSYWCMM&start_radio=1",1);
INSERT INTO works (id,work_id,title,description,url,movie_url,security) VALUES ("1","w19","初めての投稿だよ19💓","もうやめましょうよ！！！","https://github.com/Funcy-ICT/Funcy_Portfolio_Android","https://www.youtube.com/watch?v=ViOzYSYWCMM&list=RDViOzYSYWCMM&start_radio=1",1);
INSERT INTO works (id,work_id,title,description,url,movie_url,security) VALUES ("1","w20","初めての投稿だよ20💓","もうやめましょうよ！！！","https://github.com/Funcy-ICT/Funcy_Portfolio_Android","https://www.youtube.com/watch?v=ViOzYSYWCMM&list=RDViOzYSYWCMM&start_radio=1",1);
-- INSERT INTO works (id,work_id,title,description,url,movie_url,security) VALUES ("1","w21","初めての投稿だよ21💓","もうやめましょうよ！！！","https://github.com/Funcy-ICT/Funcy_Portfolio_Android","https://www.youtube.com/watch?v=ViOzYSYWCMM&list=RDViOzYSYWCMM&start_radio=1",1);
-- INSERT INTO works (id,work_id,title,description,url,movie_url,security) VALUES ("1","w22","初めての投稿だよ22💓","もうやめましょうよ！！！","https://github.com/Funcy-ICT/Funcy_Portfolio_Android","https://www.youtube.com/watch?v=ViOzYSYWCMM&list=RDViOzYSYWCMM&start_radio=1",1);
-- INSERT INTO works (id,work_id,title,description,url,movie_url,security) VALUES ("1","w23","初めての投稿だよ23💓","もうやめましょうよ！！！","https://github.com/Funcy-ICT/Funcy_Portfolio_Android","https://www.youtube.com/watch?v=ViOzYSYWCMM&list=RDViOzYSYWCMM&start_radio=1",1);
-- INSERT INTO works (id,work_id,title,description,url,movie_url,security) VALUES ("1","w24","初めての投稿だよ24💓","もうやめましょうよ！！！","https://github.com/Funcy-ICT/Funcy_Portfolio_Android","https://www.youtube.com/watch?v=ViOzYSYWCMM&list=RDViOzYSYWCMM&start_radio=1",1);
-- INSERT INTO works (id,work_id,title,description,url,movie_url,security) VALUES ("1","w25","初めての投稿だよ25💓","もうやめましょうよ！！！","https://github.com/Funcy-ICT/Funcy_Portfolio_Android","https://www.youtube.com/watch?v=ViOzYSYWCMM&list=RDViOzYSYWCMM&start_radio=1",1);
-- INSERT INTO works (id,work_id,title,description,url,movie_url,security) VALUES ("1","w26","初めての投稿だよ26💓","もうやめましょうよ！！！","https://github.com/Funcy-ICT/Funcy_Portfolio_Android","https://www.youtube.com/watch?v=ViOzYSYWCMM&list=RDViOzYSYWCMM&start_radio=1",1);
-- INSERT INTO works (id,work_id,title,description,url,movie_url,security) VALUES ("1","w27","初めての投稿だよ27💓","もうやめましょうよ！！！","https://github.com/Funcy-ICT/Funcy_Portfolio_Android","https://www.youtube.com/watch?v=ViOzYSYWCMM&list=RDViOzYSYWCMM&start_radio=1",1);
-- INSERT INTO works (id,work_id,title,description,url,movie_url,security) VALUES ("1","w28","初めての投稿だよ28💓","もうやめましょうよ！！！","https://github.com/Funcy-ICT/Funcy_Portfolio_Android","https://www.youtube.com/watch?v=ViOzYSYWCMM&list=RDViOzYSYWCMM&start_radio=1",1);

-- work images
INSERT INTO `work_images` VALUES ("w1","i1","https://avatars.githubusercontent.com/u/40165303?v=4");
INSERT INTO `work_images` VALUES ("w1","i2","https://avatars.githubusercontent.com/u/40165303?v=4");
INSERT INTO `work_images` VALUES ("w1","i3","https://avatars.githubusercontent.com/u/40165303?v=4");
INSERT INTO `work_images` VALUES ("w2","i4","https://avatars.githubusercontent.com/u/40165303?v=4");
INSERT INTO `work_images` VALUES ("w2","i5","https://avatars.githubusercontent.com/u/40165303?v=4");
INSERT INTO `work_images` VALUES ("w2","i6","https://avatars.githubusercontent.com/u/40165303?v=4");
INSERT INTO `work_images` VALUES ("w3","i7","https://avatars.githubusercontent.com/u/40165303?v=4");
INSERT INTO `work_images` VALUES ("w3","i9","https://avatars.githubusercontent.com/u/40165303?v=4");
INSERT INTO `work_images` VALUES ("w4","i10","https://avatars.githubusercontent.com/u/40165303?v=4");
INSERT INTO `work_images` VALUES ("w4","i11","https://avatars.githubusercontent.com/u/40165303?v=4");
INSERT INTO `work_images` VALUES ("w5","i12","https://avatars.githubusercontent.com/u/40165303?v=4");
INSERT INTO `work_images` VALUES ("w5","i13","https://avatars.githubusercontent.com/u/40165303?v=4");
INSERT INTO `work_images` VALUES ("w6","i14","https://avatars.githubusercontent.com/u/40165303?v=4");
INSERT INTO `work_images` VALUES ("w6","i15","https://avatars.githubusercontent.com/u/40165303?v=4");
INSERT INTO `work_images` VALUES ("w7","i16","https://avatars.githubusercontent.com/u/40165303?v=4");
INSERT INTO `work_images` VALUES ("w7","i17","https://avatars.githubusercontent.com/u/40165303?v=4");
INSERT INTO `work_images` VALUES ("w8","i18","https://avatars.githubusercontent.com/u/40165303?v=4");
INSERT INTO `work_images` VALUES ("w9","i19","https://avatars.githubusercontent.com/u/40165303?v=4");
INSERT INTO `work_images` VALUES ("w10","i20","https://avatars.githubusercontent.com/u/40165303?v=4");
INSERT INTO `work_images` VALUES ("w11","i21","https://avatars.githubusercontent.com/u/40165303?v=4");
INSERT INTO `work_images` VALUES ("w12","i22","https://avatars.githubusercontent.com/u/40165303?v=4");
INSERT INTO `work_images` VALUES ("w13","i23","https://avatars.githubusercontent.com/u/40165303?v=4");
INSERT INTO `work_images` VALUES ("w14","i24","https://avatars.githubusercontent.com/u/40165303?v=4");
INSERT INTO `work_images` VALUES ("w15","i25","https://avatars.githubusercontent.com/u/40165303?v=4");
INSERT INTO `work_images` VALUES ("w16","i26","https://avatars.githubusercontent.com/u/40165303?v=4");
INSERT INTO `work_images` VALUES ("w17","i27","https://avatars.githubusercontent.com/u/40165303?v=4");
INSERT INTO `work_images` VALUES ("w18","i28","https://avatars.githubusercontent.com/u/40165303?v=4");
INSERT INTO `work_images` VALUES ("w19","i29","https://avatars.githubusercontent.com/u/40165303?v=4");
INSERT INTO `work_images` VALUES ("w20","i30","https://avatars.githubusercontent.com/u/40165303?v=4");
-- work tags
INSERT INTO `work_tags` VALUES ("w1","i1","android");
INSERT INTO `work_tags` VALUES ("w1","i2","go");
INSERT INTO `work_tags` VALUES ("w1","i3","天下統一");