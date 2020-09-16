CREATE TABLE `axures`
(
    `id`             bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_at`     datetime(3)     DEFAULT NULL,
    `updated_at`     datetime(3)     DEFAULT NULL,
    `deleted_at`     datetime(3)     DEFAULT NULL,
    `name`           varchar(255),
    `secret_key`     varchar(255),
    `axure_group_id` bigint unsigned DEFAULT '0',
    PRIMARY KEY (`id`),
    KEY `idx_axures_axure_group_id` (`axure_group_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;