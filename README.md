# go-orm

REQUIREMENTS
1. go get gorm.io/driver/mysql: mysql driver
2. go get gorm.io/gorm: orm for go-mysql
3. go get github.com/gorilla/mux: for rest-api json
4. Mysql database with user "root", password "mysql", dbname "gorm"
    Table Person (id, name, address, account_id)
    Table Account (id, username, password)

DOCUMENTATION
1. https://gorm.io/id_ID/docs/index.html (go orm package)

FOLDER STRUCTURE
1. connection: contain logic to connect with mysql database
2. model: is filled with table mapping to go struct
3. service: directory where functions are implemented
4. mysql101: documentation of basic mysql CRUD
5. jsonRequest: json request mapping to go struct
6. jsonResponse: mapping json response to go struct

NOTES
1. gorm by default pluralize your code, see model folder for detail