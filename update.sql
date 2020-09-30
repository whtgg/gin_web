-- 建表SQL

drop table if exists sites;
CREATE TABLE `sites` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '网站所有人',
  `category_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '分类ID',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '网站名称',
  `keywords` varchar(255) NOT NULL DEFAULT '' COMMENT '网站关键字',
  `description` varchar(255) NOT NULL DEFAULT '' COMMENT '网站内容描述',
  `domain` varchar(255) NOT NULL DEFAULT '' COMMENT '域名',
  `site_url` varchar(255) NOT NULL DEFAULT '' COMMENT '访问地址',
  `qr` varchar(255) NOT NULL DEFAULT '' COMMENT '二维码',
  `db` varchar(255) NOT NULL DEFAULT '' COMMENT '数据库配置信息',
	`qq` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'QQ',
	`phone` varchar(255) NOT NULL DEFAULT '' COMMENT '联系电话',
	`icon` varchar(255) NOT NULL DEFAULT '' COMMENT '网站ICON',
  `logo` varchar(255) NOT NULL DEFAULT '' COMMENT '网站LOGO',
  `image` varchar(255) NOT NULL DEFAULT '' COMMENT '封面图片地址',
  `image_alt` varchar(255) NOT NULL DEFAULT '' COMMENT '封面图片描述',
  `thumbnail` varchar(255) NOT NULL DEFAULT '' COMMENT '封面图片缩略图',
	`notice` varchar(255) NOT NULL DEFAULT '' COMMENT '警示通知',
  `safe_tips` text COMMENT '风险提示',
  `script_top` text COMMENT '顶部脚本',
  `script_append` text COMMENT '统计代码',
  `script_marketing` text COMMENT '营销QQ代码',
  `tpl_id` int(11) NOT NULL DEFAULT '0' COMMENT '模板',
  `sort` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `web_type` int(11) unsigned NOT NULL DEFAULT '1' COMMENT '网站类型: 1=>PC; 0=>移动端',
  `state` int(11) unsigned NOT NULL DEFAULT '1' COMMENT '网站状态: 1=>启用; 0=>禁用',
  `status` int(11) unsigned NOT NULL DEFAULT '1' COMMENT '平台管理状态: 1=>启用; 0=>禁用',
  `created_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `updated_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_category_id` (`category_id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='网站信息';


drop table if exists users;
CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `login_name` varchar(64) DEFAULT '' COMMENT '登录名',
  `login_pwd` varchar(256) DEFAULT '' COMMENT '密码',
  `access_token` varchar(128) DEFAULT '' COMMENT '令牌',
  `expiry` int(11) unsigned DEFAULT '0' COMMENT '令牌有效期时间戳',
  `nickname` varchar(64) DEFAULT '' COMMENT '昵称',
  `mobile` varchar(64) DEFAULT '' COMMENT '手机号',
  `avatar` varchar(255) DEFAULT '' COMMENT '头像',
  `mail` varchar(64) DEFAULT '' COMMENT 'mail',
  `qq` varchar(64) DEFAULT '' COMMENT 'qq',
  `sex` int(11) unsigned DEFAULT '0' COMMENT '性别: 0=>女; 1=>男; 2=>保密',
  `amount` int(11) unsigned DEFAULT '0' COMMENT '余额',
  `vip` int(11) DEFAULT '0' COMMENT 'VIP等级: 0=>0级; 1=>1级; 2=>2级',
  `point` int(11) unsigned DEFAULT '0' COMMENT '积分',
  `ua` varchar(255) DEFAULT '' COMMENT '浏览器 User-Agent',
  `ip` varchar(255) DEFAULT '' COMMENT '来源 IP',
	`safe_code` varchar(255) DEFAULT '' COMMENT '安全码',
  `status` int(11) unsigned NOT NULL DEFAULT '1' COMMENT '状态: 0=>禁用; 1=>启用; 2=>黑名单; 3=>禁言',
  `created_at` int(11) unsigned DEFAULT '0' COMMENT '创建时间戳',
  `updated_at` int(11) unsigned DEFAULT '0' COMMENT '更新时间戳',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';


DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `parent_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '父ID',
  `category_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '分类ID',
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '作者UID',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '页面标题',
  `keywords` varchar(255) NOT NULL DEFAULT '' COMMENT '页面关键字',
  `description` varchar(255) NOT NULL DEFAULT '' COMMENT '页面内容描述',
  `ads_pic` varchar(255) NOT NULL DEFAULT '' COMMENT '广告图片地址',
  `ads_target` varchar(255) NOT NULL DEFAULT '' COMMENT '广告图片链接地址',
  `image` varchar(255) NOT NULL DEFAULT '' COMMENT '封面图片地址',
  `image_alt` varchar(255) NOT NULL DEFAULT '' COMMENT '封面图片描述',
  `thumbnail` varchar(255) NOT NULL DEFAULT '' COMMENT '封面图片缩略图',
  `subtitle` varchar(255) NOT NULL DEFAULT '' COMMENT '子标题',
  `body` longtext NOT NULL COMMENT '正文',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
  `identity` varchar(255) NOT NULL DEFAULT '' COMMENT '身份ID',
  `publish_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '发布时间戳',
  `subscript` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '角标: 0=>无角标; 1=>标新; 2=>热门',
  `sort` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `istop` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '置顶: 1=>置顶; 0=>不置顶',
  `pv` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '浏览量',
  `status` int(11) unsigned NOT NULL DEFAULT '1' COMMENT '状态: 1=>启用; 0=>禁用',
  `created_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `updated_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_category_id` (`category_id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章表';

DROP TABLE IF EXISTS `categorys`;
CREATE TABLE `categorys` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '文章分类ID',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '分类标题',
  `parent_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '分类',
  `category_type` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '分类类型',
  `keywords` varchar(255) NOT NULL DEFAULT '' COMMENT '关键字',
  `slug` varchar(255) NOT NULL DEFAULT '' COMMENT '栏目路径',
  `description` varchar(255) NOT NULL COMMENT '详细描述',
  `body` longtext NOT NULL COMMENT '内容',
  `ads_pic` varchar(255) NOT NULL DEFAULT '' COMMENT '广告图片地址',
  `ads_target` varchar(255) NOT NULL DEFAULT '' COMMENT '广告图片链接地址',
  `image` varchar(255) NOT NULL DEFAULT '' COMMENT '封面图片地址',
  `image_alt` varchar(255) NOT NULL DEFAULT '' COMMENT '封面图片描述',
  `thumbnail` varchar(255) NOT NULL DEFAULT '' COMMENT '封面图片缩略图',
  `subtitle` varchar(255) NOT NULL DEFAULT '' COMMENT '子标题',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
  `identity` varchar(255) NOT NULL DEFAULT '' COMMENT '身份ID',
  `publish_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '发布时间戳',
  `subscript` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '角标: 0=>无角标; 1=>标新; 2=>热门',
  `sort` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `istop` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '置顶: 1=>置顶; 0=>不置顶',
  `pv` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '浏览量',
  `status` int(11) unsigned NOT NULL DEFAULT '1' COMMENT '状态: 1=>启用; 0=>禁用',
  `created_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `updated_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='分类表';


DROP TABLE IF EXISTS `relate_article`;
CREATE TABLE `relate_article` (
  `article_id` int(20) NOT NULL COMMENT '文章ID',
  `collection` varchar(255) NOT NULL DEFAULT '' COMMENT '相关文章ID集合',
  `created_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `updated_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  PRIMARY KEY (`article_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章相关文章';


DROP TABLE IF EXISTS `relate_tag_article`;
CREATE TABLE `relate_tag_article` (
  `tag_id` int(20) unsigned NOT NULL COMMENT '词ID',
  `collection` varchar(255) NOT NULL COMMENT '相关词文章ID集合',
  `degree` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '关联度',
  `created_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `updated_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  PRIMARY KEY (`tag_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='词相关文章表';


DROP TABLE IF EXISTS `relate_tags`;
CREATE TABLE `relate_tags` (
  `tag_id` int(20) unsigned NOT NULL COMMENT '词ID',
  `collection` varchar(255) NOT NULL COMMENT '相关词与词ID集合',
  `created_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `updated_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  PRIMARY KEY (`tag_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='词相关的词';

DROP TABLE IF EXISTS `tags`;
CREATE TABLE `tags` (
  `id` int(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '词ID',
  `batch` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '批次',
  `score` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '积分',
  `tag` varchar(100) NOT NULL DEFAULT '' COMMENT '关键词',
  `pic_hash` varchar(255) NOT NULL DEFAULT '' COMMENT '图片hash,一张或多张图片',
  `created_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `updated_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  PRIMARY KEY (`id`),
  UNIQUE key (`tag`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='词库表';

-- END TIME 2020-06-13 16:01:46
