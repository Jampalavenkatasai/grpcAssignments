// tests/user_service_test.go
package tests

import (
	"context"
	"grpcAssign/pkg/user"
	"testing"
)

func TestUserService_GetUserById(t *testing.T) {
	// Create a new instance of the UserService
	svc := user.NewUserService()

	t.Run("Valid user ID", func(t *testing.T) {
		// Define a request with a valid user ID
		req := &api.UserRequest{Id: 1}

		// Call the GetUserById method
		user, err := svc.GetUserById(context.Background(), req)

		// Check if the result is as expected
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if user == nil {
			t.Error("Expected user, got nil")
		}
		// Add more specific checks based on your implementation
	})

	t.Run("Invalid user ID", func(t *testing.T) {
		// Define a request with an invalid user ID (e.g., -1)
		req := &api.UserRequest{Id: -1}

		// Call the GetUserById method
		_, err := svc.GetUserById(context.Background(), req)

		// Check if the result is as expected
		if err == nil {
			t.Error("Expected an error, got nil")
		}
		// Add more specific checks based on your implementation
	})
}

func TestUserService_GetUsersByIds(t *testing.T) {
	// Create a new instance of the UserService
	svc := user.NewUserService()

	t.Run("Valid user IDs", func(t *testing.T) {
		// Define a request with valid user IDs
		req := &api.UsersRequest{Ids: []int32{1}}

		// Create a mock gRPC stream for testing
		stream := &user.MockUserService_GetUsersByIdsServer{}

		// Call the GetUsersByIds method
		err := svc.GetUsersByIds(req, stream)

		// Check if the result is as expected
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		// Add more specific checks based on your implementation
	})

	t.Run("Invalid user IDs", func(t *testing.T) {
		// Define a request with invalid user IDs, e.g., including -1
		req := &api.UsersRequest{Ids: []int32{1, -1}}

		// Create a mock gRPC stream for testing
		stream := &user.MockUserService_GetUsersByIdsServer{}

		// Call the GetUsersByIds method
		err := svc.GetUsersByIds(req, stream)

		// Check if the result is as expected
		if err == nil {
			t.Error("Expected an error, got nil")
		}
		// Add more specific checks based on your implementation
	})

	t.Run("Empty list of user IDs", func(t *testing.T) {
		// Define a request with an empty list of user IDs
		req := &api.UsersRequest{Ids: []int32{}}

		// Create a mock gRPC stream for testing
		stream := &user.MockUserService_GetUsersByIdsServer{}

		// Call the GetUsersByIds method
		err := svc.GetUsersByIds(req, stream)

		// Check if the result is as expected
		if err == nil {
			t.Error("Expected an error, got nil")
		}
		// Add more specific checks based on your implementation
	})
}
