# fibo-api

A API for calculating Fibonacci Sequence.

#Deployment

*  Download a recent Go release from page https://golang.org/dl/ and install package.

Open the command line and type `go version` to verify installation.

* Create a Go work-space

Ensure Path under `System Variables` has the ~`C:\Go\bin` variable in it.

Then create a Go work-space. For example, Go installation files were saved

under the path `C:/Go` and work-space was created under `C:\Projects\Go`

* Create the GOPATH environment variable

Open `Control Panel` and navigate to `System` and then `Environment Variables`.

Click `New`.
 
Next to Variable Name, enter `GOPATH`, and next to `Variable Value` enter `C:\Projects\Go` 

After all type `echo %GOPATH%` in terminal to verify installation.

* Clone the repository  

Open the directory `C:\Projects\Go\src` and type 

`git clone https://github.com/AbstractBreazy/fibo-api.git` in local terminal.

* Final steps 

Open the directory `C:\Projects\Go\src\fibo-api` and type `go install` to install the packages named by the import paths.

Now is the time to verify that all is working correctly by typing `go build`.

After these steps just type `fibo-api.exe` in terminal.

If the installation was successful, you should get the following message:

`Start listening http at port 8090...`

        




