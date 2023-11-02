package repo

import (
	"errors"
	"graphqlProject/graph/model"
	"log"

	"gorm.io/gorm"
)

type Storer interface {
	CreateUser(input *model.NewUser) (*model.User, error)
	CreateCompany(input *model.NewCompany) (*model.Company, error)
	CreateJob(input *model.NewJob) (*model.Job, error)

	GetAllUser() ([]*model.User, error)
	UserById(userId int) (*model.User, error)

	GetAllCompany() ([]*model.Company, error)
	CompanyById(companyId int) (*model.Company, error)

	GetAllJob() ([]*model.Job, error)
	JobById(jobId int) (*model.Job, error)
}

type Conn struct {
	db *gorm.DB
}

func NewConn(db *gorm.DB) (*Conn, error) {

	if db == nil {
		return nil, errors.New("provide a valid connection")
	}

	return &Conn{db: db}, nil
}

type Store struct {
	Storer
}

func NewStore(storer Storer) Store {

	return Store{Storer: storer}

}

func (c Conn) CreateUser(input *model.NewUser) (*model.User, error) {

	user := User{
		Name:  input.Name,
		Email: input.Email,
	}

	tx := c.db.Create(&user)

	if tx.Error != nil {
		log.Fatalln(tx.Error)
		return nil, tx.Error
	}

	return &model.User{Name: input.Name, Email: input.Email}, nil
}

func (c Conn) GetAllUser() ([]*model.User, error) {

	var users []*model.User

	tx := c.db.Find(&users)

	if tx.Error != nil {
		return []*model.User{}, tx.Error
	}

	return users, nil

}

func (c *Conn) UserById(userId int) (*model.User, error) {

	var user model.User

	tx := c.db.Where("ID=?", userId)

	err := tx.Find(&user).Error

	if err != nil {
		return &model.User{}, err
	}

	return &user, nil
}

func (c *Conn) CreateCompany(input *model.NewCompany) (*model.Company, error) {

	comp := Company{
		Name:     input.Name,
		Location: input.Location,
	}

	tx := c.db.Create(&comp)

	if tx.Error != nil {
		log.Fatalln(tx.Error)
		return &model.Company{}, tx.Error
	}

	return &model.Company{Name: input.Name, Location: input.Location}, nil

}

func (c *Conn) GetAllCompany() ([]*model.Company, error) {

	var companies []*model.Company

	tx := c.db.Find(&companies)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return companies, nil

}

func (c *Conn) CompanyById(companyId int) (*model.Company, error) {

	var company model.Company

	tx := c.db.Where("ID=?", companyId)

	err := tx.Find(&company).Error

	if err != nil {
		return &model.Company{}, err
	}

	return &company, nil
}

func (c *Conn) CreateJob(input *model.NewJob) (*model.Job, error) {

	job := Job{
		Title:       input.Title,
		Description: input.Description,
		CompanyId:   input.CompanyID,
	}

	tx := c.db.Create(&job)

	if tx.Error != nil {
		log.Fatalln(tx.Error)
		return nil, tx.Error
	}
	return &model.Job{Title: input.Title, Description: input.Description, CompanyID: input.CompanyID}, nil
}

func (c *Conn) GetAllJob() ([]*model.Job, error) {

	var jobs []*model.Job

	tx := c.db.Find(&jobs)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return jobs, nil
}

func (c *Conn) JobById(jobId int) (*model.Job, error) {

	var job model.Job

	tx := c.db.Where("JobId=?", jobId)

	err := tx.Find(&job).Error

	if err != nil {
		return &model.Job{}, err
	}

	return &job, nil

}
