CREATE TABLE `user_tab` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(45) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名',
  `password` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL,
  `nickname` varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户昵称',
  `email` varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户邮箱',
  `avatar` varchar(256) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username_UNIQUE` (`username`),
  UNIQUE KEY `username_UNIQUE` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;