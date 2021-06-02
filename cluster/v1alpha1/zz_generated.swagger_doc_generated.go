package v1alpha1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_Affinity = map[string]string{
	"":                    "Affinity defines the soft constraint on how cluster should be p",
	"clusterAffinity":     "ClusterAffinity defines that placement decisions are preferred to be put in the certain set of clusters. Each ClusterAffinity has a weight and the cluster with the maximum sum weight is selected at first.",
	"clusterAntiAffinity": "ClusterAntiAffinity defines that placement decisions are not preferred to be put in the certain set of clusters. Each ClusterAntiAffinity has a weight and the cluster with the maximum sum weight is selected at last.",
}

func (Affinity) SwaggerDoc() map[string]string {
	return map_Affinity
}

var map_ClusterAffinityTerm = map[string]string{
	"":                  "ClusterAffinityTerm defines an affinity terminology",
	"weight":            "Weight defines the importance of this affinity terminology",
	"whenUnsatisfiable": "WhenUnsatisfiable represents when clusterSelector cannot be met, what action should be taken to a cluster: DoNotSelect tells the placement controller not to select the cluster. It's a hard constraint. SelectAnyway tells the placement controller to still select it while prioritizing clusters. It's a soft constraint.",
}

func (ClusterAffinityTerm) SwaggerDoc() map[string]string {
	return map_ClusterAffinityTerm
}

var map_ClusterClaim = map[string]string{
	"":     "ClusterClaim represents cluster information that a managed cluster claims ClusterClaims with well known names include,\n  1. id.k8s.io, it contains a unique identifier for the cluster.\n  2. clusterset.k8s.io, it contains an identifier that relates the cluster\n     to the ClusterSet in which it belongs.\nClusterClaims created on a managed cluster will be collected and saved into the status of the corresponding ManagedCluster on hub.",
	"spec": "Spec defines the attributes of the ClusterClaim.",
}

func (ClusterClaim) SwaggerDoc() map[string]string {
	return map_ClusterClaim
}

var map_ClusterClaimList = map[string]string{
	"":         "ClusterClaimList is a collection of ClusterClaim.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
	"items":    "Items is a list of ClusterClaim.",
}

func (ClusterClaimList) SwaggerDoc() map[string]string {
	return map_ClusterClaimList
}

var map_ClusterClaimSelector = map[string]string{
	"":                 "ClusterClaimSelector is a claim query over a set of ManagedClusters. An empty cluster claim selector matches all objects. A null cluster claim selector matches no objects.",
	"matchExpressions": "matchExpressions is a list of cluster claim selector requirements. The requirements are ANDed.",
}

func (ClusterClaimSelector) SwaggerDoc() map[string]string {
	return map_ClusterClaimSelector
}

var map_ClusterClaimSpec = map[string]string{
	"value": "Value is a claim-dependent string",
}

func (ClusterClaimSpec) SwaggerDoc() map[string]string {
	return map_ClusterClaimSpec
}

var map_ClusterDecision = map[string]string{
	"":            "ClusterDecision represents a decision from a placement An empty ClusterDecision indicates it is not scheduled yet.",
	"clusterName": "ClusterName is the name of the ManagedCluster. If it is not empty, its value should be unique cross all placement decisions for the Placement.",
	"reason":      "Reason represents the reason why the ManagedCluster is selected.",
}

func (ClusterDecision) SwaggerDoc() map[string]string {
	return map_ClusterDecision
}

var map_ClusterPredicate = map[string]string{
	"":                        "ClusterPredicate represents a predicate to select ManagedClusters.",
	"requiredClusterSelector": "RequiredClusterSelector represents a selector of ManagedClusters by label and claim. If specified, 1) Any ManagedCluster, which does not match the selector, should not be selected by this ClusterPredicate; 2) If a selected ManagedCluster (of this ClusterPredicate) ceases to match the selector (e.g. due to\n   an update) of any ClusterPredicate, it will be eventually removed from the placement decisions;\n3) If a ManagedCluster (not selected previously) starts to match the selector, it will either\n   be selected or at least has a chance to be selected (when NumberOfClusters is specified);",
}

func (ClusterPredicate) SwaggerDoc() map[string]string {
	return map_ClusterPredicate
}

var map_ClusterSelector = map[string]string{
	"":              "ClusterSelector represents the AND of the containing selectors. An empty cluster selector matches all objects. A null cluster selector matches no objects.",
	"labelSelector": "LabelSelector represents a selector of ManagedClusters by label",
	"claimSelector": "ClaimSelector represents a selector of ManagedClusters by clusterClaims in status",
}

func (ClusterSelector) SwaggerDoc() map[string]string {
	return map_ClusterSelector
}

