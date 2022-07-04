package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ExposedSecretReportsCRName    = "exposedsecretreports.aquasecurity.github.io"
	ExposedSecretReportsCRVersion = "v1alpha1"
	ExposedSecretReportKind       = "ExposedSecretReport"
	ExposedSecretReportListKind   = "ExposedSecretReportList"
)

// ExposedSecretSummary is a summary of ExposedSecret counts grouped by Severity.
type ExposedSecretSummary struct {
	// CriticalCount is the number of exposed secrets with Critical Severity.
	//+kubebuilder:validation:Minimum=0
	CriticalCount int `json:"criticalCount"`

	// HighCount is the number of exposed secrets with High Severity.
	//+kubebuilder:validation:Minimum=0
	HighCount int `json:"highCount"`

	// MediumCount is the number of exposed secrets with Medium Severity.
	//+kubebuilder:validation:Minimum=0
	MediumCount int `json:"mediumCount"`

	// LowCount is the number of exposed secrets with Low Severity.
	//+kubebuilder:validation:Minimum=0
	LowCount int `json:"lowCount"`
}

// ExposedSecret is the spec for a exposed secret record.
type ExposedSecret struct {
	// Target is where the exposed secret was found.
	Target string `json:"target"`

	// RuleID is rule the identifier.
	RuleID string `json:"ruleID"`

	Title    string   `json:"title"`
	Category string   `json:"category"`	
	//+kubebuilder:validation:Enum={SeverityCritical,SeverityHigh,SeverityMedium,SeverityLow}
	Severity Severity `json:"severity"`
	Match    string   `json:"match"`
}

//+kubebuilder:object:root=true
//+kubebuilder:printcolumn:name="Repository",type=string,JSONPath=`.report.artifact.repository`
//+kubebuilder:printcolumn:name="Tag",type=string,JSONPath=`.report.artifact.tag`
//+kubebuilder:printcolumn:name="Scanner",type=string,JSONPath=`.report.scanner.name`
//+kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
//+kubebuilder:printcolumn:name="Critical",type=integer,JSONPath=`.report.summary.criticalCount`,priority=1
//+kubebuilder:printcolumn:name="High",type=integer,JSONPath=`.report.summary.highCount`,priority=1
//+kubebuilder:printcolumn:name="Medium",type=integer,JSONPath=`.report.summary.mediumCount`,priority=1
//+kubebuilder:printcolumn:name="Low",type=integer,JSONPath=`.report.summary.lowCount`,priority=1

// ExposedSecretReport is a specification for the ExposedSecretReport resource.
type ExposedSecretReport struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Report is the actual exposed secret report data.
	Report ExposedSecretReportData `json:"report"`
}

// ExposedSecretReportData is the spec for the exposed secret scan result.
//
// The spec follows the Pluggable Scanners API defined for Harbor.
// @see https://github.com/goharbor/pluggable-scanner-spec/blob/master/api/spec/scanner-adapter-openapi-v1.0.yaml
type ExposedSecretReportData struct {
	// UpdateTimestamp is a timestamp representing the server time in UTC when this report was updated.
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Format=date-time
	UpdateTimestamp metav1.Time `json:"updateTimestamp"`

	// Scanner is the scanner that generated this report.
	Scanner Scanner `json:"scanner"`

	// Registry is the registry the Artifact was pulled from.
	Registry Registry `json:"registry"`

	// Artifact is a container image scanned for exposed secrets.
	Artifact Artifact `json:"artifact"`

	// Summary is a summary of ExposedSecret counts grouped by Severity.
	Summary ExposedSecretSummary `json:"summary"`

	// Secrets is a list of passwords, api keys, tokens and others items found in the Artifact.
	Secrets []ExposedSecret `json:"secrets"`
}

//+kubebuilder:object:root=true

// ExposedSecretReportList is a list of ExposedSecretReport resources.
type ExposedSecretReportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ExposedSecretReport `json:"items"`
}
