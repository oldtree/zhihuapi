package zhihu

import (
	"fork/tools/log"

	"github.com/astaxie/beego/orm"
)

var o orm.Ormer

type School struct {
	ID           int64  `json:"-" orm:"column(id);pk;auto"`
	Name         string `json:"name,omitempty" orm:"column(name)"`
	Introduction string `json:"introduction,omitempty" orm:"column(introduction);type(text)"`
	Excerpt      string `json:"excerpt,omitempty" orm:"column(excerpt);type(text)"`
	Url          string `json:"url,omitempty" orm:"column(url)"`
	Type         string `json:"type,omitempty" orm:"column(type)"`
	SID          string `json:"id,omitempty" orm:"column(sid)"`
}

func (s *School) Insert() {
	if s == nil {
		return
	}
	err := o.Read(s, "sid")
	if err == orm.ErrNoRows {
		_, err := o.Insert(s)
		if err != nil {
			log.Info(err)
			return
		}
		return
	}
	return
}

type Major struct {
	ID           int64  `json:"-" orm:"column(id);pk;auto"`
	Name         string `json:"name,omitempty" orm:"column(name)"`
	Introduction string `json:"introduction,omitempty" orm:"column(introduction);type(text)"`
	Excerpt      string `json:"excerpt,omitempty" orm:"column(excerpt);type(text)"`
	Url          string `json:"url,omitempty" orm:"column(url)"`
	Type         string `json:"type,omitempty" orm:"column(type)"`
	MID          string `json:"id,omitempty" orm:"column(mid)"`
}

func (m *Major) Insert() {
	if m == nil {
		return
	}
	err := o.Read(m, "mid")
	if err == orm.ErrNoRows {
		_, err := o.Insert(m)
		if err != nil {
			log.Info(err)
			return
		}
		return
	}
	return
}

type Education struct {
	School *School `json:"school,omitempty"`
	Major  *Major  `json:"major,omitempty"`
}

type EducationRelationShip struct {
	ID     int64  `json:"-" orm:"column(id);pk;auto"`
	SID    string `json:"sid,omitempty" orm:"column(sid)"`
	School string `json:"school,omitempty" orm:"column(school)"`
	MID    string `json:"mid,omitempty" orm:"column(mid)"`
	Major  string `json:"major,omitempty" orm:"column(major)"`
	UID    string `json:"uid,omitempty" orm:"column(uid)"`
	Name   string `json:"name,omitempty" orm:"column(name)"`
}

func (m *EducationRelationShip) Insert() {
	if m == nil {
		return
	}
	err := o.Read(m, "sid", "mid", "uid")
	if err == orm.ErrNoRows {
		_, err := o.Insert(m)
		if err != nil {
			log.Info(err)
			return
		}
		return
	}
	return
}

type Business struct {
	ID           int64  `json:"-" orm:"column(id);pk;auto"`
	URL          string `json:"url,omitempty" orm:"column(url)"`
	AvatarUrl    string `json:"avatar_url,omitempty" orm:"column(avatar_url)"`
	Introduction string `json:"introduction,omitempty" orm:"column(introduction);type(text)"`
	Type         string `json:"type,omitempty" orm:"column(type)"`
	Excerpt      string `json:"excerpt,omitempty" orm:"column(excerpt);type(text)"`
	BID          string `json:"id,omitempty" orm:"column(bid)"`
}

func (m *Business) Insert() {
	if m == nil {
		return
	}
	err := o.Read(m, "bid")
	if err == orm.ErrNoRows {
		_, err := o.Insert(m)
		if err != nil {
			log.Info(err)
			return
		}
		return
	}
	return
}

type Location struct {
	ID           int64  `json:"-" orm:"column(id);pk;auto"`
	Name         string `json:"name,omitempty" orm:"column(name)"`
	URL          string `json:"url,omitempty" orm:"column(url)"`
	AvatarUrl    string `json:"avatar_url,omitempty" orm:"column(avatar_url)"`
	Introduction string `json:"introduction,omitempty" orm:"column(introduction);type(text)"`
	Type         string `json:"type,omitempty" orm:"column(type)"`
	Excerpt      string `json:"excerpt,omitempty" orm:"column(excerpt);type(text)"`
	LID          string `json:"id,omitempty" orm:"column(lid)"`
}

