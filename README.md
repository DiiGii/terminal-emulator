# terminal-emulator
A lightweight UNIX-like terminal emulator built in Go. 

## How to run
After making sure the necessary dependencies are installed, you can simply run: 
```go
go run main.go
```

## The UI
The UI is a popup that looks pretty similar to the terminal popup. To implement it, I used the Fyne UI toolkit, which is pretty well documented. 

## Pseudoterminal
Next, we create a pseudoterminal that interacts with the PTY master and the PTY slave. 
![terminal_diagram](https://github.com/user-attachments/assets/6cdcc87b-1413-4d40-b760-89cfe07005e2)

The PTY master receives input (like key presses) and sends them to the PTY slave, which communicates with bash. Finally, the PTY slave sends the output from bash (like command results or error messages) back to the PTY master, which then sends it to the terminal emulator for display to the user.

## The End Result
<img width="571" alt="image" src="https://github.com/user-attachments/assets/1495f75c-3095-4ff3-bdc7-331bf3016f06" />

