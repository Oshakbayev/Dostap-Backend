package service

import "hellowWorldDeploy/pkg/entity"

type FriendInterface interface {
	CreateFriendRequest(entity.FriendRequest) error
	EditFriendRequestStatus(int64, bool) error
	GetFriendRequestsByRecipientID(int64) ([]entity.FriendRequest, error)
	DeleteFriend(int64, int64) error
}

func (s *Service) CreateFriendRequest(freq entity.FriendRequest) error {
	err := s.repo.CreateFriendRequest(freq)
	if err != nil {
		s.log.Printf("error during CreateFriendRequest(service): %v", err)
		return err
	}
	return nil
}

func (s *Service) EditFriendRequestStatus(requestID int64, status bool) error {
	if status {
		if err := s.repo.AcceptFriendRequest(requestID); err != nil {
			s.log.Printf("error during AcceptFriendRequest EditFriendRequestStatus(service): %v", err)
			return err
		}
		req, err := s.repo.GetFriendRequestByID(requestID)
		if err != nil {
			s.log.Printf("error during GetFriendRequestByID EditFriendRequestStatus(service): %v", err)
			return err
		}
		if err := s.repo.CreateFriends(req.SenderID, req.RecipientID); err != nil {
			s.log.Printf("error during CreateFriends EditFriendRequestStatus(service): %v", err)
			return err
		}
	} else {
		if err := s.repo.DeleteFriendRequest(requestID); err != nil {
			s.log.Printf("error during DeleteFriendRequest EditFriendRequestStatus(service): %v", err)
			return err
		}
	}
	return nil
}

func (s *Service) GetFriendRequestsByRecipientID(recipientID int64) ([]entity.FriendRequest, error) {
	friendRequests, err := s.repo.GetFriendRequestsByRecipientID(recipientID)
	if err != nil {
		s.log.Printf("error during  GetFriendRequestsByRecipientID(service): %v", err)
		return nil, err
	}
	return friendRequests, nil
}

func (s *Service) DeleteFriend(friendID1, friendID2 int64) error {
	err := s.repo.DeleteFriend(friendID1, friendID2)
	if err != nil {
		s.log.Printf("error during  GetFriendRequestsByRecipientID(service): %v", err)
		return err
	}
	return nil
}
