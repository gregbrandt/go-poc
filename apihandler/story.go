package apihandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	story "github.com/gregbrandt/Go-POC/domain"
	eh "github.com/looplab/eventhorizon"
    "github.com/gregbrandt/Go-POC/infrastructure"
)

type StoryModel struct {
	id eh.UUID
	name string
    content string
}

func CreateStory(w http.ResponseWriter, r *http.Request) {

	newStory := StoryModel{}

	newStory.name = r.FormValue("name")
	newStory.content = r.FormValue("content")
    newStory.id = eh.NewUUID()

	output, err := json.Marshal(newStory)

	fmt.Println(string(output))

	if err != nil {

		fmt.Println("Something went wrong!")

	}

	if err != nil {

		fmt.Println(err)

	}

	infrastructure.GetCommandBus().HandleCommand(&story.CreateStory{StoryId: newStory.id, Name: newStory.name, Content: newStory.content})
}

func GetStory(w http.ResponseWriter, r *http.Request) {

	// w.Header().Set("Pragma", "no-cache")

	// urlParams := mux.Vars(r)

	// id := urlParams["id"]

	// ReadUser := User{}

	// err := database.QueryRow("select * from users where user_id=?", id).Scan(&ReadUser.ID, &ReadUser.Name, &ReadUser.First, &ReadUser.Last, &ReadUser.Email)

	// switch {

	// case err == sql.ErrNoRows:

	// 	fmt.Fprintf(w, "No such user")

	// case err != nil:

	// 	log.Fatal(err)

	// default:

	// 	output, _ := json.Marshal(ReadUser)

	// 	fmt.Fprintf(w, string(output))

//	}

    }
