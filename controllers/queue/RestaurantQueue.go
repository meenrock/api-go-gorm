package queue

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"

	broker "restapi/database"
)

type Comment struct {
	Name    string `form:"name" json:"name"`
	Energy  int    `form:"energy"`
	Protein int    `form:"protein"`
}

func PushCommentToQueue(topic string, message []byte) error {

	host1 := os.Getenv("KAFKA_BROKER_HOST")
	port1 := os.Getenv("KAFKA_BROKER_PORT")

	dsn := host1 + ":" + port1

	brokersUrl := []string{dsn, dsn}
	producer, err := broker.ConnectProducer(brokersUrl)
	if err != nil {
		return err
	}
	defer producer.Close()
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
	return nil
}

func CreateComment(c *gin.Context) {
	// Instantiate new Message struct
	cmt := new(Comment)
	cmt.Name = c.PostForm("name")
	cmt.Energy, _ = strconv.Atoi(c.PostForm("energy"))
	cmt.Protein, _ = strconv.Atoi(c.PostForm("protein"))
	if err := c.ShouldBindJSON(cmt); err != nil {
		c.JSON(400, &gin.H{
			"success": false,
			"message": err,
		})
	}
	// convert body into bytes and send it to kafka
	cmtInBytes, err := json.Marshal(cmt)
	PushCommentToQueue("comments", cmtInBytes)
	c.JSON(200, &gin.H{
		"success": true,
		"message": "Comment pushed successfully",
		"comment": cmt,
	})
	// Return Comment in JSON format
	if err != nil {
		c.JSON(500, &gin.H{
			"success": false,
			"message": "Error creating product",
		})
	}
}
