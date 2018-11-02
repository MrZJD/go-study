package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Servers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []Server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type Server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

func ReadXML() {
	file, err := os.Open("D:/repo/code/golang/src/learning/webs/text/xml/target.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	v := Servers{}

	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	fmt.Println(v)
}

func WriteXML() {
	v := &Servers{
		Version: "1",
	}
	v.Svs = append(v.Svs, Server{
		ServerName: "SHANG_VPN",
		ServerIP:   "127.0.0.1",
	})
	v.Svs = append(v.Svs, Server{
		ServerName: "GUANGZHOU_VPN",
		ServerIP:   "127.0.0.2",
	})
	output, err := xml.MarshalIndent(v, " ", "  ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}

func main() {

	ReadXML()

	WriteXML()

}
