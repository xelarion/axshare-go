CREATE TABLE `attachments`
(
    `id`             bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_at`     datetime(3)     DEFAULT NULL,
    `updated_at`     datetime(3)     DEFAULT NULL,
    `deleted_at`     datetime(3)     DEFAULT NULL,
    `desc`           longtext,
    `link`           varchar(255),
    `file_hash`      varchar(255),
    `release_status` tinyint(1)      DEFAULT '0',
    `release_error`  longtext,
    `axure_id`       bigint unsigned DEFAULT '0',
    `account_id`     bigint unsigned DEFAULT '0',
    PRIMARY KEY (`id`),
    KEY `idx_attachments_axure_id` (`axure_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;