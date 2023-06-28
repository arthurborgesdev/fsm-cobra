/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/looplab/fsm"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type Option struct {
	Name  string
	Event string
}

type PromptContent struct {
	ErrorMsg string
	Label    string
	Options  []Option
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init the FSM",
	Long:  "Init the FSM with options",
	Run: func(cmd *cobra.Command, args []string) {
		myFsm := fsm.NewFSM(
			"Menu",
			fsm.Events{
				{Name: "ChooseItem", Src: []string{"Menu"}, Dst: "Confirmation"},
				{Name: "Accept", Src: []string{"Confirmation"}, Dst: "Checkout"},
				{Name: "Reject", Src: []string{"Confirmation"}, Dst: "Menu"},
				{Name: "Pay", Src: []string{"Checkout"}, Dst: "Finish"},
				{Name: "Cancel", Src: []string{"Checkout"}, Dst: "Menu"},
			},
			fsm.Callbacks{},
		)

		for {
			menuContent := PromptContent{
				ErrorMsg: "Err - opção inválida",
				Label:    "Cardápio de hoje",
				Options: []Option{
					{
						Name:  "Coke",
						Event: "ChooseItem",
					},
					{
						Name:  "Pizza",
						Event: "ChooseItem",
					},
					{
						Name:  "Hamburger",
						Event: "ChooseItem",
					},
					{
						Name:  "Ice cream",
						Event: "ChooseItem",
					},
				},
			}
			menuItems := []string{"coke", "pizza", "hamburger", "icecream"}
			var foodName, event string
			foodName, event = promptGetSelect(menuContent, menuItems, myFsm)

			fmt.Println(myFsm.Current())

			confirmationContent := PromptContent{
				ErrorMsg: "Err - opção inválida",
				Label:    "Está certo disso?",
				Options: []Option{
					{
						Name:  "Sim",
						Event: "Accept",
					},
					{
						Name:  "Não",
						Event: "Reject",
					},
				},
			}
			confirmationOptions := []string{"Accept", "Reject"}
			_, event = promptGetSelect(confirmationContent, confirmationOptions, myFsm)
			if event == "Reject" {
				continue
			}

			fmt.Println(myFsm.Current())

			checkoutContent := PromptContent{
				ErrorMsg: "Err - opção inválida",
				Label:    "Deseja pagar agora?",
				Options: []Option{
					{
						Name:  "Sim",
						Event: "Pay",
					},
					{
						Name:  "Cancelar",
						Event: "Cancel",
					},
				},
			}
			checkoutOptions := []string{"Pay", "Cancel"}

			_, event = promptGetSelect(checkoutContent, checkoutOptions, myFsm)

			fmt.Println("Você comprou: " + foodName)

			if event == "Cancel" {
				continue
			}

			fmt.Println(myFsm.Current())

			fmt.Println("Pedido concluído!")
			break
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func promptGetSelect(pc PromptContent, items []string, myFsm *fsm.FSM) (string, string) {
	var selectedOption Option

	optionNames := make([]string, len(pc.Options))
	for i, option := range pc.Options {
		optionNames[i] = option.Name
	}

	for {
		prompt := promptui.Select{
			Label: pc.Label,
			Items: optionNames,
		}

		index, _, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			fmt.Printf("Let's try again...\n")
			continue
		}

		selectedOption = pc.Options[index]

		fmt.Printf("Input: %s\n", selectedOption.Name)

		err = myFsm.Event(context.Background(), selectedOption.Event)
		if err != nil {
			fmt.Println(err)
			continue
		}

		break
	}

	return selectedOption.Name, selectedOption.Event
}
