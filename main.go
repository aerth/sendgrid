package main

import (
  "bufio"
        "fmt"
        "github.com/sendgrid/sendgrid-go"
        "os"
)

func main() {
    sendgridKey := os.Getenv("SENDGRID_API_KEY")
    if sendgridKey == "" {
                fmt.Println("Environment variable SENDGRID_API_KEY is undefined. Did you forget to source sendgrid.env?")
                os.Exit(1);
    }


    fmt.Println("Mail to who?")
		maildestination := getTypin()
    if maildestination == "" {
      maildestination = getTypin()
    }
    if maildestination == "" {
      maildestination = getTypin()
    }
    if maildestination == "" {
     os.Exit(1)
    }

    fmt.Println("Destination Name? Press ENTER for "+maildestination)
		maildestinationname := getTypin()
    if maildestinationname == "" {
      maildestinationname = maildestination
    }

    fmt.Println("SUBJECT: Press ENTER three times for no subject.")
		mailsubject := getTypin()
    if mailsubject == "" {
      mailsubject = getTypin()
    }
    if mailsubject == "" {
      mailsubject = getTypin()
    }

    fmt.Println("Message Text: Press ENTER when finished.")
		mailbody := getTypin()
    if mailbody == "" {
      mailbody = getTypin()
    }
    if mailbody == "" {
      mailbody = getTypin()
    }
    if mailbody == "" {
     os.Exit(1)
    }

    fmt.Println("From Address")
		mailfrom := getTypin()
    if mailfrom == "" {
      mailfrom = getTypin()
    }
    if mailfrom == "" {
      mailfrom = getTypin()
    }
    if mailfrom == "" {
     os.Exit(1)
    }

    sg := sendgrid.NewSendGridClientWithApiKey(sendgridKey)
    message := sendgrid.NewMail()
    message.AddTo(maildestination)
    message.AddToName(maildestinationname)
    message.SetSubject(mailsubject)
    message.SetText(mailbody)
    message.SetFrom(mailfrom)
    fmt.Println(message)
    if !askForConfirmation() {
    os.Exit(1)
  }
    if r := sg.Send(message); r == nil {
                fmt.Println("Email sent!")
        } else {
                fmt.Println(r)
        }
}



// Does x contain y?
func containsString(slice []string, element string) bool {
	return !(posString(slice, element) == -1)
}

// Ask user to confirm the action.
func askForConfirmation() bool {
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	okayResponses := []string{"y", "Y", "yes", "Yes", "YES"}
	nokayResponses := []string{"n", "N", "no", "No", "NO"}
	quitResponses := []string{"q", "Q", "exit", "quit"}
	if containsString(okayResponses, response) {
		return true
	} else if containsString(nokayResponses, response) {
		return false
	} else if containsString(quitResponses, response) {
		return false
	} else {
		fmt.Println("\nNot valid answer, try again. [y/n] [yes/no]")
		return askForConfirmation()
	}
}

// For use only in containsString()
func posString(slice []string, element string) int {
	for index, elem := range slice {
		if elem == element {
			return index
		}
	}
	return -1
}

// Receive non-hidden input from user.
func getTypin() string {
	fmt.Printf("\nPress ENTER when you are finished typing.\n\n")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		//	fmt.Println(line)
		return line
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return ""
}
