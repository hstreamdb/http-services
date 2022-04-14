package httpClient

import (
	"context"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestClient(t *testing.T) {
	suite.Run(t, new(testClientSuite))
}

type testClientSuite struct {
	suite.Suite
	client *Client
}

func (s *testClientSuite) SetupTest() {
	var err error
	cfg := DefaultConfig()
	s.client, err = NewHTTPClient(cfg)
	s.NoError(err)
}

func (s *testClientSuite) TearDownTest() {
	s.client.Close()
}

func (s *testClientSuite) TestCreateStream() {
	req, err := s.client.Get().SetResource("streams").BuildRequest(context.Background())
	s.NoError(err)
	resp, code, err := s.client.Send(req)
	s.NoError(err)
	s.Equal(200, code)
	s.T().Log(resp)
}
