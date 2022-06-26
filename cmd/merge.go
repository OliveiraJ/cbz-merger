/*
Copyright © 2022 JORDAN SILVA OLIVEIRA <JORDANSILVA102@GMAIL.COM>

*/
package cmd

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
)

// mergeCmd represents the merge command
var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge cbz files inside the same folders into one",
	Long: `Receives the path to the folder where the files are and the name of the final cbz file. Exemple:
	
			./cbz-merger merge "/home/user/.../cbzfiles" "Comic"

		The command above will merge the files inside the "/home/user/.../cbzfiles" directory into the Comic.cbz file
		created inside the "/home/user/.../cbzfiles" directory as well.
	`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Merging cbz files inside of " + args[0])

		var rootFolderPath = args[0]
		var destinyFolder = args[1]
		var directorys []string
		var pages []string
		var pagenumber = 0

		//Cimnhando pelos arquivos e pegando os nomes de cada página e pasta
		err := filepath.WalkDir(rootFolderPath, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				directorys = append(directorys, d.Name())

			} else {
				pages = append(pages, d.Name())
			}
			return nil
		})

		directorys = directorys[1:]
		if err != nil {
			fmt.Println(err)
		}
		//return the ammount of pages in the new .cbz file
		fmt.Println("The final cbz file will have: " + strconv.Itoa(len(pages)) + " pages")

		//create a destiny folder
		os.Mkdir(rootFolderPath+"/"+destinyFolder, 0755)

		//Copy files to the detiny folder and rename them to keep the right order
		for _, comicFolder := range directorys {
			err := filepath.WalkDir(rootFolderPath+"/"+comicFolder, func(path string, d os.DirEntry, err error) error {

				if err != nil {
					return err
				}
				if !d.IsDir() {
					pagenumber++

					originalPage, err := os.Open(rootFolderPath + "/" + comicFolder + "/" + d.Name())
					if err != nil {
						return err
					}
					defer originalPage.Close()

					if pagenumber < 10 {
						copyPage, err := os.Create(rootFolderPath + "/" + destinyFolder + "/" + "00" + strconv.Itoa(pagenumber) + ".jpg")
						if err != nil {
							return err
						}
						defer copyPage.Close()

						_, err = io.Copy(copyPage, originalPage)
						if err != nil {
							return err
						}
					} else if pagenumber < 100 {
						copyPage, err := os.Create(rootFolderPath + "/" + destinyFolder + "/" + "0" + strconv.Itoa(pagenumber) + ".jpg")
						if err != nil {
							return err
						}
						defer copyPage.Close()

						_, err = io.Copy(copyPage, originalPage)
						if err != nil {
							return err
						}
					} else {
						copyPage, err := os.Create(rootFolderPath + "/" + destinyFolder + "/" + strconv.Itoa(pagenumber) + ".jpg")
						if err != nil {
							return err
						}
						defer copyPage.Close()

						_, err = io.Copy(copyPage, originalPage)
						if err != nil {
							return err
						}
					}
				}
				return nil
			})
			if err != nil {
				fmt.Println(err)
			}
		}

		//Compress folder into a .cbz
		finalFile, err := os.Create(rootFolderPath + "/" + destinyFolder + ".zip")
		if err != nil {
			panic(err)
		}
		defer finalFile.Close()

		renamedFiles, err := ioutil.ReadDir(rootFolderPath + "/" + destinyFolder)
		if err != nil {
			panic(err)
		}

		zipWriter := zip.NewWriter(finalFile)

		for _, file := range renamedFiles {

			f, err := os.Open(rootFolderPath + "/" + destinyFolder + "/" + file.Name())
			if err != nil {
				panic(err)
			}
			defer f.Close()

			w, err := zipWriter.Create(destinyFolder + "/" + file.Name())
			if err != nil {
				panic(err)
			}
			if _, err := io.Copy(w, f); err != nil {
				panic(err)
			}
		}
		zipWriter.Close()

		err = os.Rename(rootFolderPath+"/"+destinyFolder+".zip", rootFolderPath+"/"+destinyFolder+".cbz")
		if err != nil {
			panic(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(mergeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mergeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mergeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
