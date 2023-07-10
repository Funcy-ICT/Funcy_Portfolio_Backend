CREATE TABLE IF NOT EXISTS `skills` (
    `skill_name` VARCHAR(128) NOT NULL COMMENT 'スキルの名前',
    `user_id` VARCHAR(64) NOT NULL COMMENT 'ユーザID',
    foreign key (user_id) references users(id),
    PRIMARY KEY (`skill_name`, `user_id`)
    );
