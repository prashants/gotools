/* generates random quotes as notifications on ubuntu machine */

/* This program is licensed under GNU GPL V3 as available on  http://www.gnu.org/licenses/gpl-3.0.html */

/*

  HOW TO USE THIS PROGRAM :

  Compile this program by executing :

	$go build notifyquotes.go

  This will create a binary executable called "notifyquotes" in the same
  directory. Add this path to the user crontab.

	$crontab -e
	* * * * * DISPLAY=:0.0 XAUTHORITY=~/.Xauthority * /path/to/notifyquotes

  The above crontab entry will run the notifyquotes every 1 minute

*/

package main

import "os/exec"
import "time"

func main() {
	cmd := exec.Command("notify-send", randomQuote())
	cmd.Run()
}

/* return a quote depending on the current minutes value */
func randomQuote() string {
	/* PUT ALL YOUR QUOTES IN THE BELOW STRING SLICE */
	var quotes = []string{"Quotes1",
		"Quotes 2",
		"Quotes 3"}
	_, min, _ := time.Now().Clock()
	cur := min % len(quotes)
	return quotes[cur]
}

