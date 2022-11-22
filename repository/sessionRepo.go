package repository

import (
	"fmt"
	"time"

	"github.com/DendiAnugerah/Todo/model"
	"gorm.io/gorm"
)

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) SessionRepository {
	return SessionRepository{db}
}

func (s *SessionRepository) AddSession(session model.Session) error {
	return s.db.Create(&session).Error
}

func (s *SessionRepository) DeleteSession(token string) error {
	return s.db.Where("token = ", token).Delete(&model.Session{}).Error
}

func (s *SessionRepository) UpdateSession(session model.Session) error {
	return s.db.Model(&session).Where("username = ?", session.Username).Updates(model.Session{Token: session.Token, Username: session.Username, Expiry: session.Expiry}).Error
}

func (s *SessionRepository) SessionNameAvail(name string) (model.Session, error) {
	sessions := model.Session{}
	err := s.db.Where("username = ", name).First(&sessions).Error
	if err != nil {
		return model.Session{}, err
	}

	return sessions, nil

}

func (s *SessionRepository) IsTokenValid(token string) (model.Session, error) {
	session, err := s.SessionTokenAvail(token)
	if err != nil {
		return model.Session{}, err
	}

	if s.TokenExpired(session) {
		err = s.DeleteSession(token)
		if err != nil {
			return model.Session{}, nil
		}
		return model.Session{}, fmt.Errorf("token is expired")
	}

	return session, nil
}

func (s *SessionRepository) SessionTokenAvail(token string) (model.Session, error) {
	sessions := model.Session{}
	err := s.db.Where("token = ", token).First(&sessions).Error
	if err != nil {
		return model.Session{}, err
	}

	return sessions, nil
}

func (s *SessionRepository) TokenExpired(session model.Session) bool {
	return session.Expiry.Before(time.Now())
}
