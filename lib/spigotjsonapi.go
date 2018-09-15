package mpspigotjsonapi

import "fmt"
import "crypto/sha256"
import "encoding/hex"
import "encoding/json"
import "log"
//import "net/http"
//import "net/url"

type requestEntries []struct {
        Name      string        `json:"name"`
        Key       string        `json:"key"`
        Username  string        `json:"username"`
        Arguments []string      `json:"arguments"`
        Tag       string        `json:"tag,omitempty"`
}

func main() {
	var api_username = "username"
	var api_password = "password"
//	var server_host	 = "localhost"
//	var server_port  = 25565

	var request_json = buildRequest(api_username, api_password)
	fmt.Printf(request_json)
}


func signRequestEntry(username, password, api_method string) string {
	req := sha256.New()
	req.Write([]byte(username + api_method + password))
	return hex.EncodeToString(req.Sum(nil))
}

func buildRequest(api_username string, api_password string) string {

        s := requestEntries{
                {
                        Name: "server.performance.memory.used",
                        Key:   signRequestEntry(api_username, api_password, "server.performance.memory.used"),
                        Username: api_username,
                        Arguments: make([]string, 0),
                },
                {
                        Name: "server.performance.tick_health",
                        Key:   signRequestEntry(api_username, api_password, "server.performance.tick_health"),
                        Username: api_username,
                        Arguments: make([]string, 0),
                },
                {
                        Name: "players.online.count",
                        Key:   signRequestEntry(api_username, api_password, "players.online.count"),
                        Username: api_username,
                        Arguments: make([]string, 0),
                },
                {
                        Name: "players.online.limit",
                        Key:   signRequestEntry(api_username, api_password, "players.online.limit"),
                        Username: api_username,
                        Arguments: make([]string, 0),
                },

        }

        buf, err := json.Marshal(s)
        if err != nil {
                log.Fatal(err)
        }

	return fmt.Sprintf("%s\n", buf)
}
