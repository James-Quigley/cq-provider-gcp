package kms

import (
	"context"
	"fmt"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cq-provider-gcp/resource"
	"go.uber.org/zap"
	kms "google.golang.org/api/cloudkms/v1"
	"strings"
)

type Client struct {
	db               *database.Database
	log              *zap.Logger
	projectID        string
	region           string
	resourceMigrated map[string]bool
	svc              *kms.Service
}

func NewClient(db *database.Database, log *zap.Logger, projectID string) (resource.ClientInterface, error) {
	ctx := context.Background()
	kmsService, err := kms.NewService(ctx)
	if err != nil {
		return nil, err
	}

	return &Client{
		db:               db,
		log:              log,
		projectID:        projectID,
		resourceMigrated: map[string]bool{},
		svc:              kmsService,
	}, nil
}

func (c *Client) CollectResource(resource string, config interface{}) error {
	switch strings.ToLower(resource) {
	case "keys":
		return c.CryptoKeys(config)
	default:
		return fmt.Errorf("unsupported resource .%s", resource)
	}

}
