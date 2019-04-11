// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/openshift/aws-account-operator/pkg/apis/aws/v1alpha1.Account":            schema_pkg_apis_aws_v1alpha1_Account(ref),
		"github.com/openshift/aws-account-operator/pkg/apis/aws/v1alpha1.AccountClaim":       schema_pkg_apis_aws_v1alpha1_AccountClaim(ref),
		"github.com/openshift/aws-account-operator/pkg/apis/aws/v1alpha1.AccountClaimSpec":   schema_pkg_apis_aws_v1alpha1_AccountClaimSpec(ref),
		"github.com/openshift/aws-account-operator/pkg/apis/aws/v1alpha1.AccountClaimStatus": schema_pkg_apis_aws_v1alpha1_AccountClaimStatus(ref),
		"github.com/openshift/aws-account-operator/pkg/apis/aws/v1alpha1.AccountPool":        schema_pkg_apis_aws_v1alpha1_AccountPool(ref),
		"github.com/openshift/aws-account-operator/pkg/apis/aws/v1alpha1.AccountPoolSpec":    schema_pkg_apis_aws_v1alpha1_AccountPoolSpec(ref),
		"github.com/openshift/aws-account-operator/pkg/apis/aws/v1alpha1.AccountPoolStatus":  schema_pkg_apis_aws_v1alpha1_AccountPoolStatus(ref),
		"github.com/openshift/aws-account-operator/pkg/apis/aws/v1alpha1.AccountSpec":        schema_pkg_apis_aws_v1alpha1_AccountSpec(ref),
		"github.com/openshift/aws-account-operator/pkg/apis/aws/v1alpha1.AccountStatus":      schema_pkg_apis_aws_v1alpha1_AccountStatus(ref),
	}
}

func schema_pkg_apis_aws_v1alpha1_Account(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Account is the Schema for the accounts API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/aws-account-operator/pkg/apis/aws/v1alpha1.AccountSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/aws-account-operator/pkg/apis/aws/v1alpha1.AccountStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/aws-account-operator/pkg/apis/aws/v1alpha1.AccountSpec", "github.com/openshift/aws-account-operator/pkg/apis/aws/v1alpha1.AccountStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_aws_v1alpha1_AccountClaim(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AccountClaim is the Schema for the accountclaims API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/aws-account-operator/pkg/apis/aws/v1alpha1.AccountClaimSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/aws-account-operator/pkg/apis/aws/v1alpha1.AccountClaimStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/aws-account-operator/pkg/apis/aws/v1alpha1.AccountClaimSpec", "github.com/openshift/aws-account-operator/pkg/apis/aws/v1alpha1.AccountClaimStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_aws_v1alpha1_AccountClaimSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AccountClaimSpec defines the desired state of AccountClaim",
				Properties: map[string]spec.Schema{
					"legalEntity": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Ref:         ref("github.com/openshift/aws-account-operator/pkg/apis/aws/v1alpha1.LegalEntity"),
						},
					},
					"awsCredentialSecret": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/aws-account-operator/pkg/apis/aws/v1alpha1.AwsCredentialSecret"),
						},
					},
					"aws": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/aws-account-operator/pkg/apis/aws/v1alpha1.Aws"),
						},
					},
					"accountLink": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"legalEntity", "awsCredentialSecret", "aws", "accountLink"},
			},
		},
		Dependencies: []string{
			"github.com/openshift/aws-account-operator/pkg/apis/aws/v1alpha1.Aws", "github.com/openshift/aws-account-operator/pkg/apis/aws/v1alpha1.AwsCredentialSecret", "github.com/openshift/aws-account-operator/pkg/apis/aws/v1alpha1.LegalEntity"},
	}
}

func schema_pkg_apis_aws_v1alpha1_AccountClaimStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AccountClaimStatus defines the observed state of AccountClaim",
				Properties: map[string]spec.Schema{
					"billingAccountID": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"conditions": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/openshift/aws-account-operator/pkg/apis/aws/v1alpha1.AccountClaimCondition"),
									},
								},
							},
						},
					},
					"state": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"billingAccountID", "conditions", "state"},
			},
		},
		Dependencies: []string{
			"github.com/openshift/aws-account-operator/pkg/apis/aws/v1alpha1.AccountClaimCondition"},
	}
}

func schema_pkg_apis_aws_v1alpha1_AccountPool(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AccountPool is the Schema for the accountpools API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/aws-account-operator/pkg/apis/aws/v1alpha1.AccountPoolSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/aws-account-operator/pkg/apis/aws/v1alpha1.AccountPoolStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/aws-account-operator/pkg/apis/aws/v1alpha1.AccountPoolSpec", "github.com/openshift/aws-account-operator/pkg/apis/aws/v1alpha1.AccountPoolStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_aws_v1alpha1_AccountPoolSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AccountPoolSpec defines the desired state of AccountPool",
				Properties: map[string]spec.Schema{
					"poolSize": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
				},
				Required: []string{"poolSize"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_aws_v1alpha1_AccountPoolStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AccountPoolStatus defines the observed state of AccountPool",
				Properties: map[string]spec.Schema{
					"poolSize": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"unclaimedAccounts": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"claimedAccounts": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
				},
				Required: []string{"poolSize", "unclaimedAccounts", "claimedAccounts"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_aws_v1alpha1_AccountSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AccountSpec defines the desired state of Account",
				Properties: map[string]spec.Schema{
					"awsAccountID": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"iamUserSecret": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"claimLink": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"awsAccountID", "iamUserSecret"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_aws_v1alpha1_AccountStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AccountStatus defines the observed state of Account",
				Properties: map[string]spec.Schema{
					"claimed": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"state": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"claimed", "state"},
			},
		},
		Dependencies: []string{},
	}
}