func (m *Location) Insert() {
	if m == nil {
		return
	}
	err := o.Read(m, "lid")
	if err == orm.ErrNoRows {
		_, err := o.Insert(m)
		if err != nil {
			log.Info(err)
			return
		}
		return
	}
	return
}

type LocalRelationShip struct {
	ID      int64  `json:"-" orm:"column(id);pk;auto"`
	LID     string `json:"lid,omitempty" orm:"column(lid)"`
	LocName string `json:"loc_name,omitempty" orm:"column(locaname)"`
	UID     string `json:"uid,omitempty" orm:"column(uid)"`
	Name    string `json:"name,omitempty" orm:"column(name)"`
}

func (l *LocalRelationShip) Insert() {
	if l == nil {
		return
	}
	err := o.Read(l, "lid", "uid")
	if err == orm.ErrNoRows {
		_, err := o.Insert(l)
		if err != nil {
			log.Info(err)
			return
		}
		return
	}
	return
}

type Company struct {
	ID           int64  `json:"-" orm:"column(id);pk;auto"`
	Name         string `json:"name,omitempty" orm:"column(name)"`
	URL          string `json:"url,omitempty" orm:"column(url)"`
	AvatarUrl    string `json:"avatar_url,omitempty" orm:"column(avatar_url)"`
	Introduction string `json:"introduction,omitempty" orm:"column(introduction);type(text)"`
	Type         string `json:"type,omitempty" orm:"column(type)"`
	Excerpt      string `json:"excerpt,omitempty" orm:"column(excerpt);type(text)"`
	CID          string `json:"id,omitempty" orm:"column(cid)"`
}

func (m *Company) Insert() {
	if m == nil {
		return
	}
	err := o.Read(m, "cid")
	if err == orm.ErrNoRows {
		_, err := o.Insert(m)
		if err != nil {
			log.Info(err)
			return
		}
		return
	}
	return
}

type Job struct {
	ID           int64  `json:"-" orm:"column(id);pk;auto"`
	Name         string `json:"name,omitempty" orm:"column(name)"`
	URL          string `json:"url,omitempty" orm:"column(url)"`
	AvatarUrl    string `json:"avatar_url,omitempty" orm:"column(avatar_url)"`
	Introduction string `json:"introduction,omitempty" orm:"column(introduction);type(text)"`
	Type         string `json:"type,omitempty" orm:"column(type)"`
	Excerpt      string `json:"excerpt,omitempty" orm:"column(excerpt);type(text)"`
	JID          string `json:"id,omitempty" orm:"column(jid)"`
}

func (m *Job) Insert() {
	if m == nil {
		return
	}
	err := o.Read(m, "jid")
	if err == orm.ErrNoRows {
		_, err := o.Insert(m)
		if err != nil {
			log.Info(err)
			return
		}
		return
	}
	return
}

type CareerRelationShip struct {
	ID      int64  `json:"-" orm:"column(id);pk;auto"`
	CID     string `json:"cid,omitempty" orm:"column(cid)"`
	Company string `json:"company,omitempty" orm:"column(company)"`
	JID     string `json:"jid,omitempty" orm:"column(jid)"`
	Job     string `json:"job,omitempty" orm:"column(job)"`
	UID     string `json:"uid,omitempty" orm:"column(uid)"`
	Name    string `json:"name,omitempty" orm:"column(name)"`
}

func (m *CareerRelationShip) Insert() {
	if m == nil {
		return
	}
	err := o.Read(m, "cid", "job", "uid")
	if err == orm.ErrNoRows {
		_, err := o.Insert(m)
		if err != nil {
			log.Info(err)
			return
		}
		return
	}
	return
}

type Career struct {
	Company *Company `json:"company,omitempty"`
	Job     *Job     `json:"job,omitempty"`
}

