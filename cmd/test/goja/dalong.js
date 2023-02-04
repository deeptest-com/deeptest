module.exports = function printInfo(name,password){
    // call golang method
    var token  = login(name,password)
    return token
}