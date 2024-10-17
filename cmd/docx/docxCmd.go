/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package docx

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
)

// timefromzoneCmd represents the timefromzone command
var DocxCmd = &cobra.Command{
	Use:   "docx",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		Path, err := cmd.Flags().GetString("path")
		if err != nil {
			log.Fatal(err)
		}
		docx, err := zip.OpenReader(Path)
		if err != nil {
			log.Fatal(err)
		}
		defer docx.Close()

		var res string
		if dir, err := cmd.Flags().GetString("output"); err == nil {
			res = path.Join(dir, "res")
		} else {
			res = "res"
		}

		os.MkdirAll(res, os.FileMode(0644))

		for _, file := range docx.File {

			dir, fileName := path.Split(file.Name)

			if match, _ := path.Match("*/media/*", dir); match {

				res, err := os.Create(path.Join(res, fileName))
				if err != nil {
					log.Fatal(err)
				}

				r, err := file.Open()
				if err != nil {
					log.Fatal(err)
				}

				io.Copy(res, r)

				continue
			}

			if match, _ := path.Match("document.xml", fileName); match {

				res, err := os.Create(path.Join(res, strings.Split(fileName, ".")[0]+".txt"))
				if err != nil {
					log.Fatal(err)
				}

				r, err := file.Open()
				if err != nil {
					log.Fatal(err)
				}

				res.Write(NewByteWRs(r))

				continue
			}
		}
	},
}

func init() {

	DocxCmd.Flags().StringP("path", "p", "", "Help message for path (required flags)")
	DocxCmd.MarkFlagRequired("path")

	DocxCmd.Flags().StringP("output", "o", ".", "Help message for output")
}
