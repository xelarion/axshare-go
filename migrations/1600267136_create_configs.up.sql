CREATE TABLE `configs`
(
    `id`                  bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_at`          datetime(3) DEFAULT NULL,
    `updated_at`          datetime(3) DEFAULT NULL,
    `is_valid`            tinyint(1)  DEFAULT NULL,
    `site_name`           varchar(255),
    `icp_record_no`       varchar(255),
    `icp_record_link`     longtext,
    `copyright`           varchar(255),
    `file_release_dir`    varchar(255),
    `web_domain`          varchar(255),
    `qiniu_access_key`    longtext,
    `qiniu_secret_key`    longtext,
    `qiniu_bucket`        varchar(255),
    `qiniu_bucket_domain` varchar(255),
    `qiniu_upload_url`    longtext,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;