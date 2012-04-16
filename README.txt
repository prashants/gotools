********************************* portscan.go *********************************

Port scanner written in Go

Run the command as :

$go run portscanner.go 127.0.0.1

This program is licensed under GNU GPL V3 as available on
http://www.gnu.org/licenses/gpl-3.0.html

****************************** notifyquotes.go ********************************

Generate random quotes as notifications

Compile this program by executing :

$go build notifyquotes.go

This will create a binary executable called "notifyquotes" in the same
directory. Add this path to the user crontab.

$crontab -e
* * * * * DISPLAY=:0.0 XAUTHORITY=~/.Xauthority /path/to/notifyquotes

The above crontab entry will run the notifyquotes every 1 minute

This program is licensed under GNU GPL V3 as available on
http://www.gnu.org/licenses/gpl-3.0.html
