package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"example/graph/generated"
	"example/graph/model"
	"fmt"
	"strconv"
)

func (r *mutationResolver) UpsertCharacter(ctx context.Context, input model.CharacterInput) (*model.Character, error) {
	id := input.ID
   var character model.Character
   character.Name = input.Name
 
   n := len(r.Resolver.CharacterStore)
   if n == 0 {
       r.Resolver.CharacterStore = make(map[string]model.Character)
   }
 
   if id != nil {
       _, ok := r.Resolver.CharacterStore[*id]
       if !ok {
           return nil, fmt.Errorf("not found")
       }
       r.Resolver.CharacterStore[*id] = character
   } else {
       // generate unique id
       nid := strconv.Itoa(n + 1)
       character.ID = nid
       r.Resolver.CharacterStore[nid] = character
   }
 
   return &character, nil
}

func (r *queryResolver) Character(ctx context.Context, id string) (*model.Character, error) {
	character, ok := r.Resolver.CharacterStore[id]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return &character, nil
}

func (r *queryResolver) Pogues(ctx context.Context) ([]*model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Kooks(ctx context.Context) ([]*model.Character, error) {
	panic(fmt.Errorf("not implemented"))
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
// //  - You have helper methods in this file. Move them out to keep these resolver files clean.
// func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
// 	panic(fmt.Errorf("not implemented"))
// }
// func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
// 	panic(fmt.Errorf("not implemented"))
// }
