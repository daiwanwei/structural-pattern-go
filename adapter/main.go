package main

func main() {
	display := DellDisplay{}
	mac := Mac{}
	adapter := MacAdapter{mac}
	display.Connect(&adapter)
}
