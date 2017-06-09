package frames

func DomainInfo() string {
	return `
        <?xml version="1.0" encoding="UTF-8" standalone="no"?>
        <epp xmlns="urn:ietf:params:xml:ns:epp-1.0">
            <command>
                <info>
                    <domain:info
                        xmlns:domain="urn:ietf:params:xml:ns:domain-1.0">
                        <domain:name hosts="all">{domain}</domain:name>
                        <domain:authInfo>
                            <domain:pw></domain:pw>
                        </domain:authInfo>
                    </domain:info>
                </info>
                <clTRID>EPP-{clTRID}</clTRID>
            </command>
        </epp>
    `
}