var map_ManagedClusterSet = map[string]string{
	"":       "ManagedClusterSet defines a group of ManagedClusters that user's workload can run on. A workload can be defined to deployed on a ManagedClusterSet, which mean:\n  1. The workload can run on any ManagedCluster in the ManagedClusterSet\n  2. The workload cannot run on any ManagedCluster outside the ManagedClusterSet\n  3. The service exposed by the workload can be shared in any ManagedCluster in the ManagedClusterSet\n\nIn order to assign a ManagedCluster to a certian ManagedClusterSet, add a label with name `cluster.open-cluster-management.io/clusterset` on the ManagedCluster to refers to the ManagedClusterSet. User is not allow to add/remove this label on a ManagedCluster unless they have a RBAC rule to CREATE on a virtual subresource of managedclustersets/join. In order to update this label, user must have the permission on both the old and new ManagedClusterSet.",
	"spec":   "Spec defines the attributes of the ManagedClusterSet",
	"status": "Status represents the current status of the ManagedClusterSet",
}

func (ManagedClusterSet) SwaggerDoc() map[string]string {
	return map_ManagedClusterSet
}

var map_ManagedClusterSetBinding = map[string]string{
	"":     "ManagedClusterSetBinding projects a ManagedClusterSet into a certain namespace. User is able to create a ManagedClusterSetBinding in a namespace and bind it to a ManagedClusterSet if they have an RBAC rule to CREATE on the virtual subresource of managedclustersets/bind. Workloads created in the same namespace can only be distributed to ManagedClusters in ManagedClusterSets bound in this namespace by higher level controllers.",
	"spec": "Spec defines the attributes of ManagedClusterSetBinding.",
}

func (ManagedClusterSetBinding) SwaggerDoc() map[string]string {
	return map_ManagedClusterSetBinding
}

var map_ManagedClusterSetBindingList = map[string]string{
	"":         "ManagedClusterSetBindingList is a collection of ManagedClusterSetBinding.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
	"items":    "Items is a list of ManagedClusterSetBinding.",
}

func (ManagedClusterSetBindingList) SwaggerDoc() map[string]string {
	return map_ManagedClusterSetBindingList
}

var map_ManagedClusterSetBindingSpec = map[string]string{
	"":           "ManagedClusterSetBindingSpec defines the attributes of ManagedClusterSetBinding.",
	"clusterSet": "ClusterSet is the name of the ManagedClusterSet to bind. It must match the instance name of the ManagedClusterSetBinding and cannot change once created. User is allowed to set this field if they have an RBAC rule to CREATE on the virtual subresource of managedclustersets/bind.",
}

func (ManagedClusterSetBindingSpec) SwaggerDoc() map[string]string {
	return map_ManagedClusterSetBindingSpec
}

var map_ManagedClusterSetList = map[string]string{
	"":         "ManagedClusterSetList is a collection of ManagedClusterSet.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
	"items":    "Items is a list of ManagedClusterSet.",
}

func (ManagedClusterSetList) SwaggerDoc() map[string]string {
	return map_ManagedClusterSetList
}

var map_ManagedClusterSetSpec = map[string]string{
	"": "ManagedClusterSetSpec describes the attributes of the ManagedClusterSet",
}

func (ManagedClusterSetSpec) SwaggerDoc() map[string]string {
	return map_ManagedClusterSetSpec
}

var map_ManagedClusterSetStatus = map[string]string{
	"":           "ManagedClusterSetStatus represents the current status of the ManagedClusterSet.",
	"conditions": "Conditions contains the different condition statuses for this ManagedClusterSet.",
}

func (ManagedClusterSetStatus) SwaggerDoc() map[string]string {
	return map_ManagedClusterSetStatus
}

var map_Placement = map[string]string{
	"":       "Placement defines a rule to select a set of ManagedClusters from the ManagedClusterSets bound to the placement namespace.\n\nHere is how the placement policy combines with other selection methods to determine a matching list of ManagedClusters: 1) Kubernetes clusters are registered with hub as cluster-scoped ManagedClusters; 2) ManagedClusters are organized into cluster-scoped ManagedClusterSets; 3) ManagedClusterSets are bound to workload namespaces; 4) Namespace-scoped Placements specify a slice of ManagedClusterSets which select a working set\n   of potential ManagedClusters;\n5) Then Placements subselect from that working set using label/claim selection.\n\nNo ManagedCluster will be selected if no ManagedClusterSet is bound to the placement namespace. User is able to bind a ManagedClusterSet to a namespace by creating a ManagedClusterSetBinding in that namespace if they have a RBAC rule to CREATE on the virtual subresource of `managedclustersets/bind`.\n\nA slice of PlacementDecisions with label cluster.open-cluster-management.io/placement={placement name} will be created to represent the ManagedClusters selected by this placement.\n\nIf a ManagedCluster is selected and added into the PlacementDecisions, other components may apply workload on it; once it is removed from the PlacementDecisions, the workload applied on this ManagedCluster should be evicted accordingly.",
	"spec":   "Spec defines the attributes of Placement.",
	"status": "Status represents the current status of the Placement",
}

func (Placement) SwaggerDoc() map[string]string {
	return map_Placement
}

