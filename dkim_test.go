package dkim

import (
	"errors"
	"strings"
	"testing"
)

type fakeDnsClient struct {
	results map[string][]string
}

func (c *fakeDnsClient) LookupTxt(hostname string) ([]string, error) {
	if result, found := c.results[hostname]; found {
		return result, nil
	} else {
		return nil, errors.New("hostname not found")
	}
}

var client = &fakeDnsClient{
	results: map[string][]string{
		"google._domainkey.vandenhooff.name.": []string{
			`v=DKIM1; k=rsa; p=MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCl2Qrp5KF1uJnQSO0YuwInVPISQRrUciXtg/5hnQl6ed+UmYvWreLyuiyaiSd9X9Zu+aZQoeKm67HCxSMpC6G2ar0NludsXW69QdfzUpB5I6fzaLW8rl/RyeGkiQ3D66kvadK1wlNfUI7Dt9WtnUs8AFz/15xvODzgTMFJDiAcAwIDAQAB`,
		},
		"ginc1024._domainkey.yahoo-inc.com.": []string{"k=rsa; t=y; p=MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC1vQTaStWYO5FZFCi5ar2pczho8AzfbC9zrAWQrWE+bdarcgE/Zf1/e3BMXvb1Vht6W+eC3JghC1vfeH+Rf/qSWYC/GjhZZulGNjy1beyuu7rLtYJpG3V1Cuqj0zkjsYn+KqQ2KnubGu4KBD1PJlBGCnVCRvAll953Ucl7hJpmgQIDAQAB;"},
	},
}

func fixupNewlines(s string) string {
	return strings.Replace(s, "\n", "\r\n", -1)
}

var complete = fixupNewlines(`Received: by igcau2 with SMTP id au2so61978408igc.0
        for <1v443yp1p8@keytree.io>; Sun, 29 Mar 2015 19:39:21 -0700 (PDT)
DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed;
        d=vandenhooff.name; s=google;
        h=mime-version:from:date:message-id:subject:to:content-type;
        bh=47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=;
        b=NCOUEepJZ6cdKYtq61hifQ9K0fimliTNcDVDBQ8C1OQToNxNGQuGifUxWQ/6odRnmm
         +TGraJoXyKu2WwVl2auHW6Hug/9QBWg6JIQrUl3TLK5Z07IZHpqBFrXjqV/fd6Yl/1+L
         ZSaJ9lwo6YW6LvwoAq4AUwPDZqXeak7i5pj2U=
X-Google-DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed;
        d=1e100.net; s=20130820;
        h=x-gm-message-state:mime-version:from:date:message-id:subject:to
         :content-type;
        bh=47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=;
        b=mIJDzFjZy3jMNQQHSn7ADick4AjIHaACjpSCxUFbDvL2i7qhIq8SXSE5uOb8bW31tf
         qKL1xvrKq8vl/YymkSpTTsY+nrQ1DCcLH0sVLXWmw3AbiaXpViCFUKGMGaZyj12Xqe4x
         jZzBEIwOpN2z/f0QDvSyRb5gq+wBRIQkay6XEI2orDrP9SrfdhiMmwNaxtDBuWI6ollS
         X3vRh0zdZxTfYIBIzHZjmgn+gwUR2d/qk5sioT64JMwEvZjbWsUF2JC8Sim3tif1Z04L
         4JpItJhazY95XgZRaae25JvgCh9rtOE7WyHjHVhek/hy7SH1dZgxa9h2u7bjSwz2iHQt
         eUZA==
X-Gm-Message-State: ALoCoQk+KvRer9AfNQDS5M2p+aje/xg2vMBICDyzBfrFJKkaM7SLGYu5umi6GDbCSbE8AJPoKSgK
X-Received: by 10.107.148.198 with SMTP id w189mr46794537iod.14.1427683161411;
        Sun, 29 Mar 2015 19:39:21 -0700 (PDT)
Return-Path: <jelle@vandenhooff.name>
Received: from mail-ie0-f172.google.com (mail-ie0-f172.google.com. [209.85.223.172])
        by mx.google.com with ESMTPSA id s7sm6539499ioi.15.2015.03.29.19.39.19
        for <1v443yp1p8@keytree.io>
        (version=TLSv1.2 cipher=ECDHE-RSA-AES128-GCM-SHA256 bits=128/128);
        Sun, 29 Mar 2015 19:39:19 -0700 (PDT)
Received: by iedm5 with SMTP id m5so106639486ied.3
        for <1v443yp1p8@keytree.io>; Sun, 29 Mar 2015 19:39:19 -0700 (PDT)
X-Received: by 10.42.89.72 with SMTP id f8mr58735189icm.24.1427683158995; Sun,
 29 Mar 2015 19:39:18 -0700 (PDT)
MIME-Version: 1.0
Received: by 10.50.3.72 with HTTP; Sun, 29 Mar 2015 19:39:03 -0700 (PDT)
From: Jelle van den Hooff <jelle@vandenhooff.name>
Date: Sun, 29 Mar 2015 22:39:03 -0400
Message-ID: <CAP=Jqubpoizbfg+Fb_+ycEkhqrgMBE=qozKrRubUuimQ717wKw@mail.gmail.com>
Subject: vnsy7km1hn4crbyp0h32m3932p38qtgbhpxf9mp01s6w40mvk2jg
To: 1v443yp1p8@keytree.io
Content-Type: text/plain; charset=UTF-8


`)

