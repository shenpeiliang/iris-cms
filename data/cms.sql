-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        8.0.21 - MySQL Community Server - GPL
-- 服务器操作系统:                      Linux
-- HeidiSQL 版本:                  11.3.0.6295
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- 导出 cms 的数据库结构
DROP DATABASE IF EXISTS `cms`;
CREATE DATABASE IF NOT EXISTS `cms` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `cms`;

-- 导出  表 cms.hs_article 结构
DROP TABLE IF EXISTS `hs_article`;
CREATE TABLE IF NOT EXISTS `hs_article` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID编号',
  `title` varchar(250) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '标题',
  `description` varchar(150) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '网页描述',
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '内容',
  `img` varchar(120) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '主图',
  `is_show` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否显示',
  `add_time` int unsigned NOT NULL DEFAULT '0' COMMENT '发布时间',
  `dateline` int unsigned NOT NULL DEFAULT '0' COMMENT '操作时间',
  `paixu` int unsigned NOT NULL DEFAULT '0' COMMENT '排序号',
  `click` int unsigned NOT NULL DEFAULT '0' COMMENT '点击',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  cms.hs_article 的数据：~2 rows (大约)
/*!40000 ALTER TABLE `hs_article` DISABLE KEYS */;
INSERT INTO `hs_article` (`id`, `title`, `description`, `content`, `img`, `is_show`, `add_time`, `dateline`, `paixu`, `click`) VALUES
	(7, '浓眉大眼的半导体芯片龙头，1000来亿的闻泰科技，3月24日', '浓眉大眼的半导体芯片龙头，1000来亿的闻泰科技，3月24日下午尾盘突然闪崩，跳水跌停。消息面目前没看到什么明显利空，投资者一脸懵。截至2021年9月底，闻泰科技股东户数12.71万户。而在盘后的龙虎榜席位上，上演机构席位的巅峰对决，前五大买卖席位当中，各有4家机构席位，实属罕见。从卖出金额来看，一', '                                    <p style="margin-top: 0px; margin-bottom: 0px; white-space: normal; padding: 0px; overflow-y: auto; max-width: 100%; line-height: 24px;">浓眉大眼的半导体芯片龙头，1000来亿的闻泰科技，3月24日下午尾盘突然闪崩，跳水跌停。</p>\n<p style="margin-top: 0px; margin-bottom: 0px; white-space: normal; padding: 0px; overflow-y: auto; max-width: 100%; line-height: 24px;"><img src="../../../static/upload/editor/1648538895199508053.jpg" alt="" width="1149" height="488"></p>\n<p style="margin-top: 0px; margin-bottom: 0px; white-space: normal; padding: 0px; overflow-y: auto; max-width: 100%; line-height: 24px;"> </p>\n<p style="margin-top: 0px; margin-bottom: 0px; white-space: normal; padding: 0px; overflow-y: auto; max-width: 100%; line-height: 24px;">消息面目前没看到什么明显利空，投资者一脸懵。截至2021年9月底，闻泰科技股东户数12.71万户。而在盘后的龙虎榜席位上，上演机构席位的巅峰对决，前五大买卖席位当中，各有4家机构席位，实属罕见。</p>\n<p style="margin-top: 0px; margin-bottom: 0px; white-space: normal; padding: 0px; overflow-y: auto; max-width: 100%; line-height: 24px;"> </p>\n<p style="margin-top: 0px; margin-bottom: 0px; white-space: normal; padding: 0px; overflow-y: auto; max-width: 100%; line-height: 24px;">是的是的是的</p>\n<p style="margin-top: 0px; margin-bottom: 0px; white-space: normal; padding: 0px; overflow-y: auto; max-width: 100%; line-height: 24px;"> </p>\n<p style="margin-top: 0px; margin-bottom: 0px; white-space: normal; padding: 0px; overflow-y: auto; max-width: 100%; line-height: 24px;">从卖出金额来看，一家机构席位卖出6.78亿元，按当天成交看，卖出股份数当超过650万股，而在去年底持股超过这个数量的基金数量不大，几只顶流知名基金位列其中。</p>\n<p style="margin-top: 0px; margin-bottom: 0px; white-space: normal; padding: 0px; overflow-y: auto; max-width: 100%; line-height: 24px;"> </p>\n<p style="margin-top: 0px; margin-bottom: 0px; white-space: normal; padding: 0px; overflow-y: auto; max-width: 100%; line-height: 24px;"> </p>\n<p style="margin-top: 0px; margin-bottom: 0px; white-space: normal; padding: 0px; overflow-y: auto; max-width: 100%; line-height: 24px;">当日晚间，闻泰科技紧急发布回应公告。公司称，目前公司生产经营正常，各项主要业务继续保持稳健的发展势头，没有影响股价的重大事件。并对公司三大业务做了具体介绍。</p>\n<p style="margin-top: 0px; margin-bottom: 0px; white-space: normal; padding: 0px; overflow-y: auto; max-width: 100%; line-height: 24px;"> </p>\n<p style="margin-top: 0px; margin-bottom: 0px; white-space: normal; padding: 0px; overflow-y: auto; max-width: 100%; line-height: 24px;"> </p>\n<p style="margin-top: 0px; margin-bottom: 0px; white-space: normal; padding: 0px; overflow-y: auto; max-width: 100%; line-height: 24px; text-align: center;"><span style="max-width: 100%; color: rgb(0, 82, 255);"><strong style="max-width: 100%;">千亿芯片大牛股：尾盘突然闪崩跌停</strong></span></p>\n<p style="margin-top: 0px; margin-bottom: 0px; white-space: normal; padding: 0px; overflow-y: auto; max-width: 100%; line-height: 24px;"> </p>\n<p style="margin-top: 0px; margin-bottom: 0px; white-space: normal; padding: 0px; overflow-y: auto; max-width: 100%; line-height: 24px;"> </p>\n<p style="margin-top: 0px; margin-bottom: 0px; white-space: normal; padding: 0px; overflow-y: auto; max-width: 100%; line-height: 24px;">3月18日，闻泰科技开盘还好好的，以103.38元的价格微跌0.12%，这个价格也成为了当天最高价，然后一路走低，尾盘更是突然闪崩跌停，跌幅10.00%，报收93.15元，当天成交额25.9亿元，收盘市值1161亿元。</p>\n<p style="white-space: normal;"> </p>\n<p> </p>\n                            ', '/static/upload/default/1648432012129767507.jpg', 1, 1648190743, 1648539071, 20, 0),
	(11, '“躺”在中国乳业背后赚钱：抽走75%利润，蒙牛伊利都为其打工', '提到中国奶业，率先考虑到的就是伊利和蒙牛两大企业，二者是中国奶业的top2。而牛奶作为营养产品，几乎是家家日常都会购买的产品，所以我国奶业的营收非常高。&nbsp;&nbsp;但是这笔钱的大头并没有被蒙牛、伊利等奶商拿走，而是被一家名为利乐的瑞典公司拿走了，不少网友可能要问为什么？今天就让我们来探索', '                                                                                                            <div class="index-module_textWrap_3ygOc ">\n<p data-from-paste="1"><span data-diagnose-id="f9b72e69bab75b6bf45812e390826cba">提到中国奶业，率先考虑到的就是伊利和蒙牛两大企业，二者是中国奶业的top2。而牛奶作为营养产品，几乎是家家日常都会购买的产品，所以我国奶业的营收非常高。</span></p>\n<p> </p>\n</div>\n<div class="index-module_textWrap_3ygOc ">\n<p> </p>\n<p data-from-paste="1"><span data-from-paste="1" data-diagnose-id="553a2b90787f3bd74ffd43536761dfaf">但是这笔钱的大头并没有被蒙牛、伊利等奶商拿走，</span><strong data-from-paste="1" data-diagnose-id="1ca516d0c7d04e387adb9928afb87174"><span data-from-paste="1" data-diagnose-id="a8549b0637b6eee855d70a2e1ea264cf">而是被一家名为利乐的瑞典公司拿走了，</span></strong><span data-from-paste="1" data-diagnose-id="6615807af7e4788d1574f7e2d113ff4f">不少网友可能要问为什么？今天就让我们来探索一下究竟是为什么？</span></p>\n<p> </p>\n</div>\n<div class="index-module_textWrap_3ygOc ">\n<h3 class="pgc-h-arrow-right" data-from-paste="1" data-diagnose-id="de08c70f69a0affcc48048505b3b150c">一、利乐</h3>\n<p> </p>\n</div>\n<div class="index-module_textWrap_3ygOc ">\n<p> </p>\n<p data-from-paste="1"><span data-from-paste="1" data-diagnose-id="53115e87aa77d1647802bcdd3cc91a91">利乐是瑞典的一家公司，创始人是鲁宾·劳辛，</span><strong data-from-paste="1" data-diagnose-id="0c4a72c92b2e3cf46eeaa4516124a8ae"><span data-from-paste="1" data-diagnose-id="ae5bb6194c307787e68acbbb282ea052">可以说鲁宾·劳辛是一个非常有远见的商人。</span></strong><span data-from-paste="1" data-diagnose-id="93fb05065f2bfe366c7a90760dd7950c">鲁宾·劳辛出生于1895年，拿到了哥伦比亚大学的硕士学位，是一位高级知识分子，他的创业灵感来源于大学时的一次偶然发现。</span></p>\n</div>\n                            \n                            \n                            ', '/static/upload/default/1649835216393044134.jpg', 1, 1648606831, 1650960948, 28, 0);
