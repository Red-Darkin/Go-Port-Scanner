### gopscanner
A simple scan performed in go to quickly scan a port range on an IP or domain using **TCP SYN** scan

### Downloading repo
```
git clone https://github.com/Red-Darkin/Go-Port-Scanner.git
cd /Go-Port-Scanner
go build gopscanner.go
```
![imagen](https://github.com/Red-Darkin/Go-Port-Scanner/assets/62677201/a9ff59ed-e359-48ef-8c89-5ffd7d6899f8)


### Use Mode
```
./gopscanner -host <host> -p <ports> (e.g., 80,443 or 1-1000)
```
![imagen](https://github.com/Red-Darkin/Go-Port-Scanner/assets/62677201/5dd66e0b-3131-4efa-bb6b-fea03ad609c4)

you can also pass a file with a list of hosts
```
./gopscanner -f file -p <ports> (e.g., 80,443 or 1-1000)
```
![imagen](https://github.com/Red-Darkin/Go-Port-Scanner/assets/62677201/bc28b519-bf15-48af-baf2-7b321ac0b21f)


