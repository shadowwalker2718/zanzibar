-- wrk -t12 -c400 -d30s -s ./benchmarks/contacts_1KB.lua http://localhost:8093/contacts/foo/contacts
-- go-torch -u http://localhost:8093/ -t5
wrk.method = "POST"
wrk.body = "{\"contacts\":[{\"fragments\":[{\"type\":\"message\",\"text\":\"foobarbaz\"}],\"attributes\":{\"firstName\":\"steve\",\"lastName\":\"stevenson\",\"hasPhoto\":true,\"numFields\":10,\"timesContacted\":5,\"lastTimeContacted\":0,\"isStarred\":false,\"hasCustomRingtone\":false,\"isSendToVoicemail\":false,\"hasThumbnail\":false,\"namePrefix\":\"\",\"nameSuffix\":\"\"}},{\"fragments\":[{\"type\":\"message\",\"text\":\"foobarbaz\"}],\"attributes\":{\"firstName\":\"steve\",\"lastName\":\"stevenson\",\"hasPhoto\":true,\"numFields\":10,\"timesContacted\":5,\"lastTimeContacted\":0,\"isStarred\":false,\"hasCustomRingtone\":false,\"isSendToVoicemail\":false,\"hasThumbnail\":false,\"namePrefix\":\"\",\"nameSuffix\":\"\"}},{\"fragments\":[],\"attributes\":{\"firstName\":\"steve\",\"lastName\":\"stevenson\",\"hasPhoto\":true,\"numFields\":10,\"timesContacted\":5,\"lastTimeContacted\":0,\"isStarred\":false,\"hasCustomRingtone\":false,\"isSendToVoicemail\":false,\"hasThumbnail\":false,\"namePrefix\":\"\",\"nameSuffix\":\"\"}},{\"fragments\":[],\"attributes\":{\"firstName\":\"steve\",\"lastName\":\"stevenson\",\"hasPhoto\":true,\"numFields\":10,\"timesContacted\":5,\"lastTimeContacted\":0,\"isStarred\":false,\"hasCustomRingtone\":false,\"isSendToVoicemail\":false,\"hasThumbnail\":false,\"namePrefix\":\"\",\"nameSuffix\":\"\"}}],\"appType\":\"MY_APP\"}"

