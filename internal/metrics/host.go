package metrics

func GetSystemInfo(provider StatProvider) (*SystemInfoResult, error) {
	info, err := provider.Info()
	if err != nil {
		return nil, err
	}

	return &SystemInfoResult{
		Info: info,
	}, nil
}
