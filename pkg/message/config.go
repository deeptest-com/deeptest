package message

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/snowlyg/multi"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

type RedisConf struct {
	Address  string `json:"address"`
	Password string `json:"password"`
	PoolSize int    `json:"pool_size"`
}

type MysqlConf struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	MaxIdleConns int    `json:"max_idle_conns"`
	MaxOpenConns int    `json:"max_open_conns"`
	Url          string `json:"url"`
	Dbname       string `json:"dbname"`
	Config       string `json:"config"`
}

func (c *Client) CheckRedis() error {

	universalOptions := &redis.UniversalOptions{
		Addrs:       strings.Split(c.Redis.Address, ","),
		Password:    c.Redis.Password,
		PoolSize:    c.Redis.PoolSize,
		IdleTimeout: 300 * time.Second,
	}
	redisClient := redis.NewUniversalClient(universalOptions)

	err := multi.InitDriver(
		&multi.Config{
			DriverType:      "redis",
			UniversalClient: redisClient,
		})

	if err != nil {
		return err
	}
	if multi.AuthDriver == nil {
		return errors.New("初始化认证驱动失败")
	}

	return nil
}

func (c *Client) InitMysql() error {
	if c.Db != nil {
		return nil
	}
	return nil
	if c.Mysql.Url == "" || c.Mysql.Dbname == "" || c.Mysql.Username == "" || c.Mysql.Password == "" {
		return errors.New("mysql的url,dbname,username,password必须赋值！")
	}
	if c.Mysql.Username == "" {
		c.Mysql.Config = "charset=utf8mb4&parseTime=True&loc=Local"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", c.Mysql.Username, c.Mysql.Password, c.Mysql.Url, c.Mysql.Dbname, c.Mysql.Config)
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig)); err != nil {
		return err
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(c.Mysql.MaxIdleConns)
		sqlDB.SetMaxOpenConns(c.Mysql.MaxOpenConns)
		c.Db = db
	}
	return nil
}

func (c *Client) CheckMessageTable() error {
	sql := `
		create table if not exists biz_message
		(
			id             bigint auto_increment
				primary key,
			message_source varchar(255)  null comment '消息模块来源',
			content        varchar(500)  null comment '内容',
			receiver_range int default 0 not null comment '接收者范围 1:全部 2:个人 3：某角色 4:某项目',
			receiver_id    int           null comment 'receiver_range=1时为0;receiver_range=2时为用户ID;receiver_range=3时为角色ID;receiver_range=4时为项目ID',
			create_at      datetime      null,
			constraint biz_message_id_uindex
				unique (id),
			INDEX receiver_id_index (receiver_id),
			INDEX receiver_range_index (receiver_range)
		);
		create table if not exists biz_message_read
		(
			id         bigint auto_increment
				primary key,
			message_id int default 0 not null comment '消息ID
		',
			user_id    int           not null comment '用户ID',
			create_at  datetime      null,
			constraint biz_message_read_id_uindex
				unique (id),
			INDEX message_id_index (message_id)，
			INDEX user_id_index (user_id)
		)
			comment '存储用户已读的消息';
`
	return c.Db.Exec(sql).Error
}
