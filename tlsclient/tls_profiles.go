package tlsclient

import (
	"github.com/bogdanfinn/fhttp/http2"
	tlsclient "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
	bfutls "github.com/bogdanfinn/utls"
)

func Chrome127() profiles.ClientProfile {
	ja3String := "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,51-5-16-65281-10-45-65037-43-17513-13-0-23-18-35-11-27-41,25497-29-23-24,0"

	supportedSignatureAlgorithms := []string{
		"ECDSAWithP256AndSHA256", // 0x0403
		"PSSWithSHA256",          // 0x0804
		"PKCS1WithSHA256",        // 0x0401
		"ECDSAWithP384AndSHA384", // 0x0503
		"PSSWithSHA384",          // 0x0805
		"PKCS1WithSHA384",        // 0x0501
		"PSSWithSHA512",          // 0x0806
		"PKCS1WithSHA512",        // 0x0601
	}

	supportedDelegatedCredentialsAlgorithms := []string{
		"ECDSAWithP256AndSHA256", // 0x0403
		"PSSWithSHA256",          // 0x0804
		"PKCS1WithSHA256",        // 0x0401
		"ECDSAWithP384AndSHA384", // 0x0503
		"PSSWithSHA384",          // 0x0805
		"PKCS1WithSHA384",        // 0x0501
		"PSSWithSHA512",          // 0x0806
		"PKCS1WithSHA512",        // 0x0601
	}
	supportedVersions := []string{"1.3", "1.2", "TLS_GREASE (0x0a0a)"}
	keyShareCurves := []string{"X25519", "X25519Kyber768", "TLS_GREASE (0xcaca)"}

	supportedProtocolsALPN := []string{"h2", "http/1.1"}
	supportedProtocolsALPS := []string{"h2"}

	echCandidateCipherSuites := []tlsclient.CandidateCipherSuites{
		{
			KdfId:  "HKDF_SHA256",
			AeadId: "AEAD_AES_128_GCM",
		},
		{
			KdfId:  "HKDF_SHA256",
			AeadId: "AEAD_CHACHA20_POLY1305",
		},
	}

	candidatePayloads := []uint16{128, 160, 192, 224}

	specFunc, err := tlsclient.GetSpecFactoryFromJa3String(ja3String, supportedSignatureAlgorithms, supportedDelegatedCredentialsAlgorithms, supportedVersions, keyShareCurves, supportedProtocolsALPN, supportedProtocolsALPS, echCandidateCipherSuites, candidatePayloads, "brotli")

	if err != nil {
		specFunc = profiles.Chrome_124.GetClientHelloSpec
	}

	settings := map[http2.SettingID]uint32{
		http2.SettingHeaderTableSize:   65536,
		http2.SettingEnablePush:        0,
		http2.SettingInitialWindowSize: 6291456,
		http2.SettingMaxHeaderListSize: 262144,
	}
	settingsOrder := []http2.SettingID{
		http2.SettingHeaderTableSize,
		http2.SettingEnablePush,
		http2.SettingInitialWindowSize,
		http2.SettingMaxHeaderListSize,
	}

	pseudoHeaderOrder := []string{
		":method",
		":authority",
		":scheme",
		":path",
	}

	connectionFlow := uint32(15663105)

	seed, err := bfutls.NewPRNGSeed()
	if err != nil {
		seed = nil
	}

	return profiles.NewClientProfile(bfutls.ClientHelloID{
		Client:               "Chrome",
		Version:              "127",
		RandomExtensionOrder: false,
		Seed:                 seed,
		Weights:              &bfutls.DefaultWeights,
		SpecFactory:          specFunc,
	}, settings, settingsOrder, pseudoHeaderOrder, connectionFlow, nil, nil)
}

