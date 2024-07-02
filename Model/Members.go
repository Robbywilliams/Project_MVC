package Model

import (
	"Project_MVC/Node"
	"Project_MVC/Project_MVC/Database"
)

func InsertMember(id int, nama string, alamat string, point float32) {
	newDataMember := Node.TableMember{
		Data: Node.Member{Id: id, Nama: nama, Alamat: alamat, Point: point},
		Next: nil,
	}
	var tempLL *Node.TableMember
	tempLL = &Database.DBmember
	if Database.DBmember.Next == nil {
		Database.DBmember.Next = &newDataMember
	} else {
		for tempLL.Next != nil {
			tempLL = tempLL.Next
		}
		tempLL.Next = &newDataMember
	}
}

func ReadAllMember() *Node.TableMember {
	var tempLL *Node.TableMember
	tempLL = &Database.DBmember
	if Database.DBmember.Next == nil {
		return nil
	} else {
		return tempLL
	}
}
