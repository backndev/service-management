package config

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/riferrei/srclient"
)

func InitProducer() (*kafka.Producer, *srclient.Schema, error) {
	var err error
	var kafkaServers = CONFIG["KAFKA_SERVERS"]
	schemaRegistryServers := CONFIG["KAFKA_SCHEMA_REGISTRY_URL"]
	topic := CONFIG["KAFKA_PRODUCED_TOPIC_AUDIT"]
	topicForGettingSchema := topic + "-value"
	saslUser := CONFIG["KAFKA_SASL_USERNAME"]
	saslPass := CONFIG["KAFKA_SASL_PASSWORD"]

	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaServers,
		"sasl.mechanisms":   "PLAIN",
		"security.protocol": "SASL_SSL",
		"sasl.username":     saslUser,
		"sasl.password":     saslPass,
	})
	if err != nil {
		log.Printf("Failed to create producer: %s", err)
		return nil, nil, err
	}

	schemaRegistryClient := srclient.CreateSchemaRegistryClient(schemaRegistryServers)
	schemaRegistryClient.SetCredentials(saslUser, saslPass)

	schema, errSchema := schemaRegistryClient.GetLatestSchema(topicForGettingSchema)
	if errSchema != nil {
		log.Println("Failed to get schema: ", errSchema.Error())
		return nil, nil, errSchema
	}
	fmt.Printf("getting schema: %+v\n", schema)

	return producer, schema, nil
}

//func GetTopicAuditTrails() string {
//	return CONFIG["KAFKA_PRODUCED_TOPIC_AUDIT"]
//}

//func ReadAvroSchema(path string) (string, error) {
//	avroSchemaAuditTrailsBytes, err := ioutil.ReadFile(path)
//	if err != nil {
//		log.Println("Failed to read avro schema: ", err)
//	}
//	schema := string(avroSchemaAuditTrailsBytes)
//
//	log.Println("Avroschema:", schema)
//
//	return schema, err
//}
