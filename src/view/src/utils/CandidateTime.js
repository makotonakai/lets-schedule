export function AddNewElement(list){
  list.push("")
}

export function DeleteLastElement(list){
  list.pop()
}

export function ConvertDateTimeToISO(time){

  // 時間が入力されければ空文字列を返す
  if (IsDateTimeNotDefined(time)) {
    return ""
  }

  let offset = time.getTimezoneOffset()
  let offsetAbs = Math.abs(offset)
  let isoString = new Date(time.getTime() - offset * 60 * 1000).toISOString()
  return `${isoString.slice(0, -1)}${offset > 0 ? '-' : '+'}${String(Math.floor(offsetAbs / 60)).padStart(2, '0')}:${String(offsetAbs % 60).padStart(2, '0')}`
}


export function ConvertDateTimeListToISO(DateTimeList) {

  let DateTimeListToISO = []

  if (IsDateTimeListNotDefined(DateTimeList)) {
    return []
  }

  for(let i = 0; i < DateTimeList.length; i++){
    if (DateTimeList[i] == "") {
      DateTimeListToISO.push(["0001-01-01 00:00:00 +0000 UTC", "0001-01-01 00:00:00 +0000 UTC"]);
    } else {
      let startTime = ConvertDateTimeToISO(DateTimeList[i][0]);
      let endTime = ConvertDateTimeToISO(DateTimeList[i][1]);
      DateTimeListToISO.push([startTime, endTime]);
    }
  }
  return DateTimeListToISO
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
      CandidateTimeDict[UserName] = [CandidateTime]
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

export function CreateAvailableTimeList(responseData){
  for(let i = 0; i < responseData.length; i++){
    responseData[i]["start_time"] = ChangeCandidateTimeFormat(responseData[i]["start_time"])
    responseData[i]["end_time"] = ChangeCandidateTimeFormat(responseData[i]["end_time"])
  }
  return responseData
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

 
function IsDateTimeNotDefined(time) {
  return time == null
}

function IsDateTimeListNotDefined(DateTimeList) {
  return DateTimeList[0] == ""
}
