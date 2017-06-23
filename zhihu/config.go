package zhihu

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	UrlUserInfo         = "/api/v4/members/%s"
	UrlUserCollections  = "/api/v4/members/%s/collections" //
	UrlUserActivities   = "/api/v4/members/%s/activities"  //
	UrlUserQuestions    = "/api/v4/members/%s/questions"
	UrlUserAnwsers      = "/api/v4/answers/%d?include=content,voteup_count"
	UrlUserAnwsersTotal = "/api/v4/members/%s/answers?data[*].is_normal,is_collapsed,collapse_reason,suggest_edit,comment_count,can_comment,content,voteup_count,reshipment_settings,comment_permission,mark_infos,created_time,updated_time,review_info,relationship.is_authorized,voting,is_author,is_thanked,is_nothelp,upvoted_followees;data[*].author.badge[?(type=best_answerer)].topics"
	UrlUserArticles     = "/api/v4/members/%s/articles" //
	UrlUserFavlists     = "/api/v4/members/%s/favlists"
	UrlUserFollowers    = "/api/v4/members/%s/followers"
	UrlUserFollowees    = "/api/v4/members/%s/followees"

	UrlUserColumn_contributions        = "/api/v4/members/%s/column-contributions"
	UrlUserFollowingtopiccontributions = "/api/v4/members/%s/following-topic-contributions"
	UrlUserFollowingquestions          = "/api/v4/members/%s/following-questions"
	UrlUserFollowingfavlists           = "/api/v4/members/%s/following-favlists"
	UrlUserFollowingcolumn             = "/api/v4/members/%s/following-column"

	UrlQuestionInfo      = "/api/v4/questions/%d"
	UrlQuestionSimiler   = "/api/v4/questions/%d/similar-questions"
	UrlQuestionFollowers = "/api/v4/questions/%d/followers"
	UrlQuestionAnswers   = "/api/v4/questions/%d/answers"

	UrlZhunalan = "https://zhuanlan.zhihu.com/p/%d"
)

const (
	PARAMS_INFO            = `?include=locations,employments,gender,educations,business,voteup_count,thanked_Count,follower_count,following_count,cover_url,following_topic_count,following_question_count,following_favlists_count,following_columns_count,avatar_hue,answer_count,articles_count,pins_count,question_count,commercial_question_count,favorite_count,favorited_count,logs_count,marked_answers_count,marked_answers_text,message_thread_token,account_status,is_active,is_force_renamed,is_bind_sina,sina_weibo_url,sina_weibo_name,show_sina_weibo,is_blocking,is_blocked,is_following,is_followed,mutual_followees_count,vote_to_count,vote_from_count,thank_to_count,thank_from_count,thanked_count,description,hosted_live_count,participated_live_count,allow_message,industry_category,org_name,org_homepage,badge`
	PARAMS_QUESTION_INFO   = `?include=detail,follower_count,answer_count,author`
	PARAMS_ANSWER_INFO     = ""
	PARAMS_TOPIC_INFO      = ""
	PARAMS_COLLECTOIN_INFO = ""
)

const (
	ANONYMOUS_UID      = "0"
	ANONYMOUS_URLTOKEN = "匿名用户"
	ANONYMOUS_NAME     = "匿名用户"
)

var Config = new(ZhiHuConfig)

type ZhiHuConfig struct {
	DbUser     string `json:"dbuser"`
	DbAddress  string `json:"dbaddress"`
	DbPassword string `json:"dbpassword"`

	RedisMasterAddress string `json:"redismaster"`
	RedisSlaveAddress  string `json:"redisslave"`

	ZhihuAccount  string `json:"zhihuaccount"`
	ZhihuPassword string `json:"zhihupassword"`
}

func (z *ZhiHuConfig) LoadConfig(data []byte) error {
	if data == nil {
		return errors.New("data is nil")
	}
	json.Unmarshal(data, z)
	return nil
}

func (z *ZhiHuConfig) ReadConfigFile(path string) error {
	if path == "" {
		return errors.New("config path is empty")
	}
	fistat, err := os.Stat(path)
	if err != nil {
		return err
	}
	if fistat == nil || fistat.IsDir() {
		return fmt.Errorf("path is dir")
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = z.LoadConfig(data)
	return err
}
