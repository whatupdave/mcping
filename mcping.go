package mcping

import (
    "bufio"
    "bytes"
    "net"
    "fmt"
    "encoding/binary"
    "encoding/json"
    "github.com/jmoiron/jsonq"
    "strings"
)

func Ping(host string, port uint16) (MCPingResponse, error) {
    //If things go south, send default struct w/ error
    defaultResp := MCPingResponse{};

    fullAddr := host + ":" + fmt.Sprint(port)
    tcpAddr, err := net.ResolveTCPAddr("tcp", fullAddr); if(err != nil) {
        return defaultResp, resolveErr
    }
    
    //Start timer 
    timer := PingTimer{}
    timer.Start();

    //Connect
    conn, err := net.DialTCP("tcp", nil, tcpAddr); if(err != nil) {
        return defaultResp, connectErr
    }

    connReader := bufio.NewReader(conn)

    var dataBuf bytes.Buffer;

    var finBuf bytes.Buffer;

    dataBuf.Write([]byte("\x00")) //Packet ID
    dataBuf.Write([]byte("\x04")) //Protocol Version 47
    
    //Write host string length + host
    hostLength := uint8(len(host))
    dataBuf.Write([]uint8{hostLength})
    dataBuf.Write([]byte(host))

    //Write port
    b := make([]byte, 2)
    binary.BigEndian.PutUint16(b, port)
    dataBuf.Write(b)

    //Next state ping
    dataBuf.Write([]byte("\x01")) 

    //Prepend packet length with data
    packetLength := []byte{uint8(dataBuf.Len())}
    finBuf.Write(append(packetLength, dataBuf.Bytes()...)) 
    
    conn.Write(finBuf.Bytes()) //Sending handshake
    conn.Write([]byte("\x01\x00")) //Status ping

    //Get situationally useless full byte length
    binary.ReadUvarint(connReader)

    //Packet type 0 means we're good to recieve ping
    packetType,_ := connReader.ReadByte()
    if (bytes.Compare([]byte{packetType}, []byte("\x00")) != 0)  {
        return defaultResp, packetTypeErr
    }

    //Get data length via Varint
    length, err := binary.ReadUvarint(connReader)
    if (err != nil) {
        return defaultResp, varintErr
    }
    if (length < 10) {
        return defaultResp, smallPacketErr
    } else if(length > 700000) {
        return defaultResp, bigPacketErr
    }

    //Recieve json buffer
    bytesRecieved := uint64(0)
    recBytes := make([]byte, length)
    for (bytesRecieved < length) {
        n,_ := connReader.Read(recBytes[bytesRecieved:length])
        bytesRecieved = bytesRecieved + uint64(n);
    }

    //Stop Timer, collect latency
    latency := timer.End();

    pingString := string(recBytes);

    //Convert buffer into jsonq instance
    pingData := map[string]interface{}{}
    dec := json.NewDecoder(strings.NewReader(pingString))
    dec.Decode(&pingData)
    jq := jsonq.NewQuery(pingData)

    //Assemble PlayerSample
    playerSampleMap, err := jq.ArrayOfObjects("players", "sample")
    playerSamples := []PlayerSample{}
    for k, _ := range playerSampleMap {
        sample := PlayerSample {}
        sample.UUID = playerSampleMap[k]["id"].(string)
        sample.Name = playerSampleMap[k]["name"].(string)
        playerSamples = append(playerSamples, sample)
    }

    //Assemble MCPingResponse
    resp := MCPingResponse {}
    resp.Latency = uint(latency)
    resp.Online,_ = jq.Int("players", "online")
    resp.Max,_ = jq.Int("players", "max")
    resp.Protocol,_ = jq.Int("version", "protocol")
    resp.Favicon,_ = jq.String("favicon")
    resp.Motd,_ = jq.String("description")
    resp.Server,_ = jq.String("server")
    resp.Version,_ = jq.String("version", "name")
    //resp.Sample = playerSamples

    return resp, nil
}