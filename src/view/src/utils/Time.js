export function AddNewElement(list){
  list.push("")
}

export function DeleteLastElement(list){
  list.pop()
}

export function ConvertDateTimeToISO(time){
  let offset = time.getTimezoneOffset()
  let offsetAbs = Math.abs(offset)
  let isoString = new Date(time.getTime() - offset * 60 * 1000).toISOString()
  return `${isoString.slice(0, -1)}${offset > 0 ? '-' : '+'}${String(Math.floor(offsetAbs / 60)).padStart(2, '0')}:${String(offsetAbs % 60).padStart(2, '0')}`
}

export function ConvertDateTimeListToISO(DateTimeList) {
  for(let i = 0; i < DateTimeList.length; i++){
    DateTimeList[i][0] = ConvertDateTimeToISO(DateTimeList[i][0]);
    DateTimeList[i][1] = ConvertDateTimeToISO(DateTimeList[i][1]);
  }
  return DateTimeList
}

export function CreateDateTimeJSONList(DateTimeList, UserId, MeetingId){
  
  let DateTimeJSONList = []
  let DateTimeISOList = ConvertDateTimeListToISO(DateTimeList);
  
  for(let i = 0; i < DateTimeISOList.length; i++){
    let DateTimeJSON = {
      "user_id": UserId,
      "meeting_id": MeetingId,
      "start_time": DateTimeISOList[i][0],
      "end_time": DateTimeISOList[i][1]
    }
    DateTimeJSONList.push(DateTimeJSON);
  }
  return DateTimeJSONList
}
