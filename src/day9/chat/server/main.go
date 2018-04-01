package main

func main()  {
	server := NewServer("0.0.0.0","7000")
	server.run()
}
