func test(){
	ct = context.Background() userId :=
	r.URL.Query.Get("id")
	query := "select * from users where user_id = " + userId
	row, _ := db.QueryContext(ct, query)
}

//When we use this code userID field can exploit easily with  1 or 1=1?
//And our query'll be like this Select * from user where user_id = 1 or 1=1
//So you,ll dump all table records with 1 or 1=1 payload always ll be true.
//Theres a way u can protect your db

func test2(){
	ct := context.Background() userId :=
	r.URL.Query.GET("id")
	query := "select * from users where user_id = ?"
	stmt, _  := db.QueryContext(ct, query, user_id)
}