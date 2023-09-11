git add .
git commit -m "Ultimo Commit"
git push
go build main.go 
rm -f main.zip
tar.exe -a -cf main.zip main 