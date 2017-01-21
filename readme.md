## Kuriboh
download cartoon from [niceoppai.net](http://www.niceoppai.net) with cli
------

#### How to install

1. install golang
2. git clone this project

#### Usage
```
--path value, -p value            cartoon name path
--startchapter value, --sc value  number start chapter that want to download
--endchapter value, --ec value    number end chapter that want to download
--destination value, -d value     output destination path
```


*** http://www.niceoppai.net/Criminale/30/ Criminale is cartoon name

#### Example 
```
    go run main.go -p Criminale -sc 3 -ec 5 -d ./mycartoon
```

start download Criminal chapter 3 to 5 save to folder ./mycartoon

