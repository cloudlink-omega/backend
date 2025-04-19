package flags

// Features

var SupportsAchievements map[string]string = map[string]string{"achievements": "This game supports earning achievements."}
var CanMakeChangesToAccount map[string]string = map[string]string{"config": "This game can make changes to your account."}
var SupportsControllers map[string]string = map[string]string{"controllers": "This game supports controllers."}
var HasDownloadableContent map[string]string = map[string]string{"dlc": "This game supports Downloadable Content."}
var SupportsLegacyProtocols map[string]string = map[string]string{"legacy": "This game supports legacy CloudLink clients."}
var HasMatchmaking map[string]string = map[string]string{"matchmaking": "This game supports matchmaking."}
var SupportsSaveData map[string]string = map[string]string{"save": "This game supports cloud save data."}
var SupportsPoints map[string]string = map[string]string{"points": "This game can earn, spend, trade or redeem points."}

// Age Ratings

var SuitableForAllAges map[string]string = map[string]string{"everyone": "This game is suitable for everyone."}
var ForAdultsOnly map[string]string = map[string]string{"mature": "This game is only for adult audiences."}
var TeensAndOlder map[string]string = map[string]string{"older": "This game is only for teens and older audiences."}

// Platform support

var GameIsMobileOnly map[string]string = map[string]string{"mobile": "This game is only available on mobile devices."}
var GameIsForAllDevices map[string]string = map[string]string{"multidev": "This game can be played on mobile or desktop devices."}

// Source Code

var GameIsOpenSource map[string]string = map[string]string{"oss": "The game is open source."}
var GameIsProprietary map[string]string = map[string]string{"proprietary": "The source code of this game is proprietary."}

// Development Platforms

var MadeWithTurbowarp map[string]string = map[string]string{"ontw": "This game was made using Turbowarp."}
var MadeWithPenguinMod map[string]string = map[string]string{"onpm": "This game was made using PenguinMod."}
var MadeWithSheeptesterMod map[string]string = map[string]string{"oneq": "This game was made using E羊icques (SheepTester's Mod)."}
var OriginalOnScratch map[string]string = map[string]string{"onscratch": "This game is also available on Scratch."}

// Advisories

var ContainsViolence map[string]string = map[string]string{"violent": "This game contains or references violent content."}
var ContainsSubstances map[string]string = map[string]string{"substances": "This game contains or references drugs, alcohol or weapons."}

// Status

var UnderReview map[string]string = map[string]string{"review": "This game is undergoing review or awaiting approval by an administrator."}

// Extra Connectivity

var SupportsBasicVoiceChat map[string]string = map[string]string{"call": "This game supports voice chat."}
var SupportsMessaging map[string]string = map[string]string{"mail": "This game can send and receive messages using your account."}
var SupportsProximityChat map[string]string = map[string]string{"vchat": "This game supports proximity voice chat."}
var SupportsVoicemail map[string]string = map[string]string{"vmail": "This game supports sending or receiving voicemail."}
