package util

import (
	"testing"
)

func TestCreateNIM(t *testing.T) {
	t.Log(NIMAndEmailGenerator("sammi Aldhi Yanto",1, 2020))
	t.Log(NIMAndEmailGenerator("Rahmatul Izzah Annisa",1, 2018))
	t.Log(NIMAndEmailGenerator("Lean suryani", 2, 2019))
	t.Log(NIMAndEmailGenerator("muhammad kurniawan",3, 2017))
	t.Log(NIMAndEmailGenerator("adit",3,  2016))
}
