package main

import (
	"context"
	"encoding/json"
	"github.com/hstreamdb/http-server/api/model"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func newStatusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Get status for the current cluster",
		RunE:  getStatus,
	}
	return cmd
}

func getStatus(cmd *cobra.Command, args []string) error {
	cli, err := newClient()
	if err != nil {
		return errors.WithMessage(err, "failed to create http client")
	}
	resp, err := cli.Get().SetResource("cluster/status").Send(context.Background())
	if err != nil {
		return errors.WithMessage(err, "failed to send request")
	}
	var res model.TableResult
	if err = json.Unmarshal(resp, &res); err != nil {
		return errors.WithMessage(err, "failed to unmarshal response")
	}
	printTableResult(&res)
	return nil
}