var justSignature = fixupNewlines(`DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed;
        d=vandenhooff.name; s=google;
        h=mime-version:from:date:message-id:subject:to:content-type;
        bh=47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=;
        b=NCOUEepJZ6cdKYtq61hifQ9K0fimliTNcDVDBQ8C1OQToNxNGQuGifUxWQ/6odRnmm
         +TGraJoXyKu2WwVl2auHW6Hug/9QBWg6JIQrUl3TLK5Z07IZHpqBFrXjqV/fd6Yl/1+L
         ZSaJ9lwo6YW6LvwoAq4AUwPDZqXeak7i5pj2U=
`)

func TestComplete(t *testing.T) {
	var email *VerifiedEmail
	var err error

	if email, err = ParseAndVerify(complete, Complete, client); err != nil {
		t.Errorf("expected success; got %s", err)
	}

	if email.CanonHeaders() != headersOnly {
		t.Errorf("unexpected canonical form")
	}

	if email.Signature.Domain != "vandenhooff.name" {
		t.Errorf("expected vandenhooff.name as domain; got %s", email.Signature.Domain)
	}

	from := email.ExtractHeader("from")
	if len(from) != 1 || from[0] != "From: Jelle van den Hooff <jelle@vandenhooff.name>\r\n" {
		t.Errorf("strange from header")
	}

	withBrokenSignature := strings.Replace(complete, "NCOUEepJZ6cdKYtq61hifQ9K0fimliTNcDVDBQ8C1OQToNxNGQuGifUxWQ", "foobar", 1)
	if complete == withBrokenSignature {
		t.Fatalf("broken test; tried to kill signature but could not find it")
	}

	if _, err := ParseAndVerify(withBrokenSignature, Complete, client); err.Error() != "no valid DKIM signature" {
		t.Errorf("expected no valid DKIM signature; got %s", err)
	}

	if _, err := ParseAndVerify(complete+"foobar", Complete, client); err.Error() != "body hash does not match" {
		t.Errorf("expected failing body hash; got %s", err)
	}

	if _, err := ParseAndVerify(justSignature, Complete, client); err.Error() != "no valid DKIM signature" {
		t.Errorf("expected no valid DKIM signature; got %s", err)
	}

	if _, err := ParseAndVerify(justSignature+justSignature, Complete, client); err.Error() != "no valid DKIM signature" {
		t.Errorf("expected no valid DKIM signature; got %s", err)
	}

	if email, err = ParseAndVerify(justSignature+complete, Complete, client); err != nil {
		t.Errorf("expected success; got %s", err)
	}

	if _, err := ParseAndVerify("", Complete, client); err.Error() != "no DKIM header found" {
		t.Errorf("expected no DKIM header found; got %s", err)
	}
}

var headersOnly = fixupNewlines(`mime-version:1.0
from:Jelle van den Hooff <jelle@vandenhooff.name>
date:Sun, 29 Mar 2015 22:39:03 -0400
message-id:<CAP=Jqubpoizbfg+Fb_+ycEkhqrgMBE=qozKrRubUuimQ717wKw@mail.gmail.com>
subject:vnsy7km1hn4crbyp0h32m3932p38qtgbhpxf9mp01s6w40mvk2jg
to:1v443yp1p8@keytree.io
content-type:text/plain; charset=UTF-8
dkim-signature:v=1; a=rsa-sha256; c=relaxed/relaxed; d=vandenhooff.name; s=google; h=mime-version:from:date:message-id:subject:to:content-type; bh=47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=; b=NCOUEepJZ6cdKYtq61hifQ9K0fimliTNcDVDBQ8C1OQToNxNGQuGifUxWQ/6odRnmm+TGraJoXyKu2WwVl2auHW6Hug/9QBWg6JIQrUl3TLK5Z07IZHpqBFrXjqV/fd6Yl/1+LZSaJ9lwo6YW6LvwoAq4AUwPDZqXeak7i5pj2U=`)

