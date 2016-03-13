// The MIT License (MIT)
//
// Copyright (c) 2016 aerth
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

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


    fmt.Println("Welcome to Sendgrid.\n\nMail to who?")
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

    fmt.Println("\nDestination Name? Press ENTER for "+maildestination)
		maildestinationname := getTypin()
    if maildestinationname == "" {
      maildestinationname = maildestination
    }

    fmt.Println("\nSUBJECT: Press ENTER for no subject.")
		mailsubject := getTypin()
    if mailsubject == "" {
      mailsubject = "<no subject>"
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
var mailfrom string
if os.Getenv("SENDGRID_FROM") != "" {
mailfrom = os.Getenv("SENDGRID_FROM")
}else{
    fmt.Println("From Address")
		mailfrom := getTypin()
    if mailfrom == "" {
      mailfrom = getTypin()
    }
    if mailfrom == "" {
      mailfrom = getTypin()
    }

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
    fmt.Println(mailfrom)
    fmt.Println(mailsubject)
    fmt.Println(maildestination + "(" + maildestinationname + ")")
    fmt.Println(mailbody)

    fmt.Println("\n\nYes [y/Y] to send\n\n")
    if !askForConfirmation() {
      fmt.Println("Mail not sent.")
    os.Exit(1)
    }
    if r := sg.Send(message); r == nil {
                fmt.Println("Email sent!")
        } else {
            fmt.Println(r)
            fmt.Println("Try again? [Y/y/Yes]")
          if !askForConfirmation() {
            fmt.Println("Mail not sent.")
          os.Exit(1)
          }
          if r := sg.Send(message); r == nil {
                      fmt.Println("Email sent!")
              } else {
                  fmt.Println(r)
                  fmt.Println("Try again? [Y/y/Yes]")
                if !askForConfirmation() {
                  fmt.Println("Mail not sent.")
                os.Exit(1)
                }
                if r := sg.Send(message); r == nil {
                            fmt.Println("Email sent!")
                    } else {
                        fmt.Println(r)
                        fmt.Println("Try again? [Y/y/Yes]")
                      if !askForConfirmation() {
                        fmt.Println("Mail not sent.")
                      os.Exit(1)
                      }
                      if r := sg.Send(message); r == nil {
                                  fmt.Println("Email sent!")
                          } else {
                              fmt.Println(r)
                              fmt.Println("Try again? [Y/y/Yes]")
                            if !askForConfirmation() {
                              fmt.Println("Mail not sent.")
                            os.Exit(1)
                            }

                            os.Exit(1)
                          }

                    }

              }

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
