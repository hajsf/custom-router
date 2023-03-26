Open a PowerShell window and navigate to the directory where you saved the post-request.ps1 file.

Execute the post-request.ps1 file by typing `.\post-request.ps1` and pressing Enter.

This will execute the PowerShell commands in the post-request.ps1 file and send a POST request with parameters to the server running on localhost at port 8080.

You don’t need to have administrator privileges on your local machine to execute a PowerShell script file. However, if you haven’t previously changed the execution policy on your machine, you may get an error message saying that the execution of scripts is disabled on your system. In this case, you can temporarily change the execution policy to allow the execution of scripts by running the following command in your PowerShell window:

`Set-ExecutionPolicy -Scope Process -ExecutionPolicy Bypass`
This command changes the execution policy for the current PowerShell process only and allows you to execute script files. After running this command, you can execute the post-request.ps1 file as described earlier.