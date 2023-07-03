package modules

import (
	"fmt"
	"github.com/purawaktra/arjuna1-go/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Arjuna1Repo struct {
	db *gorm.DB
}

func CreateArjuna1Repo(gorm *gorm.DB) Arjuna1Repo {
	return Arjuna1Repo{
		db: gorm,
	}
}

func (br Arjuna1Repo) SelectCredentialByAccountId(query entities.Credentials, offset uint) ([]entities.Credentials, error) {
	var credentials = make([]entities.Credentials, 0)
	tx := br.db.Raw(
		fmt.Sprint("SELECT * FROM cities WHERE credentials = ", query.AccountId, " LIMIT 10 OFFSET ", offset)).Scan(
		&credentials)
	err := tx.Error
	if err != nil {
		return nil, err
	}
	return credentials, nil
}

func (br Arjuna1Repo) InsertSingleCredential(query entities.Credentials) (entities.Credentials, error) {
	var credentials entities.Credentials
	tx := br.db.Model(&credentials).Clauses(clause.Returning{}).Create(query)
	err := tx.Error
	if err != nil {
		return entities.Credentials{}, err
	}
	return credentials, nil
}

func (br Arjuna1Repo) InsertMultipleCredential(queries []entities.Credentials) ([]entities.Credentials, []error) {
	var credential entities.Credentials
	var credentials []entities.Credentials
	var errors []error
	for _, query := range queries {
		tx := br.db.Model(&credential).Clauses(clause.Returning{}).Create(query)
		credentials = append(credentials, credential)
		err := tx.Error
		if err != nil {
			errors = append(errors, err)
		} else {
			errors = append(errors, nil)
		}
	}
	return credentials, errors
}

func (br Arjuna1Repo) UpdateSingleCredentialByAccountId(query entities.Credentials) (entities.Credentials, error) {
	var credential entities.Credentials
	tx := br.db.Model(&credential).Clauses(clause.Returning{}).Updates(query)
	err := tx.Error
	if err != nil {
		return entities.Credentials{}, err
	}
	return credential, nil
}

func (br Arjuna1Repo) UpdateMultipleCredentialByAccountId(queries []entities.Credentials) ([]entities.Credentials, []error) {
	var credential entities.Credentials
	var credentials []entities.Credentials
	var errors []error
	for _, query := range queries {
		tx := br.db.Model(&credential).Clauses(clause.Returning{}).Updates(query)
		credentials = append(credentials, credential)
		err := tx.Error
		if err != nil {
			errors = append(errors, err)
		} else {
			errors = append(errors, nil)
		}
	}
	return credentials, errors
}

func (br Arjuna1Repo) DeleteSingleCredentialByAccountId(query entities.Credentials) (entities.Credentials, error) {
	tx := br.db.Clauses(clause.Returning{}).Delete(query)
	credential := query
	err := tx.Error
	if err != nil {
		return entities.Credentials{}, err
	}
	return credential, nil
}

func (br Arjuna1Repo) DeleteMultipleCredentialByAccountId(queries []entities.Credentials) ([]entities.Credentials, []error) {
	var errors []error
	var credentials []entities.Credentials
	for _, query := range queries {
		tx := br.db.Clauses(clause.Returning{}).Delete(&query)
		credentials = append(credentials, query)
		err := tx.Error
		if err != nil {
			errors = append(errors, err)
		} else {
			errors = append(errors, nil)
		}
	}
	return credentials, errors
}
