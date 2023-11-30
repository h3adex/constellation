/*
Copyright (c) Edgeless Systems GmbH

SPDX-License-Identifier: AGPL-3.0-only
*/

// Package testdata contains testing data for an attestation process.
package testdata

import _ "embed"

// AttestationReport is an example attestation report from a Constellation VM.
//
//go:embed attestation.bin
var AttestationReport []byte

// AttestationReportVLEK is an example attestation report signed by a VLEK.
const AttestationReportVLEK = "02000000000000000000030000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000300000000000ace0300000000000000040000000000000044a93ab043ad14ece9bfa97305d95302c9cc6ed95e17efaf7348ed7a7603e1ca89d12758e089d2abcf5a4dd16a99e3cb4cba8f0b8e8cb8eac3e926f1d2b5cfecc2c84b9364fc9f0f54b04534768c860c6e0e386ad98b96e8b98eca46ac8971d05c531ba48373f054c880cfd1f4a0a84e00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008c5d6770df734a203cd061a3698e702caed25e7f744dc060eb9dcba0f2e4bdb2ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0300000000000a73000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000300000000000a7301360100013601000300000000000a73000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000b9853dac65f127574c6a578c11885e1887d4c7ae446237d4273715dd8c05cfe4bd49facc1392f2ca7354c8f0d34d65500000000000000000000000000000000000000000000000004013481e9c6a6bb112818aeba3bd178d788dedf62600b8c7892a8d3df4d880265010e7d833201156364a001e62f47b570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"

// AzureThimVCEK is an example VCEK certificate (PEM, as returned from Azure THIM) for the AttestationReport.
//
//go:embed vcek.pem
var AzureThimVCEK []byte

// AmdKdsVCEK is an example VCEK certificate (DER, as returned from AMD KDS) for the AttestationReport.
//
//go:embed vcek.cert
var AmdKdsVCEK []byte

// RuntimeData is an example runtime data from the TPM for the AttestationReport.
//
//go:embed runtimedata.bin
var RuntimeData []byte

// CertChain is a valid certificate chain (PEM, as returned from Azure THIM) for the VCEK certificate.
//
//go:embed certchain.pem
var CertChain []byte

// VlekCertChain is a valid certificate chain (PEM) for the VLEK certificate.
//
//go:embed vlekcertchain.pem
var VlekCertChain []byte

// Vlek is a valid VLEK certificate (PEM).
//
//go:embed vlek.pem
var Vlek []byte