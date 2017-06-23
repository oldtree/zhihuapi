package zhihu

import (
	"encoding/json"
	"fmt"
	"fork/tools/log"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"
)

func GetUserInfo(name string, usingCache bool) {
	defer func() {
		if re := recover(); re != nil {
			log.Error("recover panic : ", re)
		}
	}()
	redisConn := redispool.Get()
	if usingCache == true {
		reply, err := redis.String(redisConn.Do("GET", fmt.Sprintf("user_info_%s", name)))
		if err == nil {
			if reply != "" {
				log.Info(reply)
				return
			}
		}
	}
	url := fmt.Sprintf(MainUrl+UrlUserInfo+PARAMS_INFO, name)
	data, err := Client.Get(url, "")
	if err != nil {
		log.Error(err)
		return
	}
	var u = new(UserInfo)
	err = json.Unmarshal(data, u)
	if err != nil {
		log.Error(err)
		if strings.Contains(err.Error(), "unexpected") {
			log.Error(string(data))
		}
		return
	}
	for _, value := range u.Badge {
		if value != nil {
			value.Insert()
		}
	}
	u.Insert()
	_, err = redisConn.Do("SET", fmt.Sprintf("user_info_%s", name), name)
	if err != nil {
		log.Error(err)
	}
	return
}
func Getcollections(name string, usingCache bool) {
	redisConn := redispool.Get()
	if usingCache == true {
		reply, err := redis.String(redisConn.Do("GET", fmt.Sprintf("user_collection_%s", name)))
		if err == nil {
			if reply != "" {
				log.Info(reply)
				return
			}
		}
	}
	url := fmt.Sprintf(MainUrl+UrlUserCollections, name)
	data, err := Client.Get(url, "")
	if err != nil {
		log.Error(err)
		return
	}
	var u = new(UserInfo)
	err = json.Unmarshal(data, u)
	if err != nil {
		log.Error(err)
		return
	}
	redisConn.Do("SET", fmt.Sprintf("user_collection_%s", name), name)
	return
}
func Getactivities(name string, usingCache bool) {
	redisConn := redispool.Get()
	if usingCache == true {
		reply, err := redis.String(redisConn.Do("GET", fmt.Sprintf("user_activities_%s", name)))
		if err == nil {
			if reply != "" {
				log.Info(reply)
				return
			}
		}
	}
	url := fmt.Sprintf(MainUrl+UrlUserActivities, name)
	data, err := Client.Get(url, "")
	if err != nil {
		log.Error(err)
		return
	}
	var u = new(UserInfo)
	err = json.Unmarshal(data, u)
	if err != nil {
		log.Error(err)
		if strings.Contains(err.Error(), "unexpected") {
			log.Error(string(data))
		}
		return
	}
	redisConn.Do("SET", fmt.Sprintf("user_activities_%s", name), name)
	return
}
func Getquestions(name string, usingCache bool) {
	defer func() {
		if re := recover(); re != nil {
			log.TraceAll()
		}
	}()
	redisConn := redispool.Get()
	if usingCache == true {
		reply, err := redis.String(redisConn.Do("GET", fmt.Sprintf("user_question_%s", name)))
		if err == nil {
			if reply != "" {
				log.Info(reply)
				return
			}
		}
	}
	url := fmt.Sprintf(MainUrl+UrlUserQuestions, name)
	data, err := Client.Get(url, "")
	if err != nil {
		log.Error(err)
		return
	}

	type QuestionResp struct {
		Paging *Paging     `json:"paging,omitempty"`
		Data   []*Question `json:"data,omitempty"`
	}
	var u = new(QuestionResp)
	err = json.Unmarshal(data, u)
	if err != nil {
		log.Error(err)
		if strings.Contains(err.Error(), "unexpected") {
			log.Error(string(data))
		}
		return
	}

	var result = *new([]*Question)
	result = append(result, u.Data...)
	var continued = true
	for continued {
		data, err = Client.Get(u.Paging.Next, "")
		if err != nil {
			log.Error(err)
			break
		}
		err = json.Unmarshal(data, u)
		if err != nil {
			log.Error(err)
			if strings.Contains(err.Error(), "unexpected") {
				log.Error(string(data))
			}
			continue
		}
		result = append(result, u.Data...)
		u.Data = *new([]*Question)
		if u.Paging.Is_start == true {
			continued = true
		} else if (u.Paging.Is_start == false) && (u.Paging.Is_end == false) {
			continued = true
		} else if (u.Paging.Is_start == false) && (u.Paging.Is_end == true) {
			continued = false
		}
	}
	if true {
		for _, value := range result {
			if value != nil {
				AddQuestion(value.QID)
			}
		}
	} else {
		for _, value := range result {
			if value != nil {
				AddQuestion(value.QID)
			}
		}
	}
	redisConn.Do("SET", fmt.Sprintf("user_question_%s", name), name)
	return
}

