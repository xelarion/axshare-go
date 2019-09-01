--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users`
(
    `id`                     bigint(20)   NOT NULL AUTO_INCREMENT,
    `email`                  varchar(255) NOT NULL DEFAULT '',
    `encrypted_password`     varchar(255) NOT NULL DEFAULT '',
    `reset_password_token`   varchar(255)          DEFAULT NULL,
    `reset_password_sent_at` datetime              DEFAULT NULL,
    `remember_created_at`    datetime              DEFAULT NULL,
    `created_at`             datetime     NOT NULL,
    `updated_at`             datetime     NOT NULL,
    `nickname`               varchar(255)          DEFAULT NULL,
    `avatar`                 varchar(255)          DEFAULT NULL,
    `description`            text,
    `status`                 int(11)               DEFAULT '1',
    `sign_in_count`          int(11)      NOT NULL DEFAULT '0',
    `current_sign_in_at`     datetime              DEFAULT NULL,
    `last_sign_in_at`        datetime              DEFAULT NULL,
    `current_sign_in_ip`     varchar(255)          DEFAULT NULL,
    `last_sign_in_ip`        varchar(255)          DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `index_users_on_email` (`email`),
    UNIQUE KEY `index_users_on_reset_password_token` (`reset_password_token`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 4
  DEFAULT CHARSET = utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users`
    DISABLE KEYS */;
INSERT INTO `users`
VALUES (1, 'admin@vicw.com', '$2a$11$wiguEDmPD0nLAoljtL5/i.ExG7i7Zk1LcVQWjnI/Wso9nXCnc2mAu', NULL, NULL, NULL,
        '2019-07-13 21:27:04', '2019-08-19 10:20:37', 'admin', NULL, NULL, 1, 14,
        '2019-08-19 10:20:37', '2019-08-18 14:14:05', '127.0.0.1', '127.0.0.1');
/*!40000 ALTER TABLE `users`
    ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `attachments`
--

DROP TABLE IF EXISTS `attachments`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `attachments`
(
    `id`             bigint(20) NOT NULL AUTO_INCREMENT,
    `reference_id`   int(11)      DEFAULT NULL,
    `reference_type` varchar(255) DEFAULT NULL,
    `key`            varchar(255) DEFAULT NULL,
    `created_at`     datetime   NOT NULL,
    `updated_at`     datetime   NOT NULL,
    `user_id`        int(11)      DEFAULT NULL,
    `link`           text,
    `desc`           text,
    PRIMARY KEY (`id`),
    KEY `index_attachments_on_reference_id` (`reference_id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 4
  DEFAULT CHARSET = utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `attachments`
--

LOCK TABLES `attachments` WRITE;
/*!40000 ALTER TABLE `attachments`
    DISABLE KEYS */;
INSERT INTO `attachments`
VALUES (1, 1, 'Axure', '621d56d0b26e416b736716f2f2c50277', '2019-08-18 14:18:47', '2019-08-18 14:18:47', 1, NULL, ''),
       (2, 2, 'Axure', '658d344d5653d247c7132103f897fe14', '2019-08-18 17:05:24', '2019-08-18 17:05:24', 1, NULL, ''),
       (3, 1, 'Axure', '495e222f945254f816c0786f9b8563bb', '2019-08-18 17:05:48', '2019-08-18 17:05:48', 1, NULL, '');
/*!40000 ALTER TABLE `attachments`
    ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `axure_groups`
--

DROP TABLE IF EXISTS `axure_groups`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `axure_groups`
(
    `id`         bigint(20) NOT NULL AUTO_INCREMENT,
    `owner_id`   int(11)      DEFAULT NULL,
    `name`       varchar(255) DEFAULT NULL,
    `desc`       text,
    `priority`   int(11)      DEFAULT NULL,
    `created_at` datetime   NOT NULL,
    `updated_at` datetime   NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 2
  DEFAULT CHARSET = utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `axure_groups`
--

LOCK TABLES `axure_groups` WRITE;
/*!40000 ALTER TABLE `axure_groups`
    DISABLE KEYS */;
INSERT INTO `axure_groups`
VALUES (1, NULL, 'simsky', 'simsky axures', NULL, '2019-08-18 14:13:50', '2019-08-18 14:13:50');
/*!40000 ALTER TABLE `axure_groups`
    ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `axures`
--

DROP TABLE IF EXISTS `axures`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `axures`
(
    `id`                bigint(20) NOT NULL AUTO_INCREMENT,
    `name`              varchar(255) DEFAULT NULL,
    `link`              text,
    `desc`              text,
    `user_id`           int(11)      DEFAULT NULL,
    `axure_category_id` int(11)      DEFAULT NULL,
    `axure_group_id`    int(11)      DEFAULT NULL,
    `created_at`        datetime   NOT NULL,
    `updated_at`        datetime   NOT NULL,
    `uuid`              varchar(255) DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `index_axures_on_user_id` (`user_id`),
    KEY `index_axures_on_axure_category_id` (`axure_category_id`),
    KEY `index_axures_on_axure_group_id` (`axure_group_id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 3
  DEFAULT CHARSET = utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `axures`
--

LOCK TABLES `axures` WRITE;
/*!40000 ALTER TABLE `axures`
    DISABLE KEYS */;
INSERT INTO `axures`
VALUES (1, NULL, NULL, '反对是否但是', 1, NULL, 1, '2019-08-18 14:18:47', '2019-08-18 14:18:47',
        '0bf66c27-821e-4412-8649-14fcf7699029'),
       (2, NULL, NULL, '方法', 1, NULL, 1, '2019-08-18 17:05:24', '2019-08-18 17:05:24',
        '62376197-df2a-4314-a735-2be220a68776');
/*!40000 ALTER TABLE `axures`
    ENABLE KEYS */;
UNLOCK TABLES;