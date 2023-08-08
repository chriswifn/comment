package comment

import (
	"bufio"
	_ "embed"
	"fmt"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/conf"
	"github.com/rwxrob/help"
	"github.com/rwxrob/vars"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	comment    = "// "
	hrulewidth = "72"
	filler     = "-"
	LineExp    = regexp.MustCompile(`(.*\S.*)`)
)

// go: embed text/comment.md
var commentDesc string

// go: embed text/uncomment.md
var uncommentDesc string

func init() {
	Z.Conf.SoftInit()
	Z.Vars.SoftInit()
	Z.Dynamic[`dcomment`] = func() string { return comment }
	Z.Dynamic[`dhrulewidth`] = func() string { return hrulewidth }
	Z.Dynamic[`filler`] = func() string { return filler }
}

var Cmd = &Z.Cmd{
	Name:      `comment`,
	Aliases:   []string{`com`},
	Usage:     `[help|PATH]`,
	Copyright: `Copyright 2023 Christian Hageloch`,
	Version:   `v0.1.0`,
	License:   `MIT`,
	Source:    `git@github.com:chriswifn/comment.git`,
	Issues:    `github.com/chriswifn/comment/issues`,
	Commands:  []*Z.Cmd{help.Cmd, vars.Cmd, conf.Cmd, commentCmd, uncommentCmd, initCmd, htitleCmd},
}

var initCmd = &Z.Cmd{
	Name:     `init`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(x *Z.Cmd, _ ...string) error {
		val, _ := x.Caller.C(`comment`)
		if val == "null" {
			val = comment
		}
		x.Caller.Set(`comment`, val)

		val, _ = x.Caller.C(`hrulewidth`)
		if val == "null" {
			val = hrulewidth
		}
		x.Caller.Set(`hrulewidth`, val)

		val, _ = x.Caller.C(`filler`)
		if val == "null" {
			val = filler
		}
		x.Caller.Set(`filler`, val)
		return nil
	},
}

var commentCmd = &Z.Cmd{
	Name:        `comment`,
	Aliases:     []string{`com`},
	Usage:       `[help|PATH]`,
	Summary:     `add a comment string to the beginning of line(s)`,
	Copyright:   `2023 Christian Hageloch`,
	License:     `MIT`,
	Source:      `git@github.com:chriswifn/comment.git`,
	Issues:      `github.com/chriswifn/comment/issues`,
	Commands:    []*Z.Cmd{help.Cmd},
	Description: commentDesc,
	Call: func(x *Z.Cmd, args ...string) error {
		com, err := x.Caller.Get("comment")
		if err != nil {
			return err
		}

		if len(args) == 1 {
			com = args[0]
		}

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := com + scanner.Text()
			fmt.Println(line)
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error: ", err)
			os.Exit(1)
		}
		return nil
	},
}

var uncommentCmd = &Z.Cmd{
	Name:        `uncomment`,
	Aliases:     []string{`ucom`},
	Usage:       `[help|PATH]`,
	Summary:     `remove a comment string on the beginning of line(s)`,
	Copyright:   `2023 Christian Hageloch`,
	License:     `MIT`,
	Source:      `git@github.com:chriswifn/comment.git`,
	Issues:      `github.com/chriswifn/comment/issues`,
	Commands:    []*Z.Cmd{help.Cmd},
	Description: uncommentDesc,
	Call: func(x *Z.Cmd, args ...string) error {
		com, err := x.Caller.Get("comment")
		if err != nil {
			return err
		}

		if len(args) == 1 {
			com = args[0]
		}

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := strings.Replace(scanner.Text(), com, "", -1)
			fmt.Println(line)
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error: ", err)
			os.Exit(1)
		}
		return nil
	},
}

var htitleCmd = &Z.Cmd{
	Name:        `htitle`,
	Aliases:     []string{`ht`},
	Usage:       `[help|PATH]`,
	Summary:     `remove a comment string on the beginning of line(s)`,
	Copyright:   `2023 Christian Hageloch`,
	License:     `MIT`,
	Source:      `git@github.com:chriswifn/comment.git`,
	Issues:      `github.com/chriswifn/comment/issues`,
	Commands:    []*Z.Cmd{help.Cmd},
	Description: uncommentDesc,
	Call: func(x *Z.Cmd, args ...string) error {
		com, err := x.Caller.Get("comment")
		if err != nil {
			return err
		}

		hrule, err := x.Caller.Get("hrulewidth")
		if err != nil {
			return err
		}

		hrulew, err := strconv.Atoi(hrule)
		if err != nil {
			return err
		}

		fil, err := x.Caller.Get("filler")
		if err != nil {
			return err
		}

		if len(args) == 1 {
			com = args[0]
		}

		if len(args) == 2 {
			com = args[0]
			fil = args[1]
		}

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			text := scanner.Text()
			length := len(text)
			side := (hrulew/2 - length/2) - len(com)
			left := side
			right := side
			if length%2 == 1 {
				right--
			}
			line := LineExp.ReplaceAllString(text, comment+strings.Repeat(fil, left)+` `+`$1`+` `+strings.Repeat(fil, right))
			fmt.Println(line)
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error: ", err)
			os.Exit(1)
		}
		return nil
	},
}
