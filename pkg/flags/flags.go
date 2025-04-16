package flags

// Features

var SupportsAchievements map[string]any = map[string]any{"achievements": "This game supports earning achievements."}
var CanMakeChangesToAccount map[string]any = map[string]any{"config": "This game can make changes to your account."}
var SupportsControllers map[string]any = map[string]any{"controllers": "This game supports controllers."}
var HasDownloadableContent map[string]any = map[string]any{"dlc": "This game supports Downloadable Content."}
var SupportsLegacyProtocols map[string]any = map[string]any{"legacy": "This game supports legacy CloudLink clients."}
var HasMatchmaking map[string]any = map[string]any{"matchmaking": "This game supports matchmaking."}
var SupportsSaveData map[string]any = map[string]any{"save": "This game supports cloud save data."}
var SupportsPoints map[string]any = map[string]any{"points": "This game can earn, spend, trade or redeem points."}

// Age Ratings

var SuitableForAllAges map[string]any = map[string]any{"everyone": "This game is suitable for everyone."}
var ForAdultsOnly map[string]any = map[string]any{"mature": "This game is only for adult audiences."}
var TeensAndOlder map[string]any = map[string]any{"older": "This game is only for teens and older audiences."}

// Platform support

var GameIsMobileOnly map[string]any = map[string]any{"mobile": "This game is only available on mobile devices."}
var GameIsForAllDevices map[string]any = map[string]any{"multidev": "This game can be played on mobile or desktop devices."}

// Source Code

var GameIsOpenSource map[string]any = map[string]any{"oss": "The game is open source."}
var GameIsProprietary map[string]any = map[string]any{"proprietary": "The source code of this game is proprietary."}

// Development Platforms

var MadeWithTurbowarp map[string]any = map[string]any{"ontw": "This game was made using Turbowarp."}
var MadeWithPenguinMod map[string]any = map[string]any{"onpm": "This game was made using PenguinMod."}
var MadeWithSheeptesterMod map[string]any = map[string]any{"oneq": "This game was made using E羊icques (SheepTester's Mod)."}
var OriginalOnScratch map[string]any = map[string]any{"onscratch": "This game is also available on Scratch."}

// Advisories

var ContainsViolence map[string]any = map[string]any{"violent": "This game contains or references violent content."}
var ContainsSubstances map[string]any = map[string]any{"substances": "This game contains or references drugs, alcohol or weapons."}

// Status

var UnderReview map[string]any = map[string]any{"review": "This game is undergoing review or awaiting approval by an administrator."}

// Extra Connectivity

var SupportsBasicVoiceChat map[string]any = map[string]any{"call": "This game supports voice chat."}
var SupportsMessaging map[string]any = map[string]any{"mail": "This game can send and receive messages using your account."}
var SupportsProximityChat map[string]any = map[string]any{"vchat": "This game supports proximity voice chat."}
var SupportsVoicemail map[string]any = map[string]any{"vmail": "This game supports sending or receiving voicemail."}