type UserCommon struct {
	AvatarUrlTemplate string `json:"avatar_url_template,omitempty"`
	Type              string `json:"type,omitempty"`
	Name              string `json:"name,omitempty"`
	Headline          string `json:"headline,omitempty"`
	UrlToken          string `json:"url_token,omitempty"`
	UserType          string `json:"user_type,omitempty"`
	IsAdvertiser      bool   `json:"is_advertiser,omitempty"`
	AvatarUrl         string `json:"avatar_url,omitempty"`
	IsOrg             bool   `json:"is_org,omitempty"`
	Gender            int    `json:"gender,omitempty"`
	Url               string `json:"url,omitempty"`
	Id                string `json:"id,omitempty"`
}

type Paging struct {
	Is_end   bool   `json:"is_end,omitempty"`
	Totals   int    `json:"totals,omitempty"`
	Previous string `json:"previous,omitempty"`
	Is_start bool   `json:"is_start,omitempty"`
	Next     string `json:"next,omitempty"`
}

type UserInfo struct {
	ID                      int64        `json:"-" orm:"column(id);pk;auto"`
	IsFollowed              bool         `json:"is_followed,omitempty" orm:"column(is_followed)"`
	Educations              []*Education `json:"educations,omitempty" orm:"-"`
	FollowingCount          int          `json:"following_count,omitempty" orm:"column(following_count)"`
	VoteFromCount           int          `json:"vote_from_count,omitempty" orm:"column(vote_from_count)"`
	UserType                string       `json:"user_type,omitempty" orm:"column(user_type)"`
	ShowSinaWeibo           bool         `json:"show_sina_weibo,omitempty" orm:"column(show_sina_weibo)"`
	PinsCount               int          `json:"pins_count,omitempty" orm:"column(pins_count)"`
	IsFollowing             bool         `json:"is_following,omitempty" orm:"column(is_following)"`
	MarkedAnswersText       string       `json:"marked_answers_text,omitempty" orm:"column(marked_answers_text)"`
	AccountStatus           interface{}  `json:"-,omitempty" orm:"-"`
	IsForceRenamed          bool         `json:"is_force_renamed,omitempty" orm:"column(is_force_renamed)"`
	UID                     string       `json:"id,omitempty" orm:"column(uid)"`
	FavoriteCount           int          `json:"favorite_count,omitempty" orm:"column(favorite_count)"`
	VoteupCount             int          `json:"voteup_count,omitempty" orm:"column(voteup_count)"`
	CommercialQuestionCount int          `json:"commercial_question_count,omitempty" orm:"column(commercial_question_count)"`
	IsBlocking              bool         `json:"is_blocking,omitempty" orm:"column(is_blocking)"`
	FollowingColumnsCount   int          `json:"following_columns_count,omitempty" orm:"column(following_columns_count)"`
	Headline                string       `json:"headline,omitempty" orm:"column(headline);type(text)"`
	UrlToken                string       `json:"url_token,omitempty" orm:"column(url_token)"`
	ParticipatedLiveCount   int          `json:"participated_live_count,omitempty" orm:"column(participated_live_count)"`
	FollowingFavlistsCount  int          `json:"following_favlists_count,omitempty" orm:"column(following_favlists_count)"`
	IsAdvertiser            bool         `json:"is_advertiser,omitempty" orm:"column(is_advertiser)"`
	IsBindSina              bool         `json:"is_bind_sina,omitempty" orm:"column(is_bind_sina)"`
	FavoritedCount          int          `json:"favorited_count,omitempty" orm:"column(favorited_count)"`
	IsOrg                   bool         `json:"is_org,omitempty" orm:"column(is_org)"`
	FollowerCount           int          `json:"follower_count,omitempty" orm:"column(follower_count)"`
	Employments             []*Career    `json:"employments,omitempty" orm:"-"`
	MarkedAnswersCount      int          `json:"marked_answers_count,omitempty" orm:"column(marked_answers_count)"`
	AvatarHue               string       `json:"avatar_hue,omitempty" orm:"column(avatar_hue)"`
	AvatarUrlTemplate       string       `json:"avatar_url_template,omitempty" orm:"column(avatar_url_template)"`
	FollowingTopicCount     int          `json:"following_topic_count,omitempty" orm:"column(following_topic_count)"`
	Description             string       `json:"description,omitempty" orm:"column(description);type(text)"`
	Business                *Business    `json:"business,omitempty" orm:"-"`
	AvatarUrl               string       `json:"avatar_url,omitempty" orm:"column(avatar_url)"`
	HostedLiveCount         int          `json:"hosted_live_count,omitempty" orm:"column(hosted_live_count)"`
	IsActive                int64        `json:"is_active,omitempty" orm:"column(is_active)"`
	ThankToCount            int          `json:"thank_to_count,omitempty" orm:"column(thank_to_count)"`
	MutualFolloweesCount    int          `json:"mutual_followees_count,omitempty" orm:"column(mutual_followees_count)"`
	CoverUrl                string       `json:"cover_url,omitempty" orm:"column(cover_url)"`
	ThankFromCount          int          `json:"thank_from_count,omitempty" orm:"column(thank_from_count)"`
	VoteToCount             int          `json:"vote_to_count,omitempty" orm:"column(vote_to_count)"`
	IsBlocked               bool         `json:"is_blocked,omitempty" orm:"column(is_blocked)"`
	AnswerCount             int          `json:"answer_count,omitempty" orm:"column(answer_count)"`
	AllowMessage            bool         `json:"allow_message,omitempty" orm:"column(allow_message)"`
	ArticlesCount           int          `json:"articles_count,omitempty" orm:"column(articles_count)"`
	Badge                   []*UserBadge `json:"badge,omitempty" orm:"-"`
	Name                    string       `json:"name,omitempty" orm:"column(name)"`
	QuestionCount           int          `json:"question_count,omitempty" orm:"column(question_count)"`
	Type                    string       `json:"type,omitempty" orm:"column(type)"`
	Url                     string       `json:"url,omitempty" orm:"column(url)"`
	MessageThreadToken      string       `json:"message_thread_token,omitempty" orm:"column(message_thread_token)"`
	LogsCount               int          `json:"logs_count,omitempty" orm:"column(logs_count)"`
	FollowingQuestionCount  int          `json:"following_question_count,omitempty" orm:"column(following_question_count)"`
	ThankedCount            int          `json:"thanked_count,omitempty" orm:"column(thanked_count)"`
	Gender                  int          `json:"gender,omitempty" orm:"column(gender)"`
	Loc                     []*Location  `json:"locations,omitempty" orm:"-"`
}

