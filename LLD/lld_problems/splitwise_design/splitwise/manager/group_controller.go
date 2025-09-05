package manager

import (
	group "splitwise_design/splitwise/entities"
)

type GroupManager struct {
	groups []group.Group
}

func NewGroupManager() *GroupManager{
	return &GroupManager{make([]group.Group, 0)}
}

func (u *GroupManager) Add(group group.Group) {
	u.groups = append(u.groups, group)
}

func (u *GroupManager) Get(userId int) group.Group {

	for i := 0; i < len(u.groups); i++ {
		if u.groups[i].GetId() == userId {
			return u.groups[i]
		}

	}

	return group.Group{}
}
