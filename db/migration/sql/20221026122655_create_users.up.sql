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
    `status` VARCHAR(32) NOT NULL COMMENT 'アカウントステータス',
    `code` VARCHAR(10) NOT NULL COMMENT 'ワンタイムパスワード',
    `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '作成時',
    `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新時',
    PRIMARY KEY (`id`),
    UNIQUE (`mail`),
    INDEX `idx_auth_token` (`id` ASC)
    );