func (u *UserInfo) Insert() {
	if u == nil {
		return
	}
	defer func() {
		if re := recover(); re != nil {
			log.Error("recover panic : ", re)
		}
	}()
	err := o.Read(u, "uid")
	if err == orm.ErrNoRows || err == orm.ErrMultiRows {
		_, err := o.Insert(u)
		if err != nil {
			log.Info(err)
			return
		}
		for _, value := range u.Badge {
			value.UID = u.UID
			value.Insert()
		}
		for _, value := range u.Loc {

			le := new(LocalRelationShip)
			if value != nil {
				value.Insert()
				le.LID = value.LID
				le.LocName = value.Name
			}
			le.UID = u.UID
			le.Name = u.UrlToken
			le.Insert()
		}
		if u.Business != nil {
			u.Business.Insert()
		}
		for _, value := range u.Educations {
			re := new(EducationRelationShip)
			if value.Major != nil {
				value.Major.Insert()
				re.MID = value.Major.MID
				re.Major = value.Major.Name
			}
			if value.School != nil {
				value.School.Insert()
				re.SID = value.School.SID
				re.School = value.School.Name
			}
			re.UID = u.UID
			re.Name = u.UrlToken
			re.Insert()
		}

		for _, value := range u.Employments {
			ca := new(CareerRelationShip)
			if value.Company != nil {
				ca.CID = value.Company.CID
				ca.Company = value.Company.Name
				value.Company.Insert()
			}
			if value.Job != nil {
				ca.JID = value.Job.JID
				ca.Job = value.Job.Name
				value.Job.Insert()
			}
			ca.Name = u.UrlToken
			ca.UID = u.UID
			ca.Insert()
		}
		return
	}
	return
}

type Question struct {
	ID           int64  `json:"-" orm:"column(id);pk;auto"`
	Created      int64  `json:"created,omitempty" orm:"column(created)"`
	URL          string `json:"url,omitempty" orm:"column(url)"`
	Title        string `json:"title,omitempty" orm:"column(title);type(text)"`
	UpdatedTime  int64  `json:"updated_time,omitempty" orm:"column(updated_time)"`
	QuestionType string `json:"question_type,omitempty" orm:"column(question_type)"`
	Type         string `json:"type,omitempty" orm:"column(type)"`
	QID          int64  `json:"id,omitempty" orm:"column(qid)"`
	Uname        string `json:"uname,omitempty" orm:"column(uname)"`
}

