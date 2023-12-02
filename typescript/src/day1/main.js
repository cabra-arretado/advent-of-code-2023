"use strict";
exports.__esModule = true;
var fs = require("fs");
fs.readFile("src/day1/test.txt", "utf-8", function (err, data) {
    if (err) {
        console.error(err);
        return;
    }
    console.log(data);
});
