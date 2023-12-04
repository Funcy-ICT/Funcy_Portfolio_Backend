CREATE TABLE IF NOT EXISTS `group_skills` (
    `group_id` VARCHAR(64) NOT NULL COMMENT 'グループID',
    `skill_name` VARCHAR(256) NOT NULL COMMENT 'スキル名',
    foreign key (group_id) references groups(id),
    PRIMARY KEY (`group_id`, `skill_name`)
    );