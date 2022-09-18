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

export function CreateCandidateTimeDict(CandidateTimeJSONList){
  let CandidateTimeDict = {}
  for (let i = 0; i < CandidateTimeJSONList.length; i++){
    let CandidateTimeJSON = CandidateTimeJSONList[i]
    let UserName = CandidateTimeJSON["user_name"]
    let StartTime =  ChangeCandidateTimeFormat(CandidateTimeJSON["start_time"])
    let EndTime =  ChangeCandidateTimeFormat(CandidateTimeJSON["end_time"])
    let CandidateTime = {
        "user_name": UserName,
        "start_time": StartTime,
        "end_time": EndTime
      }
    if (UserName in CandidateTimeDict) {
      CandidateTimeDict[UserName].push(CandidateTime)
    }else{
      CandidateTimeDict[UserName] = CandidateTime
    }
  }
  return CandidateTimeDict
}

export function ChangeCandidateTimeFormat(CandidateTime){
  let [Date, NotDate] = CandidateTime.split("T")
  let Hour = NotDate.split("+")[0]
  return Date + " " + Hour
}
