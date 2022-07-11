package models

import "github.com/beego/beego/client/orm"

//`id` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'id',
//`work_no` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '工号',
//`name` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '花名',
//`gender` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '性别:男 Male,女 Female,其它 Other',
//`depart_id` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '部门Id',
//`depart_name` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '部门名称',
//`merger_name` varchar(256) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '部门全称',
//`position_id` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '职位Id',
//`position_name` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '职位',
//`rank_id` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '职级id',
//`rank_name` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '职级',
//`project` varchar(256) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '参与的项目',
//`area` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '地区:菲律宾 PH,迪拜 UAE',
//`used_name` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '曾用名',
//`leader_id` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '直属领导Id',
//`leader` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '直属领导',
//`jixin` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '吉信账号',
//`slack` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'slack账号',
//`telegram` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'telegram账号',
//`whatsapp` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'whatsapp账号',
//`recruit_type` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '招聘类型:招聘 Recruit,内推 Push,猎头 Headhunt,公司 Company',
//`referee_id` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '内推人Id',
//`referee` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '内推人',
//`entry_date` datetime DEFAULT NULL COMMENT '入职日期',
//`quit_date` datetime DEFAULT NULL COMMENT '离职日期',
//`regular_date` datetime DEFAULT NULL COMMENT '转正日期',
//`return_date` datetime DEFAULT NULL COMMENT '回国日期',
//`reback_date` datetime DEFAULT NULL COMMENT '回岗日期',
//`out_date` datetime DEFAULT NULL COMMENT '外宿日期',
//`usdt_address` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'USDT地址',
//`status` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '状态:转正 Regular,离职 Quit,试用期 Trial,回国 Return,开除 Fire,邓退 Persuade,停薪留职 Leave',
//`work_style` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '工作方式:本地 Local,远程 Remote',
//`stay_status` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '宿住状态:外宿 Out,内宿 In',
//`create_time` datetime DEFAULT NULL COMMENT '创建时间',
//`update_time` datetime DEFAULT NULL COMMENT '修改时间',
//`operator_id` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '操作人id',
//`operator` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '操作人',
//`summary` varchar(128) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
//`regular_status` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '转正状态 转正员工:Regular 实习员工:Trial',
//`nationality` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '国籍',
//`account_id` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '权限系统账户ID',

type Staff struct {
	Id            string
	WorkNo        string
	Name          string
	Status        string
	RegularStatus string
	WorkStyle     string
	UsdtAddress   string
	EntryDate     string
	QuitDate      string
	CreateTime    string
	UpdateTime    string
	Summary       string
	Nationality   string
}

func init() {
	orm.RegisterModel(new(Staff))
}
