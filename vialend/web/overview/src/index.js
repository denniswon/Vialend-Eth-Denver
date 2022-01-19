/*
index.js: webpack entry
develop:
webpack ./src/index.js -o ./build/dev --mode=development
production:
webpack ./src/index.js -o ./build/ --mode=production
*/
function hello(){
 console.log("hello world");
}