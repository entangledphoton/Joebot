package main 

import ( "fmt"
		"go-ircevent"
		"database/sql"
		_"github.com/lunny/godbc"
		"flag"
		"time"
		"runtime"
		"log"
		"github.com/andybons/hipchat"

		
)

var roomName = "#joebot-test"

func main() {
	// Start up SQL connection
	
		var (
		mssrv = flag.String("mssrv", "NCC1701D\\KEPLER", "NCC1701D\\KEPLER")
		msdb  = flag.String("msdb", "JoeBotDB", "JoeBotDB")
		//	dsn		= flag.String("dsn", "", "NCC1701D")
		msuser   = flag.String("msuser", "", "")
		mspass   = flag.String("mspass", "", "")
		msdriver = flag.String("msdriver", defaultDriver(), "NCC1701D")
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

	// Set up IRC Connection
    con := irc.IRC("drd_1812", "drd_1812")
    err1 := con.Connect("irc.freenode.net:6667")
    if err1 != nil {
        fmt.Println("Failed connecting")
        return
    }
    
    //Callbacks are actions based on certain event triggers
    
        con.AddCallback("001", func (e *irc.Event) {
        con.Join(roomName)
    })
    con.AddCallback("JOIN", func (e *irc.Event) {
        con.Privmsg(roomName, "ALL HAIL THE GLOW CLOUD.")
    })
        con.AddCallback("JOIN", func (e *irc.Event) {
        con.Privmsg("nickserv", "identify entangledphoton PASSWORD")
    })
    con.AddCallback("PRIVMSG", func (e *irc.Event) {
//        con.Privmsg(roomName, e.Nick + ": " + e.Message())

		stmt, err := db.Prepare("INSERT INTO EVENT_LOG(Code, Nick, Message, Time, Host, Source, Channel, Raw) VALUES(?,?,?,?,?,?,?,?)")
//		stmt, err := db.Prepare("INSERT INTO EVENT_LOG VALUES(?,?,?,?,?,?,?,?)")
if err != nil {
    	fmt.Println("Cannot connect: ", err.Error())
			return
}
dt := time.Now().Round(time.Second)


res, err := stmt.Exec(e.Code, e.Nick, e.Message(), dt, e.Host, e.Source, roomName, e.Raw )
//res, err := stmt.Exec(e.Code, e.Raw, e.Nick, e.Host, e.Source, e.User, dt , e.Message())
if err != nil {
    			fmt.Println("Cannot connect: ", err.Error())
			return
}
lastId, err := res.LastInsertId()
if err != nil {
    			fmt.Println("Cannot connect: ", err.Error())
			return
}
rowCnt, err := res.RowsAffected()
if err != nil {
    		fmt.Println("Cannot connect: ", err.Error())
			return
			}
	fmt.Printf("ID = %d, affected = %d\n", lastId, rowCnt)

    }) 
    con.AddCallback("PRIVMSG" , func (e *irc.Event) {
        	var msg string
			msg =  e.Nick + ": "+ e.Message()
			hep, err := HepCat(msg)
			if err != nil {	
			fmt.Println("Hip Returns: %s", err)
			return
			}
			fmt.Println("Hep is: ", hep)
        })
    con.Loop()
}


func defaultDriver() string {
	if runtime.GOOS == "windows" {
		return "sql server"
	} else {
		return "freetds"
	}
}

func HepCat(msg string)(hep string, err error){
	c := hipchat.Client{AuthToken: ""}
	req := hipchat.MessageRequest{
		RoomId:        "",
		From:          "",
		Message:       msg,
		Color:         hipchat.ColorPurple,
		MessageFormat: hipchat.FormatText,
		Notify:        true,
	}
	if err := c.PostMessage(req); err != nil {
		log.Printf("Expected no error, but got %q", err)
		
	}
	return "crap", err
}
