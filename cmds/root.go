// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmds

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"log"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "esdump",
	Short: "es import export",
	Long: `es import export `,
	// Uncomment the following line if your bare application
	// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
		PersistentPreRun:func(cmd *cobra.Command, args []string){
			log.SetOutput(os.Stderr)
			log.Println("binary version:",Version,"build Time:",BuildTime)
			log.Println("execute ",cmd.Use)
			timeStart=time.Now()
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			log.Printf("%s time spend %s",cmd.Use,time.Now().Sub(timeStart).String())
			return nil
		},
}
var timeStart time.Time

var EsUrl string
var IndexName string
func init() {
	log.SetOutput(os.Stderr)
	//log.Println("cobra root init")
	 RootCmd.PersistentFlags().StringVar(&EsUrl,"es","http://localhost:9200","es url")
	RootCmd.PersistentFlags().StringVar(&IndexName,"index","my_index","index name")
	RootCmd.MarkFlagRequired("es")
	RootCmd.MarkFlagRequired("index")
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print version",
	Long:  `print version`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

