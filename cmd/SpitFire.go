/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/IcaroTARique/stress-cli/internal/api"
	"time"

	"github.com/spf13/cobra"
)

func newCreateCmd(request api.Request) *cobra.Command {
	return &cobra.Command{
		Use:   "SpitFire",
		Short: "SpitFire executes calls to the given URL",
		Long:  `SpitFire executes calls to given URLs using all elementes also informed by you`,
		RunE:  runCreate(request),
	}
}

func runCreate(request api.Request) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {

		url, _ := cmd.Flags().GetString("url")
		jobs, _ := cmd.Flags().GetInt("jobs")
		workers, _ := cmd.Flags().GetInt("workers")
		verbose, _ := cmd.Flags().GetBool("verbose")

		request.SetUrl(url)
		request.SetJobs(jobs)
		request.SetWorkers(workers)

		done := make(chan struct{})
		go showLoading(done)

		res := request.GoRequest(verbose)

		close(done)

		fmt.Print("\r                    \r")
		fmt.Println("Duration: ", res.TotalTime)
		fmt.Println("Amount of requests: ", res.ReqAmmount)
		fmt.Println("Code \t Ammount")
		for k, v := range res.Responses {
			fmt.Println(k, "\t", v)
		}

		return nil
	}
}

func showLoading(done <-chan struct{}) {
	loadingChars := []rune{'|', '/', '-', '\\'}
	i := 0
	for {
		select {
		case <-done:
			fmt.Print("\r")
			return
		default:
			fmt.Printf("\rLoading... %c", loadingChars[i])
			i = (i + 1) % len(loadingChars)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func init() {
	createCmd := newCreateCmd(GetActionApi())
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("url", "u", "", "URL to be requested")
	createCmd.Flags().IntP("jobs", "j", 1, "Number of jobs to be executed")
	createCmd.Flags().IntP("workers", "w", 1, "Number of workers to be executed")
	createCmd.Flags().BoolP("verbose", "v", false, "Verbose mode")
	createCmd.MarkFlagRequired("url")

}
