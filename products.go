package main

var productUrls = map[string]string{
	"Marketplace":               "/marketplace",
	"Billing":                   "/billing",
	"APIs & Services":           "/apis",
	"Support":                   "/support",
	"IAM & Admin":               "/iam-admin",
	"Getting started":           "/getting-started",
	"Security":                  "/security",
	"Compliance":                "/compliance",
	"Anthos":                    "/anthos",
	"Compute Engine":            "/compute",
	"Kubernetes Engine":         "/kubernetes",
	"VMware Engine":             "/gve",
	"Workload Manager":          "/workloads",
	"Batch":                     "/batch",
	"Cloud Run":                 "/run",
	"Cloud Functions":           "/functions",
	"App Engine":                "/appengine",
	"API Gateway":               "/api-gateway",
	"Endpoints":                 "/endpoints",
	"Filestore":                 "/filestore",
	"Cloud Storage":             "/storage",
	"Storage Transfer":          "/transfer",
	"AlloyDB":                   "/alloydb",
	"Bigtable":                  "/bigtable",
	"Database Migration":        "/dbmigration",
	"Datastore":                 "/datastore",
	"Firestore":                 "/firestore",
	"Memorystore":               "/memorystore",
	"Spanner":                   "/spanner",
	"SQL":                       "/sql",
	"Eventarc":                  "/eventarc",
	"Cloud Scheduler":           "/cloudscheduler",
	"Cloud Tasks":               "/cloudtasks",
	"Workflows":                 "/workflows",
	"Apigee API Management":     "/apigee",
	"Integration Connectors":    "/connectors",
	"Application Integration":   "/integrations",
	"VPC network":               "/networking",
	"Network services":          "/net-services",
	"Hybrid Connectivity":       "/hybrid",
	"Network Service Tiers":     "/net-tier",
	"Network Security":          "/net-security",
	"Network Intelligence":      "/net-intelligence",
	"Backup and DR":             "/backupdr",
	"Monitoring":                "/monitoring",
	"Debugger":                  "/debug",
	"Logging":                   "/logs",
	"Error Reporting":           "/errors",
	"Trace":                     "/traces",
	"Profiler":                  "/profiler",
	"Cloud Build":               "/cloud-build",
	"Cloud Deploy":              "/deploy",
	"Container Registry":        "/gcr",
	"Artifact Registry":         "/artifacts",
	"Deployment Manager":        "/dm",
	"Identity Platform":         "/customer-identity",
	"Service Catalog":           "/catalog",
	"Carbon Footprint":          "/carbon",
	"Cloud Shell Editor":        "/cloudshelleditor",
	"Cloud Workstations":        "/workstations",
	"Migration Center":          "/migration",
	"Composer":                  "/composer",
	"Dataproc":                  "/dataproc",
	"Pu":                        "/Sub /cloudpubsub",
	"Dataflow":                  "/dataflow",
	"Datastream":                "/datastream",
	"IoT Core":                  "/iot",
	"BigQuery":                  "/bigquery",
	"Looker":                    "/looker",
	"Dataplex":                  "/dataplex",
	"Data Catalog":              "/datacatalog",
	"Healthcare":                "/healthcare",
	"Financial Services":        "/financial_services",
	"Dataprep":                  "/dataprep",
	"Data Fusion":               "/data-fusion",
	"Life Sciences":             "/lifesciences",
	"Vertex AI":                 "/vertex-ai",
	"Gen App Builder":           "/gen-app-builder",
	"AI Platform":               "/ai-platform",
	"Document AI":               "/ai/document-ai",
	"Natural Language":          "/natural-language",
	"Discovery Engine":          "/ai/discovery",
	"Recommendations AI":        "/recommendation",
	"Retail":                    "/ai/retail",
	"Speech":                    "/speech",
	"Tables":                    "/automl-tables",
	"Translation":               "/translation",
	"Talent Solution":           "/talent-solution",
	"DocAI Warehouse":           "/ai/docai-warehouse",
	"Vertex AI Vision":          "/ai/vision-ai",
	"Enterprise KG":             "/ai/knowledge-graph",
	"Google Maps Platform":      "/google/maps-apis",
	"Immersive Stream":          "/immersive-stream/xr",
	"Google Workspace":          "/workspace-api",
	"Redis Enterprise":          "/marketplace/product/redis-marketplace-isaas/redis-enterprise-cloud-flexible-plan",
	"MongoDB Atlas":             "/marketplace/product/mongodb/mdb-atlas-self-service",
	"Cloud Volumes":             "/netapp/cloud-volumes",
	"Splunk Cloud":              "/marketplace/product/gcp-marketplace-lve-1/splunk-cloud",
	"Apache Kafka on Confluent": "/marketplace/product/confluent-prod/apache-kafka-on-confluent-cloud",
	"Databricks":                "/marketplace/product/databricks-prod/databricks",
	"Neo4j Aura":                "/marketplace/product/endpoints/prod.n4gcp.neo4j.io",
	"PowerScale":                "/dell-emc/cloud-onefs",
	"Elastic Cloud":             "/marketplace/product/elastic-prod/elastic-cloud",
	"Citrix DaaS":               "/marketplace/product/citrix-public/citrix-daas-for-google-cloud.endpoints.citrix-master-project.cloud.goog",
	"Appliances":                "/appliances",
	"Edge":                      "/distributed-cloud",
}