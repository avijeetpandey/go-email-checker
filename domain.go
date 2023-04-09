package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func checkDomain(domain string) {
	var hasDMARC, hasSPF, hasMX bool
	var dmarcRecord, spfRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v", err)
	}

	// checking if mxRecords exists
	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error: %v", err)
	}

	// grabbing up the SPF records
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	// looking up dmarcRecords
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		log.Printf("Error: %v", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Println("The anaylysis of the domain %v is done", domain)
	if hasMX {
		fmt.Println("MXRecords: %v", mxRecords)
	}

	if hasSPF {
		fmt.Println("SPFRecords: %v", spfRecord)
	}

	if hasDMARC {
		fmt.Println("DMARCRecords: %v", dmarcRecord)
	}
}
