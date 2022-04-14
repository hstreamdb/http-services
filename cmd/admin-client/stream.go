package main

import (
	"context"
	"encoding/json"
	"github.com/hstreamdb/http-server/api/model"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var stream = model.Stream{}

func newStreamCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stream",
		Short: "Stream related commands",
	}

	cmd.AddCommand(
		createStreamCmd(),
		listStreamsCmd(),
	)
	return cmd
}

func createStreamCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a stream",
		RunE:  createStream,
	}

	cmd.Flags().StringVarP(&stream.StreamName, "name", "n", "", "Stream name")
	cmd.MarkFlagRequired("name")
	cmd.Flags().Uint32VarP(&stream.ReplicationFactor, "replication-factor", "r", 1, "Replication factor")
	cmd.Flags().Uint32VarP(&stream.BacklogDuration, "backlog-duration", "b", 0, "Backlog duration")
	return cmd
}

func createStream(cmd *cobra.Command, args []string) error {
	cli, err := newClient()
	if err != nil {
		return errors.WithMessage(err, "failed to create http client")
	}
	body, err := json.Marshal(stream)
	if err != nil {
		return errors.WithMessage(err, "failed to marshal request body")
	}
	_, err = cli.Post().SetResource("streams").Body(body).Send(context.Background())
	if err != nil {
		return errors.WithMessage(err, "failed to send request")
	}
	return nil
}

func listStreamsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all streams",
		RunE:  listStreams,
	}
	return cmd
}

func listStreams(cmd *cobra.Command, args []string) error {
	cli, err := newClient()
	if err != nil {
		return errors.WithMessage(err, "failed to create http client")
	}
	resp, err := cli.Get().SetResource("streams").Send(context.Background())
	if err != nil {
		return errors.WithMessage(err, "failed to send request")
	}
	var streams []model.Stream
	if err = json.Unmarshal(resp, &streams); err != nil {
		return errors.WithMessage(err, "failed to unmarshal response")
	}
	printStreams(streams)
	return nil
}