//"collapsed_counts,reviewing_comments_count,content,voteup_count,created,updated"
func Getanswers(name string, usingCache bool) {
	redisConn := redispool.Get()
	if usingCache == true {
		reply, err := redis.String(redisConn.Do("GET", fmt.Sprintf("user_answer_%s", name)))
		if err == nil {
			if reply != "" {
				log.Info(reply)
				return
			}
		}
	}
	url := fmt.Sprintf(MainUrl+UrlUserAnwsersTotal, name)

	data, err := Client.Get(url, "")
	if err != nil {
		log.Error(err)
		return
	}
	type AnswersResp struct {
		Paging *Paging       `json:"paging,omitempty"`
		Data   []*UserAnswer `json:"data,omitempty"`
	}
	var u = new(AnswersResp)
	err = json.Unmarshal(data, u)
	if err != nil {
		log.Error(err)
		return
	}

	var result []*UserAnswer
	result = append(result, u.Data...)
	var continued = true
	for continued {
		if u.Paging == nil || u.Paging.Next == "" {
			break
		}
		data, err = Client.Get(u.Paging.Next, "")
		if err != nil {
			log.Error(err)
			break
		}
		err = json.Unmarshal(data, u)
		if err != nil {
			log.Error(err)
			if strings.Contains(err.Error(), "unexpected") {
				log.Error(string(data))
			}
			continue
		}
		result = append(result, u.Data...)
		u.Data = *new([]*UserAnswer)
		if u.Paging.Is_start == true {
			continued = true
		} else if (u.Paging.Is_start == false) && (u.Paging.Is_end == false) {
			continued = true
		} else if (u.Paging.Is_start == false) && (u.Paging.Is_end == true) {
			continued = false
		}
	}
	for _, value := range result {
		value.QID = value.Question.QID
		value.UID = value.Author.Name
		value.Insert()
	}
	redisConn.Do("SET", fmt.Sprintf("user_answer_%s", name), name)
	return
}
func Getarticles(name string, usingCache bool) {
	redisConn := redispool.Get()
	if usingCache == true {
		reply, err := redis.String(redisConn.Do("GET", fmt.Sprintf("user_article_%s", name)))
		if err == nil {
			if reply != "" {
				log.Info(reply)
				return
			}
		}
	}
	url := fmt.Sprintf(MainUrl+UrlUserArticles, name)
	data, err := Client.Get(url, "")
	if err != nil {
		log.Error(err)
		return
	}
	type ArticlesResp struct {
		Paging *Paging         `json:"paging,omitempty"`
		Data   []*UserArticles `json:"data,omitempty"`
	}
	var u = new(ArticlesResp)
	err = json.Unmarshal(data, u)
	if err != nil {
		log.Error(err)
		if strings.Contains(err.Error(), "unexpected") {
			log.Error(string(data))
		}
		return
	}

	var result []*UserArticles
	result = append(result, u.Data...)
	var continued = true
	for continued {
		data, err = Client.Get(u.Paging.Next, "")
		if err != nil {
			log.Error(err)
			break
		}
		err = json.Unmarshal(data, u)
		if err != nil {
			log.Error(err)
			if strings.Contains(err.Error(), "unexpected") {
				log.Error(string(data))
			}
			continue
		}
		result = append(result, u.Data...)
		u.Data = *new([]*UserArticles)
		if u.Paging.Is_start == true {
			continued = true
		} else if (u.Paging.Is_start == false) && (u.Paging.Is_end == false) {
			continued = true
		} else if (u.Paging.Is_start == false) && (u.Paging.Is_end == true) {
			continued = false
		}
	}
	for _, value := range result {
		if value != nil {
			value.UID = value.Author.Id
			value.UrlToken = value.Author.UrlToken
			value.Insert()

		}
	}
	redisConn.Do("SET", fmt.Sprintf("user_article_%s", name), name)
	return
}
func Getfavlists(name string, usingCache bool) {
	redisConn := redispool.Get()
	if usingCache == true {
		reply, err := redis.String(redisConn.Do("GET", fmt.Sprintf("user_favlist_%s", name)))
		if err == nil {
			if reply != "" {
				log.Info(reply)
				return
			}
		}
	}
	url := fmt.Sprintf(MainUrl+UrlUserFavlists, name)

	data, err := Client.Get(url, "")
	if err != nil {
		log.Error(err)
		return
	}
	type FavlistsResp struct {
		Paging *Paging         `json:"paging,omitempty"`
		Data   []*UserFavlists `json:"data,omitempty"`
	}
	var u = new(FavlistsResp)
	err = json.Unmarshal(data, u)
	if err != nil {
		log.Error(err)
		if strings.Contains(err.Error(), "unexpected") {
			log.Error(string(data))
		}
		return
	}
	var result []*UserFavlists
	result = append(result, u.Data...)
	var continued = true
	for continued {
		data, err = Client.Get(u.Paging.Next, "")
		if err != nil {
			log.Error(err)
			break
		}
		err = json.Unmarshal(data, u)
		if err != nil {
			log.Error(err)
			if strings.Contains(err.Error(), "unexpected") {
				log.Error(string(data))
			}
			return
		}
		result = append(result, u.Data...)
		u.Data = *new([]*UserFavlists)
		if u.Paging.Is_start == true {
			continued = true
		} else if (u.Paging.Is_start == false) && (u.Paging.Is_end == false) {
			continued = true
		} else if (u.Paging.Is_start == false) && (u.Paging.Is_end == true) {
			continued = false
		}
	}
	for _, value := range result {
		value.UID = name
		value.Insert()

	}
	redisConn.Do("SET", fmt.Sprintf("user_favlist_%s", name), name)
	return
}
func Getfollowers(name string, usingCache bool) {
	redisConn := redispool.Get()
	if usingCache == true {
		reply, err := redis.String(redisConn.Do("GET", fmt.Sprintf("user_follower_%s", name)))
		if err == nil {
			if reply != "" {
				log.Info(reply)
				return
			}
		}
	}
	url := fmt.Sprintf(MainUrl+UrlUserFollowers, name)

	data, err := Client.Get(url, "")
	if err != nil {
		log.Error(err)
		return
	}
	type FollowersResp struct {
		Paging *Paging          `json:"paging,omitempty"`
		Data   []*UserFollowers `json:"data,omitempty"`
	}
	var u = new(FollowersResp)
	err = json.Unmarshal(data, u)
	if err != nil {
		log.Error(err)
		if strings.Contains(err.Error(), "unexpected") {
			log.Error(string(data))
		}
		return
	}
	var result []*UserFollowers
	result = append(result, u.Data...)
	var continued = true
	for continued {
		data, err = Client.Get(u.Paging.Next, "")
		if err != nil {
			log.Error(err)
			break
		}
		err = json.Unmarshal(data, u)
		if err != nil {
			log.Error(err)
			if strings.Contains(err.Error(), "unexpected") {
				log.Error(string(data))
			}
			return
		}
		result = append(result, u.Data...)
		u.Data = *new([]*UserFollowers)
		if u.Paging.Is_start == true {
			continued = true
		} else if (u.Paging.Is_start == false) && (u.Paging.Is_end == false) {
			continued = true
		} else if (u.Paging.Is_start == false) && (u.Paging.Is_end == true) {
			continued = false
		}
	}
	for _, value := range result {
		token := value.Url_token
		go AddNewUser(token)
		value.FID = value.Url_token
		value.UID = name
		value.Insert()
	}
	redisConn.Do("SET", fmt.Sprintf("user_follower_%s", name), name)
	return
}
func Getfollowees(name string, usingCache bool) {
	redisConn := redispool.Get()
	if usingCache == true {
		reply, err := redis.String(redisConn.Do("GET", fmt.Sprintf("user_followee_%s", name)))
		if err == nil {
			if reply != "" {
				log.Info(reply)
				return
			}
		}
	}
	url := fmt.Sprintf(MainUrl+UrlUserFollowees, name)

	data, err := Client.Get(url, "")
	if err != nil {
		log.Error(err)
		return
	}
	type FolloweesResp struct {
		Paging *Paging          `json:"paging,omitempty"`
		Data   []*UserFollowees `json:"data,omitempty"`
	}
	var u = new(FolloweesResp)
	err = json.Unmarshal(data, u)
	if err != nil {
		log.Error(err)
		return
	}
	var result []*UserFollowees
	result = append(result, u.Data...)
	var continued = true
	for continued {
		data, err = Client.Get(u.Paging.Next, "")
		if err != nil {
			log.Error(err)
			break
		}
		err = json.Unmarshal(data, u)
		if err != nil {
			log.Error(err)
			if strings.Contains(err.Error(), "unexpected") {
				log.Error(string(data))
			}
			return
		}
		result = append(result, u.Data...)
		u.Data = *new([]*UserFollowees)
		if u.Paging.Is_start == true {
			continued = true
		} else if (u.Paging.Is_start == false) && (u.Paging.Is_end == false) {
			continued = true
		} else if (u.Paging.Is_start == false) && (u.Paging.Is_end == true) {
			continued = false
		}
	}
	for _, value := range result {
		token := value.Url_token
		go AddNewUser(token)
		value.FID = name
		value.UID = value.Url_token
		value.Insert()
	}
	redisConn.Do("SET", fmt.Sprintf("user_followee_%s", name), name)
	return
}
func Getcolumncontributions(name string, usingCache bool) {
	redisConn := redispool.Get()
	if usingCache == true {
		reply, err := redis.String(redisConn.Do("GET", fmt.Sprintf("user_columncontribution_%s", name)))
		if err == nil {
			if reply != "" {
				log.Info(reply)
				return
			}
		}
	}
	url := fmt.Sprintf(MainUrl+UrlUserColumn_contributions, name)

	data, err := Client.Get(url, "")
	if err != nil {
		log.Error(err)
		return
	}

	type Cell struct {
		ContributionsCount int         `json:"contributions_count,omitempty"`
		Column             *UserColumn `json:"column,omitempty"`
	}

	type ColumnContributionsResp struct {
		Paging *Paging `json:"paging,omitempty"`
		Data   []*Cell `json:"data,omitempty"`
	}
	var u = new(ColumnContributionsResp)
	err = json.Unmarshal(data, u)
	if err != nil {
		log.Error(err)
		if strings.Contains(err.Error(), "unexpected") {
			log.Error(string(data))
		}
		return
	}
	var result []*Cell
	result = append(result, u.Data...)
	var continued = true
	for continued {
		data, err = Client.Get(u.Paging.Next, "")
		if err != nil {
			log.Error(err)
			break
		}
		err = json.Unmarshal(data, u)
		if err != nil {
			log.Error(err)
			if strings.Contains(err.Error(), "unexpected") {
				log.Error(string(data))
			}
			return
		}
		result = append(result, u.Data...)
		u.Data = *new([]*Cell)
		if u.Paging.Is_start == true {
			continued = true
		} else if (u.Paging.Is_start == false) && (u.Paging.Is_end == false) {
			continued = true
		} else if (u.Paging.Is_start == false) && (u.Paging.Is_end == true) {
			continued = false
		}
	}
	for _, value := range result {
		if value != nil {
			if value.Column != nil {
				value.Column.ContributionsCount = value.ContributionsCount
				value.Column.UID = value.Column.Author.Id
				value.Column.UrlToken = value.Column.Author.UrlToken
				value.Column.URL = fmt.Sprintf(UrlZhunalan, value.Column.CID)
				value.Column.Insert()
			}
		}
	}
	redisConn.Do("SET", fmt.Sprintf("user_columncontribution_%s", name), name)
	return
}
func Getfollowingtopiccontributions(name string, usingCache bool) {
	redisConn := redispool.Get()
	if usingCache == true {
		reply, err := redis.String(redisConn.Do("GET", fmt.Sprintf("user_followingtopiccontribution_%s", name)))
		if err == nil {
			if reply != "" {
				log.Info(reply)
				return
			}
		}
	}
	url := fmt.Sprintf(MainUrl+UrlUserFollowingtopiccontributions, name)
	data, err := Client.Get(url, "")
	if err != nil {
		log.Error(err)
		return
	}
	var u = new(UserInfo)
	err = json.Unmarshal(data, u)
	if err != nil {
		log.Error(err)
		if strings.Contains(err.Error(), "unexpected") {
			log.Error(string(data))
		}
		return
	}
	redisConn.Do("SET", fmt.Sprintf("user_followingtopiccontribution_%s", name), name)
	return
}
func Getfollowingquestions(name string, usingCache bool) {
	redisConn := redispool.Get()
	if usingCache == true {
		reply, err := redis.String(redisConn.Do("GET", fmt.Sprintf("user_followingquestion_%s", name)))
		if err == nil {
			if reply != "" {
				log.Info(reply)
				return
			}
		}
	}
	url := fmt.Sprintf(MainUrl+UrlUserFollowingquestions, name)
	data, err := Client.Get(url, "")
	if err != nil {
		log.Error(err)
		return
	}
	var u = new(UserInfo)
	err = json.Unmarshal(data, u)
	if err != nil {
		log.Error(err)
		if strings.Contains(err.Error(), "unexpected") {
			log.Error(string(data))
		}
		return
	}
	redisConn.Do("SET", fmt.Sprintf("user_followingquestion_%s", name), name)
	return
}
func Getfollowingfavlists(name string, usingCache bool) {
	redisConn := redispool.Get()
	if usingCache == true {
		reply, err := redis.String(redisConn.Do("GET", fmt.Sprintf("user_followingfavlist_%s", name)))
		if err == nil {
			if reply != "" {
				log.Info(reply)
				return
			}
		}
	}
	url := fmt.Sprintf(MainUrl+UrlUserFollowingfavlists, name)
	data, err := Client.Get(url, "")
	if err != nil {
		log.Error(err)
		return
	}
	var u = new(UserInfo)
	err = json.Unmarshal(data, u)
	if err != nil {
		log.Error(err)
		if strings.Contains(err.Error(), "unexpected") {
			log.Error(string(data))
		}
		return
	}
	redisConn.Do("SET", fmt.Sprintf("user_followingfavlist_%s", name), name)
	return
}

