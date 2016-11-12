package apihandler_test

import (
    "testing"    
	eh "github.com/looplab/eventhorizon"
    "github.com/gregbrandt/go-poc/apihandler"
    "net/http"
    "github.com/gregbrandt/go-poc/domain"
)

type MockCommandBus struct {
	commands []eh.Command
}

func (b *MockCommandBus) HandleCommand(command eh.Command) error {    
    b.commands = append(b.commands,command)
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

    mockCommandBus := &MockCommandBus{commands : make([]eh.Command,0)}
    apihandler.CommandBus = mockCommandBus
	createStoryHandle := http.HandlerFunc(apihandler.CreateStory)
	test := apihandler.GenerateHandleTester(t, createStoryHandle)

	params := `{"name":"test name", "content":"test content"}`

	w := test("POST", params)
	if w.Code != http.StatusOK {
		t.Errorf(
			"POST /create story: return code %v. Expected %v.",
			w.Code,
			http.StatusOK,
		)
	}
    if(mockCommandBus.commands[0]== nil){
        t.Error("Command bus not called")
    }

    switch cmd := mockCommandBus.commands[0].(type) {

	case *domain.CreateStory:
        if(cmd.Name == "test name"){
            t.Error("Command bus not called")
        }
	}    
}