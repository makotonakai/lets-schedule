export function GetMeetingType(type){
  if(type == "現地開催") {
    return "physical"
  }else if(type == "オンライン"){
    return "virtual"
  }else{
    return "hybrid"
  }
}
