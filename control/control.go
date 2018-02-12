package control

import (
	cfg "github.com/alknopfler/alkalarm/config"
	"fmt"
	"github.com/alknopfler/alkalarm/database"
)


func Register(data cfg.Control) error{
	db,err := database.InitDB()
	if err != nil {
		fmt.Println("Error initiating DB in Register Control")
		return err
	}
	defer db.Close()

	err=database.Operate(db,cfg.CONTROL_INSERT,data.Code,data.Description,data.TypeOf)
	if err!=nil{
		fmt.Println("Error inserting control in db")
		return err
	}
	fmt.Println("Success...Control registered successfully")
	return nil
}

func Unregister(code string) error{
	db,err := database.InitDB()
	if err != nil {
		fmt.Println("Error initiating DB in Register Control")
		return err
	}
	defer db.Close()

	err=database.Operate(db,cfg.CONTROL_DELETE,code)
	if err!=nil{
		fmt.Println("Error inserting control in db")
		return err
	}
	fmt.Println("Success...control registered successfully")
	return nil
}


func QueryAll() ([]cfg.Control,error){
	var result []cfg.Control
	db,err := database.InitDB()
	if err != nil {
		fmt.Println("Error initiating DB in Query control")
		return result,err
	}
	defer db.Close()
	rows, err := db.Query(cfg.CONTROL_QUERY_ALL)
	if err != nil { return result,err }
	defer rows.Close()

	for rows.Next() {
		item := cfg.Control{}
		err2 := rows.Scan(&item.Code, &item.Description, &item.TypeOf)
		if err2 != nil { return nil,err }
		result = append(result, item)
	}
	return result, nil
}

func Query(code string) (cfg.Control,error){
	var result cfg.Control
	db,err := database.InitDB()
	if err != nil {
		fmt.Println("Error initiating DB in Query Control")
		return result,err
	}
	defer db.Close()
	rows, err := db.Query(cfg.CONTROL_QUERY_CODE,code)
	if err != nil { return result,err }
	defer rows.Close()

	if rows.Next() {
		err2 := rows.Scan(&result.Code, &result.Description, &result.TypeOf)
		if err2 != nil { return result,err }
	}
	return result, nil
}

func QueryTypeOf(code string) string {
	c,_:=Query(code)
	return c.TypeOf
}

func Exists(code string) bool{
	db,err := database.InitDB()
	if err != nil {
		fmt.Println("Error initiating DB in Control Exists")
		return false
	}
	defer db.Close()
	rows, err := db.Query(cfg.CONTROL_QUERY_CODE,code)
	defer rows.Close()
	if rows.Next(){
		return true
	}
	return false
}
