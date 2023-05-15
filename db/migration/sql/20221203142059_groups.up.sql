CREATE TABLE IF NOT EXISTS `groups` (
    `id` VARCHAR(64) NOT NULL COMMENT 'グループID',
    `name` VARCHAR(128) NOT NULL COMMENT 'グループ名',
    PRIMARY KEY (`id`)
    );

CREATE TABLE IF NOT EXISTS `group_member` (
    `group_id` VARCHAR(64) NOT NULL COMMENT 'グループID',
    `user_id` VARCHAR(64) NOT NULL COMMENT 'ユーザID',
    `role` VARCHAR(64) NOT NULL COMMENT '役職',
    `status` boolean NOT NULL COMMENT 'アクティブユーザかどうか',
    foreign key (user_id) references users(id),
    foreign key (group_id) references groups(id),
    PRIMARY KEY (`group_id`, `user_id`)
    );