func (u *Question) Insert() {
	if u == nil {
		return
	}
	err := o.Read(u, "qid")
	if err == orm.ErrNoRows || err == orm.ErrMultiRows {
		_, err := o.Insert(u)
		if err != nil {
			log.Info(err)
			return
		}
		return
	}
	return
}

type UserQuestion struct {
	ID            int64      `json:"-" orm:"column(id);pk;auto"`
	Author        UserCommon `json:"author" orm:"-"`
	Created       int64      `json:"created,omitempty" orm:"column(created)"`
	Url           string     `json:"url,omitempty" orm:"column(url)"`
	Title         string     `json:"title,omitempty" orm:"column(title)"`
	Detail        string     `json:"detail,omitempty" orm:"column(detail);type(text)"`
	AnswerCount   int        `json:"answer_count" orm:"column(answer_count)"`
	FollowerCount int        `json:"follower_count" orm:"column(follower_count)"`
	Updated_time  int64      `json:"updated_time,omitempty" orm:"column(updated_time)"`
	Question_type string     `json:"question_type,omitempty" orm:"column(question_type)"`
	Type          string     `json:"type,omitempty" orm:"column(type)"`
	QID           int64      `json:"id,omitempty" orm:"column(qid)"`
	UrlToken      string     `json:"url_token,omitempty" orm:"column(url_token)"`
	UID           string     `json:"uid,omitempty" orm:"column(uid)"`
}

func (u *UserQuestion) Insert() {
	if u == nil {
		return
	}
	err := o.Read(u, "qid")
	if err == orm.ErrNoRows || err == orm.ErrMultiRows {
		_, err := o.Insert(u)
		if err != nil {
			log.Info(err)
			return
		}
		return
	}
	return
}

type QuestionFollowRelationShip struct {
	ID           int64  `json:"-" orm:"column(id);pk;auto"`
	QID          int64  `json:"qid,omitempty" orm:"column(qid)"`
	UID          string `json:"uid,omitempty" orm:"column(uid)"`
	UrlToken     string `json:"url_token,omitempty" orm:"column(url_token)"`
	FollowStatus int    `json:"follow_status,omitempty" orm:"column(follow_status)"` //0: not follow 1,followed
}

func (u *QuestionFollowRelationShip) Insert() {
	if u == nil {
		return
	}
	err := o.Read(u, "qid", "uid")
	if err == orm.ErrNoRows {
		_, err := o.Insert(u)
		if err != nil {
			log.Info(err)
			return
		}
		return
	}
	return
}

/*
is_collapsed: false,
author: {},
url: "http://www.zhihu.com/api/v4/answers/155589453",
excerpt: "女儿三岁那年的某个下午，趁着女儿午睡的空档，婆婆去菜场买菜。半个多小时后，婆婆突然在菜场看到一位陌生女人领着哭哭啼啼的孙女在寻找她，惊得心脏突突直跳。原来女儿中途醒来发现家长无人，哭一阵还是没人，就自己开门到小区滑滑梯那里去找奶奶，但还是…",
id: 155589453,
question: {},
updated_time: 1491446789,
content: "女儿三岁那年的某个下午，趁着女儿午睡的空档，婆婆去菜场买菜。半个多小时后，婆婆突然在菜场看到一位陌生女人领着哭哭啼啼的孙女在寻找她，惊得心脏突突直跳。原来女儿中途醒来发现家长无人，哭一阵还是没人，就自己开门到小区滑滑梯那里去找奶奶，但还是没人，于是她就站在那里伤心、恐惧、无助地哭。她的哭声引来了一位陌生女人，也不知她们是如何沟通的，反正这位善良的女人领着女儿走了好几个路口，在偌大的菜场找到了我婆婆，把小孩安全交到家人手中。婆婆慌得甚至忘记道声谢谢。八年过去了，每当想起这件细思极恐的事，我都心怀感激。此后，我也尽力帮助那些需要帮助的陌生人，因为我曾受到过上天和陌生人如此大的恩惠。",
extras: "",
created_time: 1491446789,
is_copyable: true,
type: "answer",
thumbnail: "",
voteup_count: 8302*/

