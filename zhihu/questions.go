package zhihu

import (
	"encoding/json"
	"fmt"
	"fork/tools/log"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"
)

func GetQuestionsSimilar(qid int64, usingCache bool) {
	if qid <= 0 {
		return
	}
	defer func() {
		if re := recover(); re != nil {
			log.TraceAll()
		}
	}()
	redisConn := redispool.Get()
	if usingCache == true {
		reply, err := redis.Int64(redisConn.Do("GET", fmt.Sprintf("question_similar_%d", qid)))
		if err == nil {
			if reply != 0 {
				return
			}
		} else {
			return
		}
	}

	url := fmt.Sprintf(MainUrl+UrlQuestionSimiler, qid)
	data, err := Client.Get(url, "")
	if err != nil {
		log.Error(err)
		return
	}
	type SimilerQuestion struct {
		Paging *Paging     `json:"paging,omitempty"`
		Data   []*Question `json:"data,omitempty"`
	}

	var ques = new(SimilerQuestion)
	err = json.Unmarshal(data, ques)
	if err != nil {
		log.Error(err)
		return
	}
	var result []*Question
	result = append(result, ques.Data...)
	var continued = true
	for continued {
		data, err = Client.Get(ques.Paging.Next, "")
		if err != nil {
			log.Error(err)
			break
		}
		err = json.Unmarshal(data, ques)
		if err != nil {
			log.Error(err)
			if strings.Contains(err.Error(), "unexpected") {
				log.Error(string(data))
			}
			continue
		}
		result = append(result, ques.Data...)
		ques.Data = *new([]*Question)
		if ques.Paging.Is_start == true {
			continued = true
		} else if (ques.Paging.Is_start == false) && (ques.Paging.Is_end == false) {
			continued = true
		} else if (ques.Paging.Is_start == false) && (ques.Paging.Is_end == true) {
			continued = false
		}
	}
	result = append(result, ques.Data...)
	if true {
		for _, value := range result {
			AddQuestion(value.QID)
		}
	} else {
		for _, value := range result {
			if value == nil {
				continue
			}
			AddQuestion(value.QID)
		}
	}
	redisConn.Do("SET", fmt.Sprintf("question_similar_%d", qid), qid)
	return
}

func GetQuestionsInfo(qid int64, usingCache bool) {
	if qid <= 0 {
		return
	}
	defer func() {
		if re := recover(); re != nil {
			log.TraceAll()
		}
	}()
	redisConn := redispool.Get()
	if usingCache == true {
		reply, err := redis.Int64(redisConn.Do("GET", fmt.Sprintf("question_info_%d", qid)))
		if err == nil {
			if reply != 0 {
				return
			}
		} else {
			return
		}
	}

	url := fmt.Sprintf(MainUrl+UrlQuestionInfo+PARAMS_QUESTION_INFO, qid)
	data, err := Client.Get(url, "")
	if err != nil {
		log.Error(err)
		return
	}
	var ques = new(UserQuestion)
	err = json.Unmarshal(data, ques)
	if err != nil {
		log.Error(err)
		return
	}

	if ques.Author.Name != ANONYMOUS_NAME {
		ques.UID = ques.Author.Id
		ques.UrlToken = ques.Author.UrlToken
		AddNewUser(ques.Author.UrlToken)
	} else {
		ques.UID = ANONYMOUS_UID
		ques.UrlToken = ANONYMOUS_URLTOKEN
	}

	ques.Insert()
	redisConn.Do("SET", fmt.Sprintf("question_info_%d", qid), qid)
	return
}

func GetQuestionsAnswers(qid int64, usingCache bool) {
	defer func() {
		if re := recover(); re != nil {
			log.TraceAll()
		}
	}()
	redisConn := redispool.Get()
	if usingCache == true {
		reply, err := redis.Int64(redisConn.Do("GET", fmt.Sprintf("question_answer_%d", qid)))
		if err == nil {
			if reply != 0 {
				return
			}
		} else {
			return
		}
	}
	url := fmt.Sprintf(MainUrl+UrlQuestionAnswers, qid)
	data, err := Client.Get(url, "")
	if err != nil {
		log.Error(err)
		return
	}
	type AnswerResp struct {
		Paging *Paging       `json:"paging,omitempty"`
		Data   []*UserAnswer `json:"data,omitempty"`
	}

	var ans = new(AnswerResp)
	err = json.Unmarshal(data, ans)
	if err != nil {
		log.Error(err)
		return
	}
	var result []*UserAnswer
	result = append(result, ans.Data...)
	var continued = true
	for continued {
		data, err = Client.Get(ans.Paging.Next, "")
		if err != nil {
			log.Error(err)
			break
		}
		err = json.Unmarshal(data, ans)
		if err != nil {
			log.Error(err)
			break
		}
		result = append(result, ans.Data...)
		ans.Data = *new([]*UserAnswer)
		if ans.Paging.Is_start == true {
			continued = true
		} else if (ans.Paging.Is_start == false) && (ans.Paging.Is_end == false) {
			continued = true
		} else if (ans.Paging.Is_start == false) && (ans.Paging.Is_end == true) {
			continued = false
		}
	}
	result = append(result, ans.Data...)
	for _, value := range result {
		if value == nil {
			continue
		}
		AddAnswer(value.AID)
	}
	redisConn.Do("SET", fmt.Sprintf("question_answer_%d", qid), qid)
	return
}

