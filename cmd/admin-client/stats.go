package main

import (
	"context"
	"encoding/json"
	"github.com/hstreamdb/http-server/api/model"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	metrics   string
	category  string
	intervals []string
)

func newStatsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stats Category Metrics [INTERVAL...]",
		Short: "Get specific stats for the current cluster",
		RunE:  getStats,
	}

	cmd.Flags().StringVarP(&category, "category", "c", "", "The stats category, e.g. stream, subscription.")
	cmd.MarkFlagRequired("category")
	cmd.Flags().StringVarP(&metrics, "method", "m", "", "The stats metrics, e.g. appends, sends.")
	cmd.MarkFlagRequired("method")
	cmd.Flags().StringSliceVarP(&intervals, "interval", "i", []string{}, "Stats intervals. Use '-i 1s -i 10s' for multi-intervals.")
	return cmd
}

func getStats(cmd *cobra.Command, args []string) error {
	cli, err := newClient()
	if err != nil {
		return errors.WithMessage(err, "failed to create http client")
	}
	resp, err := cli.Get().SetResource("cluster/stats").
		Param("category", category).
		Param("metrics", metrics).
		Params("interval", intervals).
		Send(context.Background())
	if err != nil {
		return errors.WithMessage(err, "failed to send request")
	}

	var res model.TableResult
	if err = json.Unmarshal(resp, &res); err != nil {
		return errors.WithMessage(err, "failed to unmarshal response")
	}
	printTableWithHeader(&res, res.Headers)
	return nil
}
