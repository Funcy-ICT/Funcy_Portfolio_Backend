CREATE TABLE IF NOT EXISTS `user_profile` (
    `user_id` VARCHAR(64) NOT NULL COMMENT 'ユーザID',
    `header_image` VARCHAR(128) COMMENT 'ヘッダ画像URL',
    `bio` VARCHAR(128) COMMENT 'ユーザの自己紹介',
    foreign key (user_id) references users(id),
    PRIMARY KEY (`user_id`)
    );
