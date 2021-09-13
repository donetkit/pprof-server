package main

func main() {
	config := Config{
		Host:     "",
		Port:     8081,
		Disabled: false,
	}
	RunServer(&config)
	for {

	}
}
