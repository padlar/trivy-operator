package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ClusterComplianceDetailReportCRName = "clustercompliancedetailreports.aquasecurity.github.io"
)

//+kubebuilder:object:root=true
//+kubebuilder:printcolumn:name="Scanner",type=string,JSONPath=`.report.scanner.name`
//+kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
//+kubebuilder:printcolumn:name="Critical",type=integer,JSONPath=`.report.summary.criticalCount`,priority=1
//+kubebuilder:printcolumn:name="High",type=integer,JSONPath=`.report.summary.highCount`,priority=1
//+kubebuilder:printcolumn:name="Medium",type=integer,JSONPath=`.report.summary.mediumCount`,priority=1
//+kubebuilder:printcolumn:name="Low",type=integer,JSONPath=`.report.summary.lowCount`,priority=1

// ClusterComplianceDetailReport is a specification for the ClusterComplianceDetailReport resource.
type ClusterComplianceDetailReport struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Report            ClusterComplianceDetailReportData `json:"report"`
}

//+kubebuilder:object:root=true

// ClusterComplianceDetailReportList is a list of compliance kinds.
type ClusterComplianceDetailReportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []ClusterComplianceReport `json:"items"`
}

type ClusterComplianceDetailReportData struct {
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Format=date-time
	UpdateTimestamp metav1.Time              `json:"updateTimestamp"`
	Type            Compliance               `json:"type"`
	Summary         ClusterComplianceSummary `json:"summary"`
	ControlChecks   []ControlCheckDetails    `json:"controlCheck"`
}

// ControlCheckDetails provides the result of conducting a single audit step.
type ControlCheckDetails struct {
	ID                 string               `json:"id"`
	Name               string               `json:"name"`
	Description        string               `json:"description,omitempty"`
	//+kubebuilder:validation:Enum={SeverityCritical,SeverityHigh,SeverityMedium,SeverityLow}
	Severity           Severity             `json:"severity"`
	ScannerCheckResult []ScannerCheckResult `json:"checkResults"`
}

type ResultDetails struct {
	Name      string        `json:"name,omitempty"`
	Namespace string        `json:"namespace,omitempty"`
	Msg       string        `json:"msg"`
	Status    ControlStatus `json:"status"`
}

type ScannerCheckResult struct {
	ObjectType  string          `json:"objectType"`
	ID          string          `json:"id,omitempty"`
	Remediation string          `json:"remediation,omitempty"`
	Details     []ResultDetails `json:"details"`
}

// Compliance is the specs for a security assessment report.
type Compliance struct {
	// Name the name of the compliance report.
	Name string `json:"name"`
	// Description of the compliance report.
	Description string `json:"description"`

	// Version the compliance report.
	Version string `json:"version"`
}
