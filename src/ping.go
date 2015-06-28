package ping

import (
	_ "code.google.com/p/odbc"
	"database/sql"
	"flag"
	"fmt"
	"runtime"
)

func main() {

	var (
		mssrv = flag.String("mssrv", "", "")
		msdb  = flag.String("msdb", "", "")
		//	dsn		= flag.String("dsn", "", "")
		msuser   = flag.String("msuser", "", "")
		mspass   = flag.String("mspass", "", "")
		msdriver = flag.String("msdriver", defaultDriver(), "")
	//	msport   = flag.String("msport", "1433", "1433")
	)

	params := map[string]string{
		"driver": *msdriver,
		//		"dsn":	*dsn,
		"server":   *mssrv,
		"database": *msdb,
	}

	if len(*msuser) == 0 {
		params["trusted_connection"] = "yes"
	} else {
		params["uid"] = *msuser
		params["pwd"] = *mspass
	}

	var c string
	for n, v := range params {
		c += n + "=" + v + ";"
	}
	db, err := sql.Open("odbc", c)
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return
	}
	
	
	
	
	
	

	rows, err := db.Query("SELECT EVENT_MESSAGE FROM EVENT_LOG WHERE EVENT_NICK = 'entangledphoton'")
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return
	}
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			fmt.Println("Cannot connect: ", err.Error())
			return
		}
		fmt.Printf("%s is\n", name)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return
	}

}

func defaultDriver() string {
	if runtime.GOOS == "windows" {
		return "sql server"
	} else {
		return "freetds"
	}
}