type UserAnswer struct {
	ID           int64        `json:"-" orm:"column(id);pk;auto"`
	AID          int64        `json:"id,omitempty" orm:"column(aid)"`
	UID          string       `json:"uid,omitempty" orm:"column(uid)"`
	UrlToken     string       `json:"url_token,omitempty" orm:"column(url_token)"`
	QID          int64        `json:"qid,omitempty" orm:"column(qid)"`
	Is_collapsed bool         `json:"is_collapsed,omitempty" orm:"column(is_collapsed)"`
	Author       UserInfo     `json:"author,omitempty" orm:"-"`
	Question     UserQuestion `json:"question,omitempty" orm:"-"`
	Url          string       `json:"url,omitempty" orm:"column(url)"`
	Excerpt      string       `json:"excerpt,omitempty" orm:"column(excerpt);type(text)"`
	Updated_time int64        `json:"updated_time,omitempty" orm:"column(updated_time)"`
	Created_time int64        `json:"created_time,omitempty" orm:"column(created_time)"`
	Extras       string       `json:"extras,omitempty" orm:"column(extras)"`
	Type         string       `json:"type,omitempty" orm:"column(type)"`
	Thumbnail    string       `json:"thumbnail,omitempty" orm:"column(thumbnail)"`
	Is_copyable  bool         `json:"is_copyable,omitempty" orm:"column(is_copyable)"`
	VoteupCount  int          `json:"voteup_count,omitempty" orm:"column(voteup_count)"`
	Content      string       `json:"content,omitempty" orm:"column(content);type(text)"`
}

func (u *UserAnswer) Insert() {
	if u == nil {
		return
	}
	err := o.Read(u, "aid")
	if err == orm.ErrNoRows {
		_, err := o.Insert(u)
		if err != nil {
			log.Info(err)
			return
		}
		return
	}
	return
}

type UserFavlists struct {
	ID        int64  `json:"-" orm:"column(id);pk;auto"`
	FID       int64  `json:"id,omitempty" orm:"column(fid)"`
	UID       string `json:"id,omitempty" orm:"column(uid)"`
	Title     string `json:"title,omitempty" orm:"column(title)"`
	Url       string `json:"url,omitempty" orm:"column(url)"`
	Is_public bool   `json:"is_public,omitempty" orm:"column(is_public)"`
	Type      string `json:"type,omitempty" orm:"column(type)"`
}

func (u *UserFavlists) Insert() {
	if u == nil {
		return
	}
	err := o.Read(u, "uid", "fid")
	if err == orm.ErrNoRows || err == orm.ErrMultiRows {
		_, err := o.Insert(u)
		if err != nil {
			log.Info(err)
			return
		}
		return
	}
	return
}

type UserFollowers struct {
	ID                  int64        `json:"-" orm:"column(id);pk;auto"`
	FID                 string       `json:"id,omitempty" orm:"column(fid)"`
	UID                 string       `json:"id,omitempty" orm:"column(uid)"`
	Is_followed         bool         `json:"is_followed,omitempty" orm:"column(is_followed)"`
	Avatar_url_template string       `json:"avatar_url_template,omitempty" orm:"column(avatar_url_template)"`
	Name                string       `json:"name,omitempty" orm:"column(name)"`
	Is_advertiser       bool         `json:"is_advertiser,omitempty" orm:"column(is_advertiser)"`
	Headline            string       `json:"headline,omitempty" orm:"column(headline);type(text)"`
	Gender              int          `json:"gender,omitempty" orm:"column(gender)"`
	User_type           string       `json:"user_type,omitempty" orm:"column(user_type)"`
	Avatar_url          string       `json:"avatar_url,omitempty" orm:"column(avatar_url)"`
	Url_token           string       `json:"url_token,omitempty" orm:"column(url_token)"`
	Url                 string       `json:"url,omitempty" orm:"column(url)"`
	Badge               []*UserBadge `json:"badge,omitempty" orm:"-"`
	Answer_count        int          `json:"answer_count,omitempty" orm:"column(answer_count)"`
	Is_org              bool         `json:"is_org,omitempty" orm:"column(is_org)"`
	Follower_count      int          `json:"follower_count,omitempty" orm:"column(follower_count)"`
	Is_folllowing       bool         `json:"is_folllowing,omitempty" orm:"column(is_folllowing)"`
	Type                string       `json:"type,omitempty" orm:"column(type)"`
	Articles_count      int          `json:"articles_count,omitempty" orm:"column(articles_counts)"`
}

