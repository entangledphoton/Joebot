package main


import (
	"fmt"
	"flag"
	"runtime"
    _ "github.com/lunny/godbc"
    "github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func main() {
// Start up SQL connection
	
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
	var err error
    engine, err = xorm.NewEngine("odbc", c)
    engine.ShowSQL = true
    	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return
		}

//	db, err := sql.Open("odbc", c)
//	if err != nil {
//		fmt.Println("Cannot connect: ", err.Error())
//		return
//	}
	}

func defaultDriver() string {
	if runtime.GOOS == "windows" {
		return "sql server"
	} else {
		return "freetds"
	}
}
