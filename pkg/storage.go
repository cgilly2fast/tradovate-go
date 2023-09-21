package tradovate

import (
	"errors"
	"time"
)

type UserID struct {
	UserId 	int
	Name 	string
}

type Storage struct {
	DeviceID 			string
	Accounts			[]Account
	AccessToken			string
	MdAccessToken		string
	ExpirationTime		string
	UserData			UserID	
}

var instance *Storage

func GetInstance() *Storage {
	if instance == nil {
		instance = &Storage{}
	}
	return instance
}

func (s *Storage) SetDeviceID(id string) {
	s.DeviceID = id
}

func (s *Storage) SetAccounts(accounts []Account) error{
	if len(accounts) == 0 {
		return errors.New("setAvailabeAccounts: Empty accounts passed")
	}
	s.Accounts = accounts
	return nil
}

func (s *Storage) GetAccounts() []Account {
	return s.Accounts
}

func (s *Storage) GetCurrentAccount() AccountMini {
	for _, account := range s.Accounts {
		if account.Active {
			return AccountMini {
				ID: 	*account.ID,
				Name: 	account.Name,
				UserID: account.UserID,
			}
		}
	}
	return AccountMini{}
}

func (s *Storage) GetDeviceID() string {
	return s.DeviceID
}

func (s *Storage) QueryAccounts(predicate func(Account) bool) Account {
	for _, account := range s.Accounts {
		if predicate(account) {
			return account
		}
	}
	return Account{}
}

func (s *Storage) SetAccessToken( token string, mdToken string, expiration string) error {
	if token == "" || expiration == "" {
		return errors.New("SetAccessToken: Attempted to set an undefined token")
	}
	s.AccessToken = token
	s.MdAccessToken = mdToken
	s.ExpirationTime = expiration
	return nil
}

func (s *Storage) GetAccessToken() AccessToken {
	return AccessToken{
		AccessToken:    s.AccessToken,
		ExpirationTime: s.ExpirationTime,
	}
}

func (s *Storage) GetMdAccessToken() MdAccessToken {
	return MdAccessToken{
		MdAccessToken:  s.MdAccessToken,
		ExpirationTime: s.ExpirationTime,
	}
}

func (s *Storage) TokenIsValid(expiration string) bool {
	expirationTime, _ := time.Parse(time.RFC3339, expiration) 
    currentTime := time.Now()
	
	return expirationTime.Sub(currentTime).Milliseconds() > 10*60*1000
}

func (s *Storage) TokenNearExpiry(expiration string) bool {
	expirationTime, _ := time.Parse(time.RFC3339, expiration) 
    currentTime := time.Now()

	return expirationTime.Sub(currentTime).Milliseconds() < 10*60*1000
}

func (s *Storage) SetUserData(data UserID) {
	s.UserData = data
}

func (s *Storage) GetUserData() UserID {
	return s.UserData
}

func (s *Storage) Clear() {
	s.DeviceID = ""
	s.Accounts = nil
	s.AccessToken = ""
	s.MdAccessToken = ""
	s.ExpirationTime = ""
	s.UserData = UserID{}
}