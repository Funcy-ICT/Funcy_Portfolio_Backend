CREATE TABLE IF NOT EXISTS `works` (
    `id` VARCHAR(64) NOT NULL COMMENT 'ユーザID',
    `work_id` VARCHAR(64) NOT NULL COMMENT '作品ID',
    `title` VARCHAR(64) NOT NULL COMMENT 'タイトル',
    `description` VARCHAR(512) NOT NULL COMMENT '作品概要',
    `url` VARCHAR(128) NOT NULL COMMENT '成果物url(github)',
    `movie_url` VARCHAR(128) NOT NULL COMMENT '成果物url(youtube)',
    `security` int NOT NULL COMMENT '公開設定',
    `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新時',
    PRIMARY KEY (`work_id`)
    );

CREATE TABLE IF NOT EXISTS `work_images` (
    `id` VARCHAR(64) NOT NULL COMMENT '作品ID',
    `image_id` VARCHAR(64) NOT NULL COMMENT '画像ID',
    `image_url` VARCHAR(128) NOT NULL COMMENT '画像URL',
    PRIMARY KEY (`image_id`)
    );

CREATE TABLE IF NOT EXISTS `work_tags` (
    `id` VARCHAR(64) NOT NULL COMMENT '作品ID',
    `tag_id` VARCHAR(64) NOT NULL COMMENT 'タグID',
    `tag` VARCHAR(128) NOT NULL COMMENT 'タグ',
    PRIMARY KEY (`tag_id`)
    );