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
<img width="774" alt="Screenshot 2024-12-25 at 10 19 55 PM" src="https://github.com/user-attachments/assets/57e5f796-f04d-445b-a639-b36fe04eded0" />


The PTY master receives input (like key presses) and sends them to the PTY slave, which communicates with bash. Finally, the PTY slave sends the output from bash (like command results or error messages) back to the PTY master, which then sends it to the terminal emulator for display to the user.

## The End Result
<img width="571" alt="image" src="https://github.com/user-attachments/assets/1495f75c-3095-4ff3-bdc7-331bf3016f06" />

