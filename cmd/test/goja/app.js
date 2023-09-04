var dalong  = require("./dalong.js")

var users = [
      {
          name:"dalong",
          age:333
      },
      {
        name:"rong",
        age:22
      },
     {
        name:"mydemo",
        age:44
     },
     {
      name: shortid.generate(),
      age:80
   }
]

var evens = _.filter(users, function(user){ return user.age % 2 == 0; });
module.exports = {
    version:"v1",
    type_info:"system",
    token: dalong("dalong","demoapp"),
    filteruser: JSON.stringify(evens),
    id:shortid.generate(),
    id2:shortid.generate()
}