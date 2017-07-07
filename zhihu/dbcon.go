package zhihu

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
)

var mysqlorm orm.Ormer
var redispool *redis.Pool

func InitMysql() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	dsn := fmt.Sprintf("%s:%s@%s?charset=utf8", Config.DbUser, Config.DbPassword, Config.DbAddress)
	err := orm.RegisterDataBase("default", "mysql", dsn, 10, 100)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	orm.RegisterModel(new(UserInfo))
	orm.RegisterModel(new(UserQuestion))
	orm.RegisterModel(new(UserAnswer))
	orm.RegisterModel(new(UserBadge))
	orm.RegisterModel(new(UserFavlists))
	orm.RegisterModel(new(UserFollowees))
	orm.RegisterModel(new(UserFollowers))

	orm.RegisterModel(new(Company))
	orm.RegisterModel(new(School))
	orm.RegisterModel(new(Location))
	orm.RegisterModel(new(Major))
	orm.RegisterModel(new(EducationRelationShip))
	orm.RegisterModel(new(Business))
	orm.RegisterModel(new(LocalRelationShip))
	orm.RegisterModel(new(CareerRelationShip))
	orm.RegisterModel(new(Job))
	orm.RegisterModel(new(QuestionFollowRelationShip))
	orm.RegisterModel(new(UserColumn))
	orm.RegisterModel(new(UserArticles))

	orm.Debug = false
	//orm.RunSyncdb("default", true, true)
	o = orm.NewOrm()
	fmt.Println("database init ok")
}

func InitRedis() {
	redispool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			reconn, err := redis.DialTimeout("tcp", "127.0.0.1:6379", 10*time.Second, 10*time.Second, 10*time.Second)
			if err != nil {
				fmt.Println("redis connect failed")
				os.Exit(-1)
				return nil, err
			}
			return reconn, err
		},
		MaxIdle:     10,
		MaxActive:   1000,
		IdleTimeout: 2 * time.Second,
		Wait:        false,
	}

	return
}

func RedisPong(c redis.Conn, t time.Time) error {
	reply, err := c.Do("ping")
	if err != nil {
		log.Println("[ERROR] ping redis fail", err)
		return err
	}
	if reply.(string) == "pong" || reply.(string) == "PONG" {
		log.Println("reply : %s ", reply)
		return nil
	}
	return err
}

func GetAnswerIdCache(aid int64) (bool, int64) {
	return true, 0
}

func GetUserIdCache(uid int64) (bool, int64) {
	return true, 0
}

func GetQuestionId(qid int64) (bool, int64) {
	return true, 0
}

type UserAnalsys struct {
	Uid int64
}

func (u *UserAnalsys) GetUserFollower() (interface{}, error) {
	return nil, nil
}

func (u *UserAnalsys) GetUserQuestion() (interface{}, error) {
	return nil, nil
}
func (u *UserAnalsys) GetUserAnwser() (interface{}, error) {
	return nil, nil
}
