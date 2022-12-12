package models

/*
	вспомогательная структура
   	для хранения связи между ребенком и родителями
*/

type StudentParentsCore struct {
	Student *StudentCore
	Parents []*ParentCore
}

func (ht *StudentParentsHTTP) FromCore(core *StudentParentsCore) {
	ht.Student.FromCore(core.Student)
	for _, parentCore := range core.Parents {
		parentHttpTemp := &ParentHTTP{
			UserHTTP: &UserHTTP{},
		}
		parentHttpTemp.FromCore(*parentCore)
		ht.Parents = append(ht.Parents, parentHttpTemp)
	}
}
