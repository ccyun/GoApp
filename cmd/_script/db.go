package main

import (
	"bbs_server/application/library/conf"
	"strings"

	"fmt"

	"github.com/astaxie/beego/orm"
)

//DB 表结构更新
type DB struct {
}

//alterTable 更新表结构
func (db *DB) alterTable() error {
	sqls := []string{

		//更新BBS表
		"ALTER TABLE `bbs_bbs` MODIFY COLUMN `publish_scope`  mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL AFTER `link`",

		//更新BBS任务表
		"ALTER TABLE `bbs_bbs_task` ADD COLUMN `reply_type`  tinyint(1) UNSIGNED NOT NULL COMMENT '反馈者类型0用户，1组织' AFTER `bbs_id`",
		"ALTER TABLE `bbs_bbs_task` ADD COLUMN `is_resubmit`  tinyint(1) UNSIGNED NOT NULL COMMENT '是否允许重复提交0否，1是' AFTER `reply_type`",
		"ALTER TABLE `bbs_bbs_task` ADD COLUMN `audit_type`  tinyint(1) UNSIGNED NOT NULL COMMENT '0无需审核，1整单审核，2每个子任务单独审核' AFTER `is_resubmit`",
		"ALTER TABLE `bbs_bbs_task` ADD COLUMN `is_flow`  tinyint(1) UNSIGNED NOT NULL COMMENT '是否流程审核（逐级审核）' AFTER `audit_type`",
		"ALTER TABLE `bbs_bbs_task` ADD COLUMN `audit_user_ids`  text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '审核者列表' AFTER `is_flow`",
		"ALTER TABLE `bbs_bbs_task` MODIFY COLUMN `end_time`  bigint(13) UNSIGNED NOT NULL COMMENT '填写（反馈）截止时间' AFTER `audit_user_ids`",
		"ALTER TABLE `bbs_bbs_task` ADD COLUMN `audit_end_time`  bigint(13) UNSIGNED NOT NULL COMMENT '审核截止时间' AFTER `allow_expired`",
		"ALTER TABLE `bbs_bbs_task` ADD COLUMN `reply_remind_at`  bigint(13) UNSIGNED NOT NULL COMMENT '反馈提醒时间' AFTER `audit_end_time`",
		"ALTER TABLE `bbs_bbs_task` ADD COLUMN `audit_remind_at`  bigint(13) UNSIGNED NOT NULL COMMENT '审核提醒时间' AFTER `reply_remind_at`",
		"ALTER TABLE `bbs_bbs_task` ADD COLUMN `is_redo`  tinyint(1) UNSIGNED NOT NULL COMMENT '审核不通过可以重做' AFTER `audit_remind_at`",
		"ALTER TABLE `bbs_bbs_task` ADD COLUMN `is_stop_reply`  tinyint(1) UNSIGNED NOT NULL COMMENT '是否停止反馈' AFTER `is_redo`",
		"ALTER TABLE `bbs_bbs_task` ADD COLUMN `stop_reply_at`  bigint(13) NOT NULL AFTER `is_stop_reply`",
		"ALTER TABLE `bbs_bbs_task` ADD COLUMN `stop_reply_user_id`  bigint(20) NOT NULL AFTER `stop_reply_at`",
		"ALTER TABLE `bbs_bbs_task` ADD COLUMN `is_done`  tinyint(1) UNSIGNED NOT NULL COMMENT '任务已完成，在提交反馈、审核时、检测并更新，所有人均已提交并审核' AFTER `stop_reply_user_id`",
		"ALTER TABLE `bbs_bbs_task` MODIFY COLUMN `close_at`  bigint(13) UNSIGNED NOT NULL COMMENT '关闭时间' AFTER `is_close`",
		"ALTER TABLE `bbs_bbs_task` DROP COLUMN `restriction`",
		"ALTER TABLE `bbs_bbs_task` DROP COLUMN `send_user_ids`",
		"ALTER TABLE `bbs_bbs_task` DROP COLUMN `is_cycle`",
		"ALTER TABLE `bbs_bbs_task` DROP COLUMN `cycle_rule`",
		"CREATE INDEX `audit_type` ON `bbs_bbs_task`(`audit_type`) USING BTREE ",
		"CREATE INDEX `end_time` ON `bbs_bbs_task`(`end_time`) USING BTREE ",
		"CREATE INDEX `audit_end_time` ON `bbs_bbs_task`(`audit_end_time`) USING BTREE ",
		"CREATE INDEX `is_stop_reply` ON `bbs_bbs_task`(`is_stop_reply`) USING BTREE ",
		"CREATE INDEX `is_redo` ON `bbs_bbs_task`(`is_redo`) USING BTREE ",
		"CREATE INDEX `is_end` ON `bbs_bbs_task`(`is_done`) USING BTREE ",
		"UPDATE `bbs_bbs_task` SET `audit_type`='2', `is_redo`='1'",

		//更新审核记录表
		"CREATE TABLE `bbs_bbs_task_audit` (`id`  bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT ,`site_id`  bigint(20) UNSIGNED NOT NULL ,`board_id`  bigint(20) UNSIGNED NOT NULL COMMENT '广播号ID' ,`bbs_id`  bigint(20) UNSIGNED NOT NULL COMMENT '广播ID，任务ID' ,`reply_id`  bigint(20) UNSIGNED NOT NULL COMMENT '反馈ID' ,`sub_task_id`  bigint(20) UNSIGNED NOT NULL COMMENT '子任务ID' ,`sub_reply_id`  bigint(20) UNSIGNED NOT NULL COMMENT '子反馈ID' ,`audit_user_id`  bigint(20) UNSIGNED NOT NULL COMMENT '审核者user_id' ,`audit_at`  bigint(13) UNSIGNED NOT NULL COMMENT '审核时间' ,`audit_score`  int(10) UNSIGNED NOT NULL COMMENT '审核评分' ,`audit_opinion`  text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '审核意见' ,`data`  text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '其他数据（单张审核意见，批注等）' ,`status`  tinyint(1) UNSIGNED NOT NULL COMMENT '审核状态，0未审核，1审核通过，2审核不通过' ,`created_at`  bigint(13) UNSIGNED NOT NULL COMMENT '创建审核任务的时间' ,PRIMARY KEY (`id`),INDEX `site_id` (`site_id`) USING BTREE ,INDEX `board_id` (`board_id`) USING BTREE ,INDEX `bbs_id` (`bbs_id`) USING BTREE ,INDEX `sub_task_id` (`sub_task_id`) USING BTREE ,INDEX `reply_id` (`reply_id`) USING BTREE ,INDEX `sub_reply_id` (`sub_reply_id`) USING BTREE ,INDEX `created_at` (`created_at`) USING BTREE ) ENGINE=InnoDB DEFAULT CHARACTER SET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=Compact",

		//更新选项表
		"CREATE TABLE `bbs_bbs_task_option` (`id`  bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT ,`site_id`  bigint(20) UNSIGNED NOT NULL ,`board_id`  bigint(20) UNSIGNED NOT NULL ,`bbs_id`  bigint(20) UNSIGNED NOT NULL ,`sub_task_id`  bigint(20) UNSIGNED NOT NULL COMMENT '子任务ID' ,`type`  tinyint(4) UNSIGNED NOT NULL COMMENT '1,单选，2多选' ,`group_name`  varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '矩阵选择组名称' ,`name`  varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '选项名称' ,`image`  varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '选项图片' ,`value`  int(10) NOT NULL COMMENT '选项值，数字' ,`is_ext_value`  tinyint(1) UNSIGNED NOT NULL COMMENT '是否有其他值' ,`selected`  tinyint(1) UNSIGNED NOT NULL COMMENT '是否默认选中' ,`show_sub_task_id`  bigint(20) UNSIGNED NOT NULL COMMENT '选中后显示哪个子任务' ,PRIMARY KEY (`id`),INDEX `site_id` (`site_id`) USING BTREE ,INDEX `board_id` (`board_id`) USING BTREE ,INDEX `bbs_id` (`bbs_id`) USING BTREE ,INDEX `sub_task_id` (`sub_task_id`) USING BTREE ,INDEX `value` (`value`) USING BTREE ) ENGINE=InnoDB DEFAULT CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=Compact",

		//更新选项值表
		"CREATE TABLE `bbs_bbs_task_option_value` (`id`  bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT ,`site_id`  bigint(20) UNSIGNED NOT NULL ,`board_id`  bigint(20) UNSIGNED NOT NULL ,`bbs_id`  bigint(20) UNSIGNED NOT NULL ,`sub_task_id`  bigint(20) UNSIGNED NOT NULL COMMENT '子任务ID' ,`reply_id`  bigint(20) UNSIGNED NOT NULL ,`sub_task_reply_id`  bigint(20) UNSIGNED NOT NULL COMMENT '子反馈ID' ,`value`  int(10) UNSIGNED NOT NULL COMMENT '选项值，按照顺序递增生成' ,`value_ext`  varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '其他值' ,PRIMARY KEY (`id`),INDEX `site_id` (`site_id`) USING BTREE ,INDEX `board_id` (`board_id`) USING BTREE ,INDEX `bbs_id` (`bbs_id`) USING BTREE ,INDEX `sub_task_id` (`sub_task_id`) USING BTREE ,INDEX `value` (`value`) USING BTREE ) ENGINE=InnoDB DEFAULT CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=Compact",

		//更新反馈表
		"ALTER TABLE `bbs_bbs_task_reply` ADD COLUMN `user_org_id`  bigint(20) NOT NULL COMMENT '反馈用户的组织ID' AFTER `user_id`",
		"ALTER TABLE `bbs_bbs_task_reply` MODIFY COLUMN `created_at`  bigint(13) UNSIGNED NOT NULL COMMENT '提交时间' AFTER `user_org_id`",
		"CREATE INDEX `user_org_id` ON `bbs_bbs_task_reply`(`user_org_id`) USING BTREE ",

		//更新子反馈表
		"ALTER TABLE `bbs_bbs_task_reply_sub` DROP INDEX `type`",
		"ALTER TABLE `bbs_bbs_task_reply_sub` MODIFY COLUMN `sub_task_id`  bigint(20) UNSIGNED NOT NULL COMMENT '子任务ID' AFTER `bbs_id`",
		"ALTER TABLE `bbs_bbs_task_reply_sub` MODIFY COLUMN `save_at`  bigint(13) UNSIGNED NOT NULL COMMENT '保存时间' AFTER `data`",
		"ALTER TABLE `bbs_bbs_task_reply_sub` MODIFY COLUMN `created_at`  bigint(13) UNSIGNED NOT NULL COMMENT '提交时间' AFTER `save_at`",
		"ALTER TABLE `bbs_bbs_task_reply_sub` MODIFY COLUMN `audit_opinion`  text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '审核描述' AFTER `audit_at`",
		"ALTER TABLE `bbs_bbs_task_reply_sub` MODIFY COLUMN `audit_user_id`  bigint(20) UNSIGNED NOT NULL COMMENT '审核用户' AFTER `audit_opinion`",
		"ALTER TABLE `bbs_bbs_task_reply_sub` ADD COLUMN `is_history`  tinyint(1) UNSIGNED NOT NULL COMMENT '是否历史子反馈，用于保存子反馈历史记录' AFTER `audit_user_id`",
		"ALTER TABLE `bbs_bbs_task_reply_sub` DROP COLUMN `type`",
		"CREATE UNIQUE INDEX `sub_reply` ON `bbs_bbs_task_reply_sub`(`sub_task_id`, `reply_id`, `user_id`, `is_history`) USING BTREE ",
		"DROP INDEX `sub_task_id` ON `bbs_bbs_task_reply_sub`",
		"CREATE INDEX `sub_task_id` ON `bbs_bbs_task_reply_sub`(`sub_task_id`) USING BTREE ",

		//子反馈扩展表
		"CREATE TABLE `bbs_bbs_task_reply_sub_ext` (`id`  bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT ,`site_id`  bigint(20) UNSIGNED NOT NULL ,`board_id`  bigint(20) UNSIGNED NOT NULL ,`bbs_id`  bigint(20) UNSIGNED NOT NULL ,`sub_task_id`  bigint(20) UNSIGNED NOT NULL ,`reply_id`  bigint(20) UNSIGNED NOT NULL ,`sub_task_reply_id`  bigint(20) UNSIGNED NOT NULL ,`value`  text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL ,`audit_opinion`  text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL ,`comments`  text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '批注' ,`status`  tinyint(1) UNSIGNED NOT NULL COMMENT '审核状态' ,PRIMARY KEY (`id`),INDEX `site_id` (`site_id`) USING BTREE ,INDEX `board_id` (`board_id`) USING BTREE ,INDEX `bbs_id` (`bbs_id`) USING BTREE ,INDEX `sub_task_id` (`sub_task_id`) USING BTREE ,INDEX `sub_task_reply_id` (`sub_task_reply_id`) USING BTREE ,INDEX `reply_id` (`reply_id`) USING BTREE ) ENGINE=InnoDB DEFAULT CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=Compact",

		//更新反馈用户tag表
		"CREATE TABLE `bbs_bbs_task_reply_tags` (`id`  bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT ,`site_id`  bigint(20) UNSIGNED NOT NULL ,`bbs_id`  bigint(20) NULL DEFAULT NULL ,`reply_id`  bigint(20) UNSIGNED NOT NULL ,`tag_id`  bigint(20) UNSIGNED NOT NULL ,`tag_name`  varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL ,`tag_code`  varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL ,`tag_enum_id`  bigint(20) UNSIGNED NOT NULL ,`tag_value`  varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL ,PRIMARY KEY (`id`),INDEX `reply_id` (`reply_id`) USING BTREE ,INDEX `tag` (`tag_id`, `tag_code`) USING BTREE ) ENGINE=InnoDB DEFAULT CHARACTER SET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=Compact",

		//更新子任务表
		"ALTER TABLE `bbs_bbs_task_sub` MODIFY COLUMN `description`  mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '任务描述' AFTER `title`",
		"ALTER TABLE `bbs_bbs_task_sub` MODIFY COLUMN `type`  enum('audio','image','checkbox','file','geo','radio','score','video','longtext','text') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '任务类型：文字、组图、音频' AFTER `description`",
		"ALTER TABLE `bbs_bbs_task_sub` ADD COLUMN `is_required`  tinyint(1) UNSIGNED NOT NULL COMMENT '是否必填' AFTER `restriction`",
		"ALTER TABLE `bbs_bbs_task_sub` ADD COLUMN `is_hide`  tinyint(1) UNSIGNED NOT NULL COMMENT '是否默认隐藏' AFTER `is_required`",
		"ALTER TABLE `bbs_bbs_task_sub` ADD COLUMN `is_audit`  tinyint(1) UNSIGNED NOT NULL COMMENT '是否需要审核' AFTER `is_hide`",
		"UPDATE `bbs_bbs_task_sub` SET `is_required`='1',`is_audit`='1'",

		//重命名msg老表
		"ALTER TABLE `bbs_msg` RENAME `bbs_msg2`",

		//更新队列表
		"ALTER TABLE `bbs_task` MODIFY COLUMN `task_type`  enum('bbs','taskAuditRemind','taskReply','taskAudit','taskClose') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'bbs广播表单任务,taskReply任务反馈提醒,taskAudit审核提醒,taskClose任务关闭' AFTER `customer_code`",
		"ALTER TABLE `bbs_task` ADD COLUMN `bbs_id`  bigint(20) NOT NULL AFTER `task_type`",
		"CREATE INDEX `bbs_id` ON `bbs_task`(`bbs_id`) USING BTREE ",

		//删除未读计数表
		"DROP TABLE `bbs_unread`",

		//更新非图片类的反馈到审核记录表
		"INSERT INTO bbs_bbs_task_audit(site_id,board_id,bbs_id,reply_id,sub_task_id,sub_reply_id,audit_user_id,audit_at,audit_opinion,status,created_at)(SELECT reply.site_id,reply.board_id,reply.bbs_id,reply.reply_id,reply.sub_task_id,reply.id sub_reply_id,reply.audit_user_id,reply.created_at audit_at,reply.audit_opinion,reply.status,reply.created_at FROM  bbs_bbs_task_reply_sub reply INNER JOIN bbs_bbs_task_sub task ON reply.sub_task_id=task.id WHERE reply.status>-1 order by reply.id asc)",

		//删除讨论组预览公告的feed
		"DELETE feed FROM bbs_feed feed,bbs_bbs bbs WHERE feed.bbs_id = bbs.id AND bbs.discuss_id > 0 AND bbs.type = 'preview'",
	}

	for _, sql := range sqls {
		if _, err := o.Raw(sql).Exec(); err != nil {
			return err
		}
	}
	//创建新的msg表
	dbShardDsns := conf.String("db_shard_dsns")
	tableNum, _ := conf.Int("db_shard_table_mun")
	dsns := strings.Split(dbShardDsns, ";")
	nodeNum := len(dsns)

	for i, dsn := range dsns {
		if err := orm.RegisterDataBase(fmt.Sprintf("msg%d", i), "mysql", dsn, 2, 2); err != nil {
			return err
		}
	}
	d := orm.NewOrm()
	for i := 0; i < tableNum; i++ {
		if err := d.Using(fmt.Sprintf("msg%d", (i / (tableNum / nodeNum)))); err != nil {
			return err
		}
		if _, err := d.Raw("CREATE TABLE `" + fmt.Sprintf("bbs_msg_%.4d", i) + "` (`id`  bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT ,`site_id`  bigint(20) UNSIGNED NOT NULL ,`board_id`  bigint(20) UNSIGNED NOT NULL ,`discuss_id`  bigint(20) UNSIGNED NOT NULL ,`bbs_id`  bigint(20) UNSIGNED NOT NULL ,`feed_type`  enum('bbs','form','task','taskReply','taskAudit','taskClose') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'bbs广播,form表单,task任务,taskReply任务反馈提醒,taskAudit审核提醒,taskClose任务关闭' ,`feed_id`  bigint(20) UNSIGNED NOT NULL ,`user_id`  bigint(20) UNSIGNED NOT NULL ,`user_org_id`  bigint(20) UNSIGNED NOT NULL ,`task_status` tinyint(1) unsigned NOT NULL,`is_read`  tinyint(1) UNSIGNED NOT NULL COMMENT '是否已读' ,`created_at`  bigint(13) UNSIGNED NOT NULL ,PRIMARY KEY (`id`),INDEX `board_id` (`board_id`) USING BTREE ,INDEX `feed_type` (`feed_type`) USING BTREE ,INDEX `user_id` (`user_id`) USING BTREE ,INDEX `bbs_id` (`bbs_id`) USING BTREE ,INDEX `is_read` (`is_read`) USING BTREE ,INDEX `discuss_id` (`discuss_id`) USING BTREE ) ENGINE=InnoDB DEFAULT CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=Compact").Exec(); err != nil {
			return err
		}
	}
	return nil
}

//后置处理
func (db *DB) clearTable() error {
	sqls := []string{
		//删除msg2表
		"DROP TABLE `bbs_msg2`",
		//删除图片类型的反馈内容
		"UPDATE bbs_bbs_task_reply_sub reply,bbs_bbs_task_sub task SET reply.data='' WHERE reply.sub_task_id=task.id and task.type='image'",
		//更新任务完成状态
		"update bbs_bbs_task task set task.is_done=1 where task.bbs_id in (select id from bbs_bbs bbs where bbs.category='task' and bbs.msg_count=(select count(reply.id) from bbs_bbs_task_reply reply where reply.bbs_id=bbs.id and reply.status=1))",
	}
	for _, sql := range sqls {
		if _, err := o.Raw(sql).Exec(); err != nil {
			return err
		}
	}
	return nil
}
