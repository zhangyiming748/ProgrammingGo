package CH4

import "testing"
var (
	iniData = []string{
		"; Cut down copy of Mozilla application.ini file",
		"",
		"[App]",
		"Vendor=Mozilla",
		"Name=Iceweasel",
		"Profile=mozilla/firefox",
		"Version=3.5.16",
		"[Gecko]",
		"MinVersion=1.9.1",
		"MaxVersion=1.9.1.*",
		"[XRE]",
		"EnableProfileMigrator=0",
		"EnableExtensionManager=1",
	}
)
func TestMaster(t *testing.T) {
	master()
}
func TestParseIni(t *testing.T) {
	ret:=ParseIni(iniData)
	t.Log(ret)
}