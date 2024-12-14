package flags

// Features

var SupportsAchievements map[string]interface{} = map[string]interface{}{"achievements": "This game supports earning achievements."}
var CanMakeChangesToAccount map[string]interface{} = map[string]interface{}{"config": "This game can make changes to your account."}
var SupportsControllers map[string]interface{} = map[string]interface{}{"controllers": "This game supports controllers."}
var HasDownloadableContent map[string]interface{} = map[string]interface{}{"dlc": "This game supports Downloadable Content."}
var SupportsLegacyProtocols map[string]interface{} = map[string]interface{}{"legacy": "This game supports legacy CloudLink clients."}
var HasMatchmaking map[string]interface{} = map[string]interface{}{"matchmaking": "This game supports matchmaking."}
var SupportsSaveData map[string]interface{} = map[string]interface{}{"save": "This game supports cloud save data."}
var SupportsPoints map[string]interface{} = map[string]interface{}{"points": "This game can earn, spend, trade or redeem points."}

// Age Ratings

var SuitableForAllAges map[string]interface{} = map[string]interface{}{"everyone": "This game is suitable for everyone."}
var ForAdultsOnly map[string]interface{} = map[string]interface{}{"mature": "This game is only for adult audiences."}
var TeensAndOlder map[string]interface{} = map[string]interface{}{"older": "This game is only for teens and older audiences."}

// Platform support

var GameIsMobileOnly map[string]interface{} = map[string]interface{}{"mobile": "This game is only available on mobile devices."}
var GameIsForAllDevices map[string]interface{} = map[string]interface{}{"multidev": "This game can be played on mobile or desktop devices."}

// Source Code

var GameIsOpenSource map[string]interface{} = map[string]interface{}{"oss": "The game is open source."}
var GameIsProprietary map[string]interface{} = map[string]interface{}{"proprietary": "The source code of this game is proprietary."}

// Development Platforms

var MadeWithTurbowarp map[string]interface{} = map[string]interface{}{"ontw": "This game was made using Turbowarp."}
var MadeWithPenguinMod map[string]interface{} = map[string]interface{}{"onpm": "This game was made using PenguinMod."}
var MadeWithSheeptesterMod map[string]interface{} = map[string]interface{}{"oneq": "This game was made using E羊icques (SheepTester's Mod)."}
var OriginalOnScratch map[string]interface{} = map[string]interface{}{"onscratch": "This game is also available on Scratch."}

// Advisories

var ContainsViolence map[string]interface{} = map[string]interface{}{"violent": "This game contains or references violent content."}
var ContainsSubstances map[string]interface{} = map[string]interface{}{"substances": "This game contains or references drugs, alcohol or weapons."}

// Status

var UnderReview map[string]interface{} = map[string]interface{}{"review": "This game is undergoing review or awaiting approval by an administrator."}

// Extra Connectivity

var SupportsBasicVoiceChat map[string]interface{} = map[string]interface{}{"call": "This game supports voice chat."}
var SupportsMessaging map[string]interface{} = map[string]interface{}{"mail": "This game can send and receive messages using your account."}
var SupportsProximityChat map[string]interface{} = map[string]interface{}{"vchat": "This game supports proximity voice chat."}
var SupportsVoicemail map[string]interface{} = map[string]interface{}{"vmail": "This game supports sending or receiving voicemail."}
