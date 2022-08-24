package timeLib

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

func ShowTime() error {
	if time, err := ntp.Time("0.beevik-ntp.pool.ntp.org"); err != nil {
		fmt.Fprintf(os.Stderr, "TimeERREOR: %v", err)
		return err
	} else {
		fmt.Println("Cur time: ", time)
	}
	return nil
}
