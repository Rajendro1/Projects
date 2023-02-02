let mysql = require('mysql');

let express = require('express');
let app = express();
let fs = require("fs");

app.get('/', function (req, res) {
    fs.readFile(__dirname + "/" + "package.json", 'utf8', function (err, data) {
        console.log(data);
        res.end(data);
    });
})
app.get('/users', GetUsersData)
app.get('/test', Test)


let server = app.listen(8082, function () {
    let host = server.address().address
    let port = server.address().port
    console.log("Example app listening at http://%s:%s", host, port)
})

let dbConn = mysql.createConnection({
host: 'localhost',
user: 'root',
password: 'Kush@789#',
database: 'MLD1'
})
dbConn.connect();

function GetUsersData(req, res){
    dbConn.query(`SELECT name FROM users WHERE user_id = 1`)
    console.log("Users")
    res.end()
}
function Test(req, res) {
    dbConn.query('SELECT user_id, name, email, birth_place, longitude, latitude, date_of_birth FROM users', function (error, results, fields) {
        if (error) throw error;
        return res.send({ error: false, data: results, message: 'users list.' });
    });
}