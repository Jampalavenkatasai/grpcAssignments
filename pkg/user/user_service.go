package user

import (
	"context"
	"errors"
	"grpcAssign/apis"
)

type UserService struct {
	users map[int32]*apis.User
}

func NewUserService() *UserService {
	users := make(map[int32]*apis.User)
	users[1] = &apis.User{
		Id:      1,
		Fname:   "Steve",
		City:    "LA",
		Phone:   1234567890,
		Height:  5.8,
		Married: true,
	}

	return &UserService{users: users}
}

// GetUserById fetches user details based on user ID.
func (s *UserService) GetUserById(ctx context.Context, req *apis.UserRequest) (*apis.User, error) {
	if req.Id <= 0 {
		return nil, errors.New("Invalid user ID")
	}

	user, found := s.users[req.Id]
	if !found {
		return nil, errors.New("User not found")
	}

	return user, nil
}

// GetUsersByIds fetches a list of user details based on a list of user IDs.
func (s *UserService) GetUsersByIds(req *apis.UsersRequest, stream apis.UserService_GetUsersByIdsServer) error {
	if len(req.Ids) == 0 {
		return errors.New("Empty list of user IDs")
	}

	for _, id := range req.Ids {
		if id <= 0 {
			return errors.New("Invalid user ID in the list")
		}

		user, found := s.users[id]
		if found {
			if err := stream.Send(user); err != nil {
				return err
			}
		}
	}
	return nil
}
