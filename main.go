package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"gopkg.in/mcuadros/go-defaults.v1"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func main() {
	WriteFile()
}

type NewTxt struct{
        S1apIP string
        PLMNID string
        MCC    string
}
var newtxt NewTxt

type Remoteserverconfig struct{
        Username        string  `default:"deploy"`
        Password        string  `default:"deploy"`
        Serverip        string  `default:"10.254.12.194:22"`
}

func SetRemoteServer() *Remoteserverconfig {
    example := new(Remoteserverconfig)
    defaults.SetDefaults(example) //<-- This set the defaults values

    return example
}
var dstpath = "/home/deploy/testtxt/test.txt"
func WriteFile(){
	remote := SetRemoteServer()

        fmt.Println("user:",remote.Username)
        sshconfig := &ssh.ClientConfig{
                User: remote.Username,
                Auth: []ssh.AuthMethod{
                        ssh.Password(remote.Password),
                },
                HostKeyCallback: ssh.InsecureIgnoreHostKey(),// optional
        }

        client,err := ssh.Dial("tcp",remote.Serverip,sshconfig)
        if err != nil {
                panic("failed to dail: "+err.Error())
        }
        fmt.Println("Succesfully connected to "+remote.Serverip+" server")
	if c.Request.Body != nil{

                        err := json.NewDecoder(c.Request.Body).Decode(&newtxt)

                        if err != nil {
				fmt.Println(err.Error())
			}
	}
	sftp,err :=sftp.NewClient(client)
        if err != nil {
                log.Fatal(err)
        }
        defer sftp.Close()
	file,err := sftp.Open(dstpath)
        if err != nil {
                fmt.Println(err)
        }
        input, err := io.ReadAll(file)
        if err != nil {
                        fmt.Println(err)
        }
        lines := strings.Split(string(input), "\n")
        Edit(lines,"s1ap",newtxt.S1apIP)
        Edit(lines,"plmn_id",newtxt.PLMNID)
        Edit(lines,"mcc",newtxt.MCC)
	  output := strings.Join(lines, "\n")
                localpath :="/home/shobhitdimri/nfstub/main/4GSanityTesting/ENBS/test.txt"
                ip:=strings.Split(string(remote.Serverip), ":")
                remotepath :=remote.Username+"@"+ip[0]+":"+dstpath
                err = ioutil.WriteFile(localpath, []byte(output), 0644)
        if err != nil {
               fmt.Println(err)
        }
                copyfile := "sshpass -p "+remote.Password+" scp -r "+localpath+" "+remotepath
                cmd := exec.Command("bash","-c",copyfile)
                fmt.Println("Remote file edited")
                stdout,err := cmd.Output()
                if err != nil {
                        fmt.Println(err)
                }
                //printing output
                fmt.Println(string(stdout))
}
func Edit(lines []string,txt string,Key string){
        for i, line := range lines {
                if strings.Contains(line, txt) {
                        lines[i] = txt+": "+Key
                }
        }
}

	
