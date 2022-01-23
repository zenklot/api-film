# RESTful API LIST FILM
Ini adalah projek sederhana yang terdiri dari Frontend & Backend beserta dengan Dokumentasi API.
Untuk Frontend menggunakan VUE.JS V3 dan Untuk Backend Menggunakan Go-language

## Clone Projek

```
$ git clone https://github.com/zenklot/api-film.git
$ cd api-film/
```

## Project setup Frontend
```
cd ./fe-film/
```
### Install Dependensi yang dibutuhkan dengan perintah
```
npm install
```
### Jalankan Server Frontend
```
npm run serve
```
### Buka IP Yang ditampilkan di Browser
```
 DONE  Compiled successfully in 43206ms


  App running at:
  - Local:   http://localhost:8080/ 
  - Network: http://192.168.59.221:8080/

  Note that the development build is not optimized.
  To create a production build, run npm run build.


```


## Project setup Backend
```
cd ./be-film/
```
### Build projek dengan perintah
```
go build
```
### Jalankan Server Backend (Linux)
```
./api-film
```
### Jalankan Server Backend (Window)
```
jalankan file `api-film.exe`
```
### Atau Jalankan Server Backend
```
go run main.go
```
### Server Berjalan
Server berjalan di localhost dengan port : 3000

---
## Dokumentasi API
Dokumentasi API dapat dilihat pada file `api-doc.json` atau pada link : 
https://app.swaggerhub.com/apis/zenklot/film-restful_api/1.0
---
# ERROR CORS
Jika terdapat error CORS solusinya install extension google chrome pada link https://chrome.google.com/webstore/detail/allow-cors-access-control/lhobafahddgcelffkeicbaginigeejlf
kemudian aktifkan CORS
