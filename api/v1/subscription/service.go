package subscription

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hstreamdb/http-server/api/model"
	"github.com/hstreamdb/http-server/pkg/errorno"
	"net/http"
	"sort"
)

type SubServices interface {
	CreateSubscription(sub model.Subscription) error
	ListSubscriptions() ([]model.Subscription, error)
	DeleteSubscription(subId string) error
}

type Service struct {
	client SubServices
}

func NewSubService(client SubServices) *Service {
	return &Service{client}
}

// Get godoc
// @ID subscriptionGet
// @Summary Get specific subscription by subscription id
// @Param subId path string true "Subscription Id"
// @Success 200 {object} model.Subscription
// @Failure 404 {object} errorno.ErrorResponse
// @Failure 500 {object} errorno.ErrorResponse
// @Router /v1/subscriptions/{subId} [get]
func (s *Service) Get(c *gin.Context) {
	target := c.Param("subId")
	subs, err := s.client.ListSubscriptions()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorno.NewErrorResponse(errorno.LIST_SUBSCRIPTIONS_ERROR, err))
		return
	}
	sort.Slice(subs, func(i, j int) bool {
		return subs[i].SubscriptionId < subs[j].SubscriptionId
	})
	idx := sort.Search(len(subs), func(i int) bool {
		return subs[i].SubscriptionId >= target
	})
	if idx < len(subs) && subs[idx].SubscriptionId == target {
		c.JSON(http.StatusOK, subs[idx])
	} else {
		c.AbortWithStatusJSON(http.StatusNotFound,
			errorno.NewErrorResponse(errorno.SUBSCRIPTION_NOT_EXIST, fmt.Errorf("subscription %s not exist", target)))
	}
}

// List godoc
// @ID subscriptionList
// @Summary List all subscriptions in the cluster
// @Success 200 {object} []model.Subscription
// @Failure 500 {object} errorno.ErrorResponse
// @Router /v1/subscriptions/ [get]
func (s *Service) List(c *gin.Context) {
	subs, err := s.client.ListSubscriptions()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorno.NewErrorResponse(errorno.LIST_SUBSCRIPTIONS_ERROR, err))
		return
	}
	c.JSON(http.StatusOK, subs)
}

// Create godoc
// @ID subscriptionCreate
// @Summary Create a subscription
// @Param request body model.Subscription true "Request body"
// @Success 200 {string} string "ok"
// @Failure 400 {object} errorno.ErrorResponse
// @Failure 500 {object} errorno.ErrorResponse
// @Router /v1/subscriptions/ [post]
func (s *Service) Create(c *gin.Context) {
	var sub model.Subscription
	if err := c.ShouldBindJSON(&sub); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorno.NewErrorResponse(errorno.INVALID_PARAMETER, err))
		return
	}

	if err := s.client.CreateSubscription(sub); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorno.NewErrorResponse(errorno.CREATE_SUBSCRIPTION_ERROR, err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

// Delete godoc
// @ID subscriptionDelete
// @Summary Delete specific subscription by subscription id
// @Param subId path string true "Subscription Id"
// @Success 200 {object} string "ok"
// @Failure 500 {object} errorno.ErrorResponse
// @Router /v1/subscriptions/{subId} [Delete]
func (s *Service) Delete(c *gin.Context) {
	subId := c.Param("subId")
	if err := s.client.DeleteSubscription(subId); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorno.NewErrorResponse(errorno.DELETE_SUBSCRIPTION_ERROR, err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
