
# API Mini POS - Golang

## 1. Set Up Project

 - Clone repository
 - Import file `mini-pos.sql` ke database
 - Import file `golang mini pos.postman_collection.json` ke postman
 - Buat file `.env` sejajar dengan file `.env.example`
 - Sesuaikan konfigurasi yang terdapat pada file `.env.example`
 - Jalankan perintah
    ```
    go mod tidy
    ```

## 2. Run Project

 - Masuk ke folder
 - Jalankan perintah
    ```
    go run .server.go
    ```
- Server akan berjalan pada `http://localhost:9000`

## 3. ERD Mini POS 
![Imgur](https://i.postimg.cc/K8hLF352/Screen-Shot-2022-08-11-at-14-13-27.png)

## 5 API List
 - **User**
    - POST login
    - POST register
    - POST user/store
    - GET user
    - GET user/show/:id
    - PUT user/update/:id
    - DELETE user/delete/:id
 - **Outlet**
   - POST outlet/store
   - GET outlet
   - GET outlet/show/:id
   - PUT outlet/update/:id
   - DELETE outlet/delete/:id
 - **Category**
    - POST category/store
    - GET category
    - GET category/show/:id
    - PUT category/update/:id
    - DELETE category/delete/:id
 - **Product**
    - POST product/store
    - GET product
    - GET product/show/:id
    - PUT product/update:id
    - DELETE product/delete/:id
 - **Transaction**
   - POST transaction/store
   - GET transaction
   - GET transaction/show/:id
   - PUT transaction/update/:id
   - DELETE transaction/delete/:id
