package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"graphqlProject/graph/model"
)

// CreateUser is the resolver for the CreateUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {

	return r.S.CreateUser(&input)

}

// CreateCompany is the resolver for the CreateCompany field.
func (r *mutationResolver) CreateCompany(ctx context.Context, input model.NewCompany) (*model.Company, error) {

	return r.S.CreateCompany(&input)

}

// CreateJob is the resolver for the CreateJob field.
func (r *mutationResolver) CreateJob(ctx context.Context, input model.NewJob) (*model.Job, error) {

	return r.S.CreateJob(&input)

}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {

	return r.S.GetAllUser()

}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {

	return r.S.UserById(id)

}

// Company is the resolver for the company field.
func (r *queryResolver) Company(ctx context.Context, id int) (*model.Company, error) {

	return r.S.CompanyById(id)

}

// Companies is the resolver for the companies field.
func (r *queryResolver) Companies(ctx context.Context) ([]*model.Company, error) {

	return r.S.GetAllCompany()

}

// Job is the resolver for the job field.
func (r *queryResolver) Job(ctx context.Context, jobID int) (*model.Job, error) {

	return r.S.JobById(jobID)

}

// Jobs is the resolver for the jobs field.
func (r *queryResolver) Jobs(ctx context.Context) ([]*model.Job, error) {

	return r.S.GetAllJob()

}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
