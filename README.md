// Create project with go


go help mod init
go mod init github.com/calculadora
go mod tidy

Compilar
go build main.go

Para ejecutarlo
./main.exe

Para ejecutarlo desde cualquie parte
go install



# Proyecto Calculadora

## git inicializacion del repositorio
```
git init 
git remote add origin https://github.com/luisfer1001/operaciones-aritmeticas.git
git branch -M main
git add .
git commit -m "firs commit"
git push -u origin main