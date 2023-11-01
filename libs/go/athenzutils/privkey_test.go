package athenzutils

import (
	"crypto/x509"
	"testing"
)

var ecdsaPrivateKeyPEM = []byte(`-----BEGIN EC PRIVATE KEY-----
MIGkAgEBBDA27vlziu7AYNJo/aaG3mS4XPK2euiTLQDxzUoDkiMpVHRXLxSbX897
Gz7dQNFo3UWgBwYFK4EEACKhZANiAARBr6GWO6EGIV09DGInLfC/JSvPOKc26mZu
jpEdar4FkJ02OsHdtZ6AM7HgLASSBETL13Mhk8LL9qfRo+PEwLcyJnvWlDsMa3eh
Pji5iP4d9rQEOm/G9PXZ3/ZZEz5DuYs=
-----END EC PRIVATE KEY-----
`)

var rsaPrivateKeyPEM = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAxq83nCd8AqH5n40dEBMElbaJd2gFWu6bjhNzyp9562dpf454
BUSN0uF+g3i1yzcwdvADTiuExKN1u/IoGURxVCa0JTzAPJw6/JIoyOZnHZCoarcg
QQqZ56/udkSQ2NssrwGSQjOwxMrgIdH6XeLgGqVN4BoEEI+gpaQZa7rSytU5RFSG
OnZWO2Vwgs1OBxiOiYg1gzA1spJXQhxcBWw/v+YrUFtjxBKsG1UrWbnHbgciiN5U
2v51Yztjo8A1T+o9eIG90jVo3EhS2qhbzd8mLAsEhjV1sP8GItjfdfwXpXT7q2QG
99W3PM75+HdwGLvJIrkED7YRj4CpMkz6F1etawIDAQABAoIBAD67C7/N56WdJodt
soNkvcnXPEfrG+W9+Hc/RQvwljnxCKoxfUuMfYrbj2pLLnrfDfo/hYukyeKcCYwx
xN9VcMK1BaPMLpX0bdtY+m+T73KyPbqT3ycqBbXVImFM/L67VLxcrqUgVOuNcn67
IWWLQF6pWpErJaVk87/Ys/4DmpJXebLDyta8+ce6r0ppSG5+AifGo1byQT7kSJkF
lyQsyKWoVN+02s7gLsln5JXXZ672y2Xtp/S3wK0vfzy/HcGSxzn1yE0M5UJtDm/Y
qECnV1LQ0FB1l1a+/itHR8ipp5rScD4ZpzOPLKthglEvNPe4Lt5rieH9TR97siEe
SrC8uyECgYEA5Q/elOJAddpE+cO22gTFt973DcPGjM+FYwgdrora+RfEXJsMDoKW
AGSm5da7eFo8u/bJEvHSJdytc4CRQYnWNryIaUw2o/1LYXRvoEt1rEEgQ4pDkErR
PsVcVuc3UDeeGtYJwJLV6pjxO11nodFv4IgaVj64SqvCOApTTJgWXF0CgYEA3gzN
d3l376mSMuKc4Ep++TxybzA5mtF2qoXucZOon8EDJKr+vGQ9Z6X4YSdkSMNXqK1j
ILmFH7V3dyMOKRBA84YeawFacPLBJq+42t5Q1OYdcKZbaArlBT8ImGT7tQODs3JN
4w7DH+V1v/VCTl2zQaZRksb0lUsQbFiEfj+SVGcCgYAYIlDoTOJPyHyF+En2tJQE
aHiNObhcs6yxH3TJJBYoMonc2/UsPjQBvJkdFD/SUWeewkSzO0lR9etMhRpI1nX8
dGbG+WG0a4aasQLl162BRadZlmLB/DAJtg+hlGDukb2VxEFoyc/CFPUttQyrLv7j
oFNuDNOsAmbHMsdOBaQtfQKBgQCb/NRuRNebdj0tIALikZLHVc5yC6e7+b/qJPIP
uZIwv++MV89h2u1EHdTxszGA6DFxXnSPraQ2VU2aVPcCo9ds+9/sfePiCrbjjXhH
0PtpxEoUM9lsqpKeb9yC6hXk4JYpfnf2tQ0gIBrrAclVsf9WdBdEDB4Prs7Xvgs9
gT0zqwKBgQCzZubFO0oTYO9e2r8wxPPPsE3ZCjbP/y7lIoBbSzxDGUubXmbvD0GO
MC8dM80plsTym96UxpKkQMAglKKLPtG2n8xB8v5H/uIB4oIegMSEx3F7MRWWIQmR
Gea7bQ16YCzM/l2yygGhAW61bg2Z2GoVF6X5z/qhKGyo97V87qTbmg==
-----END RSA PRIVATE KEY-----
`)

// generated by:
// openssl genpkey -algorithm EC -pkeyopt ec_paramgen_curve:P-256 -pkeyopt ec_param_enc:named_curve
var ecdsaPrivateKeyPEMInPKCS8Format = []byte(`-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgTd/SNTDvns5/2c0i
5NWCXQrwC4OdfpjUU6WDCGEFUYmhRANCAAQKrQAm7ST0rDqki9G36DankKT+LM36
m/YZAaN5KC9KYjFFbT5EwGwRUiDkp7LwnrMb5ib4ppFHpbo0/kUtoxSN
-----END PRIVATE KEY-----
`)

// generated by:
// openssl genpkey -algorithm RSA -pkeyopt rsa_keygen_bits:2048 -pkeyopt rsa_keygen_pubexp:65537
var rsaPrivateKeyPEMInPKCS8Format = []byte(`-----BEGIN PRIVATE KEY-----
MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDltv1XRkGcVFdI
1uhAo52A68qwjNnia6yQLw53HgQsFbmG2Byxsqaeab0keHFD9BEdr6FF2BLeyh1H
pq+btxlnOXUgX3uxev4FFwx9Xjt+FARapl25J3Esy4eUSoqbo0ZBtn+HKlOchhHF
rc2ieL/YvtFvWov0uWhjZnf/ilFIRrGGTPBfcP/VMZTvAyx+yInEoxbyFTfSGwz3
qvl2zp6TNbYnzLpkqq5sa3+Y0IdPuMfVQcuolFEFUT2t8FdM6LqEPnN85uplAzXN
jY5Gmyl8UrL2PDgqlZyqA0eVOBcGrtzrJH6qx9PVx0osi/RT+lJL6HuOV1UMp154
NfASUDlzAgMBAAECggEAHe95V1DABxvUnhjnflj0ExNnQBey4zdN7yI6u9otCAOy
wDhUkPGrlfRIokKR3B2nx1sWZLAyUVc8dpRpyRyU0mdh9JyM1YWmKcqlpYbMsPLx
2FBa4WCa9o/1dKU8J+kgpDqgpuAksj6kfULXi+c5dQj06RJ/L56j1GRLmgEP85+H
jk4kx668TF81lt8/7toCUG9sK5gTq5DpJ4MpL7eowRnL4Z0GKrTK0AeVhulTnSSp
kkFKK6RER3orP3/owvAP8oKl3LzI5GvI+w4ayYeyxN4Ej0UIUms1GIQRH3VYM8um
zFVjBJySdn5qb2hb+fzsM/4PKJ8FTjwvSdPRal99dQKBgQD7ucCnC+2Nap8l++bq
Y7TYhNjmrO2KVJm5BqRXJZbzD4wYApyn8xpPsy9S+BxBvunEvYga+M/Xf0U/Wyb2
14y9Ov1+CMC/BbnXnzBNGJjiXCiWm8j4MKzPX84DpMAbhiWis6ACaZjO92hFbgfF
SL+/xkH3uYs2fKNm19s5qyQ+rQKBgQDpnY51ATeUy1DGUuZW4BK+QopWyYvgKI3Z
y+CuEi9eKA7X4SXTxjqkNedIry+46OGij6x9Dpg/zL54x8269P9XNVzTAEbpvHVB
BTilJEhR7fpgxvimbn7e1Vx/jvn+UMaJwWlB9/erE7rg5kyRpo+1yj1vBfc13D3k
l421F+f8nwKBgEgNgGaQVHvhJBLUSuGWjqJXTFqi7w9kbef3Tb0gJlgGgDwzKzIr
tMFRcd9W44eyJOnKspW92Ig/hsu+xKVtR3y20O5thPZopixhBYtb2g8ZAAk0KE9a
Z2yoaKjEVLTMLiOnNMrb/QBo8vDEsPa4fyJelm1ZL8712DPM35RfN221AoGAdtk2
CSZ2XVdWH587GcVjI7H8aQyeAYsAJ2ZGRqhvuqoMax1avjNhz/qwUFT3pU2sxKPt
L64GHKcP26hibJOJd5dpQtsoOG8tA8ghOjqMJEo6j5OKGjmqh7jqFubpHc2AQ8LG
xs3dDQa7kwD2wT6IbAaYXGwfiSIjxrCnYhLobacCgYA26nh4bjNBFHHcpD+l3dXo
G1ThRdxU6m6yyIUEwnLgqVkmijd2UauDUJmr8G3A/k5KwI0QX+MKi2LK16aIdCEl
r9OhcvHFPNh36qfXsP29VFaPktwfOmCeizwbZu/aWzgSCVA/Qe90XZBujK8rP7HE
Ja3LgyuTMIL7deAllV5wPg==
-----END PRIVATE KEY-----
`)

var unsupportedPrivateKeyPEM = []byte(`-----BEGIN UNSUPPORTED PRIVATE KEY-----
MIGkAgEBBDA27vlziu7AYNJo/aaG3mS4XPK2euiTLQDxzUoDkiMpVHRXLxSbX897
Gz7dQNFo3UWgBwYFK4EEACKhZANiAARBr6GWO6EGIV09DGInLfC/JSvPOKc26mZu
jpEdar4FkJ02OsHdtZ6AM7HgLASSBETL13Mhk8LL9qfRo+PEwLcyJnvWlDsMa3eh
Pji5iP4d9rQEOm/G9PXZ3/ZZEz5DuYs=
-----END UNSUPPORTED PRIVATE KEY-----
`)

func TestExtractSignerInfo(t *testing.T) {
	tests := []struct {
		name          string
		privateKeyPEM []byte
		algo          x509.SignatureAlgorithm
		wantErr       bool
	}{
		{"ecdsa PrivateKey PEM", ecdsaPrivateKeyPEM, x509.ECDSAWithSHA256, false},
		{"rsa PrivateKey PEM", rsaPrivateKeyPEM, x509.SHA256WithRSA, false},
		{"ecdsa PrivateKey PEM in PKCS#8 format", ecdsaPrivateKeyPEM, x509.ECDSAWithSHA256, false},
		{"rsa PrivateKey PEM in PKCS#8 format", rsaPrivateKeyPEM, x509.SHA256WithRSA, false},
		{"unsupported PEM", unsupportedPrivateKeyPEM, x509.UnknownSignatureAlgorithm, true},
		{"malformed PEM", []byte("abc"), x509.UnknownSignatureAlgorithm, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			signer, algorithm, err := ExtractSignerInfo(tt.privateKeyPEM)
			if algorithm != tt.algo {
				t.Errorf("invalid algorithm %s", algorithm)
			}
			if tt.wantErr {
				if err == nil {
					t.Errorf("no error, expected error")
				}
				if signer != nil {
					t.Errorf("non nil signer, expected nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error %v", err)
				}
				if signer == nil {
					t.Errorf("nil signer, expected non nil")
				}
			}
		})
	}
}
