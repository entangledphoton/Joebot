package main 

import (
		"fmt"
		"go-ircevent"
		//"database/sql"
	//	_"code.google.com/p/odbc"
		"flag"
	//	"time"
	//	"runtime"
		"log"
		"github.com/andybons/hipchat"
)

var roomName = "#joebot-test"
const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func main() {
    flag.Parse() // Scan the arguments list 
    
    //First queury the database and store the latest message id to initialize the data
    //Get messages from the room and recod the last id as the new first id to query 
    //run this check every .5 seconds
    con := irc.IRC("drd_1812", "drd_1812")
    err1 := con.Connect("irc.freenode.net:6667")
    if err1 != nil {
        fmt.Println("Failed connecting")
        return
    }
    
    
	c := hipchat.Client{AuthToken: "3e3ca2b2bcee915cd46c3cd133727f"}	
	if repls, err := c.RoomHistory("Stoopkids", "recent", "UTC" ); err != nil {
		log.Printf("Expected no error, but got %q", err)
	
	var dat map[string]interface{}
			
	con.Privmsg(roomName, []repls.Message )
	}
    if *versionFlag {
        fmt.Println("Version:", APP_VERSION)
    }
}

