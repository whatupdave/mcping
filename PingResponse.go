package mcping


//PingResponse contains all known fields of the ping response packet
type PingResponse struct {
	Latency  uint   //Latency in ms
	Online   int    //Amount of online players
	Max      int    //Maximum amount of players
	Protocol int    //E.g '4'
	Favicon  string //Base64 encoded favicon in data URI format
	Motd     string
	Server   string //E.g 'PaperSpigot'
	Version  string //E.g "1.7.10"
	Sample   []PlayerSample
}
