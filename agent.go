package scout_apm_go

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gogo/protobuf/proto"
)

// Workflow:
//   StartAgent(OrgKey, AppName) ?
//   Wire up channel to start receiving transactions
//   HTTP Post to known endpoint w/ {OrgKey, Name}
//   Parse response (as json) and check for validity
//   Store connection detail (ip/port)
//   Start the background reporter loop

var scoutOrgKey string
var scoutAppName string
var scoutReportHost *string
var scoutAgentStarted bool
var scoutAgentChannel chan *ScoutTransaction

const scoutRegistrationUrl = "http://localhost:8090/register"

type AuthRequest struct {
	OrgKey  string `json:"org_key"`
	AppName string `json:"app_name"`
}

type AuthResponse struct {
	OrgId      int    `json:"org_id"`
	AppName    string `json:"app_name"`
	Authorized bool   `json:"authorized"`
	ReportHost string `json:"report_host"`
}

func StartAgent(orgKey string, appName string) {
	scoutAgentChannel = make(chan *ScoutTransaction, 100)
	scoutAgentStarted = false
	scoutOrgKey = orgKey
	scoutAppName = appName

	go func() {
		// ********** Create the JSON AuthRequest Data

		registrationRequest, err := json.Marshal(AuthRequest{
			OrgKey:  scoutOrgKey,
			AppName: scoutAppName,
		})
		if err != nil {
			return
		}

		fmt.Printf("Starting Agent (OrgKey: %s), (AppName: %s)\n", scoutOrgKey, scoutAppName)

		// ********** Make the HTTP Request

		resp, err := http.Post(scoutRegistrationUrl, "application/json", bytes.NewReader(registrationRequest))
		if err != nil {
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}

		// **********  Parse out the AuthResponse

		authResponse := AuthResponse{}
		if err := json.Unmarshal(body, &authResponse); err != nil {
			return
		}

		// **********  Check if we're authorized, and set the connection ip

		if authResponse.Authorized {
			scoutReportHost = &authResponse.ReportHost
		} else {
			return
		}

		// ********** And now start the agent

		scoutAgentStarted = true
		go reporter(scoutAgentChannel)
	}()
}

func ShipTransaction(t *ScoutTransaction) {
	if scoutAgentChannel != nil {
		// Channel wasn't full, but might be by now
		select {
		case scoutAgentChannel <- t: // Put 2 in the channel unless it is full
		default:
			fmt.Println("Channel was full when we attempted to write. Discarding transaction")
		}
	} else {
		log.Fatal("Nil channel")
	}
}

// Enter an infinite loop
func reporter(c chan *ScoutTransaction) {
	for true {
		if scoutReportHost == nil {
			fmt.Println("Stopping ScoutAgent due to missing ReportHost")
			break
		}

		fmt.Println("Scout Agent Dialing")
		conn, err := net.Dial("tcp", *scoutReportHost)
		if err != nil {
			// Failed to connect. Sleep, then fall through to the loop to retry
			fmt.Println("Failed to connect to localhost:8001, sleeping, then retrying")
			time.Sleep(1 * time.Second)
		} else {
			fmt.Println("Scout Agent Connected")
			innerLoop(conn, c)
			conn = nil
		}
	}

}

func innerLoop(conn net.Conn, c chan *ScoutTransaction) {
	// While we have a valid connection, stay here.
	for true {
		var transaction = <-c

		// Turn our type into something that can be encoded to protobuf later
		protobufStruct := transaction.EncodeToProtobuf()

		// Serialize to protobuf
		data, err := proto.Marshal(protobufStruct)
		if err != nil {
			log.Fatal("marshaling error: ", err)
		}

		// Calculate & serialize the length
		datalen := make([]byte, 4)
		binary.BigEndian.PutUint32(datalen, uint32(len(data)))

		// Write the length to the connection
		// On a failure, conneciton is invalid, break out of this inner loop
		_, err = conn.Write(datalen)
		if err != nil {
			break
		}

		// Write the proto data itself
		// On a failure, conneciton is invalid, break out of this inner loop
		_, err = conn.Write(data)
		if err != nil {
			break
		}
	}
}
