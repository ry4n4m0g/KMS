package api

// APIService is a service that manages the API access
type APIService struct{}

// New creates an instance of the service
func New() *APIService {
	return &APIService{}
}

// DoAuth checks if a user's auth credentials are valid
func (as *APIService) DoAuth(username, password string) (isValid bool, err error) {
	authorizedUsers := make(map[string]string)

	authorizedUsers["EPOINT"] = "Oc3UmUZsK0"

	if val, ok := authorizedUsers[username]; ok && val == password {
		isValid = true
	}

	return
}
