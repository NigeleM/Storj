package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"storj.io/uplink"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {

	APIKey := ""
	satelliteAddress := ""
	bucketName := ""
	rootPassphrase := ""
	localFILEorDIR := ""
	RemoteFileorDIR := ""

	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tok := scanner.Text()
		if strings.Contains(tok, "APIKey") {
			api := strings.Split(tok, "=")
			APIKey = api[1]
		} else if strings.Contains(tok, "satelliteAddress") {
			sat := strings.Split(tok, "=")
			satelliteAddress = sat[1]
		} else if strings.Contains(tok, "bucketName") {
			buck := strings.Split(tok, "=")
			bucketName = buck[1]
		} else if strings.Contains(tok, "rootPassphrase") {
			root := strings.Split(tok, "=")
			rootPassphrase = root[1]
		} else if strings.Contains(tok, "localFILEorDIR") {
			loca := strings.Split(tok, "=")
			localFILEorDIR = loca[1]
		} else if strings.Contains(tok, "RemoteFileorDIR") {
			Remo := strings.Split(tok, "=")
			RemoteFileorDIR = Remo[1]
			if RemoteFileorDIR == " " {
				RemoteFileorDIR = ""
			}
		}
	}

	access, err := uplink.RequestAccessWithPassphrase(context.Background(), satelliteAddress, APIKey, rootPassphrase)
	check(err)
	project, err := uplink.OpenProject(context.Background(), access)
	check(err)
	defer project.Close()
	object, err := project.DownloadObject(context.Background(), bucketName, RemoteFileorDIR, nil)
	check(err)
	defer object.Close()
	Contents, err := io.ReadAll(object)
	localfile, err := os.Create(localFILEorDIR)
	check(err)
	defer localfile.Close()
	localfile.Write(Contents)
	localfile.Close()
	// fmt.Println(Contents)
	// if !bytes.Equal(receivedContents, dataToUpload) {
	// 	return fmt.Errorf("got different object back: %q != %q", dataToUpload, receivedContents)
	// }
	// localfile, err := os.Open(localFILEorDIR)
	// check(err)
	// defer file.Close()
	// fileInfo, err := localfile.Stat()
	// check(err)

	// if fileInfo.IsDir() {
	// 	c, err := os.ReadDir(localFILEorDIR)
	// 	check(err)
	// 	for _, data := range c {
	// 		fmt.Println(data.Name())
	// 		if RemoteFileorDIR == "" {
	// 			object, err := project.UploadObject(context.Background(), bucketName, data.Name(), nil)
	// 			check(err)
	// 			file, err := os.Open(localFILEorDIR + data.Name())
	// 			check(err)
	// 			defer file.Close()
	// 			scanner := bufio.NewScanner(file)
	// 			content := make([]byte, 0)
	// 			for scanner.Scan() {
	// 				tok := scanner.Text()
	// 				byter := []byte(tok)
	// 				newline := []byte("\n")
	// 				content = append(content, byter...)
	// 				content = append(content, newline...)
	// 			}
	// 			object.Write(content)
	// 			object.Commit()
	// 			fmt.Println(localFILEorDIR, "file committed to storj")

	// 		} else {
	// 			object, err := project.UploadObject(context.Background(), bucketName, RemoteFileorDIR+"/"+data.Name(), nil)
	// 			check(err)
	// 			file, err := os.Open(localFILEorDIR + data.Name())
	// 			check(err)
	// 			defer file.Close()
	// 			scanner := bufio.NewScanner(file)
	// 			content := make([]byte, 0)
	// 			for scanner.Scan() {
	// 				tok := scanner.Text()
	// 				byter := []byte(tok)
	// 				newline := []byte("\n")
	// 				content = append(content, byter...)
	// 				content = append(content, newline...)
	// 			}
	// 			object.Write(content)
	// 			object.Commit()
	// 			fmt.Println(localFILEorDIR, "files committed to storj")
	// 		}

	// 	}
	// } else {
	// 	if RemoteFileorDIR == "" {
	// 		localFILE := ""
	// 		if strings.Contains(localFILEorDIR, "/") {
	// 			localFILE = localFILEorDIR[strings.LastIndex(localFILEorDIR, "/")+1 : len(localFILEorDIR)]
	// 		}
	// 		object, err := project.UploadObject(context.Background(), bucketName, localFILE, nil)
	// 		check(err)
	// 		file, err := os.Open(localFILEorDIR)
	// 		check(err)
	// 		defer file.Close()
	// 		scanner := bufio.NewScanner(file)
	// 		content := make([]byte, 0)
	// 		for scanner.Scan() {
	// 			tok := scanner.Text()
	// 			byter := []byte(tok)
	// 			newline := []byte("\n")
	// 			content = append(content, byter...)
	// 			content = append(content, newline...)
	// 		}
	// 		object.Write(content)
	// 		object.Commit()
	// 		fmt.Println(localFILEorDIR, "file committed to storj")

	// 	} else if strings.Contains(RemoteFileorDIR, "/") {
	// 		object, err := project.UploadObject(context.Background(), bucketName, RemoteFileorDIR+"/"+localFILEorDIR, nil)
	// 		check(err)
	// 		file, err := os.Open(localFILEorDIR)
	// 		check(err)
	// 		defer file.Close()
	// 		scanner := bufio.NewScanner(file)
	// 		content := make([]byte, 0)
	// 		for scanner.Scan() {
	// 			tok := scanner.Text()
	// 			byter := []byte(tok)
	// 			newline := []byte("\n")
	// 			content = append(content, byter...)
	// 			content = append(content, newline...)
	// 		}
	// 		object.Write(content)
	// 		object.Commit()
	// 		fmt.Println(localFILEorDIR, "file committed to storj")

	// 	}

	// }

}
