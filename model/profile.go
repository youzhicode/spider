package model

import "fmt"

type Profile struct {
	LoveId    string
	Age       string
	Sex       string
	Height    string
	Weight    string
	IsMarry   string
	Education string
	WorkAddr  string
	Income    string
}

func (p Profile) String() string {
	return fmt.Sprintf("[ID: %s, 年龄:%s, 性别: %s, 身高: %s, 体重: %s, 婚况: %s, 学历: %s, 工作地址: %s, 年收入: %s]",
		p.LoveId, p.Age, p.Sex, p.Height, p.Weight, p.IsMarry, p.Education, p.WorkAddr, p.Income)
}
