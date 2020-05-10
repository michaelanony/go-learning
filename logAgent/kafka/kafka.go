package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)
var client sarama.SyncProducer //声明一个全局的连接kafka的生产者client
func Init(addrs []string)(err error)  {
	//1.生产者配置
	config:=sarama.NewConfig()
	config.Producer.RequiredAcks=sarama.WaitForAll//发送完数据需要leader和follow都确认  ACK
	config.Producer.Partitioner=sarama.NewRandomPartitioner//新选出一个partition		PARTITION
	config.Producer.Return.Successes=true//成功交付的消息将在success channel返回		CONFIRM

	//2、连接kafka
	client,err=sarama.NewSyncProducer(addrs,config)
	if err!=nil{
		fmt.Println(err)
		return
	}
	return 
}
func SendToKafka(topic,data string)  {
	//1、封装消息
	msg:=&sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value=sarama.StringEncoder(data)

	//2、发送消息
	pid,offset,err:=client.SendMessage(msg)
	if err!=nil{
		fmt.Println("send msg failed,err:",err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n",pid,offset)
}
