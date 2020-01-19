package featureflags

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	optly "github.com/optimizely/go-sdk"
	"github.com/optimizely/go-sdk/pkg/client"
	"github.com/optimizely/go-sdk/pkg/entities"
)

// OptiService struct to hold the client
type OptiService struct {
	Client *client.OptimizelyClient
}

// GetClient sends you back optmilzely client.
func GetClient(sdkKey string) *client.OptimizelyClient {
	var optClient, err = optly.Client(sdkKey)
	checkError(err)

	return optClient
}

// GetEnabledFeatures get which features are enabled for users
func (optiService *OptiService) GetEnabledFeatures(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: api/feature-flags")

	attributes := map[string]interface{}{
		"test_audience": isTestUser(r),
	}

	userID := getOrGenerateUID(r)
	w.Header().Add("UUID", userID)

	user := entities.UserContext{
		ID:         userID,
		Attributes: attributes,
	}
	enabledFeatures, err := optiService.Client.GetEnabledFeatures(user)

	checkError(err)

	w.Header().Add("Feature-Flags", string(strings.Join(enabledFeatures, ";")))
	fmt.Fprintf(w, "Feature-Flags: \nThose are the enabled features: %s", enabledFeatures)
}

func getOrGenerateUID(r *http.Request) string {
	qUserID, ok := r.URL.Query()["user_id"]

	if !ok || len(qUserID[0]) < 1 {
		log.Println("Url Param 'userID' is missing, generating UUID")
		return uuid.Must(uuid.NewRandom()).String()
	}

	log.Println("user_id sent %s, \n", qUserID[0])
	return qUserID[0]
}

func isTestUser(r *http.Request) bool {
	cookie, err := r.Cookie("__test_user")
	if err != nil || cookie == nil || cookie.Value != "true" {
		log.Printf("Cant find cookie __test_user :\n")
		log.Println(err, cookie)
		return false
	}

	log.Println("Cookie exists and its true")
	return true
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
