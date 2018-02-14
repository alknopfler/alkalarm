package states

import (
	cfg "github.com/alknopfler/alkalarm/config"
	"log"
	"github.com/alknopfler/alkalarm/database"
)



func Update(newstate string) error{
	db,err := database.InitDB()
	if err != nil {
		log.Println("Error initiating DB in update Global State")
		return err
	}
	defer db.Close()

	err=database.Operate(db,cfg.GLOBAL_STATE_UPDATE,newstate)
	if err!=nil{
		log.Println("Error updating the global state in db")
		return err
	}
	log.Println("Success...Global State updated to ",newstate)
	return nil
}

func Query() string{
	var result cfg.GlobalState
	db,err := database.InitDB()
	if err != nil {
		log.Println("Error initiating DB in Query State")
		return ""
	}
	defer db.Close()
	rows, err := db.Query(cfg.GLOBAL_STATE_QUERY)
	if err != nil { return "" }
	defer rows.Close()

	if rows.Next() {
		err2 := rows.Scan(&result.Id,&result.GState)
		if err2 != nil { return ""}
	}
	return result.GState
}

