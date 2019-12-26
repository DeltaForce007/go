package processors

import (
	"testing"

	"github.com/stellar/go/services/horizon/internal/db2/history"
	. "github.com/stellar/go/services/horizon/internal/test/transactions"
	"github.com/stellar/go/xdr"
	"github.com/stretchr/testify/assert"
)

func TestOperationEffects(t *testing.T) {
	testCases := []struct {
		desc          string
		envelopeXDR   string
		resultXDR     string
		metaXDR       string
		feeChangesXDR string
		hash          string
		index         uint32
		sequence      uint32
		expected      []map[string]interface{}
	}{
		{
			desc:          "createAccount",
			envelopeXDR:   "AAAAAGL8HQvQkbK2HA3WVjRrKmjX00fG8sLI7m0ERwJW/AX3AAAAZAAAAAAAAAAaAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAoZftFP3p4ifbTm6hQdieotu3Zw9E05GtoSh5MBytEpQAAAACVAvkAAAAAAAAAAABVvwF9wAAAEDHU95E9wxgETD8TqxUrkgC0/7XHyNDts6Q5huRHfDRyRcoHdv7aMp/sPvC3RPkXjOMjgbKJUX7SgExUeYB5f8F",
			resultXDR:     "AAAAAAAAAGQAAAAAAAAAAQAAAAAAAAABAAAAAAAAAAA=",
			metaXDR:       "AAAAAQAAAAIAAAADAAAAOQAAAAAAAAAAYvwdC9CRsrYcDdZWNGsqaNfTR8bywsjubQRHAlb8BfcLGrZY9dZxbAAAAAAAAAAZAAAAAAAAAAEAAAAAYvwdC9CRsrYcDdZWNGsqaNfTR8bywsjubQRHAlb8BfcAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAABAAAAOQAAAAAAAAAAYvwdC9CRsrYcDdZWNGsqaNfTR8bywsjubQRHAlb8BfcLGrZY9dZxbAAAAAAAAAAaAAAAAAAAAAEAAAAAYvwdC9CRsrYcDdZWNGsqaNfTR8bywsjubQRHAlb8BfcAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAABAAAAAwAAAAMAAAA5AAAAAAAAAABi/B0L0JGythwN1lY0aypo19NHxvLCyO5tBEcCVvwF9wsatlj11nFsAAAAAAAAABoAAAAAAAAAAQAAAABi/B0L0JGythwN1lY0aypo19NHxvLCyO5tBEcCVvwF9wAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAEAAAA5AAAAAAAAAABi/B0L0JGythwN1lY0aypo19NHxvLCyO5tBEcCVvwF9wsatlahyo1sAAAAAAAAABoAAAAAAAAAAQAAAABi/B0L0JGythwN1lY0aypo19NHxvLCyO5tBEcCVvwF9wAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAA5AAAAAAAAAAChl+0U/eniJ9tObqFB2J6i27dnD0TTka2hKHkwHK0SlAAAAAJUC+QAAAAAOQAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAA==",
			feeChangesXDR: "AAAAAgAAAAMAAAA3AAAAAAAAAABi/B0L0JGythwN1lY0aypo19NHxvLCyO5tBEcCVvwF9wsatlj11nHQAAAAAAAAABkAAAAAAAAAAQAAAABi/B0L0JGythwN1lY0aypo19NHxvLCyO5tBEcCVvwF9wAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAEAAAA5AAAAAAAAAABi/B0L0JGythwN1lY0aypo19NHxvLCyO5tBEcCVvwF9wsatlj11nFsAAAAAAAAABkAAAAAAAAAAQAAAABi/B0L0JGythwN1lY0aypo19NHxvLCyO5tBEcCVvwF9wAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAA==",
			hash:          "0e5bd332291e3098e49886df2cdb9b5369a5f9e0a9973f0d9e1a9489c6581ba2",
			index:         0,
			sequence:      57,
			expected: []map[string]interface{}{
				map[string]interface{}{
					"address":     "GCQZP3IU7XU6EJ63JZXKCQOYT2RNXN3HB5CNHENNUEUHSMA4VUJJJSEN",
					"operationID": int64(244813139969),
					"details": map[string]interface{}{
						"starting_balance": "1000.0000000",
					},
					"effectType": history.EffectAccountCreated,
					"order":      uint32(1),
				},
				map[string]interface{}{
					"address":     "GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H",
					"operationID": int64(244813139969),
					"details": map[string]interface{}{
						"amount":     "1000.0000000",
						"asset_type": "native",
					},
					"effectType": history.EffectAccountDebited,
					"order":      uint32(2),
				},
				map[string]interface{}{
					"address":     "GCQZP3IU7XU6EJ63JZXKCQOYT2RNXN3HB5CNHENNUEUHSMA4VUJJJSEN",
					"operationID": int64(244813139969),
					"details": map[string]interface{}{
						"public_key": "GCQZP3IU7XU6EJ63JZXKCQOYT2RNXN3HB5CNHENNUEUHSMA4VUJJJSEN",
						"weight":     1,
					},
					"effectType": history.EffectSignerCreated,
					"order":      uint32(3),
				},
			},
		},
		{
			desc:          "payment",
			envelopeXDR:   "AAAAABpcjiETZ0uhwxJJhgBPYKWSVJy2TZ2LI87fqV1cUf/UAAAAZAAAADcAAAABAAAAAAAAAAAAAAABAAAAAAAAAAEAAAAAGlyOIRNnS6HDEkmGAE9gpZJUnLZNnYsjzt+pXVxR/9QAAAAAAAAAAAX14QAAAAAAAAAAAVxR/9QAAABAK6pcXYMzAEmH08CZ1LWmvtNDKauhx+OImtP/Lk4hVTMJRVBOebVs5WEPj9iSrgGT0EswuDCZ2i5AEzwgGof9Ag==",
			resultXDR:     "AAAAAAAAAGQAAAAAAAAAAQAAAAAAAAABAAAAAAAAAAA=",
			metaXDR:       "AAAAAQAAAAIAAAADAAAAOAAAAAAAAAAAGlyOIRNnS6HDEkmGAE9gpZJUnLZNnYsjzt+pXVxR/9QAAAACVAvjnAAAADcAAAAAAAAAAAAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAABAAAAOAAAAAAAAAAAGlyOIRNnS6HDEkmGAE9gpZJUnLZNnYsjzt+pXVxR/9QAAAACVAvjnAAAADcAAAABAAAAAAAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAABAAAAAA==",
			feeChangesXDR: "AAAAAgAAAAMAAAA3AAAAAAAAAAAaXI4hE2dLocMSSYYAT2ClklSctk2diyPO36ldXFH/1AAAAAJUC+QAAAAANwAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAEAAAA4AAAAAAAAAAAaXI4hE2dLocMSSYYAT2ClklSctk2diyPO36ldXFH/1AAAAAJUC+OcAAAANwAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAA==",
			hash:          "2a805712c6d10f9e74bb0ccf54ae92a2b4b1e586451fe8133a2433816f6b567c",
			index:         0,
			sequence:      56,
			expected: []map[string]interface{}{
				map[string]interface{}{
					"address": "GANFZDRBCNTUXIODCJEYMACPMCSZEVE4WZGZ3CZDZ3P2SXK4KH75IK6Y",
					"details": map[string]interface{}{
						"amount":     "10.0000000",
						"asset_type": "native",
					},
					"effectType":  history.EffectAccountCredited,
					"operationID": int64(240518172673),
					"order":       uint32(1),
				},
				map[string]interface{}{
					"address": "GANFZDRBCNTUXIODCJEYMACPMCSZEVE4WZGZ3CZDZ3P2SXK4KH75IK6Y",
					"details": map[string]interface{}{
						"amount":     "10.0000000",
						"asset_type": "native",
					},
					"effectType":  history.EffectAccountDebited,
					"operationID": int64(240518172673),
					"order":       uint32(2),
				},
			},
		},
		{
			desc:          "pathPaymentStrictReceive",
			envelopeXDR:   "AAAAAONt/6wGI884Zi6sYDYC1GOV/drnh4OcRrTrqJPoOTUKAAAAZAAAABAAAAADAAAAAAAAAAAAAAABAAAAAAAAAAIAAAAAAAAAADuaygAAAAAABAjoBMEUiZNLUjsWXL1iK59D90Li4w56076b8HKxZfIAAAABRVVSAAAAAAAuwvNzNk9twbuJHUBqnX26GYI3MbCdpQU9t4n6EVRXsQAAAAA7msoAAAAAAAAAAAAAAAAB6Dk1CgAAAEB+7jxesBKKrF343onyycjp2tiQLZiGH2ETl+9fuOqotveY2rIgvt9ng+QJ2aDP3+PnDsYEa9ZUaA+Zne2nIGgE",
			resultXDR:     "AAAAAAAAAGQAAAAAAAAAAQAAAAAAAAACAAAAAAAAAAEAAAAALsLzczZPbcG7iR1Aap19uhmCNzGwnaUFPbeJ+hFUV7EAAAAAAAAAAgAAAAFFVVIAAAAAAC7C83M2T23Bu4kdQGqdfboZgjcxsJ2lBT23ifoRVFexAAAAADuaygAAAAAAAAAAADuaygAAAAAABAjoBMEUiZNLUjsWXL1iK59D90Li4w56076b8HKxZfIAAAABRVVSAAAAAAAuwvNzNk9twbuJHUBqnX26GYI3MbCdpQU9t4n6EVRXsQAAAAA7msoAAAAAAA==",
			metaXDR:       "AAAAAQAAAAIAAAADAAAAFAAAAAAAAAAA423/rAYjzzhmLqxgNgLUY5X92ueHg5xGtOuok+g5NQoAAAACVAvi1AAAABAAAAACAAAAAQAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAABAAAAFAAAAAAAAAAA423/rAYjzzhmLqxgNgLUY5X92ueHg5xGtOuok+g5NQoAAAACVAvi1AAAABAAAAADAAAAAQAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAABAAAACAAAAAMAAAATAAAAAQAAAAAECOgEwRSJk0tSOxZcvWIrn0P3QuLjDnrTvpvwcrFl8gAAAAFFVVIAAAAAAC7C83M2T23Bu4kdQGqdfboZgjcxsJ2lBT23ifoRVFexAAAAAHc1lAB//////////wAAAAEAAAAAAAAAAAAAAAEAAAAUAAAAAQAAAAAECOgEwRSJk0tSOxZcvWIrn0P3QuLjDnrTvpvwcrFl8gAAAAFFVVIAAAAAAC7C83M2T23Bu4kdQGqdfboZgjcxsJ2lBT23ifoRVFexAAAAALLQXgB//////////wAAAAEAAAAAAAAAAAAAAAMAAAATAAAAAgAAAAAuwvNzNk9twbuJHUBqnX26GYI3MbCdpQU9t4n6EVRXsQAAAAAAAAACAAAAAUVVUgAAAAAALsLzczZPbcG7iR1Aap19uhmCNzGwnaUFPbeJ+hFUV7EAAAAAAAAAADuaygAAAAABAAAAAQAAAAAAAAAAAAAAAAAAAAIAAAACAAAAAC7C83M2T23Bu4kdQGqdfboZgjcxsJ2lBT23ifoRVFexAAAAAAAAAAIAAAADAAAAFAAAAAAAAAAA423/rAYjzzhmLqxgNgLUY5X92ueHg5xGtOuok+g5NQoAAAACVAvi1AAAABAAAAADAAAAAQAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAABAAAAFAAAAAAAAAAA423/rAYjzzhmLqxgNgLUY5X92ueHg5xGtOuok+g5NQoAAAACGHEY1AAAABAAAAADAAAAAQAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAADAAAAEwAAAAAAAAAALsLzczZPbcG7iR1Aap19uhmCNzGwnaUFPbeJ+hFUV7EAAAACVAvi1AAAABAAAAADAAAAAgAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAQAAAAA7msoAAAAAAHc1lAAAAAAAAAAAAAAAAAEAAAAUAAAAAAAAAAAuwvNzNk9twbuJHUBqnX26GYI3MbCdpQU9t4n6EVRXsQAAAAKPpqzUAAAAEAAAAAMAAAABAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAABAAAAAAAAAAAAAAAAdzWUAAAAAAAAAAAA",
			feeChangesXDR: "AAAAAgAAAAMAAAATAAAAAAAAAADjbf+sBiPPOGYurGA2AtRjlf3a54eDnEa066iT6Dk1CgAAAAJUC+M4AAAAEAAAAAIAAAABAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAEAAAAUAAAAAAAAAADjbf+sBiPPOGYurGA2AtRjlf3a54eDnEa066iT6Dk1CgAAAAJUC+LUAAAAEAAAAAIAAAABAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAA==",
			hash:          "96415ac1d2f79621b26b1568f963fd8dd6c50c20a22c7428cefbfe9dee867588",
			index:         0,
			sequence:      20,
			expected: []map[string]interface{}{
				map[string]interface{}{
					"address": "GACAR2AEYEKITE2LKI5RMXF5MIVZ6Q7XILROGDT22O7JX4DSWFS7FDDP",
					"details": map[string]interface{}{
						"amount":       "100.0000000",
						"asset_code":   "EUR",
						"asset_issuer": "GAXMF43TGZHW3QN3REOUA2U5PW5BTARXGGYJ3JIFHW3YT6QRKRL3CPPU",
						"asset_type":   "credit_alphanum4",
					},
					"effectType":  history.EffectAccountCredited,
					"operationID": int64(85899350017),
					"order":       uint32(1),
				}, map[string]interface{}{
					"address": "GDRW375MAYR46ODGF2WGANQC2RRZL7O246DYHHCGWTV2RE7IHE2QUQLD",
					"details": map[string]interface{}{
						"amount":     "100.0000000",
						"asset_type": "native",
					},
					"effectType":  history.EffectAccountDebited,
					"operationID": int64(85899350017),
					"order":       uint32(2),
				}, map[string]interface{}{
					"address": "GDRW375MAYR46ODGF2WGANQC2RRZL7O246DYHHCGWTV2RE7IHE2QUQLD",
					"details": map[string]interface{}{
						"bought_amount":       "100.0000000",
						"bought_asset_code":   "EUR",
						"bought_asset_issuer": "GAXMF43TGZHW3QN3REOUA2U5PW5BTARXGGYJ3JIFHW3YT6QRKRL3CPPU",
						"bought_asset_type":   "credit_alphanum4",
						"offer_id":            xdr.Int64(2),
						"seller":              "GAXMF43TGZHW3QN3REOUA2U5PW5BTARXGGYJ3JIFHW3YT6QRKRL3CPPU",
						"sold_amount":         "100.0000000",
						"sold_asset_type":     "native",
					},
					"effectType":  history.EffectTrade,
					"operationID": int64(85899350017),
					"order":       uint32(3),
				}, map[string]interface{}{
					"address": "GAXMF43TGZHW3QN3REOUA2U5PW5BTARXGGYJ3JIFHW3YT6QRKRL3CPPU",
					"details": map[string]interface{}{
						"bought_amount":     "100.0000000",
						"bought_asset_type": "native",
						"offer_id":          xdr.Int64(2),
						"seller":            "GDRW375MAYR46ODGF2WGANQC2RRZL7O246DYHHCGWTV2RE7IHE2QUQLD",
						"sold_amount":       "100.0000000",
						"sold_asset_code":   "EUR",
						"sold_asset_issuer": "GAXMF43TGZHW3QN3REOUA2U5PW5BTARXGGYJ3JIFHW3YT6QRKRL3CPPU",
						"sold_asset_type":   "credit_alphanum4",
					},
					"effectType":  history.EffectTrade,
					"operationID": int64(85899350017),
					"order":       uint32(4),
				},
			},
		},
		{
			desc:          "manageSellOffer - without claims",
			envelopeXDR:   "AAAAAC7C83M2T23Bu4kdQGqdfboZgjcxsJ2lBT23ifoRVFexAAAAZAAAABAAAAACAAAAAAAAAAAAAAABAAAAAAAAAAMAAAAAAAAAAVVTRAAAAAAALsLzczZPbcG7iR1Aap19uhmCNzGwnaUFPbeJ+hFUV7EAAAAA7msoAAAAAAEAAAACAAAAAAAAAAAAAAAAAAAAARFUV7EAAABALuai5QxceFbtAiC5nkntNVnvSPeWR+C+FgplPAdRgRS+PPESpUiSCyuiwuhmvuDw7kwxn+A6E0M4ca1s2qzMAg==",
			resultXDR:     "AAAAAAAAAGQAAAAAAAAAAQAAAAAAAAADAAAAAAAAAAAAAAAAAAAAAC7C83M2T23Bu4kdQGqdfboZgjcxsJ2lBT23ifoRVFexAAAAAAAAAAEAAAAAAAAAAVVTRAAAAAAALsLzczZPbcG7iR1Aap19uhmCNzGwnaUFPbeJ+hFUV7EAAAAA7msoAAAAAAEAAAACAAAAAAAAAAAAAAAA",
			metaXDR:       "AAAAAQAAAAIAAAADAAAAEgAAAAAAAAAALsLzczZPbcG7iR1Aap19uhmCNzGwnaUFPbeJ+hFUV7EAAAACVAvi1AAAABAAAAABAAAAAAAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAABAAAAEgAAAAAAAAAALsLzczZPbcG7iR1Aap19uhmCNzGwnaUFPbeJ+hFUV7EAAAACVAvi1AAAABAAAAACAAAAAAAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAABAAAAAwAAAAMAAAASAAAAAAAAAAAuwvNzNk9twbuJHUBqnX26GYI3MbCdpQU9t4n6EVRXsQAAAAJUC+LUAAAAEAAAAAIAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAEAAAASAAAAAAAAAAAuwvNzNk9twbuJHUBqnX26GYI3MbCdpQU9t4n6EVRXsQAAAAJUC+LUAAAAEAAAAAIAAAABAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAABAAAAAAAAAAAAAAAA7msoAAAAAAAAAAAAAAAAAAAAABIAAAACAAAAAC7C83M2T23Bu4kdQGqdfboZgjcxsJ2lBT23ifoRVFexAAAAAAAAAAEAAAAAAAAAAVVTRAAAAAAALsLzczZPbcG7iR1Aap19uhmCNzGwnaUFPbeJ+hFUV7EAAAAA7msoAAAAAAEAAAACAAAAAAAAAAAAAAAA",
			feeChangesXDR: "AAAAAgAAAAMAAAASAAAAAAAAAAAuwvNzNk9twbuJHUBqnX26GYI3MbCdpQU9t4n6EVRXsQAAAAJUC+OcAAAAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAEAAAASAAAAAAAAAAAuwvNzNk9twbuJHUBqnX26GYI3MbCdpQU9t4n6EVRXsQAAAAJUC+M4AAAAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAA==",
			hash:          "ca756d1519ceda79f8722042b12cea7ba004c3bd961adb62b59f88a867f86eb3",
			index:         0,
			sequence:      56,
			expected:      []map[string]interface{}{},
		},
		{
			desc:          "manageSellOffer - with claims",
			envelopeXDR:   "AAAAAPrjQnnOn4RqMmOSDwYfEMVtJuC4VR9fKvPfEtM7DS7VAAAAZAAMDl8AAAADAAAAAAAAAAAAAAABAAAAAAAAAAMAAAAAAAAAAVNUUgAAAAAASYK2XlJiUiNav1waFVDq1fzoualYC4UNFqThKBroJe0AAAACVAvkAAAAAGMAAADIAAAAAAAAAAAAAAAAAAAAATsNLtUAAABABmA0aLobgdSrjIrus94Y8PWeD6dDfl7Sya12t2uZasJFI7mZ+yowE1enUMzC/cAhDTypK8QuH2EVXPQC3xpYDA==",
			resultXDR:     "AAAAAAAAAGQAAAAAAAAAAQAAAAAAAAADAAAAAAAAAAEAAAAADkfaGg9y56NND7n4CRcr4R4fvivwAcMd4ZrCm4jAe5AAAAAAAI0f+AAAAAFTVFIAAAAAAEmCtl5SYlIjWr9cGhVQ6tX86LmpWAuFDRak4Sga6CXtAAAAAS0Il1oAAAAAAAAAAlQL4/8AAAACAAAAAA==",
			metaXDR:       "AAAAAQAAAAIAAAADAAxMfwAAAAAAAAAA+uNCec6fhGoyY5IPBh8QxW0m4LhVH18q898S0zsNLtUAAAAU9GsC1QAMDl8AAAACAAAAAQAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAABAAxMfwAAAAAAAAAA+uNCec6fhGoyY5IPBh8QxW0m4LhVH18q898S0zsNLtUAAAAU9GsC1QAMDl8AAAADAAAAAQAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAABAAAACgAAAAMADEx+AAAAAgAAAAAOR9oaD3Lno00PufgJFyvhHh++K/ABwx3hmsKbiMB7kAAAAAAAjR/4AAAAAVNUUgAAAAAASYK2XlJiUiNav1waFVDq1fzoualYC4UNFqThKBroJe0AAAAAAAAAA2L6BdYAAABjAAAAMgAAAAAAAAAAAAAAAAAAAAEADEx/AAAAAgAAAAAOR9oaD3Lno00PufgJFyvhHh++K/ABwx3hmsKbiMB7kAAAAAAAjR/4AAAAAVNUUgAAAAAASYK2XlJiUiNav1waFVDq1fzoualYC4UNFqThKBroJe0AAAAAAAAAAjXxbnwAAABjAAAAMgAAAAAAAAAAAAAAAAAAAAMADEx+AAAAAAAAAAAOR9oaD3Lno00PufgJFyvhHh++K/ABwx3hmsKbiMB7kAAAABnMMdMvAAwOZQAAAAIAAAACAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAABAAAABrSdIAkAAAAAAAAAAAAAAAAAAAAAAAAAAQAMTH8AAAAAAAAAAA5H2hoPcuejTQ+5+AkXK+EeH74r8AHDHeGawpuIwHuQAAAAHCA9ty4ADA5lAAAAAgAAAAIAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAEAAAAEYJE8CgAAAAAAAAAAAAAAAAAAAAAAAAADAAxMfgAAAAEAAAAADkfaGg9y56NND7n4CRcr4R4fvivwAcMd4ZrCm4jAe5AAAAABU1RSAAAAAABJgrZeUmJSI1q/XBoVUOrV/Oi5qVgLhQ0WpOEoGugl7QAAABYDWSXWf/////////8AAAABAAAAAQAAAAAAAAAAAAAAA2L6BdYAAAAAAAAAAAAAAAEADEx/AAAAAQAAAAAOR9oaD3Lno00PufgJFyvhHh++K/ABwx3hmsKbiMB7kAAAAAFTVFIAAAAAAEmCtl5SYlIjWr9cGhVQ6tX86LmpWAuFDRak4Sga6CXtAAAAFNZQjnx//////////wAAAAEAAAABAAAAAAAAAAAAAAACNfFufAAAAAAAAAAAAAAAAwAMDnEAAAABAAAAAPrjQnnOn4RqMmOSDwYfEMVtJuC4VR9fKvPfEtM7DS7VAAAAAVNUUgAAAAAASYK2XlJiUiNav1waFVDq1fzoualYC4UNFqThKBroJe0AAAAYdX9/Wn//////////AAAAAQAAAAAAAAAAAAAAAQAMTH8AAAABAAAAAPrjQnnOn4RqMmOSDwYfEMVtJuC4VR9fKvPfEtM7DS7VAAAAAVNUUgAAAAAASYK2XlJiUiNav1waFVDq1fzoualYC4UNFqThKBroJe0AAAAZoogWtH//////////AAAAAQAAAAAAAAAAAAAAAwAMTH8AAAAAAAAAAPrjQnnOn4RqMmOSDwYfEMVtJuC4VR9fKvPfEtM7DS7VAAAAFPRrAtUADA5fAAAAAwAAAAEAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAQAMTH8AAAAAAAAAAPrjQnnOn4RqMmOSDwYfEMVtJuC4VR9fKvPfEtM7DS7VAAAAEqBfHtYADA5fAAAAAwAAAAEAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAA",
			feeChangesXDR: "AAAAAgAAAAMADA5xAAAAAAAAAAD640J5zp+EajJjkg8GHxDFbSbguFUfXyrz3xLTOw0u1QAAABT0awM5AAwOXwAAAAIAAAABAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAEADEx/AAAAAAAAAAD640J5zp+EajJjkg8GHxDFbSbguFUfXyrz3xLTOw0u1QAAABT0awLVAAwOXwAAAAIAAAABAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAA==",
			hash:          "ef62da32b6b3eb3c4534dac2be1088387fb93b0093b47e113073c1431fac9db7",
			index:         0,
			sequence:      56,
			expected: []map[string]interface{}{
				map[string]interface{}{
					"address": "GD5OGQTZZ2PYI2RSMOJA6BQ7CDCW2JXAXBKR6XZK6PPRFUZ3BUXNLFKP",
					"details": map[string]interface{}{
						"seller":              "GAHEPWQ2B5ZOPI2NB647QCIXFPQR4H56FPYADQY54GNMFG4IYB5ZAJ5H",
						"offer_id":            xdr.Int64(9248760),
						"sold_amount":         "999.9999999",
						"bought_amount":       "505.0505050",
						"sold_asset_type":     "native",
						"bought_asset_code":   "STR",
						"bought_asset_type":   "credit_alphanum4",
						"bought_asset_issuer": "GBEYFNS6KJRFEI22X5OBUFKQ5LK7Z2FZVFMAXBINC2SOCKA25AS62PUN",
					},
					"effectType":  history.EffectTrade,
					"operationID": int64(240518172673),
					"order":       uint32(1),
				},
				map[string]interface{}{
					"address": "GAHEPWQ2B5ZOPI2NB647QCIXFPQR4H56FPYADQY54GNMFG4IYB5ZAJ5H",
					"details": map[string]interface{}{
						"seller":            "GD5OGQTZZ2PYI2RSMOJA6BQ7CDCW2JXAXBKR6XZK6PPRFUZ3BUXNLFKP",
						"offer_id":          xdr.Int64(9248760),
						"sold_amount":       "505.0505050",
						"bought_amount":     "999.9999999",
						"sold_asset_code":   "STR",
						"sold_asset_type":   "credit_alphanum4",
						"bought_asset_type": "native",
						"sold_asset_issuer": "GBEYFNS6KJRFEI22X5OBUFKQ5LK7Z2FZVFMAXBINC2SOCKA25AS62PUN",
					},
					"effectType":  history.EffectTrade,
					"operationID": int64(240518172673),
					"order":       uint32(2),
				},
			},
		},
		{
			desc:          "manageBuyOffer - with claims",
			envelopeXDR:   "AAAAAEotqBM9oOzudkkctgQlY/PHS0rFcxVasWQVnSytiuBEAAAAZAANIfEAAAADAAAAAAAAAAAAAAABAAAAAAAAAAwAAAAAAAAAAlRYVGFscGhhNAAAAAAAAABKLagTPaDs7nZJHLYEJWPzx0tKxXMVWrFkFZ0srYrgRAAAAAB3NZQAAAAAAQAAAAEAAAAAAAAAAAAAAAAAAAABrYrgRAAAAEAh57TBifjJuUPj1TI7zIvaAZmyRjWLY4ktc0F16Knmy4Fw07L7cC5vCwjn4ZXyrgr9bpEGhv4oN6znbPpNLQUH",
			resultXDR:     "AAAAAAAAAGQAAAAAAAAAAQAAAAAAAAAMAAAAAAAAAAEAAAAAgbI9jY68fYXd6+DwMcZQQIYCK4HsKKvqnR5o+1IdVoUAAAAAAJovcgAAAAJUWFRhbHBoYTQAAAAAAAAASi2oEz2g7O52SRy2BCVj88dLSsVzFVqxZBWdLK2K4EQAAAAAdzWUAAAAAAAAAAAAdzWUAAAAAAIAAAAA",
			metaXDR:       "AAAAAQAAAAIAAAADAA0pGAAAAAAAAAAASi2oEz2g7O52SRy2BCVj88dLSsVzFVqxZBWdLK2K4EQAAAAXSHbm1AANIfEAAAACAAAAAAAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAABAA0pGAAAAAAAAAAASi2oEz2g7O52SRy2BCVj88dLSsVzFVqxZBWdLK2K4EQAAAAXSHbm1AANIfEAAAADAAAAAAAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAABAAAACAAAAAMADSkYAAAAAAAAAABKLagTPaDs7nZJHLYEJWPzx0tKxXMVWrFkFZ0srYrgRAAAABdIdubUAA0h8QAAAAMAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAEADSkYAAAAAAAAAABKLagTPaDs7nZJHLYEJWPzx0tKxXMVWrFkFZ0srYrgRAAAABbRQVLUAA0h8QAAAAMAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAMADSjEAAAAAgAAAACBsj2Njrx9hd3r4PAxxlBAhgIrgewoq+qdHmj7Uh1WhQAAAAAAmi9yAAAAAlRYVGFscGhhNAAAAAAAAABKLagTPaDs7nZJHLYEJWPzx0tKxXMVWrFkFZ0srYrgRAAAAAAAAAAAstBeAAAAAAEAAAABAAAAAAAAAAAAAAAAAAAAAQANKRgAAAACAAAAAIGyPY2OvH2F3evg8DHGUECGAiuB7Cir6p0eaPtSHVaFAAAAAACaL3IAAAACVFhUYWxwaGE0AAAAAAAAAEotqBM9oOzudkkctgQlY/PHS0rFcxVasWQVnSytiuBEAAAAAAAAAAA7msoAAAAAAQAAAAEAAAAAAAAAAAAAAAAAAAADAA0oxAAAAAAAAAAAgbI9jY68fYXd6+DwMcZQQIYCK4HsKKvqnR5o+1IdVoUAAAAZJU0xXAANGSMAAAARAAAABAAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAQADMowLgdQAAAAAAAAAAAAAAAAAAAAAAAAAAAEADSkYAAAAAAAAAACBsj2Njrx9hd3r4PAxxlBAhgIrgewoq+qdHmj7Uh1WhQAAABmcgsVcAA0ZIwAAABEAAAAEAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAABAAMyi5RMQAAAAAAAAAAAAAAAAAAAAAAAAAAAAwANKMQAAAABAAAAAIGyPY2OvH2F3evg8DHGUECGAiuB7Cir6p0eaPtSHVaFAAAAAlRYVGFscGhhNAAAAAAAAABKLagTPaDs7nZJHLYEJWPzx0tKxXMVWrFkFZ0srYrgRAAACRatNxoAf/////////8AAAABAAAAAQAAAAAAAAAAAAAAALLQXgAAAAAAAAAAAAAAAAEADSkYAAAAAQAAAACBsj2Njrx9hd3r4PAxxlBAhgIrgewoq+qdHmj7Uh1WhQAAAAJUWFRhbHBoYTQAAAAAAAAASi2oEz2g7O52SRy2BCVj88dLSsVzFVqxZBWdLK2K4EQAAAkWNgGGAH//////////AAAAAQAAAAEAAAAAAAAAAAAAAAA7msoAAAAAAAAAAAA=",
			feeChangesXDR: "AAAAAgAAAAMADSSgAAAAAAAAAABKLagTPaDs7nZJHLYEJWPzx0tKxXMVWrFkFZ0srYrgRAAAABdIduc4AA0h8QAAAAIAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAEADSkYAAAAAAAAAABKLagTPaDs7nZJHLYEJWPzx0tKxXMVWrFkFZ0srYrgRAAAABdIdubUAA0h8QAAAAIAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAA==",
			hash:          "9caa91eec6e29730f4aabafb60898a8ecedd3bf67b8628e6e32066fbba9bec5d",
			index:         0,
			sequence:      56,
			expected: []map[string]interface{}{
				map[string]interface{}{
					"address": "GBFC3KATHWQOZ3TWJEOLMBBFMPZ4OS2KYVZRKWVRMQKZ2LFNRLQEIRCV",
					"details": map[string]interface{}{
						"seller":              "GCA3EPMNR26H3BO55PQPAMOGKBAIMARLQHWCRK7KTUPGR62SDVLIL7D6",
						"offer_id":            xdr.Int64(10104690),
						"sold_amount":         "200.0000000",
						"bought_amount":       "200.0000000",
						"sold_asset_type":     "native",
						"bought_asset_code":   "TXTalpha4",
						"bought_asset_type":   "credit_alphanum12",
						"bought_asset_issuer": "GBFC3KATHWQOZ3TWJEOLMBBFMPZ4OS2KYVZRKWVRMQKZ2LFNRLQEIRCV",
					},
					"effectType":  history.EffectTrade,
					"operationID": int64(240518172673),
					"order":       uint32(1),
				},
				map[string]interface{}{
					"address": "GCA3EPMNR26H3BO55PQPAMOGKBAIMARLQHWCRK7KTUPGR62SDVLIL7D6",
					"details": map[string]interface{}{
						"seller":            "GBFC3KATHWQOZ3TWJEOLMBBFMPZ4OS2KYVZRKWVRMQKZ2LFNRLQEIRCV",
						"offer_id":          xdr.Int64(10104690),
						"sold_amount":       "200.0000000",
						"bought_amount":     "200.0000000",
						"sold_asset_code":   "TXTalpha4",
						"sold_asset_type":   "credit_alphanum12",
						"bought_asset_type": "native",
						"sold_asset_issuer": "GBFC3KATHWQOZ3TWJEOLMBBFMPZ4OS2KYVZRKWVRMQKZ2LFNRLQEIRCV",
					},
					"effectType":  history.EffectTrade,
					"operationID": int64(240518172673),
					"order":       uint32(2),
				},
			},
		},
		{
			desc:          "createPassiveSellOffer",
			envelopeXDR:   "AAAAAAHwZwJPu1TJhQGgsLRXBzcIeySkeGXzEqh0W9AHWvFDAAAAZAAN3tMAAAACAAAAAQAAAAAAAAAAAAAAAF4FBqwAAAAAAAAAAQAAAAAAAAAEAAAAAAAAAAFDT1AAAAAAALly/iTceP/82O3aZAmd8hyqUjYAANfc5RfN0/iibCtTAAAAADuaygAAAAAJAAAACgAAAAAAAAABB1rxQwAAAEDz2JIw8Z3Owoc5c2tsiY3kzOYUmh32155u00Xs+RYxO5fL0ApYd78URHcYCbe0R32YmuLTfefWQStR3RfhqKAL",
			resultXDR:     "AAAAAAAAAGQAAAAAAAAAAQAAAAAAAAADAAAAAAAAAAEAAAAAMgQ65fmCczzuwmU3oQLivASzvZdhzjhJOQ6C+xTSDu8AAAAAAKMvZgAAAAFDT1AAAAAAALly/iTceP/82O3aZAmd8hyqUjYAANfc5RfN0/iibCtTAAAA6NSlEAAAAAAAAAAAADuaygAAAAACAAAAAA==",
			metaXDR:       "AAAAAQAAAAIAAAADAA3fGgAAAAAAAAAAAfBnAk+7VMmFAaCwtFcHNwh7JKR4ZfMSqHRb0Ada8UMAAAAXSHbnOAAN3tMAAAABAAAAAQAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAABAA3fGgAAAAAAAAAAAfBnAk+7VMmFAaCwtFcHNwh7JKR4ZfMSqHRb0Ada8UMAAAAXSHbnOAAN3tMAAAACAAAAAQAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAABAAAACgAAAAMADd72AAAAAgAAAAAyBDrl+YJzPO7CZTehAuK8BLO9l2HOOEk5DoL7FNIO7wAAAAAAoy9mAAAAAUNPUAAAAAAAuXL+JNx4//zY7dpkCZ3yHKpSNgAA19zlF83T+KJsK1MAAAAAAAAA6NSlEAAAAAABAAAD6AAAAAAAAAAAAAAAAAAAAAIAAAACAAAAADIEOuX5gnM87sJlN6EC4rwEs72XYc44STkOgvsU0g7vAAAAAACjL2YAAAADAA3fGQAAAAAAAAAAMgQ65fmCczzuwmU3oQLivASzvZdhzjhJOQ6C+xTSDu8AAAAXSHbkfAAIGHsAAAAJAAAAAwAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAQAAAAB3NZQAAAAAAAAAAAAAAAAAAAAAAAAAAAEADd8aAAAAAAAAAAAyBDrl+YJzPO7CZTehAuK8BLO9l2HOOEk5DoL7FNIO7wAAABeEEa58AAgYewAAAAkAAAACAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAABAAAAADuaygAAAAAAAAAAAAAAAAAAAAAAAAAAAwAN3xkAAAABAAAAADIEOuX5gnM87sJlN6EC4rwEs72XYc44STkOgvsU0g7vAAAAAUNPUAAAAAAAuXL+JNx4//zY7dpkCZ3yHKpSNgAA19zlF83T+KJsK1MAABI3mQjsAH//////////AAAAAQAAAAEAAAAAAAAAAAAAAdGpSiAAAAAAAAAAAAAAAAABAA3fGgAAAAEAAAAAMgQ65fmCczzuwmU3oQLivASzvZdhzjhJOQ6C+xTSDu8AAAABQ09QAAAAAAC5cv4k3Hj//Njt2mQJnfIcqlI2AADX3OUXzdP4omwrUwAAEU7EY9wAf/////////8AAAABAAAAAQAAAAAAAAAAAAAA6NSlEAAAAAAAAAAAAAAAAAMADd7UAAAAAQAAAAAB8GcCT7tUyYUBoLC0Vwc3CHskpHhl8xKodFvQB1rxQwAAAAFDT1AAAAAAALly/iTceP/82O3aZAmd8hyqUjYAANfc5RfN0/iibCtTAAAAAAAAAAB//////////wAAAAEAAAAAAAAAAAAAAAEADd8aAAAAAQAAAAAB8GcCT7tUyYUBoLC0Vwc3CHskpHhl8xKodFvQB1rxQwAAAAFDT1AAAAAAALly/iTceP/82O3aZAmd8hyqUjYAANfc5RfN0/iibCtTAAAA6NSlEAB//////////wAAAAEAAAAAAAAAAAAAAAMADd8aAAAAAAAAAAAB8GcCT7tUyYUBoLC0Vwc3CHskpHhl8xKodFvQB1rxQwAAABdIduc4AA3e0wAAAAIAAAABAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAEADd8aAAAAAAAAAAAB8GcCT7tUyYUBoLC0Vwc3CHskpHhl8xKodFvQB1rxQwAAABcM3B04AA3e0wAAAAIAAAABAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAA==",
			feeChangesXDR: "AAAAAgAAAAMADd7UAAAAAAAAAAAB8GcCT7tUyYUBoLC0Vwc3CHskpHhl8xKodFvQB1rxQwAAABdIduecAA3e0wAAAAEAAAABAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAEADd8aAAAAAAAAAAAB8GcCT7tUyYUBoLC0Vwc3CHskpHhl8xKodFvQB1rxQwAAABdIduc4AA3e0wAAAAEAAAABAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAA==",
			hash:          "e4b286344ae1c863ab15773ddf6649b08fe031383135194f8613a3a475c41a5a",
			index:         0,
			sequence:      56,
			expected: []map[string]interface{}{
				map[string]interface{}{
					"address": "GAA7AZYCJ65VJSMFAGQLBNCXA43QQ6ZEUR4GL4YSVB2FXUAHLLYUHIO5",
					"details": map[string]interface{}{
						"bought_amount":       "100000.0000000",
						"bought_asset_code":   "COP",
						"bought_asset_issuer": "GC4XF7RE3R4P77GY5XNGICM56IOKUURWAAANPXHFC7G5H6FCNQVVH3OH",
						"bought_asset_type":   "credit_alphanum4",
						"offer_id":            xdr.Int64(10694502),
						"seller":              "GAZAIOXF7GBHGPHOYJSTPIIC4K6AJM55S5Q44OCJHEHIF6YU2IHO6VHU",
						"sold_amount":         "100.0000000",
						"sold_asset_type":     "native",
					},
					"effectType":  history.EffectTrade,
					"operationID": int64(240518172673),
					"order":       uint32(1),
				},
				map[string]interface{}{
					"address": "GAZAIOXF7GBHGPHOYJSTPIIC4K6AJM55S5Q44OCJHEHIF6YU2IHO6VHU",
					"details": map[string]interface{}{
						"bought_amount":     "100.0000000",
						"bought_asset_type": "native",
						"offer_id":          xdr.Int64(10694502),
						"seller":            "GAA7AZYCJ65VJSMFAGQLBNCXA43QQ6ZEUR4GL4YSVB2FXUAHLLYUHIO5",
						"sold_amount":       "100000.0000000",
						"sold_asset_code":   "COP",
						"sold_asset_issuer": "GC4XF7RE3R4P77GY5XNGICM56IOKUURWAAANPXHFC7G5H6FCNQVVH3OH",
						"sold_asset_type":   "credit_alphanum4",
					},
					"effectType":  history.EffectTrade,
					"operationID": int64(240518172673),
					"order":       uint32(2),
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			tt := assert.New(t)
			transaction := BuildLedgerTransaction(
				t,
				TestTransaction{
					Index:         1,
					EnvelopeXDR:   tc.envelopeXDR,
					ResultXDR:     tc.resultXDR,
					MetaXDR:       tc.metaXDR,
					FeeChangesXDR: tc.feeChangesXDR,
					Hash:          tc.hash,
				},
			)

			operation := transactionOperationWrapper{
				index:          tc.index,
				transaction:    transaction,
				operation:      transaction.Envelope.Tx.Operations[tc.index],
				ledgerSequence: tc.sequence,
			}

			effects, err := operation.Effects()
			tt.NoError(err)
			tt.Equal(tc.expected, effects)
		})
	}
}
