package subscription

import (
	"github.com/gin-gonic/gin"
	"github.com/hstreamdb/http-server/api/model"
	"github.com/hstreamdb/http-server/pkg/errorno"
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

func (s *Service) List(c *gin.Context) {
	streams, err := s.client.ListSubscriptions()
	if err != nil {
		c.JSON(errorno.LIST_SUBSCRIPTIONS_ERROR, gin.H{"error": err.Error()})
		return
	}
	c.JSON(errorno.SUCCESS, gin.H{"streams": streams})
}

func (s *Service) Create(c *gin.Context) {
	var sub model.Subscription
	if err := c.ShouldBindJSON(&sub); err != nil {
		c.JSON(errorno.INVALID_PARAMS, gin.H{"error": err.Error()})
		return
	}

	if err := s.client.CreateSubscription(sub); err != nil {
		c.JSON(errorno.CREATE_SUBSCRIPTION_ERROR, gin.H{"error": err.Error()})
		return
	}
	c.JSON(errorno.SUCCESS, gin.H{
		"msg": "success",
	})
}

func (s *Service) Delete(c *gin.Context) {}
