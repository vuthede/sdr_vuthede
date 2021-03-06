package model

func getGroupById(id int) Group {
	var group Group
	db.First(&group, id)
	return group
}

func getUsersInGroup(group Group) []User{
	var users []User
	db.Model(&group).Association("Users").Find(&users)
	return users
}

func getUsersByGroupId(id int) []User{
	var users []User
	group := getGroupById(id)
	db.Model(&group).Association("Users").Find(&users)
	return users
}

func getReportsByGroupId(id int) []Report{
	var reports []Report
	var group Group = getGroupById(id)
	db.Model(&group).Association("Reports").Find(&reports)
	return reports
}

func getReportsByGroupName(name string) []Report{
	var reports []Report
	var group Group = getGroupByName(name)
	db.Model(&group).Association("Reports").Find(&reports)
	return reports
}

func isNameGroupExisted(name string) bool{
	var group Group
	var count int
	db.Where("name = ?", name).Find(&group).Count(&count)
	if (count > 0){
		return true
	}
	return false
}

func createGroup(name string) bool{
	db.Create(&Group{Name:name})
	return true;
}

func getGroupByName(name string) Group {
	var group Group
	db.Where("name = ?",name).First(&group)
	return group
}

func addUserToGroup(email string, name_group string) bool{
	user := getUserByEmail(email)
	group:= getGroupByName(name_group)
	db.Model(&user).Association("Groups").Append(group)
	return true;
}

func deleteUserInGroup(email string, name string) bool{
	//Todo
	return false
}

