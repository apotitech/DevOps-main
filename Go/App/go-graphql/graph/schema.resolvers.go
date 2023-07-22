package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31
import (
	"context"
	"go-grapgql/db"
	"go-grapgql/graph/model"
	"math/rand"
	"strconv"
)

var database = db.NewMongoDB()

// AddVideo is the resolver for the addVideo field.
func (r *mutationResolver) AddVideo(ctx context.Context, input model.NewVideo) (*model.Video, error) {
	video := &model.Video{
		ID:     strconv.Itoa(rand.Int()),
		Title:  input.Title,
		URL:    input.URL,
		Author: &model.Author{AuthorID: input.AuthorID, Name: "user" + input.AuthorID},
	}
	database.Save(video)

	return video, nil
}

// Videos is the resolver for the videos field.
func (r *queryResolver) Videos(ctx context.Context) ([]*model.Video, error) {
	return database.FindAll(), nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
