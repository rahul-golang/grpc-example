package models

import "interview/grpc/proto"

type MemberDetails struct {
	UserKey  UserKey
	Contact  Contact
	Interest Interests
}

func NewMemberDetails(details *proto.MemberDetails) MemberDetails {
	return MemberDetails{
		UserKey: UserKey{
			Key: details.UserKey.Key,
		},
		Contact: Contact{
			Contact: details.CommonContacts.Contacts,
		},
		Interest: Interests{
			Interests: details.CommonInterests.Interests,
		},
	}
}
func (md MemberDetails) MapToMemberDetails() *proto.MemberDetails {
	return &proto.MemberDetails{
		UserKey: &proto.UserKey{
			Key: md.UserKey.Key,
		},
		CommonContacts: &proto.Contacts{
			Contacts: md.Contact.Contact,
		},
		CommonInterests: &proto.Interests{
			Interests: md.Interest.Interests,
		},
	}
}

type Contact struct {
	Contact []string
}

func NewContact(contacts *proto.Contacts) Contact {
	return Contact{
		Contact: contacts.Contacts,
	}
}

type UserViewingNetwork struct {
	User     User
	Interest Interests
}

type NetworkKey struct {
	key int64
}

type NetworkMembersView struct {
	MembersDetails []MemberDetails
}

func NewNetworkMembersView(req *proto.NetworkMembersView) NetworkMembersView {
	details := []MemberDetails{}

	for _, v := range req.Members {
		details = append(details, NewMemberDetails(v))
	}
	return NetworkMembersView{
		MembersDetails: details,
	}
}

func (nmv NetworkMembersView) MapToProtoPb() *proto.NetworkMembersView {
	var test = []*proto.MemberDetails{}
	for _, v := range nmv.MembersDetails {
		test = append(test, v.MapToMemberDetails())
	}
	return &proto.NetworkMembersView{
		Members: test,
	}
}
