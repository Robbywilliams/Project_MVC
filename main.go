package main

import (
	"fmt"
	"src/Model"
)

func main() {
	Model.InsertMember(1, "Kurniawan", "surabaya", 100)
	Model.InsertMember(2, "Aan", "surabaya", 50)
	members := Model.ReadAllMember()
	if members != nil {
		for members.Next != nil {
			fmt.Println(members.Next.Data)
			members = members.Next
		}
	}
}
