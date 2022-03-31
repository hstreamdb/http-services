package subscription

import (
	"github.com/gin-gonic/gin"
	"github.com/hstreamdb/http-server/api/model"
	"github.com/hstreamdb/http-server/pkg/errorno"
	"net/http"
)

type SubServices interface {
	CreateSubscription(sub model.Subscription) error
	ListSubscriptions() ([]model.Subscription, error)
}

type Service struct {
	client SubServices
}

func NewSubService(client SubServices) *Service {
	return &Service{client}
}

func (s *Service) Get(c *gin.Context) {

}

// List godoc
// @ID subscriptionList
// @Summary List all subscriptions in the cluster
// @Success 200 {object} []model.Subscription
// @Failure 400 {object} errorno.ErrorResponse
// @Router /v1/subscriptions/ [get]
func (s *Service) List(c *gin.Context) {
	subs, err := s.client.ListSubscriptions()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorno.NewErrorResponse(errorno.LIST_SUBSCRIPTIONS_ERROR, err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"subscriptions": subs})
}

// Create godoc
// @ID subscriptionCreate
// @Summary Create a subscription
// @Param request body model.Subscription true "Request body"
// @Success 200 {string} string "ok"
// @Failure 400 {object} errorno.ErrorResponse
// @Router /v1/subscriptions/ [post]
func (s *Service) Create(c *gin.Context) {
	var sub model.Subscription
	if err := c.ShouldBindJSON(&sub); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorno.NewErrorResponse(errorno.INVALID_PARAMETER, err))
		return
	}

	if err := s.client.CreateSubscription(sub); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorno.NewErrorResponse(errorno.CREATE_SUBSCRIPTION_ERROR, err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

func (s *Service) Delete(c *gin.Context) {}
