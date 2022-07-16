package utils

import (
	"context"

	firebase "firebase.google.com/go"
	"github.com/ShariqT/csfood4/pkg/core"
	"google.golang.org/api/option"
)

type FirebaseStore struct {
	app *firebase.App
	ctx context.Context
}

func (f FirebaseStore) QuerySingle(query string) (core.Modelable, error) {
	return nil, nil
}

func (f FirebaseStore) QueryMany(query string) ([]core.Modelable, error) {
	return nil, nil
}

func (f FirebaseStore) ListAll(resource core.ModelResource) ([]core.Modelable, error) {
	client, err := f.app.Firestore(f.ctx)
	if err != nil {
		return nil, err
	}
	entries := client.Collection(resource.String())
	docs, err := entries.DocumentRefs(f.ctx).GetAll()
	if err != nil {
		return nil, err
	}
	var res []core.Modelable
	for _, doc := range docs {
		d, _ := doc.Get(f.ctx)
		res = append(res, core.NewMarketFromData(d.Data()))
	}

	return res, nil
}

func (f FirebaseStore) AuthUser(username string, password string) (*core.User, error) {
	return nil, nil
}

func (f FirebaseStore) GetById(mr core.ModelResource, uuid string) (core.Modelable, error) {
	return nil, nil
}

func NewDB() (core.DB, error) {
	opt := option.WithCredentialsFile("./commonstock-dev-firebase-adminsdk-dads1-914582515e.json")
	ctx := context.Background()
	config := &firebase.Config{}
	config.ProjectID = "commonstock-dev"
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		return nil, err
	}
	return &FirebaseStore{app, ctx}, nil
}
