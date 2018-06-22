package main

const (
	RDF  = "http://www.w3.org/1999/02/22-rdf-syntax-ns#"
	RDFS = "http://www.w3.org/2000/01/rdf-schema#"
	OWL  = "http://www.w3.org/2002/07/owl#"

	RDF_nil       = RDF + "nil"
	RDF_type      = RDF + "type"
	RDF_Statement = RDF + "Statement"

	RDFS_Datatype      = RDFS + "Datatype"
	RDFS_label         = RDFS + "label"
	RDFS_comment       = RDFS + "comment"
	RDFS_domain        = RDFS + "domain"
	RDFS_subClassOf    = RDFS + "subClassOf"
	RDFS_subPropertyOf = RDFS + "subPropertyOf"
	RDFS_range         = RDFS + "range"

	OWL_Ontology           = OWL + "Ontology"
	OWL_Class              = OWL + "Class"
	OWL_ObjectProperty     = OWL + "ObjectProperty"
	OWL_DatatypeProperty   = OWL + "DatatypeProperty"
	OWL_FunctionalProperty = OWL + "FunctionalProperty"
	OWL_DeprecatedProperty = OWL + "DeprecatedProperty"
	OWL_Restriction        = OWL + "Restriction"
	OWL_equivalentProperty = OWL + "equivalentProperty"
)
