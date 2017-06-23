package zhihu

import (
	"fork/tools/log"
	"time"
)

var topicChan = make(chan *Topic, 256)

func AddTopic(t *Topic) {
	topicChan <- t
}

func TopicLoop() {

	for {
		select {
		case a := <-topicChan:
			log.Info(a.ID)
			time.Sleep(time.Second * 30)
		}
	}
}
