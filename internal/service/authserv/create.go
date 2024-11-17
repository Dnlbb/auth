package authserv

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Dnlbb/auth/internal/models"
	"github.com/IBM/sarama"
)

func (s service) Create(ctx context.Context, user models.User) (*int64, error) {
	var id int64
	if err := validatePass(user.Password); err != nil {
		return nil, fmt.Errorf("invalid password: %w", err)
	}

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.storage.Save(ctx, user)
		if errTx != nil {
			return errTx
		}

		errTx = s.storage.Log(ctx, models.SAVE)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("user creation error: %w", err)
	}

	if err = s.cache.Create(ctx, id, user); err != nil {
		log.Printf("failed to cache user: %v", err)
	}

	data, err := json.Marshal(user)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal user: %w", err)
	}

	msg := &sarama.ProducerMessage{
		Topic: "notify",
		Value: sarama.StringEncoder(data),
	}

	_, _, err = s.kafkaProducer.SendMessage(msg)
	if err != nil {
		log.Printf("failed to send message: %v", err)
	}

	return &id, nil
}

func validatePass(password string) error {
	if len(password) < 8 || len(password) > 255 {
		return fmt.Errorf("password must be at least 8 characters but no more 255")
	}
	return nil
}