func GetQuestionFollowers(qid int64, usingCache bool) {
	if qid <= 0 {
		return
	}
	defer func() {
		if re := recover(); re != nil {
			log.TraceAll()
		}
	}()
	redisConn := redispool.Get()
	if usingCache == true {
		reply, err := redis.Int64(redisConn.Do("GET", fmt.Sprintf("question_follower_%d", qid)))
		if err == nil {
			if reply != 0 {
				return
			}
		} else {
			return
		}
	}
	url := fmt.Sprintf(MainUrl+UrlQuestionFollowers, qid)
	data, err := Client.Get(url, "")
	if err != nil {
		log.Error(err)
		return
	}
	type FollowersResp struct {
		Paging *Paging       `json:"paging,omitempty"`
		Data   []*UserCommon `json:"data,omitempty"`
	}
	var folques = new(FollowersResp)
	err = json.Unmarshal(data, folques)
	if err != nil {
		log.Error(err)
		return
	}
	var result []*UserCommon
	result = append(result, folques.Data...)
	var continued = true
	for continued {
		data, err = Client.Get(folques.Paging.Next, "")
		if err != nil {
			log.Error(err)
			break
		}
		err = json.Unmarshal(data, folques)
		if err != nil {
			log.Error(err)
			if strings.Contains(err.Error(), "unexpected") {
				log.Error(string(data))
			}
			continue
		}
		time.Sleep(time.Second * 5)
		result = append(result, folques.Data...)
		folques.Data = *new([]*UserCommon)
		if folques.Paging.Is_start == true {
			continued = true
		} else if (folques.Paging.Is_start == false) && (folques.Paging.Is_end == false) {
			continued = true
		} else if (folques.Paging.Is_start == false) && (folques.Paging.Is_end == true) {
			continued = false
		}
	}
	result = append(result, folques.Data...)
	if true {
		for _, value := range result {
			if value.Name != ANONYMOUS_NAME {
				AddNewUser(value.UrlToken)
				qf := new(QuestionFollowRelationShip)
				qf.FollowStatus = 1
				qf.QID = qid
				qf.UID = value.Id
				qf.UrlToken = value.UrlToken
				qf.Insert()
				redisConn.Do("SET", fmt.Sprintf("questoin_follower_%d", qid), qid)
			} else {
				qf := new(QuestionFollowRelationShip)
				qf.FollowStatus = 1
				qf.QID = qid
				qf.UID = ANONYMOUS_UID
				qf.UrlToken = ANONYMOUS_URLTOKEN
				qf.Insert()
				redisConn.Do("SET", fmt.Sprintf("questoin_follower_%d", qid), qid)
			}
		}
	} else {
		for _, value := range result {
			if value == nil {
				continue
			}
			if value.Name != ANONYMOUS_NAME {
				AddNewUser(value.UrlToken)
				qf := new(QuestionFollowRelationShip)
				qf.FollowStatus = 1
				qf.QID = qid
				qf.UID = value.Id
				qf.UrlToken = value.UrlToken
				qf.Insert()
			} else {
				qf := new(QuestionFollowRelationShip)
				qf.FollowStatus = 1
				qf.QID = qid
				qf.UID = ANONYMOUS_UID
				qf.UrlToken = ANONYMOUS_URLTOKEN
				qf.Insert()
			}
		}
	}
	redisConn.Do("SET", fmt.Sprintf("question_follower_%d", qid), qid)
	return
}

func FollowQuestion(qid int64, usingCache bool) {
	if qid <= 0 {
		return
	}
	redisConn := redispool.Get()
	if usingCache == true {
		reply, err := redis.Int64(redisConn.Do("GET", fmt.Sprintf("follow_question_%d", qid)))
		if err == nil {
			if reply != 0 {
				return
			}
		} else {
			return
		}
	}
	url := fmt.Sprintf(MainUrl+UrlQuestionFollowers, qid)
	_, err := client.Post(url, "", nil)
	if err != nil {
		log.Error(err)
		return
	}
	redisConn.Do("SET", fmt.Sprintf("follow_question_%d", qid), qid)
	return
}

type FollowQuestionCondition func() bool

var questoinChan = make(chan int64, 1024)

func AddQuestion(qid int64) {
	questoinChan <- qid
}

type QuestionMark struct {
	Mark int   `json:"mark,omitempty"`
	Qid  int64 `json:"qid,omitempty"`
	Func func(int64, bool)
}

type Questions struct {
	Vector []*QuestionMark `json:"vector,omitempty"`
}

func (q *Questions) Process() {

	for index, value := range q.Vector {
		if value != nil {
			if value.Mark == 1 {
				go func(idx int) {
					defer func() {
						if re := recover(); re != nil {
							log.TraceAll()
						}
					}()
					q.Vector[idx].Func(q.Vector[idx].Qid, true)
				}(index)
			}
		}
	}
}

func NewQuestion(qid int64) *Questions {
	return &Questions{
		Vector: []*QuestionMark{
			&QuestionMark{1, qid, GetQuestionsInfo},
			&QuestionMark{1, qid, GetQuestionsSimilar},
			&QuestionMark{1, qid, GetQuestionsAnswers},
			&QuestionMark{1, qid, GetQuestionFollowers},
			&QuestionMark{1, qid, FollowQuestion},
		},
	}
}

func QuestionsLoop() {
	for {
		select {
		case id := <-questoinChan:
			qid := id
			go func(idx int64) {
				NewQuestion(idx).Process()
			}(qid)
			time.Sleep(time.Second * 5)
		}
	}
}
