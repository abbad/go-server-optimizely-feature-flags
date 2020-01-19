# go-server-optimizely-feature-flags
Dabbling with go and optimizely, created a simple server to return enabled features for users.

Please read optimizely documentation. 
https://docs.developers.optimizely.com/full-stack/docs

The Api `/api/feature-flags` takes one optional arugment which is `user_id`and it returns 
the list of features that are enabled for that user. If you don't specify a user_id, it will autogenerate one and send it back in the header. 

Example response header.  
```
Feature-Flag: Feature-1;Feature-2
UUID: 123231-2122112-12212121 (UUID v4)
```

If you set the following cookie `__test_user` to `"true"`
It will pass it as an attribute to `getEnabledFeatures` attribute. 

Here is the main function that does the heavy lifting, this snippet take from this repo `pkg/featureflags/featureflags.go`

```
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
```

## TO Run. 

1. Install docker. 
2. Run `docker-compose run golang bash`
3. For some reason `go.mod` is not working. Installing dependencies is done manually. 

    1. `go get github.com/google/uuid`
    2. `go get github.com/optimizely/go-sdk`
    This is not ideal. 
4. cp .env.template .env and get optimizely SDK key.  
5. Exit docker and run docker-compose up. Your server should be running on localhost:10000


## TODO:
1. Fix go.mod file and make it work. 
2. Inside pkg/featureflags, fix multiple imports from the same package. 
3. Load test and see how many requests per sec this server can handle.	
