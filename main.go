package main

import (
    "fmt"
    "os"
    "time"

    "github.com/atotto/clipboard"
    "github.com/rivo/tview"
)

func main() {
    history := []string{}
    seen := map[string]bool{}
    var last string

    go func() {
        for {
            text, err := clipboard.ReadAll()
            if err == nil && text != "" && text != last {
                last = text
                if !seen[text] {
                    history = append([]string{text}, history...)
                    seen[text] = true
                }
            }
            time.Sleep(1000 * time.Millisecond)
        }
    }()

    app := tview.NewApplication()
    list := tview.NewList().
        ShowSecondaryText(false)

    list.SetBorder(true).SetTitle(" Clopy - Clipboard History ")

    go func() {
        for {
            app.QueueUpdateDraw(func() {
                list.Clear()
                for _, item := range history {
                    i := item
                    list.AddItem(i, "", 0, func() {
                        clipboard.WriteAll(i)
                    })
                }
            })
            time.Sleep(1 * time.Second)
        }
    }()

    if err := app.SetRoot(list, true).Run(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
