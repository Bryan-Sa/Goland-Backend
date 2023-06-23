FROM golang:latest
#ici on utilise la dernière version de Go

WORKDIR /app
#ici on définit le dossier de travail

COPY . .
#on copie les fichiers de notre app dans le conteneur

RUN go build -o main .
#compile l'application

EXPOSE 8080
#le port sur lequel écoute l'app

CMD ["./main"]
#Commande pour exec l'app