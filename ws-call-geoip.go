// Example How To Call Web Service from Go
//
// Copyright (C), Henryk Paluch, BSD license
//
// Note: We are Calling REST like variant of Web Service 
//       (GET requeust, data are returned as XML without SOAP Envelope)
// See:  http://www.webservicex.net/WS/WSDetails.aspx?CATID=12&WSID=64
//       for WS specification
//
// Compile and run using:
//      go run ws-call-geoip.go
//
package main

import (
	"encoding/xml"
	"net/http"
	"net/url"
	"io/ioutil"
	"fmt"
)

type  GeoIPResponse struct {
  XMLName xml.Name `xml:"GeoIP"`
  ReturnCode int
  IP string
  ReturnCodeDetails string
  CountryName string
  CountryCode string
}

func deSerializeXML(data []byte) (*GeoIPResponse,error) {
    v := new(GeoIPResponse)	 
    err := xml.Unmarshal(data, v)
    return v,err
}

func HPGetGeoIP( ip string) (*GeoIPResponse,error){
    url := "http://www.webservicex.net/geoipservice.asmx/GetGeoIP?IPAddress="+url.QueryEscape(ip)
    resp,err := http.Get(url)
    if err!=nil {
	    return nil,err
    }
    body,err2:= ioutil.ReadAll(resp.Body)
    if err2!=nil {
	    return nil,err2
    }
    v,err3 := deSerializeXML(body)
    if err3!=nil {
	    return nil,err3
    }
    return v,nil
}

func main(){
	addresses := []string{"8.8.4.4","194.79.52.192","134.76.12.3","213.180.204.46"}
	for _,ip := range addresses {
		v,err := HPGetGeoIP(ip)
		if err!=nil {
			panic(fmt.Sprintln("Error on GetGeoIP",err))
		}
		fmt.Printf("IP %15s (%d,%s) %s %s\n",v.IP,v.ReturnCode,v.ReturnCodeDetails,v.CountryName,v.CountryCode)
       }
}

