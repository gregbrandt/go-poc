package apihandler_test

import (
    "testing"    
	eh "github.com/looplab/eventhorizon"
    "github.com/gregbrandt/go-poc/apihandler"
    "net/http"
    "github.com/gregbrandt/go-poc/domain"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

type MockCommandBus struct {
	mock.Mock
}

func (b *MockCommandBus) HandleCommand(command eh.Command) error { 
    b.Called(command)
	return nil
}

func (b *MockCommandBus) SetHandler(handler eh.CommandHandler, commandType eh.CommandType) error {
	return nil
}

type MockRequest struct{}
func (r *MockRequest) FormValue(key string) string {
	return "test"
}

func TestCreateStory(t *testing.T) {
    mockCommandBus := new(MockCommandBus)    
    mockCommandBus.On("HandleCommand",mock.AnythingOfType("*domain.CreateStory")).Return(nil)

    apihandler.CommandBus = mockCommandBus
	createStoryHandle := http.HandlerFunc(apihandler.CreateStory)
	test := apihandler.GenerateHandleTester(t, createStoryHandle)

	params := "{'name':'test name', 'content':'test content'}"

	w := test("POST", params)
	if w.Code != http.StatusOK {
		t.Errorf(
			"POST /create story: return code %v. Expected %v.",
			w.Code,
			http.StatusOK,
		)
	}


   mockCommandBus.AssertCalled(t,"HandleCommand",mock.AnythingOfType("*domain.CreateStory") )
   assert.NotNil(t,mockCommandBus.Calls[0].Arguments[0])
   
    switch cmd := mockCommandBus.Calls[0].Arguments[0].(type) {

	case *domain.CreateStory:
        assert.Equal(t, "test name1", cmd.Name)
	
    default:
        t.Error("Invalid command type")
    }

}