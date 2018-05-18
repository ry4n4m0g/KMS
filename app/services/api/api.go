package api

type APIService struct{}

func New() *APIService {
	return &APIService{}
}

func (as *APIService) DoAuth(username, password string) (isValid bool, err error) {
	authorizedUsers := make(map[string]string)

	authorizedUsers["EPOINT"] = "Oc3UmUZsK0"

	if val, ok := authorizedUsers[username]; ok && val == password {
		isValid = true
	}

	return
}
