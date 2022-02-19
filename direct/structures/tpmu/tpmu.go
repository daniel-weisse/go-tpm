// package tpmu contains TPM 2.0 structures prefixed with "TPMU_"
package tpmu

import "github.com/google/go-tpm/direct/structures/internal"

// HA represents a TPMU_HA.
// See definition in Part 2: Structures, section 10.3.1.
type HA = internal.TPMUHA

// Capabilities represents a TPMU_CAPABILITIES.
// See definition in Part 2: Structures, section 10.10.1.
type Capabilities = internal.TPMUCapabilities

// Attest represents a TPMU_ATTEST.
// See definition in Part 2: Structures, section 10.12.11.
type Attest = internal.TPMUAttest

// SymKeyBits represents a TPMU_SYM_KEY_BITS.
// See definition in Part 2: Structures, section 11.1.3.
type SymKeyBits = internal.TPMUSymKeyBits

// SymMode represents a TPMU_SYM_MODE.
// See definition in Part 2: Structures, section 11.1.4.
type SymMode = internal.TPMUSymMode

// SymDetails represents a TPMU_SYM_DETAILS.
// See definition in Part 2: Structures, section 11.1.5.
type SymDetails = internal.TPMUSymDetails

// SchemeKeyedHash represents a TPMU_SCHEME_KEYEDHASH.
// See definition in Part 2: Structures, section 11.1.22.
type SchemeKeyedHash = internal.TPMUSchemeKeyedHash

// SigScheme represents a TPMU_SIG_SCHEME.
// See definition in Part 2: Structures, section 11.2.1.4.
type SigScheme = internal.TPMUSigScheme

// KDFScheme represents a TPMU_KDF_SCHEME.
// See definition in Part 2: Structures, section 11.2.3.2.
type KDFScheme = internal.TPMUKDFScheme

// AsymScheme represents a TPMU_ASYM_SCHEME.
// See definition in Part 2: Structures, section 11.2.3.5.
type AsymScheme = internal.TPMUAsymScheme

// Signature represents a TPMU_SIGNATURE.
// See definition in Part 2: Structures, section 11.3.3.
type Signature = internal.TPMUSignature

// PublicID represents a TPMU_PUBLIC_ID.
// See definition in Part 2: Structures, section 12.2.3.2.
type PublicID = internal.TPMUPublicID

// PublicParms represents a TPMU_PUBLIC_PARMS.
// See definition in Part 2: Structures, section 12.2.3.7.
type PublicParms = internal.TPMUPublicParms
