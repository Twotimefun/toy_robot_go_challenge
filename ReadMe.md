1. 
Clone the repo, or download the zip.

2. 
Run this command in the CLI: go mod init parentFolderName. 
For example: 'go mod init toy_robot_go_challenge'

3.
Ensure the imports are correct.
Example: Since I ran, 'toy_robot_go_challenge' the parent module will be 'toy_robot_go_challenge/folder'.

4.
Run the command: go build
This will build the application into an .exe

5.
Run in the CLI: ./file.exe
Where 'file'.exe is the name of your .exe created by using go build.

5.1
While there are examples of acceptable commands provided during runtime, a list of commands will be provided here:
- 'PLACE X,Y,F' where X and Y denote its placement on said axes, and F denotes the direction the robot is facing. These directions can be one of the following: NORTH, SOUTH, EAST, WEST.
- MOVE which moves the robot 1 unit in the current direction its facing
- LEFT which translates the direction the robot is facing by 90 degrees counterclockwise
- RIGHT which translates the direction the robot is facing by 90 degrees clockwise
- REPORT which prints the current position of the robot.
- EXIT which exits the program.

5.2
Side Notes:
- Before any other commands are run, PLACE or QUIT are the only avaliable commands. The robot must exist before its behaviour is changed.
- The inputs are case sensitive to capitalization.