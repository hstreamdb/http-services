package main

import (
	"context"
	"encoding/json"
	"fmt"
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

	headers := make([]string, 0, len(intervals)+1)
	headers = append(headers, "stream_name")
	for _, interval := range intervals {
		headers = append(headers, fmt.Sprintf("%s_%s", method, interval))
	}
	printTableWithHeader(&res, headers)
	return nil
}
