const jsonData = require("./data.json")
//console.log(jsonData)

function inMeetingRoom(a){
  a.forEach(element => {
    if(element.placement.name == "Meeting Room"){
      console.log(element)
    }
  });
}

function electronic(a){
  a.forEach(element => {
    if(element.type == "electronic"){
      console.log(element)
    }
  });
}

function furniture(a){
  a.forEach(element => {
    if(element.type == "furniture"){
      console.log(element)
    }
  });
}

function purchase(a){
  const date = new Date("2020-01-16").getTime()/1000
  a.forEach(element => {
    if(element.purchased_at == date){
      console.log(element)
    }
  });
}

function brownColor(a){
  a.forEach(element => {
    if(element.tags.includes('brown')){
      console.log(element)
    }
  });
}

console.log("in Meeting room")
inMeetingRoom(jsonData)
console.log("")

console.log("electronic type")
electronic(jsonData)
console.log("")

console.log("furniture type")
furniture(jsonData)
console.log("")

console.log("purchased at")
purchase(jsonData)
console.log("")

console.log("brown color")
brownColor(jsonData)
console.log("")