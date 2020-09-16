CREATE TABLE `accounts`
(
    `id`                 bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_at`         datetime(3)  DEFAULT NULL,
    `updated_at`         datetime(3)  DEFAULT NULL,
    `deleted_at`         datetime(3)  DEFAULT NULL,
    `email`              varchar(100)    NOT NULL,
    `username`           varchar(100)    NOT NULL,
    `encrypted_password` longtext        NOT NULL,
    `nickname`           varchar(100) DEFAULT NULL,
    `avatar`             longtext,
    `status`             tinyint(1)   DEFAULT '0',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_accounts_email` (`email`),
    UNIQUE KEY `idx_accounts_username` (`username`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;
