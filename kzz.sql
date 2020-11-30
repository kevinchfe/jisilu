CREATE TABLE `blog_kzz` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `bond_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '可转债id',
  `bond_nm` varchar(20) NOT NULL COMMENT '可转债name',
  `pb` decimal(5,2) NOT NULL DEFAULT '0.00' COMMENT 'pb',
  `price` decimal(8,3) unsigned NOT NULL DEFAULT '0.000' COMMENT '现价',
  `premium_rt` varchar(10) NOT NULL DEFAULT '0.00' COMMENT '溢价率',
  `orig_iss_amt` decimal(7,3) unsigned NOT NULL DEFAULT '0.000' COMMENT '剩余规模(亿元)',
  `volume` decimal(8,2) unsigned NOT NULL COMMENT '成交额(万元)',
  `turnover_rt` decimal(6,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '换手率',
  `ytm_rt` varbinary(10) NOT NULL DEFAULT '0.00' COMMENT '到期税前收益',
  `dblow` decimal(6,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '双低值',
  `ytm_premium_rt` decimal(10,3) NOT NULL DEFAULT '0.000' COMMENT '50%ytm+50%溢价率',
  PRIMARY KEY (`id`),
  KEY `bond_id` (`bond_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='可转债总表';