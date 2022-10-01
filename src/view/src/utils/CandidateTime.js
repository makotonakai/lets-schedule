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

export function CreateCandidateTimeList(responseData) {
  let candidateTimeList = []
  for(let i = 0; i < responseData.length; i++){
    let startTime = new Date(responseData[i]["start_time"])
    let endTime = new Date(responseData[i]["end_time"])
    let candidateTime = [startTime, endTime]
    candidateTimeList.push(candidateTime)
  }
  return candidateTimeList
}

export function ChangeCandidateTimeFormat(CandidateTime){
  let [Date, NotDate] = CandidateTime.split("T")
  let Hour = NotDate.split("+")[0]
  return Date + " " + Hour
}

export function FindAvailableTime(CandidateTimeDict){

  let StartTime = null
  let EndTime = null

  let User1 = Object.keys(CandidateTimeDict)[0]
  let User2 = Object.keys(CandidateTimeDict)[1]

  let User1StartTime = new Date(CandidateTimeDict[User1]["start_time"])
  let User1EndTime = new Date(CandidateTimeDict[User1]["end_time"])

  let User2StartTime = new Date(CandidateTimeDict[User2]["start_time"])
  let User2EndTime = new Date(CandidateTimeDict[User2]["end_time"])

  if (User1StartTime <= User2EndTime && User2StartTime <= User1EndTime){
    StartTime = GetStringFromDate(User2StartTime)
    EndTime = GetStringFromDate(User1EndTime)
  }

  return [StartTime, EndTime]
}

//日付から文字列に変換する関数
function GetStringFromDate(date) {
 
  let year_str = date.getFullYear();
  //月だけ+1すること
  let month_str = 1 + date.getMonth();
  let day_str = date.getDate();
  let hour_str = date.getHours();
  let minute_str = date.getMinutes();
  let second_str = date.getSeconds();
  let format_str = 'YYYY-MM-DD hh:mm:ss';
  
  format_str = format_str.replace(/YYYY/g, year_str.toString().padStart(4, '0'));
  format_str = format_str.replace(/MM/g, month_str.toString().padStart(2, '0'));
  format_str = format_str.replace(/DD/g, day_str.toString().padStart(2, '0'));
  format_str = format_str.replace(/hh/g, hour_str.toString().padStart(2, '0'));
  format_str = format_str.replace(/mm/g, minute_str.toString().padStart(2, '0'));
  format_str = format_str.replace(/ss/g, second_str.toString().padStart(2, '0'));
  
  return format_str;
 };
