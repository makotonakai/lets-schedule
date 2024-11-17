package config

import (
	"errors"
)

var (
	EmailAddressIsEmpty = "メールアドレスが入力されていません"
	UserNameIsEmpty = "ユーザー名が入力されていません"
	PasswordIsEmpty = "パスワードが入力されていません"
	UserAlreadyExists = "ユーザーがすでに登録されています"
	LoginFailed = "ログインに失敗しました"
	AvailableTimeNotFound = "ミーティング可能な時間が見つかりません"
	AvailableTimeTooLong = "入力されたミーティング時間が長すぎます"
	AvailableTimeOutOfLimit = "入力された時間帯はミーティング可能な時間帯の範囲外です"
	TitleIsEmpty = "タイトルが入力されていません"
	HourIsEmpty = "時間が入力されていません"
	MeetingNotFound = "ミーティングが見つかりません"
	PlaceIsNotSpecified = "開催場所が指定されていません"
	URLIsNotSpecified = "URLが指定されていません"
	NeitherPlaceOrURLIsSpecified = "開催場所もURLも指定されていません"
	CandidateTimeIsEmpty = "空のミーティング候補時間があります"
	CandidateTimeIsPast = "過去のミーティング候補時間があります"
	ErrIdConversionFailed = errors.New("The given id is failed to be converted to an integer")
	ErrUserIsNil = errors.New("The pointer of the given User object is nil")
	ErrListOfErrorsDoesntExist = errors.New("The list of error messages doesn't exist")
	ErrEmailAddressIsEmpty = errors.New("The given email address is empty")
	ErrEmailAddressNotFound = errors.New("Email address not found")
	ErrUserNotFound = errors.New("The given user object doesn't exist")
	ErrUserNameNotFound = errors.New("Username not found")
	ErrUserWithUserNameNotFound = errors.New("User with the given username does not exist")
	ErrUserWithUserIdNotFound = errors.New("User with the given user id does not exist")
	ErrUserIdWithEmailAddressNotFound = errors.New("User id with the given email address does not exist")
	ErrMeetingNotFound = errors.New("The given Meeting object doesn't exist")
	ErrMeetingHourIsNegative = errors.New("The meeting hour needs to be more than 0")
	ErrRecordNotFound = errors.New("record not found")
	ErrArrayIsNil = errors.New("The given array is nil")
	ErrArrayIsEmpty = errors.New("The given array is empty")
	ErrCandidateTimeNotFound = errors.New("The given CandidateTime object doesn't exist")
	ErrAvailableTimeIsNil = errors.New("The given AvailableTime object is nil")
	ErrParticipantIsNil = errors.New("The given Participant object is nil")
	ErrParticipantWithUserNameIsNil = errors.New("The given ParticipantWithUserName object is nil")
	ErrParticipantListIsEmpty = errors.New("The given Participant list is empty")
	ErrParticipantListIsNil = errors.New("The given Participant list is nil")
	ErrParticipantWithUserNameListIsEmpty = errors.New("The given ParticipantWithUserName list is empty")
	ErrParticipantWithUserNameListIsNil = errors.New("The given ParticipantWithUserName list is nil")
	ErrParticipantWithUserNameListFailedToRegister = errors.New("Failed to register new participantWithUserName list")
	ErrIntegerIsNil = errors.New("The given integer is nil")
)
