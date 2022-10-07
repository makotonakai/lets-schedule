export function CreateParticipantJSONList(Host, ParticipantList, MeetingId) {

  let ParticipantJSONList = []

  let HostJSON = {
    "user_name": Host,
    "meeting_id": MeetingId,
    "is_host": true,
    "has_responded": true
  }
  ParticipantJSONList.push(HostJSON)

  for(let i = 0; i < ParticipantList.length; i++){
    let ParticipantJSON = {
      "user_name": ParticipantList[i],
      "meeting_id": MeetingId,
      "is_host": false,
      "has_responded": false
    }
    ParticipantJSONList.push(ParticipantJSON);
  }

  return ParticipantJSONList

}

export function GetHost(allParticipantList) {
  for (let i = 0; i < allParticipantList.length; i++) {
    let participant = allParticipantList[i]
    if(participant["is_host"]){
      return participant["user_name"]
    }
  }
}

export function GetParticipantList(allParticipantList) {
  let participantList = []
  for (let i = 0; i < allParticipantList.length; i++) {
    let participant = allParticipantList[i]
    if(participant["is_host"] == false){
      let username = participant["user_name"]
      participantList.push(username)
    }
  }
  return participantList
}
