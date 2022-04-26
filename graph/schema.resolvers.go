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

var db = database.Connect("mongodb://localhost:27017/")

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
func (r *queryResolver) Friend(ctx context.Context, id string) (*model.Friends, error) {
	return db.FindFriendsById(id), nil
}

func (r *queryResolver) Friends(ctx context.Context, id string) ([]*model.Friends, error) {
	return db.FriendsAll(), nil
}

func (r *queryResolver) Product(ctx context.Context, id string) (*model.Product, error) {
	return db.FindProductById(id), nil
}

func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	return db.All(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