func (u *UserFollowers) Insert() {
	if u == nil {
		return
	}
	err := o.Read(u, "uid", "fid")
	if err == orm.ErrNoRows || err == orm.ErrMultiRows {
		_, err := o.Insert(u)
		if err != nil {
			log.Info(err)
			return
		}
		for _, value := range u.Badge {
			value.UID = u.UID
			value.Insert()
		}
		return
	}
	return
}

type UserFollowees struct {
	ID                  int64        `json:"-" orm:"column(id);pk;auto"`
	FID                 string       `json:"id,omitempty" orm:"column(fid)"`
	UID                 string       `json:"id,omitempty" orm:"column(uid)"`
	Is_followed         bool         `json:"is_followed,omitempty" orm:"column(is_followed)"`
	Avatar_url_template string       `json:"avatar_url_template,omitempty" orm:"column(avatar_url_template)"`
	Name                string       `json:"name,omitempty" orm:"column(name)"`
	Is_advertiser       bool         `json:"is_advertiser,omitempty" orm:"column(is_advertiser)"`
	Headline            string       `json:"headline,omitempty" orm:"column(headline);type(text)"`
	Gender              int          `json:"gender,omitempty" orm:"column(gender)"`
	User_type           string       `json:"user_type,omitempty" orm:"column(user_type)"`
	Avatar_url          string       `json:"avatar_url,omitempty" orm:"column(avatar_url)"`
	Url_token           string       `json:"url_token,omitempty" orm:"column(url_token)"`
	Url                 string       `json:"url,omitempty" orm:"column(url)"`
	Badge               []*UserBadge `json:"badge,omitempty" orm:"-"`
	Answer_count        int          `json:"answer_count,omitempty" orm:"column(answer_count)"`
	Is_org              bool         `json:"is_org,omitempty" orm:"column(is_org)"`
	Follower_count      int          `json:"follower_count,omitempty" orm:"column(follower_count)"`
	Is_folllowing       bool         `json:"is_folllowing,omitempty" orm:"column(is_folllowing)"`
	Type                string       `json:"type,omitempty" orm:"column(type)"`
	Articles_count      int          `json:"articles_count,omitempty" orm:"column(articles_counts)"`
}

func (u *UserFollowees) Insert() {
	if u == nil {
		return
	}
	err := o.Read(u, "uid", "fid")
	if err == orm.ErrNoRows || err == orm.ErrMultiRows {
		_, err := o.Insert(u)
		if err != nil {
			log.Info(err)
			return
		}
		for _, value := range u.Badge {
			value.UID = u.UID
			value.Insert()
		}
		return
	}
	return
}

type Topic struct {
	ID           int64  `json:"-" orm:"column(id);pk;auto"`
	Name         string `json:"name,omitempty" orm:"column(name)"`
	Introduction string `json:"introduction,omitempty" orm:"column(introduction)"`
	Excerpt      string `json:"excerpt,omitempty" orm:"column(excerpt)"`
	URL          string `json:"url,omitempty" orm:"column(url)"`
	AvaterURL    string `json:"avater_url,omitempty" orm:"column(avatar_url)"`
	Type         string `json:"type,omitempty" orm:"column(type)"`
	TID          string `json:"id,omitempty" orm:"column(tid)"`
}

func (u *Topic) Insert() {
	if u == nil {
		return
	}
	err := o.Read(u, "tid")
	if err == orm.ErrNoRows || err == orm.ErrMultiRows {
		_, err := o.Insert(u)
		if err != nil {
			log.Info(err)
			return
		}
		return
	}
	return
}

type UserFollowingTopicContributions struct {
	Topics             *Topic `json:"topic,omitempty"`
	ContributionsCount int    `json:"contributions_count,omitempty"`
}

