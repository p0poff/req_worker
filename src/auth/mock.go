package auth

type MockToken struct {
	token string
}

func (t *MockToken) IsValid() bool {
	return true
}