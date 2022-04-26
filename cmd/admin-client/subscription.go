package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hstreamdb/http-server/api/model"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var sub = model.Subscription{}

func newSubscriptionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sub",
		Short: "Subscription related commands",
	}

	cmd.AddCommand(
		createSubscriptionCmd(),
		listSubscriptionCmd(),
	)
	return cmd
}

func createSubscriptionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a Subscription",
		RunE:  createSubscription,
	}

	cmd.Flags().StringVarP(&sub.StreamName, "name", "n", "", "Stream name to subscribe to")
	cmd.MarkFlagRequired("name")
	cmd.Flags().StringVarP(&sub.SubscriptionId, "id", "i", "", "Subscription id")
	cmd.MarkFlagRequired("id")
	cmd.Flags().Int32VarP(&sub.AckTimeoutSeconds, "timeout", "t", 60, "Ack timeout in seconds")
	cmd.Flags().Int32VarP(&sub.MaxUnackedRecords, "unacked", "m", 10000, "Max unacked records")
	return cmd
}

func createSubscription(cmd *cobra.Command, args []string) error {
	cli, err := newClient()
	if err != nil {
		return errors.WithMessage(err, "failed to create http client")
	}
	body, err := json.Marshal(sub)
	if err != nil {
		return errors.WithMessage(err, "failed to marshal request body")
	}
	_, err = cli.Post().SetResource("subscriptions").Body(body).Send(context.Background())
	if err != nil {
		return errors.WithMessage(err, "failed to send request")
	}
	fmt.Println("OK")
	return nil
}

func listSubscriptionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all subscriptions",
		RunE:  listSubscription,
	}
	return cmd
}

func listSubscription(cmd *cobra.Command, args []string) error {
	cli, err := newClient()
	if err != nil {
		return errors.WithMessage(err, "failed to create http client")
	}
	resp, err := cli.Get().SetResource("subscriptions").Send(context.Background())
	if err != nil {
		return errors.WithMessage(err, "failed to send request")
	}
	var subs []model.Subscription
	if err = json.Unmarshal(resp, &subs); err != nil {
		return errors.WithMessage(err, "failed to unmarshal response")
	}
	printSubs(subs)
	return nil
}
