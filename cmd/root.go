/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"crypto/sha512"
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var char_list string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_\"#$%&'*+,./:;=?!@\\^`|~[]{}()<>"
var length string
var passphrase string
var alias string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bytevault",
	Short: "Mathimagically Stores Passwords.",
	Long:  `Mathimagically Stores Passwords..`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(calcPass(passphrase, alias, length))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bytevault.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVarP(&length, "length", "l", "32", "Length of password")
	rootCmd.Flags().StringVarP(&passphrase, "passphrase", "p", " ", "Pass Phrase to unlock passwords")
	rootCmd.MarkFlagRequired("passphrase")
	rootCmd.Flags().StringVarP(&alias, "alias", "a", " ", "Pass Alias to unlock passwords")
	rootCmd.MarkFlagRequired("alias")
}

func calcPass(seedPhrase string, passAlias string, length string) string {
	hash := sha512.Sum512([]byte(seedPhrase + passAlias + length))
	var userPass string
	// Convert the hash bytes to an integer.
	var seed1 int64
	binary.Read(bytes.NewReader(hash[:]), binary.BigEndian, &seed1)
	rand.Seed(seed1)

	input_length_int, err := strconv.Atoi(length)
	if err != nil {
		fmt.Println("Error: not a number!: ", err)
		os.Exit(1)
	}

	for i := 0; i < input_length_int; i++ {
		//fmt.Printf("%c", char_list[rand.Intn(len(char_list))])
		userPass = userPass + string(char_list[rand.Intn(len(char_list))])
	}
	return userPass
}
