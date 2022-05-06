package cmd

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/edgelesssys/constellation/cli/azure"
	"github.com/edgelesssys/constellation/cli/cloudprovider"
	"github.com/edgelesssys/constellation/cli/gcp"
	"github.com/spf13/cobra"
)

// warnAWS warns that AWS isn't supported.
func warnAWS(providerPos int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if cloudprovider.FromString(args[providerPos]) == cloudprovider.AWS {
			return errors.New("AWS isn't supported")
		}
		return nil
	}
}

func isCloudProvider(arg int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if provider := cloudprovider.FromString(args[arg]); provider == cloudprovider.Unknown {
			return fmt.Errorf("argument %s isn't a valid cloud provider", args[arg])
		}
		return nil
	}
}

func validInstanceTypeForProvider(insType string, provider cloudprovider.Provider) error {
	switch provider {
	case cloudprovider.GCP:
		for _, instanceType := range gcp.InstanceTypes {
			if insType == instanceType {
				return nil
			}
		}
		return fmt.Errorf("%s isn't a valid GCP instance type", insType)
	case cloudprovider.Azure:
		for _, instanceType := range azure.InstanceTypes {
			if insType == instanceType {
				return nil
			}
		}
		return fmt.Errorf("%s isn't a valid Azure instance type", insType)
	default:
		return fmt.Errorf("%s isn't a valid cloud platform", provider)
	}
}

func validateEndpoint(endpoint string, defaultPort int) (string, error) {
	_, _, err := net.SplitHostPort(endpoint)
	if err == nil {
		return endpoint, nil
	}

	if strings.Contains(err.Error(), "missing port in address") {
		return net.JoinHostPort(endpoint, strconv.Itoa(defaultPort)), nil
	}

	return "", err
}
