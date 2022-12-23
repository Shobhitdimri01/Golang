package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func main() {
	ModifyTxtFile()
}

type NewTxt struct{
	S1ap_ip string `default:"10.258.33.99"`
}
var new NewTxt

type remoteserverconfig struct{
	username	string	`default:"user12"`
	password	string	`default:'userpassword'`
	serverip	string	`default:"127.0.0.1"`
}
 
var remote remoteserverconfig
var dstpath = "/home/user12/mytxtfile/test.txt"
func ModifyTxtFile(){
	sshconfig := &ssh.ClientConfig{
		User: remote.username,
		Auth: []ssh.AuthMethod{
			ssh.Password(remote.password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),// optional
	}

	client,err := ssh.Dial("tcp",remote.serverip,sshconfig)
	if err != nil {
		panic("failed to dail: "+err.Error())
	}
	fmt.Println("Succesfully connected to "+remote.serverip+" server")

	sftp,err :=sftp.NewClient(client)
	if err != nil {
		log.Fatal(err)
	}
	defer sftp.Close()

	file,err := sftp.Open(dstpath)
	if err != nil {
		log.Fatal(err)
	}
	input, err := io.ReadAll(file)
	if err != nil {
			log.Fatalln(err)
	}
	lines := strings.Split(string(input), "\n")
	Edit(lines,"s1ap",new.S1ap_ip)
	
	output := strings.Join(lines, "\n")
		localpath :="/home/user/test.txt"
		remotepath :=remote.username+"@"+remote.serverip+":"+dstpath
		err = ioutil.WriteFile(localpath, []byte(output), 0644)
        if err != nil {
                log.Fatalln(err)
        }  
	
		copyfile := "sshpass -p "+remote.password+" scp -r "+localpath+" "+remotepath
		cmd := exec.Command("bash","-c",copyfile)
		stdout,err := cmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		//printing output
		fmt.Println(string(stdout))
}

func Edit(lines []string,txt string,Key string){
	for i, line := range lines {
		if strings.Contains(line, txt) {
			lines[i] = "s1ap: "+Key
		}
	}
}
