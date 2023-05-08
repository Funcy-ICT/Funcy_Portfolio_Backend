CREATE TABLE IF NOT EXISTS `sns` (
    `id` VARCHAR(64) NOT NULL COMMENT 'SNS 登録 ID',
    `user_id` VARCHAR(64) NOT NULL COMMENT 'ユーザID',
    `sns` VARCHAR(64) NOT NULL COMMENT 'snsのリンク',
    foreign key (user_id) references users(id),
    PRIMARY KEY (`user_id`, `sns`)
    );
