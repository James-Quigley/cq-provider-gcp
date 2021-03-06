package compute

import (
	"context"
	"fmt"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cq-provider-gcp/resources/resource"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/api/compute/v1"
)

type Client struct {
	db               *database.Database
	log              hclog.Logger
	projectID        string
	region           string
	resourceMigrated map[string]bool
	svc              *compute.Service
}

func NewClient(db *database.Database, log hclog.Logger,
	projectID string) (resource.ClientInterface, error) {
	ctx := context.Background()
	computeService, err := compute.NewService(ctx)
	if err != nil {
		return nil, err
	}

	return &Client{
		db:               db,
		log:              log,
		projectID:        projectID,
		resourceMigrated: map[string]bool{},
		svc:              computeService,
	}, nil
}

func (c *Client) CollectResource(resource string, config interface{}) error {
	switch resource {
	case "instances":
		return c.instances(config)
	case "images":
		return c.images(config)
	case "addresses":
		return c.addresses(config)
	case "disk_types":
		return c.diskTypes(config)
	case "autoscalers":
		return c.autoscalers(config)
	case "interconnects":
		return c.interconnects(config)
	case "networks":
		return c.networks(config)
	case "subnetworks":
		return c.subnetworks(config)
	case "ssl_certificates":
		return c.sslCertificates(config)
	case "vpn_gateways":
		return c.vpnGateways(config)
	case "forwarding_rules":
		return c.forwardingRules(config)
	case "backend_services":
		return c.backendServices(config)
	case "firewalls":
		return c.firewalls(config)
	default:
		return fmt.Errorf("unsupported resource compute.%s", resource)
	}

}
