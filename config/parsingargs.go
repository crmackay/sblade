package config

import (
//"github.com/codegangsta/cli"
//"os"
)

/*
func ParseArgs() {
	app := cli.NewApp()
	app.Name = "Switch-Blade"
	app.Usage = "Correct and Efficient FASTQ trimming \n\n find more at: github.com/crmackay/switch-blade"
	app.Version = "0.0.1"
	app.Author = "Christopher R. MacKay"
	app.Email = "christopher.mackay@umassmed.edu"
	app.Flags = []cli.Flag{
//		cli.BoolFlag{
//		    Name: "config, c",
//		    Usage: "if using a config file instead of command line parameters //\n\t\t ~note: the 'config.txt' file must be in the \n\t\t current working  //directory",
		},
		cli.IntFlag{
			Name:  "numbercores, n",
			Value: 1,
			Usage: "the number of cores to run in parallel: \n\t\t ~the default value is 1",
		},
		cli.StringFlag{
			Name:  "input, i",
			Value: "",
			Usage: "input FASTQ file",
		},
		cli.StringFlag{
			Name:  "output, o",
			Usage: "output path and filename",
		},
		cli.StringFlag{
			Name:  "linker, l",
			Usage: "the linker sequence to be removed",
		},
		cli.IntFlag{
			Name:  "end, e",
			Value: 3,
			Usage: "select the which end of the read to remove the supplied linker: \n\t\t ~ the default is the 3' end \n\t\t ~ supply an integer: either 3 or 5",
		},
	}
	app.Action = func(c *cli.Context) {
		println("Hello friend!")
	}

    //return a map of arguments
        // NumCores: n
        // Input: i
        // Output: 0
        // Linker: l
        // End: e
	app.Run(os.Args)
}
*/
