CREATE TABLE `comment` (
  `id` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'コメントID',
  `user_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'ユーザーID',
  `works_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '作品ID',
  `text` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '本文',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
