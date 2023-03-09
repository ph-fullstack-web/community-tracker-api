package models

type Community struct {			
	CommunityID    int    			`gorm:"column:communityid" json:"community_id"`
	CommunityName  string 			`gorm:"column:communityname" json:"community_name"`
	CommunityDesc  string 			`gorm:"column:communitydesc" json:"community_description"`
	CommunityMgrID string 			`gorm:"column:communitymgrid" json:"community_manager"`
	Icon           string 			`gorm:"column:communityicon" json:"icon"`
	IsActive		 	 bool					`gorm:"column:isactive" json:"is_active"`
	Manager        AdminManager `gorm:"foreignKey:ID;references:CommunityMgrID" json:"manager"`
}

func (Community) TableName() string {
	return "community"
}

type CommunityMembers struct {
	CommunityID              string       `gorm:"column:communityid" json:"community_id"`
	CommunityName            string       `gorm:"column:communityname" json:"community_name"`
	CommunityManagerPeopleID string       `gorm:"column:communitymgrid" json:"-"`
	Members                  []People     `gorm:"foreignKey:Communityid;references:CommunityID" json:"members"`
	Manager                  AdminManager `gorm:"foreignKey:ID;references:CommunityManagerPeopleID" json:"manager"`
}

func (CommunityMembers) TableName() string {
	return "community"
}

type CreateCommunity struct {
	CommunityID      int    `gorm:"primaryKey;column:communityid" json:"community_id"`
	CommunityName    string `gorm:"column:communityname" json:"community_name"`
	CommunityManager int    `gorm:"column:communitymgrid" json:"community_manager"`
	CommunityDesc    string `gorm:"column:communitydesc" json:"community_description"`
	Icon             string `gorm:"column:communityicon" json:"icon"`
	IsActive         bool   `gorm:"column:isactive" json:"is_active"`
}

func (CreateCommunity) TableName() string {
	return "community"
}

type UpdateCommunity struct {
	CommunityID      int    `gorm:"primaryKey;column:communityid" json:"community_id"`
	CommunityName    string `gorm:"column:communityname" json:"community_name"`
	CommunityManager int    `gorm:"column:communitymgrid" json:"community_manager"`
	CommunityDesc    string `gorm:"column:communitydesc" json:"community_description"`
	Icon             string `gorm:"column:communityicon" json:"icon"`
	IsActive         bool   `gorm:"column:isactive" json:"is_active"`
}

func (UpdateCommunity) TableName() string {
	return "community"
}

type CommunityWithMembersPercentage struct {
	CommunityID     int    `gorm:"column:communityid" json:"community_id"`
	CommunityName   string `gorm:"column:communityname" json:"community_name"`
	CommunityDesc   string `gorm:"column:communitydesc" json:"community_description"`
	CommunityIcon   string `gorm:"column:communityicon" json:"icon"`
	Percentage      int    `gorm:"column:percentage" json:"percentage"`
	Members         int    `gorm:"column:members" json:"members"`
	ManagerFullName string `gorm:"column:communityadminandmanagername" json:"manager_full_name"`
}

func (CommunityWithMembersPercentage) TableName() string {
	return "community"
}

type CommunityMembersSearch struct {
	CommunityID int `gorm:"column:communityid" json:"community_id"`
	CommunityName string `gorm:"column:communityname" json:"community_name"`
	CommunityManagerPeopleID string `gorm:"column:communitymgrid" json:"-"`
	Manager AdminManager `gorm:"foreignKey:ID;references:CommunityManagerPeopleID" json:"manager"`
	Members []PeopleUnderCommunitySearch `gorm:"foreignKey:Communityid;references:CommunityID" json:"members"`
}

func (CommunityMembersSearch) TableName() string {
	return "community"
}