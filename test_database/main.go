package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	//connect to database
	conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=test_connect user=postgres password=postgres")
	if err != nil {
		log.Fatalf(  "Unable to connect: %v\n",err.Error() )
	}
	defer conn.Close()

	log.Println("Connected to Database ...")

	//test my connection
	err = conn.Ping()
	if err!=nil{
		log.Fatal("Cannot Ping the Database!!",err.Error())
	}

	log.Println("Pinged Database !!! --->  Database is Connected !!!")


	//get rows from table
	err = getAllRows(conn)
	if err !=nil{
		log.Fatal("Getting Error in getting the rows")
	}

	//insert a row in a table
	query := `insert into users(first_name,last_name) values ($1,$2)`
	_,err  = conn.Exec(query,"Jack","Brown")
	if err!=nil{
		log.Fatal(err)
	}

	//get row from table again
	err = getAllRows(conn)
	if err !=nil{
		log.Fatal("Getting Error in getting the rows")
	}

	//update a row in a table
	updt := `update users set first_name=$1 where last_name=$2`
	_,err = conn.Exec(updt,"Rajesh","Bisht")
	if err !=nil{
		log.Fatal(err.Error())
	}

	log.Println("Updated one or more rows")

	//get row from table again
	err = getAllRows(conn)
	if err !=nil{
		log.Fatal("Getting Error in getting the rows")
	}

	//get one row by id
	query = `select first_name,last_name from users where first_name=$1`
	var FirstName,LastName string
	row := conn.QueryRow(query,"shobhit")

	err = row.Scan(&FirstName,&LastName)
	if err!=nil{
		log.Fatal(err.Error())
	}
	fmt.Printf("****************Query row returns*************************\nFirstName = %s\t\tLastName = %s\t\n",FirstName,LastName)
	fmt.Println("************************END*******************************")


	//delete a row
	query = `delete from users where first_name=$1`
	_,err   = conn.Exec(query,"Rahul")
	if err!=nil{
		log.Fatal(err.Error())
	}
	fmt.Println("DELETED row!! ")
	
	//get row from table again
	err = getAllRows(conn)
	if err !=nil{
		log.Fatal("Getting Error in getting the rows")
	}

}

//func to call all the row 
func getAllRows(conn *sql.DB)error{
	rows , err := conn.Query("select first_name,last_name from users")
	if err!=nil{
		log.Println("Can't fetch data!!!")
		return err
	}
	defer rows.Close()

	var FirstName, LastName string
	fmt.Println("###############Fetching User data from Database###############")
	for rows.Next(){
		err := rows.Scan(&FirstName , &LastName)
		if err!=nil{
			log.Println("Got error ")
			return err
		}
		
		fmt.Println("Data-Record is ",FirstName,LastName)
		fmt.Println("------------------------------------------------------")
	}
	if err=rows.Err();err !=nil{
		log.Fatal("Error in scanning rows")
			return err
	}

	
	return nil
}