func Edge117() profiles.ClientProfile {
	ja3String := "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,13-10-51-18-65281-27-0-23-35-43-16-11-5-45-17513-21,29-23-24,0"

	// find the following in tls.go
	_ = tlsclient.H2SettingsMap
	_ = bfutls.PKCS1WithSHA256

	supportedSignatureAlgorithms := []string{
		"ECDSAWithP256AndSHA256", // 0x0403
		"PSSWithSHA256",          // 0x0804
		"PKCS1WithSHA256",        // 0x0401
		"ECDSAWithP384AndSHA384", // 0x0503
		"PSSWithSHA384",          // 0x0805
		"PKCS1WithSHA384",        // 0x0501
		"PSSWithSHA512",          // 0x0806
		"PKCS1WithSHA512",        // 0x0601
	}

	supportedDelegatedCredentialsAlgorithms := []string{
		"ECDSAWithP256AndSHA256", // 0x0403
		"PSSWithSHA256",          // 0x0804
		"PKCS1WithSHA256",        // 0x0401
		"ECDSAWithP384AndSHA384", // 0x0503
		"PSSWithSHA384",          // 0x0805
		"PKCS1WithSHA384",        // 0x0501
		"PSSWithSHA512",          // 0x0806
		"PKCS1WithSHA512",        // 0x0601
	}
	supportedVersions := []string{"1.3", "1.2"}
	keyShareCurves := []string{"X25519"}

	supportedProtocolsALPN := []string{"h2", "http/1.1"}
	supportedProtocolsALPS := []string{"h2"}

	echCandidateCipherSuites := []tlsclient.CandidateCipherSuites{
		{
			KdfId:  "HKDF_SHA256",
			AeadId: "AEAD_AES_128_GCM",
		},
		{
			KdfId:  "HKDF_SHA256",
			AeadId: "AEAD_CHACHA20_POLY1305",
		},
	}

	candidatePayloads := []uint16{128, 160, 192, 224}

	specFunc, err := tlsclient.GetSpecFactoryFromJa3String(ja3String, supportedSignatureAlgorithms, supportedDelegatedCredentialsAlgorithms, supportedVersions, keyShareCurves, supportedProtocolsALPN, supportedProtocolsALPS, echCandidateCipherSuites, candidatePayloads, "zlib")

	if err != nil {
		specFunc = profiles.Chrome_124.GetClientHelloSpec
	}

	settings := map[http2.SettingID]uint32{
		http2.SettingHeaderTableSize:   65536,
		http2.SettingEnablePush:        0,
		http2.SettingInitialWindowSize: 6291456,
		http2.SettingMaxHeaderListSize: 262144,
	}
	settingsOrder := []http2.SettingID{
		http2.SettingHeaderTableSize,
		http2.SettingEnablePush,
		http2.SettingInitialWindowSize,
		http2.SettingMaxHeaderListSize,
	}

	pseudoHeaderOrder := []string{
		":method",
		":authority",
		":scheme",
		":path",
	}

	connectionFlow := uint32(15663105)

	seed, err := bfutls.NewPRNGSeed()
	if err != nil {
		seed = nil
	}

	return profiles.NewClientProfile(bfutls.ClientHelloID{
		Client:               "Edge",
		Version:              "117",
		RandomExtensionOrder: false,
		Seed:                 seed,
		Weights:              &bfutls.DefaultWeights,
		SpecFactory:          specFunc,
	}, settings, settingsOrder, pseudoHeaderOrder, connectionFlow, nil, nil)
}
func Firefox129() profiles.ClientProfile {
	ja3String := "771,4865-4867-4866-49195-49199-52393-52392-49196-49200-49162-49161-49171-49172-156-157-47-53,0-23-65281-10-11-35-16-5-34-51-43-13-45-28-65037,29-23-24-25-256-257,0"

	supportedSignatureAlgorithms := []string{
		"ECDSAWithP256AndSHA256", // 0x0403
		"ECDSAWithP384AndSHA384", // 0x0503
		"ECDSAWithP521AndSHA512", // 0x0603
		"PSSWithSHA256",          // 0x0804
		"PSSWithSHA384",          // 0x0805
		"PSSWithSHA512",          // 0x0806
		"PKCS1WithSHA256",        // 0x0401
		"PKCS1WithSHA384",        // 0x0501
		"PKCS1WithSHA512",        // 0x0601
		"ECDSAWithSHA1",          // 0x0203
		"PKCS1WithSHA1",          // 0x0201
	}

	supportedDelegatedCredentialsAlgorithms := []string{
		"ECDSAWithP256AndSHA256", // 0x0403
		"ECDSAWithP384AndSHA384", // 0x0503
		"ECDSAWithP521AndSHA512", // 0x0603
		"ECDSAWithSHA1",          // 0x0203
	}

	supportedVersions := []string{"1.3", "1.2"}
	keyShareCurves := []string{"X25519", "P-256", "P-384", "P-521", "ffdhe2048", "ffdhe3072"}

	supportedProtocolsALPN := []string{"h2", "http/1.1"}
	supportedProtocolsALPS := []string{"h2"}

	echCandidateCipherSuites := []tlsclient.CandidateCipherSuites{
		{
			KdfId:  "HKDF_SHA256",
			AeadId: "AEAD_AES_128_GCM",
		},
		{
			KdfId:  "HKDF_SHA256",
			AeadId: "AEAD_CHACHA20_POLY1305",
		},
	}

	candidatePayloads := []uint16{128, 160, 192, 224}

	specFunc, err := tlsclient.GetSpecFactoryFromJa3String(ja3String, supportedSignatureAlgorithms, supportedDelegatedCredentialsAlgorithms, supportedVersions, keyShareCurves, supportedProtocolsALPN, supportedProtocolsALPS, echCandidateCipherSuites, candidatePayloads, "brotli")

	if err != nil {
		specFunc = profiles.Firefox_123.GetClientHelloSpec
	}

	settings := map[http2.SettingID]uint32{
		http2.SettingHeaderTableSize:   65536,
		http2.SettingInitialWindowSize: 131072,
		http2.SettingMaxFrameSize:      16384,
	}
	settingsOrder := []http2.SettingID{
		http2.SettingHeaderTableSize,
		http2.SettingInitialWindowSize,
		http2.SettingMaxFrameSize,
	}

	pseudoHeaderOrder := []string{
		":method",
		":authority",
		":scheme",
		":path",
	}

	connectionFlow := uint32(12517377)

	seed, err := bfutls.NewPRNGSeed()
	if err != nil {
		seed = nil
	}

	return profiles.NewClientProfile(bfutls.ClientHelloID{
		Client:               "Firefox",
		Version:              "129",
		RandomExtensionOrder: false,
		Seed:                 seed,
		Weights:              &bfutls.DefaultWeights,
		SpecFactory:          specFunc,
	}, settings, settingsOrder, pseudoHeaderOrder, connectionFlow, nil, nil)
}
