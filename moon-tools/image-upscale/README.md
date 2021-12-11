Image-Upscale
===

Image Size를 키우는 Windows Application으로   
Walk란 "Windows Application Library Kit"를 이용하여    
Go언어를 통해 만든, Windows GUI Application이다.

공식 깃허브 주소는 다음과 같다.   
[walk Github 주소](https://github.com/lxn/walk)

## 필요한 Import
    go get github.com/lxn/walk
    go get github.com/akavel/rsrc

## build 명령어
    //mainfest 생성
    rsrc -manifest test.manifest -o rsrc.syso
    rsrc -manifest MoonUpScale.exe.manifest -o rsrc.syso

    //ico 옵션도 함께 주자
    rsrc -manifest MoonUpScale.exe.manifest -ico asset/icon/MoonUpScale.ico -o rsrc.syso

    // console창 안나오게 build 
    go build -ldflags="-H windowsgui"
    go build -ldflags="-H windowsgui" -o MoonUpScale.exe