var map_PlacementDecision = map[string]string{
	"":       "PlacementDecision indicates a decision from a placement PlacementDecision should has a label cluster.open-cluster-management.io/placement={placement name} to reference a certain placement.\n\nIf a placement has spec.numberOfClusters specified, the total number of decisions contained in status.decisions of PlacementDecisions should always be NumberOfClusters; otherwise, the total number of decisions should be the number of ManagedClusters which match the placement requirements.\n\nSome of the decisions might be empty when there are no enough ManagedClusters meet the placement requirements.",
	"status": "Status represents the current status of the PlacementDecision",
}

func (PlacementDecision) SwaggerDoc() map[string]string {
	return map_PlacementDecision
}

var map_PlacementDecisionList = map[string]string{
	"":         "ClusterDecisionList is a collection of PlacementDecision.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
	"items":    "Items is a list of PlacementDecision.",
}

func (PlacementDecisionList) SwaggerDoc() map[string]string {
	return map_PlacementDecisionList
}

var map_PlacementDecisionStatus = map[string]string{
	"":          "PlacementDecisionStatus represents the current status of the PlacementDecision.",
	"decisions": "Decisions is a slice of decisions according to a placement The number of decisions should not be larger than 100",
}

func (PlacementDecisionStatus) SwaggerDoc() map[string]string {
	return map_PlacementDecisionStatus
}

var map_PlacementList = map[string]string{
	"":         "PlacementList is a collection of Placements.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
	"items":    "Items is a list of Placements.",
}

func (PlacementList) SwaggerDoc() map[string]string {
	return map_PlacementList
}

var map_PlacementSpec = map[string]string{
	"":                 "PlacementSpec defines the attributes of Placement. An empty PlacementSpec selects all ManagedClusters from the ManagedClusterSets bound to the placement namespace. The containing fields are ANDed.",
	"clusterSets":      "ClusterSets represent the ManagedClusterSets from which the ManagedClusters are selected. If the slice is empty, ManagedClusters will be selected from the ManagedClusterSets bound to the placement namespace, otherwise ManagedClusters will be selected from the intersection of this slice and the ManagedClusterSets bound to the placement namespace.",
	"numberOfClusters": "NumberOfClusters represents the desired number of ManagedClusters to be selected which meet the placement requirements. 1) If not specified, all ManagedClusters which meet the placement requirements (including ClusterSets,\n   and Predicates) will be selected;\n2) Otherwise if the nubmer of ManagedClusters meet the placement requirements is larger than\n   NumberOfClusters, a random subset with desired number of ManagedClusters will be selected;\n3) If the nubmer of ManagedClusters meet the placement requirements is equal to NumberOfClusters,\n   all of them will be selected;\n4) If the nubmer of ManagedClusters meet the placement requirements is less than NumberOfClusters,\n   all of them will be selected, and the status of condition `PlacementConditionSatisfied` will be\n   set to false;",
	"predicates":       "Predicates represent a slice of predicates to select ManagedClusters. The predicates are ORed.",
	"affinity":         "Affinity defines the affinity of the placement decision with cluster.",
	"spreadPolicy":     "SpreadPolicy defines how placement decisions should be spreaded among the clusters in the scope of clustersets and predicates.",
}

func (PlacementSpec) SwaggerDoc() map[string]string {
	return map_PlacementSpec
}

var map_PlacementStatus = map[string]string{
	"numberOfSelectedClusters": "NumberOfSelectedClusters represents the number of selected ManagedClusters",
	"conditions":               "Conditions contains the different condition statuses for this Placement.",
}

func (PlacementStatus) SwaggerDoc() map[string]string {
	return map_PlacementStatus
}

var map_SpreadConstraintsTerm = map[string]string{
	"":                  "SpreadConstraintsTerm defines a terminology to spread placement decisions",
	"maxSkew":           "MaxSkew is the maximum degree of which the placement decision can be evenly distributed",
	"topologyKey":       "TopologyKey is either a label key or a cluster claim name of ManagedClusters",
	"topologyKeyType":   "TopologyKeyType indicates the type of TopologyKey. It could be Label or Claim.",
	"whenUnsatisfiable": "WhenUnsatisfiable represents when maxSkew cannot be met, what action should be taken to a cluster: DoNotSelect tells the placement controller not to select the cluster. It's a hard constraint. SelectAnyway tells the placement controller to still select it while prioritizing clusters. It's a soft constraint.",
}

func (SpreadConstraintsTerm) SwaggerDoc() map[string]string {
	return map_SpreadConstraintsTerm
}

var map_SpreadPolicy = map[string]string{
	"":                  "SpreadPolicy defines how placement decisions should be spreaded among the clusters in the scope of clustersets and predicates.",
	"spreadConstraints": "SpreadConstraints defines how placement decision should be distributed among a set of clusters.",
}

func (SpreadPolicy) SwaggerDoc() map[string]string {
	return map_SpreadPolicy
}

// AUTO-GENERATED FUNCTIONS END HERE
