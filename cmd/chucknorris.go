/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

type ChuckNorris struct {
	Joke string `json:"value"`
}

// chucknorrisCmd represents the chucknorris command
var chucknorrisCmd = &cobra.Command{
	Use:   "chucknorris",
	Short: "Get a random chuck joke",
	Long: `Chuck Norris facts are satirical factoids about martial artist 
	and actor Chuck Norris that have become an Internet phenomenon and as a 
	result have become widespread in popular culture. The 'facts' are normally 
	absurd hyperbolic claims about Norris' toughness, attitude, virility, 
	sophistication, and masculinity.`,

	Run: func(cmd *cobra.Command, args []string) {

		URL := "https://api.chucknorris.io/jokes/random"

		response, err := http.Get(URL)

		if err != nil {
			fmt.Println(err)
			return
		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			byte_body, err := io.ReadAll(response.Body)

			if err != nil {
				fmt.Println(err)
				return
			}

			var chucknorris ChuckNorris
			json.Unmarshal([]byte(byte_body), &chucknorris)

			fmt.Println(chucknorris.Joke)
		}

	},
}

func init() {
	rootCmd.AddCommand(chucknorrisCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chucknorrisCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chucknorrisCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
