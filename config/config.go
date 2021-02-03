package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/fsnotify/fsnotify"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {}

var Conf Config

func HotLoadConfig(path string) error {
	if err := NewWatcher(path); err != nil {
		return err
	}
	if err := LoadConfig(path); err != nil {
		return err
	}
	return nil
}

func LoadConfig(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &Conf)
	if err != nil {
		return err
	}
	return nil
}

func NewWatcher(path string) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Errorf("newWatcher failed, error: %s", err)
	}
	go func() {
		for {
			select {
			case event :=  <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Remove == fsnotify.Remove {
					if err := LoadConfig(event.Name); err != nil {
						log.Println("hot load config error: ", err)
					}
				}

				if event.Op&fsnotify.Remove == fsnotify.Remove {
					err := watcher.Add(event.Name)
					if err != nil {
						log.Printf("remove file event:%s, err:%s", event, err)
					}
				}
			case err := <-watcher.Errors:
				if err != nil {
					log.Printf("watch error:%s", err)
				}

			}
		}
	}()

	err = watcher.Add(path)
	if err != nil {
		log.Printf("add watcher failed, error:%s", err)
	}
	return nil
}