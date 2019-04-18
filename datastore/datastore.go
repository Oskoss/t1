package datastore

import (
	"encoding/json"
	"fmt"

	rClient "github.com/go-redis/redis"
	"github.com/pantomath-io/demo-grpc/controllers"
)

type DataProvider interface {
	CreateStore(storeName string) error
	AddToStore(storeName string, value interface{}) error
	RemoveFromStore(storeName string, value interface{}) error
	ReadFromStore(storeName string, value interface{}) ([]byte, error)
}

func StoreDeployment(dp DataProvider, deployment controllers.Deployment) error {
	store := "deployment"
	if err := dp.CreateStore(store); err != nil {
		return err
	}
	if err := dp.AddToStore(store, deployment); err != nil {
		return err
	}
	return nil
}

func RemoveDeployment(dp DataProvider, deployment controllers.Deployment) error {
	store := "deployment"
	if err := dp.RemoveFromStore(store, deployment); err != nil {
		return err
	}
	return nil
}

func StatusDeployment(dp DataProvider, deployment controllers.Deployment) (*controllers.Deployment, error) {
	store := "deployment"
	var result controllers.Deployment
	bytes, err := dp.ReadFromStore(store, deployment)
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

type Redis struct {
	FQDN     string
	Port     string
	Password string
	Database int
}

func (r Redis) CreateStore(storeName string) error {

	// We will use <CONTROLLER_TYPE>:<CONTROLLER_NAME> as the key in Redis
	return nil
}

func (r Redis) AddToStore(storeName string, value interface{}) error {

	var key string

	encodedValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	switch v := value.(type) {
	case controllers.Deployment:
		key = storeName + ":" + v.Name
	default:
		return fmt.Errorf("Error %+v is of unknown type", v)
	}

	client := rClient.NewClient(&rClient.Options{
		Addr:     r.FQDN + ":" + r.Port,
		Password: r.Password,
		DB:       r.Database,
	})

	_, err = client.Ping().Result()
	if err != nil {
		return err
	}

	err = client.Set(key, encodedValue, 0).Err()
	if err != nil {
		return fmt.Errorf("Error %+v adding %s with %+v", err, key, encodedValue)
	}

	return nil
}

func (r Redis) RemoveFromStore(storeName string, value interface{}) error {

	var key string

	switch v := value.(type) {
	case controllers.Deployment:
		key = storeName + ":" + v.Name
	default:
		return fmt.Errorf("Error %+v is of unknown type", v)
	}

	client := rClient.NewClient(&rClient.Options{
		Addr:     r.FQDN + ":" + r.Port,
		Password: r.Password,
		DB:       r.Database,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return err
	}
	err = client.Del(key).Err()
	if err != nil {
		return fmt.Errorf("Error %+v removing %s from Redis", err, key)
	}

	return nil
}

func (r Redis) ReadFromStore(storeName string, value interface{}) ([]byte, error) {

	var key string

	switch v := value.(type) {
	case controllers.Deployment:
		key = storeName + ":" + v.Name
	default:
		return nil, fmt.Errorf("Error %+v is of unknown type", v)
	}
	client := rClient.NewClient(&rClient.Options{
		Addr:     r.FQDN + ":" + r.Port,
		Password: r.Password,
		DB:       r.Database,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	encodedData, err := client.Get(key).Result()
	if err != nil {
		return nil, fmt.Errorf("Error %+v getting key %s:%s from Redis", err, storeName, key)
	}

	return []byte(encodedData), nil
}
