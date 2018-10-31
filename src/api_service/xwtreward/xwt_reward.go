package xwtreward

import (
	"api_service/models"
	"os"
	"strconv"
	"time"
)

func init() {

}

const (
	XWT_DAY_DISTRI = 320000
	XWT_DAY_CREATE_ESSYA = XWT_DAY_DISTRI * 0.5 *0.85 * 0.2
	XWT_DAY_ESSAY_LIKE = XWT_DAY_DISTRI * 0.5 * 0.85 * 0.8
)


type EssayCountPower struct {
	EssayId    int64
	UserId     int64
	Power      int64
}

func CreatEssayReward() {
	var ee []EssayCountPower
	c := models.GetEssaySnapshot()
	var e EssayCountPower
	for _,v := range c {
		e.EssayId = v.EssayId
		e.UserId = v.EssayId
		e.Power = EssayCountToValue(v.Type, v.EssayCount)
		ee = append(ee, e)
	}
    var sum int64 = 0
	for _, v := range ee {
		sum += v.Power
	}
	for _,v := range ee {
		u,_ := models.GetUser(v.UserId)
		e,_ := models.GetEssay(v.EssayId)
		reward := XWT_DAY_CREATE_ESSYA * v.Power / sum
		u.XwtBalance += reward
		models.UpdateUser(v.UserId, u)
		content := "<" + u.NickName + ">" + "create_article_reward" + "<" +
			        strconv.FormatInt(v.EssayId, 10)+ ">" + "<" + e.Title + ">" +
		            "<" + strconv.FormatInt(reward, 10) + ">"
		var log models.Xwtlog
		log.Content = content
		log.Time = e.CreateTime
		models.Addlog(log)
	}
}

func EssayLikeReward() {
	ue := make(map[int64][]int64)
	essay_value := make(map[int64]int64)
	now := time.Now()
	user_essay := models.GetAllLikeEssayByData(now)
	for _,v := range user_essay {
		ue[v.EssayId] = append(ue[v.EssayId], v.UserId)
	}
	for k,v := range ue {
		essay_value[k] = 0
		for _, v1 := range v {
			user,_ := models.GetUser(v1)
			essay_value[k] += user.XwtPowerValue
		}
	}
	var total int64 = 0
	for _,v := range essay_value {
		total += v
	}
	for k,v := range essay_value {
		reward := XWT_DAY_ESSAY_LIKE * v / total
		essay,_ := models.GetEssay(k)
		user,_ := models.GetUser(essay.UserId)
		user.XwtBalance += reward
		models.UpdateUser(user.UserId, user)
		content := "<" + user.NickName + ">" + "article_liked_reward" + "<" +
			strconv.FormatInt(k, 10)+ ">" + "<" + essay.Title + ">" +
			"<" + strconv.FormatInt(reward, 10) + ">"
		var log models.Xwtlog
		log.Content = content
		log.Time = time.Now()
		models.Addlog(log)
	}
}

func WriteLogToFile() {
	logs := models.GetLogNotWrited()
	today := time.Now()
	fd,_ := os.OpenFile(today.Format("2018-05-09"), os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	for _,v := range logs {
		line := v.Content + v.Time.Format("2018-05-11") + "\n"
		buf := []byte(line)
		fd.Write(buf)
	}
	fd.Close()
}


func EssayCountToValue(t int64, wordcount int64) (int64) {
	var value int64
	if (t == 1) {
        value = 1
    } else if wordcount < 300 {
		value = 3
	} else if wordcount >= 300 && wordcount < 1000 {
		value = 4
	} else if wordcount >= 1000 && wordcount < 2000 {
		value = 8
	} else if wordcount > 2000 && wordcount < 3000 {
		value = 15
	} else {
		value = 20
	}
	return value
}


