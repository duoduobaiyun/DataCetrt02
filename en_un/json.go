package main

import (
	"DataCertProject_Me/models"
	"encoding/json"
	"fmt"
)

func main()  {
	user1:=models.User{
		Id: 1,
		Phone: "1515616111",
		Password: "abdadda",
	}
	fmt.Println("内存当中的数据user1:",user1)

	jsonBvtes,_:=json.Marshal(user1)
	fmt.Println("内存中的数据User1:",string(jsonBvtes))


	var  user2 models.User
	json.Unmarshal(jsonBvtes,&user2)
	fmt.Println("反序列的user2:",user2)



}
