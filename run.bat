set "BOT_TOKEN=5945815849:AAEz1DnAUCW4G2DGKacwdHvy602T-7l8cZ4" 
set "GROUP_ID=-1001923600726" 

echo "$BOT_TOKEN $GROUP_ID"
go build -o .\build\main.exe  .\cmd\main.go
.\build\main.exe
