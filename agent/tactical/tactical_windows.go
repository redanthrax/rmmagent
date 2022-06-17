package tactical

func NewAgentConfig() *rmm.AgentConfig {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\TacticalRMM`, registry.ALL_ACCESS)

	if err != nil {
		return &rmm.AgentConfig{}
	}

	baseurl, _, _ := k.GetStringValue("BaseURL")
	agentid, _, _ := k.GetStringValue("AgentID")
	apiurl, _, _ := k.GetStringValue("ApiURL")
	token, _, _ := k.GetStringValue("Token")
	agentpk, _, _ := k.GetStringValue("AgentPK")
	pk, _ := strconv.Atoi(agentpk)
	cert, _, _ := k.GetStringValue("Cert")
	proxy, _, _ := k.GetStringValue("Proxy")
	customMeshDir, _, _ := k.GetStringValue("MeshDir")

	return &rmm.AgentConfig{
		BaseURL:       baseurl,
		AgentID:       agentid,
		APIURL:        apiurl,
		Token:         token,
		AgentPK:       agentpk,
		PK:            pk,
		Cert:          cert,
		Proxy:         proxy,
		CustomMeshDir: customMeshDir,
	}
}