func FollowSomeUser(urltoken string, usingCache bool) {
	redisConn := redispool.Get()
	if usingCache == true {
		reply, err := redis.String(redisConn.Do("GET", fmt.Sprintf("user_following_%s", urltoken)))
		if err == nil {
			if reply != "" {
				log.Info(reply)
				return
			}
		}
	}
	url := fmt.Sprintf(MainUrl+UrlUserFollowers, urltoken)
	_, err := client.Post(url, "", nil)
	if err != nil {
		log.Error(err)
		return
	}
	redisConn.Do("SET", fmt.Sprintf("user_following_%s", urltoken), urltoken)
	return
}

type Mark struct {
	Mark int    `json:"mark,omitempty"`
	Name string `json:"name,omitempty"`
	Func func(string, bool)
}

type User struct {
	Vector []*Mark `json:"vector,omitempty"`
}

func NewUser(username string) *User {
	return &User{
		Vector: []*Mark{
			&Mark{1, username, GetUserInfo},
			&Mark{0, username, Getcollections},
			&Mark{0, username, Getactivities},
			&Mark{1, username, Getquestions},
			&Mark{0, username, Getarticles},
			&Mark{1, username, Getanswers},
			&Mark{0, username, Getfavlists},
			&Mark{1, username, Getfollowers},
			&Mark{1, username, Getfollowees},
			&Mark{0, username, Getcolumncontributions},
			&Mark{0, username, Getfollowingtopiccontributions},
			&Mark{0, username, Getfollowingquestions},
			&Mark{0, username, Getfollowingfavlists},
			&Mark{1, username, FollowSomeUser},
		},
	}
}

// TODO:s sfsdafds

func (s *User) Process() {
	defer func() {
		if re := recover(); re != nil {
			log.Error("recover panic : ", re)
		}
	}()
	for index, value := range s.Vector {
		if value != nil {
			if value.Mark == 1 {
				id := index
				go func(idx int) {
					defer func() {
						if re := recover(); re != nil {
							log.Error("recover panic : ", re, idx)
							log.TraceAll()
						}
					}()
					s.Vector[idx].Func(s.Vector[idx].Name, true)
				}(id)
			}
		}
	}
}

type ShellResp struct {
	Paging *Paging     `json:"paging,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

var userChan = make(chan string, 2048)

func AddNewUser(uname string) {
	log.Info(uname)
	userChan <- uname
}

func UserLoop() {
	for {
		select {
		case u := <-userChan:
			name := u
			go func() {
				defer func() {
					if re := recover(); re != nil {
						log.Error("recover panic : ", re)
					}
				}()
				NewUser(name).Process()
			}()
			time.Sleep(time.Second * 30)
		}
	}
}
