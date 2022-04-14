package main

import (
	"context"
	"encoding/json"
	"github.com/hstreamdb/http-server/api/model"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	method    string
	intervals []string
)

func newStatsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stats METHOD [INTERVAL...]",
		Short: "Get specific stats for the current cluster",
		RunE:  getStats,
	}

	cmd.Flags().StringVarP(&method, "method", "c", "", "The stats method")
	cmd.Flags().StringSliceVarP(&intervals, "interval", "i", []string{}, "Stats intervals. Use '-i 1s -i 10s' for multi-intervals.")
	cmd.MarkFlagRequired("method")
	return cmd
}

func getStats(cmd *cobra.Command, args []string) error {
	cli, err := newClient()
	if err != nil {
		return errors.WithMessage(err, "failed to create http client")
	}
	resp, err := cli.Get().SetResource("admin/stats").
		Param("method", method).
		Params("interval", intervals).
		Send(context.Background())
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
