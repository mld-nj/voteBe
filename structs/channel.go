package structs

type VoteChannel struct {
	MikeKind  string `gorm:"column:mikeKind" json:"mikeKind"`
	ChannelId int64  `gorm:"column:channelId" json:"channelId"`
}

func (VoteChannel) TableName() string {
	return "voteChannel"
}

type VoteOption struct {
	MikeFlavour string `gorm:"column:mikeFlavour" json:"mikeFlavour"`
	ChannelId   int64  `gorm:"column:channelId" json:"channelId"`
	OptionId    int64  `gorm:"column:optionId" json:"optionId"`
	Count       int64  `gorm:"column:count" json:"count"`
}

func (VoteOption) TableName() string {
	return "voteOption"
}