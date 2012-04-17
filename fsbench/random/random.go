/* This package generates random file and directory names */

package random

import "crypto/rand"
import mrand "math/rand"

const letters = "abcdefghijklmnopqrstuvwxyz"

var num int

func GenFileName() string {
	fname := make([]byte, 8)
	fext := make([]byte, 3)

	/* fill up buffer with random data */
	rand.Read(fname)
	rand.Read(fext)

	/* normalize the random data within the valid letter range */
	for i, j := 0, 0; j < len(fname); j++ {
		i = int(fname[j]) % len(letters)
		fname[j] = letters[i]
	}
	for i, j := 0, 0; j < len(fext); j++ {
		i = int(fext[j]) % len(letters)
		fext[j] = letters[i]
	}

	filename := string(fname) + "." + string(fext)
	return filename
}

func GenDirName() string {
	dname := make([]byte, 8)

	/* fill up buffer with random data */
	rand.Read(dname)

	for i, j := 0, 0; j < len(dname); j++ {
		i = int(dname[j]) % len(letters)
		dname[j] = letters[i]
	}

	return string(dname)
}

func GetSmallSize() int {
	return mrand.Intn(1024)
}

func GetMediumSize() int {
	num = mrand.Intn(1048576) /* 1MB = 1048576 */
	switch {
		case num < 10: return num * 1000
		case num < 100: return num * 100
		case num < 1000: return num * 10
	}
	return num
}

func GetLargeSize() int {
	num = mrand.Intn(524288000) /* 1MB = 1048576; 1048576 * 500 = 524288000 */
	switch {
		case num < 10: return num * 10000
		case num < 100: return num * 1000
		case num < 1000: return num * 100
		case num < 10000: return num * 10
	}
	return num
}

