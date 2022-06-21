package software

import (
	"fmt"

	"github.com/amidaware/rmmagent/agent/utils"
	wapi "github.com/iamacarpet/go-win64api"
)

func GetInstalledSoftware() ([]Software, error) {
	ret := make([]Software, 0)
	sw, err := wapi.InstalledSoftwareList()
	if err != nil {
		return ret, err
	}

	for _, s := range sw {
		t := s.InstallDate
		ret = append(ret, Software{
			Name:        utils.CleanString(s.Name()),
			Version:     utils.CleanString(s.Version()),
			Publisher:   utils.CleanString(s.Publisher),
			InstallDate: fmt.Sprintf("%02d-%d-%02d", t.Year(), t.Month(), t.Day()),
			Size:        utils.ByteCountSI(s.EstimatedSize * 1024),
			Source:      utils.CleanString(s.InstallSource),
			Location:    utils.CleanString(s.InstallLocation),
			Uninstall:   utils.CleanString(s.UninstallString),
		})
	}

	return ret, nil
}