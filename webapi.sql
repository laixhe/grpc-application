CREATE DATABASE `webapi` CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_unicode_ci';

CREATE TABLE `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `type_id` int NOT NULL DEFAULT '0' COMMENT '类型 1普通',
  `account` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '账号',
  `mobile` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '手机号',
  `email` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `password` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '密码',
  `nickname` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '头像地址',
  `sex` int NOT NULL DEFAULT '0' COMMENT '性别 0未填写 1男 2女',
  `states` int NOT NULL DEFAULT '0' COMMENT '状态 0封禁 1正常',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `user_account_idx` (`account`),
  KEY `user_mobile_idx` (`mobile`),
  KEY `user_email_idx` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户';

CREATE TABLE `user_extend` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `uid` int unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `birthday` int NOT NULL DEFAULT '0' COMMENT '生日(年月日)',
  `height` int NOT NULL DEFAULT '0' COMMENT '身高(cm)',
  `weight` int NOT NULL DEFAULT '0' COMMENT '体重(kg)',
  PRIMARY KEY (`id`),
  KEY `user_extend_uid_idx` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户扩展';

CREATE TABLE `user_third_party` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `uid` int unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `wechat_unionid` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '微信unionid',
  `wechat_openid` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '微信openid',
  PRIMARY KEY (`id`),
  KEY `user_third_party_uid_idx` (`uid`),
  KEY `user_third_party_wechat_openid_idx` (`wechat_openid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户第三方';