func TestHeadersOnly(t *testing.T) {
	var email *VerifiedEmail
	var err error

	if email, err = ParseAndVerify(headersOnly, HeadersOnly, client); err != nil {
		t.Errorf("expected success; got %s", err)
	}

	if email, err = ParseAndVerify(headersOnly+"\r\n\r\n\r\nfoo bar", HeadersOnly, client); err != nil {
		t.Errorf("expected success; got %s", err)
	}

	if email.CanonHeaders() != headersOnly {
		t.Errorf("unexpected canonical form")
	}
}

var yahooIncDKIMtest = strings.Replace(`X-Apparently-To: andreser@yahoo-inc.com; Mon, 17 Aug 2015 22:49:25 +0000
Return-Path: <andreser@yahoo-inc.com>
Received-SPF: pass (domain of yahoo-inc.com designates 216.145.54.109 as permitted sender)
X-YMailISG: mjZhcf4WLDsyM65yWRfgyfO_lZT.dRW6ZkL0mQ36QKSZ1wt8
 norPyPfS_RaocAsatZMUc76bWB9uuFubtxIu.6wHOaop_IvkFzIxMpIj0qV.
 Lrx.L7iOLJ2Y5WVt6viLV7QS58O_2NzGwj3OIQL5EkGvSAZntHzX6fwew2_o
 mtpmgrO9DKSOmSxs0mI1hgXdqr2U2oqrtF9ibc4Z2cFMaZ4R1JeYcprQW9Xu
 X0YqkidSky.VEpst35uNTE.OMGZrIFHPzaKfF5GarnIJGSqhk.5NMjq_Bywg
 5LYpX9AoXaCFOQd0Tzp4raM0IUmhBRaGPPXUBbzqovVvuLdJ.clh6.kYtv_F
 5aNQtHP5cNqhPTooi1c_mZlh6phP12PMUVdx9WdfEmvVaN1Jumay.SzOtTPh
 89IA7pgAferCuLh5f_9lEkYLkFomW4SRwexAbpdfwm1R1CYprsZMQ1YhFZI3
 GinHyEiPUo48hxgTJgWIuv0oiCoDzd8exD5.u0ZW6Ztvy3UVvogbGCJ6KvXy
 7CT1iwdHcoCiGcoE9e7zEqZdH7GftkZGobaX83r3bzhhc0GVMmY29fB4BnZj
 suHtpK.Cx7vY.hJvV_R_.QH5npxcM8ptVFLgkNW6tBzqF9GnbWtr7v2ERGjn
 hewHjiEQAGbay6c19tw.3s0SEEhb0BdbxeGajeqNJhYLC8j18hRQR67oWyFF
 LON7S1cfRM2sQKVWW4K0I7KMad7FrxEi6VJdfIVD8gLMW7uhlkowqOE9rhtj
 042FEnYc7kcrvL58Bj8v9TY3Z2Nl8HXifr6dGK_Kw9HK79We3O00cdZSWASu
 R8pA_AB40d80d82.0crHu0oFFX6KFT8xkAipyIvPhK4bZz7r.NnBD1ZKq7ZF
 TpBmxt0hbxWy_Qkz1M9BrzrGbbeSAFhAyyZqoPYsWy8FN5U3jzU.ZQygaK.E
 DT18hIHBF2qN4R3JLVxA7zX1OfxL24UlPvuPaAERm9Wq4WRagcK7ysJt7.9b
 WskH.vySl_.3mtF7yBFXOR_7_aIM54djcILP_MGhqEjJVbPp12KmbQ51cD_o
 76mHVraxIkOZV0eVal8V9QwIaAbbb9caFJcySJdUSIVvojxd6fneN83jCsD.
 Df0Iz4J2pF0BYiVnY4.MIhUSZZtCjxBAK4roNSvdyVDEQdYPiJpQHoBIDCnj
 mgQCRPWOCRXaaDMnFvBzJJ7_z04R5rB3vFg65xsBN0wyeDX1veLLsMHChAbp
 8RPEQFrsqmFFrXXRODtacXX1ZOZV1tTI
X-Originating-IP: [216.145.54.109]
Authentication-Results: mta2007.corp.mail.ne1.yahoo.com  from=yahoo-inc.com; domainkeys=neutral (no sig);  from=yahoo-inc.com; dkim=pass (ok)
Received: from 127.0.0.1  (EHLO mrout4.yahoo.com) (216.145.54.109)
  by mta2007.corp.mail.ne1.yahoo.com with SMTPS; Mon, 17 Aug 2015 22:49:25 +0000
Received: from omp1017.mail.ne1.yahoo.com (omp1017.mail.ne1.yahoo.com [98.138.89.161])
	by mrout4.yahoo.com (8.14.9/8.14.9/y.out) with ESMTP id t7HMn6pT004450
	(version=TLSv1/SSLv3 cipher=DHE-RSA-CAMELLIA256-SHA bits=256 verify=NO)
	for <andreser@yahoo-inc.com>; Mon, 17 Aug 2015 15:49:06 -0700 (PDT)
Received: (qmail 15334 invoked by uid 1000); 17 Aug 2015 22:49:06 -0000
DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed; d=yahoo-inc.com; s=ginc1024; t=1439851746; bh=Zg0pSZvCcMHE9S9qpkoEKeacBIM4T3Xu4TUSMEL4rXw=; h=Date:From:Reply-To:To:Message-ID:Subject:MIME-Version:Content-Type; b=Ut+tXUluIOFrGnFm6m0fvXuIQDIEulXFkWmj9bQSO0JN3gPiWfuh1bFhZBdnu2C4SREtTfrxHI8q5DGPjD8yg4LnxFh3HOuaf4Ttm8w72QGO1HxJCdwkNvu5W4mnFTEB8hdl2u5naE4JqjJtM291ZYIJGvxFA2J3+Snj/N2aG40=
X-YMail-OSG: G1B4VdwVM1lYA9kmxoxrGwEODiHeae6vbYVeBm754R2VWrC5KBM9pyd4ojSurOA
 q0um_rXRvGr1aqpHntt5GL5mcITy4qZFZWIBKRlGdOvQKNsKMSzsglbrG0Io._.0dI8XBQ.DNWG3
 Z5uVt9prZqJLlJG.FcGrNnYQTiX.Q0HDTID4rDKM.sA6Z_CUAPOto0IFqnA9buS5R8Rjy3xqs5qf
 krxUdQCFbVG.ML8Kl0WJfy8ZKxjg1mT7Nma.ZOA--
Received: by 98.138.105.251; Mon, 17 Aug 2015 22:49:05 +0000 
Date: Mon, 17 Aug 2015 22:49:05 +0000 (UTC)
From: Andres Erbsen <andreser@yahoo-inc.com>
Reply-To: Andres Erbsen <andreser@yahoo-inc.com>
To: Andres Erbsen Erbsen <andreser@yahoo-inc.com>
Message-ID: <408588803.6263873.1439851745104.JavaMail.yahoo@mail.yahoo.com>
Subject: end-to-end public key verification [test]
MIME-Version: 1.0
Content-Type: multipart/alternative; 
	boundary="----=_Part_6263872_19047179.1439851745102"
Content-Length: 622

------=_Part_6263872_19047179.1439851745102
Content-Type: text/plain; charset=UTF-8
Content-Transfer-Encoding: 7bit

fdsfasdfdasgawgasdgdsgadfgadsgdgadga
------=_Part_6263872_19047179.1439851745102
Content-Type: text/html; charset=UTF-8
Content-Transfer-Encoding: 7bit

<html><body><div style="color:#000; background-color:#fff; font-family:HelveticaNeue-Light, Helvetica Neue Light, Helvetica Neue, Helvetica, Arial, Lucida Grande, sans-serif;font-size:16px"><div id="yui_3_16_0_1_1439835732243_26145" dir="ltr">fdsfasdfdasgawgasdgdsgadfgadsgdgadga</div></div></body></html>
------=_Part_6263872_19047179.1439851745102--`, "\n", "\r\n", -1)

func TestYahooIncDKIM(t *testing.T) {
	_, err := ParseAndVerify(yahooIncDKIMtest, HeadersOnly, client)
	if err != nil {
		t.Fatal(err)
	}
}
