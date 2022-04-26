package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gogql/database"
	"gogql/graph/generated"
	"gogql/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProduct) (*model.Product, error) {
	return db.InsertProductById(input), nil
}

func (r *mutationResolver) CreateFriends(ctx context.Context, input model.NewFriends) (*model.Friends, error) {
	return db.InsertFriendsById(input), nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Product(ctx context.Context, id string) (*model.Product, error) {
	return db.FindProductById(id), nil
}

func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	return db.All(), nil
}

func (r *queryResolver) Friends(ctx context.Context) ([]*model.Friends, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllFriends(ctx context.Context) ([]*model.Friends, error) {
	return db.FriendsAll(), nil
}

func (r *queryResolver) Friend(ctx context.Context, id string) (*model.Friends, error) {
	return db.FindFriendsById(id), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var db = database.Connect("mongodb://localhost:27017/")