type Favlists struct {
	ID       int64  `json:"-" orm:"column(id);pk;auto"`
	UID      string `json:"id,omitempty" orm:"column(uid)"`
	URL      string `json:"url,omitempty" orm:"column(url)"`
	IsPublic bool   `json:"is_public,omitempty" orm:"column(is_public)"`
	FID      int64  `json:"id,omitempty" orm:"column(fid)"`
	Type     string `json:"type,omitempty" orm:"column(type)"`
	Title    string `json:"title,omitempty" orm:"column(title)"`
}

func (u *Favlists) Insert() {
	if u == nil {
		return
	}
	err := o.Read(u, "fid")
	if err == orm.ErrNoRows || err == orm.ErrMultiRows {
		_, err := o.Insert(u)
		if err != nil {
			log.Info(err)
			return
		}
		return
	}
	return
}

type UserBadge struct {
	ID          int64  `json:"-" orm:"column(id);pk;auto"`
	UID         string `json:"-" orm:"column(uid)"`
	Type        string `json:"type,omitempty" orm:"column(type)"`
	Description string `json:"description,omitempty" orm:"column(description)"`
}

func (u *UserBadge) Insert() {
	if u == nil {
		return
	}
	_, err := o.Insert(u)
	if err != nil {
		log.Info(err)
		return
	}
	return

}

type UserArticles struct {
	ID                int64       `json:"-" orm:"column(id);pk;auto"`
	Updated           int64       `json:"updated,omitempty" orm:"column(updated)"`
	Created           int64       `json:"created,omitempty" orm:"column(created)"`
	Url               string      `json:"url,omitempty" orm:"column(url)"`
	CommentPermission string      `json:"comment_permission,omitempty" orm:"column(comment_permission)"`
	UID               string      `json:"uid,omitempty" orm:"column(uid)"`
	UrlToken          string      `json:"url_token,omitempty" orm:"column(url_token)"`
	Excerpt           string      `json:"excerpt,omitempty" orm:"column(excerpt);type(text)"`
	ImageUrl          string      `json:"image_url,omitempty" orm:"column(image_url)"`
	ExcerptTitle      string      `json:"excerpt_title,omitempty" orm:"column(excerpt_title);type(text)"`
	Title             string      `json:"title,omitempty" orm:"column(title);type(text)"`
	Type              string      `json:"type,omitempty" orm:"column(type)"`
	AID               int         `json:"aid,omitempty" orm:"column(aid)"`
	Author            *UserCommon `json:"author,omitempty" orm:"-"`
}

func (u *UserArticles) Insert() {
	if u == nil {
		return
	}
	err := o.Read(u, "aid")
	if err == orm.ErrNoRows || err == orm.ErrMultiRows {
		_, err := o.Insert(u)
		if err != nil {
			log.Info(err)
			return
		}
		return
	}
	return
}

type UserColumn struct {
	ID                 int64       `json:"-" orm:"column(id);pk;auto"`
	Updated            int64       `json:"updated,omitempty" orm:"column(updated)"`
	Title              string      `json:"title,omitempty" orm:"column(title)"`
	URL                string      `json:"url,omitempty" orm:"column(url)"`
	CommentPermission  string      `json:"comment_permission,omitempty" orm:"column(comment_permission)"`
	Author             *UserCommon `json:"author,omitempty" orm:"-"`
	UID                string      `json:"uid,omitempty" orm:"column(uid)"`
	UrlToken           string      `json:"url_token,omitempty" orm:"column(url_token)"`
	ImageUrl           string      `json:"image_url,omitempty" orm:"column(image_url)"`
	Type               string      `json:"type,omitempty" orm:"column(type)"`
	CID                string      `json:"cid,omitempty" orm:"column(cid)"`
	ContributionsCount int         `json:"contributions_count,omitempty" orm:"column(contributions_count)"`
}

func (u *UserColumn) Insert() {
	if u == nil {
		return
	}
	err := o.Read(u, "cid")
	if err == orm.ErrNoRows || err == orm.ErrMultiRows {
		_, err := o.Insert(u)
		if err != nil {
			log.Info(err)
			return
		}
		return
	}
	return
}
