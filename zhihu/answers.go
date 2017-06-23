package zhihu

import (
	"encoding/json"
	"fmt"
	"fork/tools/log"
	"time"

	"github.com/garyburd/redigo/redis"
)

var answersChan = make(chan int64, 4096)

func GetAnswerinfo(aid int64, usingCache bool) {
	if aid <= 0 {
		return
	}
	redisConn := redispool.Get()
	if usingCache == true {
		reply, err := redis.String(redisConn.Do("GET", fmt.Sprintf("answer_%d", aid)))
		if err == nil {
			if reply != "" {
				return
			}
		} else {
			return
		}
	}
	url := fmt.Sprintf(MainUrl+UrlUserAnwsers, aid)
	data, err := Client.Get(url, "")
	if err != nil {
		log.Error(err)
		return
	}

	var ans = new(UserAnswer)
	err = json.Unmarshal(data, ans)
	if err != nil {
		log.Error(err)
		return
	}
	if ans != nil {
		ans.QID = ans.Question.QID
		ans.UID = ans.Author.UID
		ans.UrlToken = ans.Author.UrlToken
		ans.Insert()
		redisConn.Do("SET", fmt.Sprintf("answer_%d", aid), aid)
	}
	return
}

type AnswerMark struct {
	Mark int   `json:"mark,omitempty"`
	Qid  int64 `json:"aid,omitempty"`
	Func func(int64, bool)
}

type Answers struct {
	Vector []*AnswerMark `json:"vector,omitempty"`
}

func (a *Answers) Process() {
	for index, value := range a.Vector {
		if value != nil {
			if value.Mark == 1 {
				go func(idx int) {
					defer func() {
						if re := recover(); re != nil {
							log.TraceAll()
						}
					}()
					a.Vector[idx].Func(a.Vector[idx].Qid, true)
				}(index)
			}
		}
	}
}

func NewAnswer(aid int64) *Questions {
	return &Questions{
		Vector: []*QuestionMark{
			&QuestionMark{1, aid, GetAnswerinfo},
		},
	}
}

func AddAnswer(aid int64) {
	answersChan <- aid
}

func AnswersLoop() {
	for {
		select {
		case aid := <-answersChan:
			go func(idx int64) {
				NewAnswer(idx).Process()
			}(aid)
			time.Sleep(time.Second * 10)
		}
	}
}
