package main

import (
    "log"
    "github.com/gotk3/gotk3/gtk"
)

func main() {
    gtk.Init(nil)

    b, err := gtk.BuilderNew()
    if err != nil {
        log.Fatal("Unable to create window:", err)
    }

    err = b.AddFromFile("templates/main.glade")
    if err != nil {
        log.Fatal("Ошибка:", err)
    }

    // Получаем объект главного окна по ID
    obj, err := b.GetObject("window_main")
    if err != nil {
        log.Fatal("Ошибка:", err)
    }

    win := obj.(*gtk.Window)
    win.Connect("destroy", func() {
        gtk.MainQuit()
    })

    win.SetTitle("Simple Example")

    win.SetDefaultSize(800, 600)

    win.ShowAll()

    gtk.Main()
}