/*!40000 ALTER TABLE `hs_article` ENABLE KEYS */;

-- 导出  表 cms.hs_user 结构
DROP TABLE IF EXISTS `hs_user`;
CREATE TABLE IF NOT EXISTS `hs_user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID编号',
  `user_name` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
  `small_img` varchar(150) NOT NULL DEFAULT '' COMMENT '头像',
  `email` varchar(60) NOT NULL DEFAULT '' COMMENT '邮箱',
  `password` varchar(72) NOT NULL DEFAULT '' COMMENT '密码',
  `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机',
  `add_time` int unsigned NOT NULL DEFAULT '0' COMMENT '注册时间',
  `dateline` int unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `is_show` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '是否审核',
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_name_UNIQUE` (`user_name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- 正在导出表  cms.hs_user 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `hs_user` DISABLE KEYS */;
INSERT INTO `hs_user` (`id`, `user_name`, `small_img`, `email`, `password`, `mobile`, `add_time`, `dateline`, `is_show`) VALUES
	(1, 'admin', '', '', '$2a$10$pB00x3W/gVuB9qeekJAXgud3sjDfcKlDgHTE5V.ClZ4dHcfX3ENty', '', 1649298501, 1649833922, 1);
/*!40000 ALTER TABLE `hs_user` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
