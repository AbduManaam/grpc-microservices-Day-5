package repository

import "database/sql"

type User struct {
	Id       string
	Username string
	Password string
}

type Repo struct {
	DB *sql.DB
}

func(c *Repo) CreateUser(username,password string)error{

    _,err:= c.DB.Exec(
		"INSERT INTO users(username,password)VALUES($1,$2)",
		username,password,
	)
	return err

}

func(c *Repo)GetByUserName(username string)(*User,error){

	row:= c.DB.QueryRow(
		"SELECT id,username,password FROM users WHERE username=$1",
		username,
	)

	var u User
	if err:= row.Scan(&u.Id,&u.Username,&u.Password);err!=nil{
		return nil, err
	}
	return &u,nil
}

func(c *Repo)GetUser(Id string)(*User,error){

    row:= c.DB.QueryRow(
		"SELECT id,username FROM users WHERE id=$1",
		Id,
	)
	var u User
	if err:= row.Scan(&u.Id,&u.Username);err!=nil{
		return nil,err
	}
    
	return &u,nil
}