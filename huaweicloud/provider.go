package huaweicloud

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/aad"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/antiddos"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/aom"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/apig"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/apm"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/as"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/bms"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/cbh"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/cbr"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/cc"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/cce"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/cci"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/cdm"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/cdn"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/ces"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/cfw"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/cloudtable"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/cmdb"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/codehub"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/cph"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/cpts"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/cse"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/css"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/cts"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/dataarts"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/dbss"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/dc"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/dcs"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/ddm"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/dds"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/deprecated"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/dew"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/dis"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/dli"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/dms"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/drs"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/dsc"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/dws"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/ecs"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/eip"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/elb"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/eps"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/er"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/evs"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/fgs"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/ga"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/gaussdb"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/hss"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/iam"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/ims"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/iotda"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/lb"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/live"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/lts"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/meeting"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/modelarts"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/mpc"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/mrs"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/nat"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/obs"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/oms"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/projectman"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/rds"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/rf"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/rms"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/scm"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/servicestage"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/servicestagev3"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/sfs"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/smn"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/sms"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/swr"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/tms"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/vod"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/vpc"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/vpcep"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/vpn"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/waf"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/workspace"
)

const (
	defaultCloud       string = "myhuaweicloud.com"
	defaultEuropeCloud string = "myhuaweicloud.eu"
)

// Provider returns a schema.Provider for HuaweiCloud.
func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"region": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  descriptions["region"],
				InputDefault: "cn-north-1",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HW_REGION_NAME",
					"OS_REGION_NAME",
				}, nil),
			},

			"access_key": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  descriptions["access_key"],
				RequiredWith: []string{"secret_key"},
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HW_ACCESS_KEY",
					"OS_ACCESS_KEY",
				}, nil),
			},

			"secret_key": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  descriptions["secret_key"],
				RequiredWith: []string{"access_key"},
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HW_SECRET_KEY",
					"OS_SECRET_KEY",
				}, nil),
			},

			"security_token": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  descriptions["security_token"],
				RequiredWith: []string{"access_key"},
				DefaultFunc:  schema.EnvDefaultFunc("HW_SECURITY_TOKEN", nil),
			},

			"domain_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["domain_id"],
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HW_DOMAIN_ID",
					"OS_DOMAIN_ID",
					"OS_USER_DOMAIN_ID",
					"OS_PROJECT_DOMAIN_ID",
				}, ""),
			},

			"domain_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["domain_name"],
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HW_DOMAIN_NAME",
					"OS_DOMAIN_NAME",
					"OS_USER_DOMAIN_NAME",
					"OS_PROJECT_DOMAIN_NAME",
				}, ""),
			},

			"user_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["user_name"],
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HW_USER_NAME",
					"OS_USERNAME",
				}, ""),
			},

			"user_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["user_id"],
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HW_USER_ID",
					"OS_USER_ID",
				}, ""),
			},

			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: descriptions["password"],
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HW_USER_PASSWORD",
					"OS_PASSWORD",
				}, ""),
			},

			"assume_role": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"agency_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: descriptions["assume_role_agency_name"],
							DefaultFunc: schema.EnvDefaultFunc("HW_ASSUME_ROLE_AGENCY_NAME", nil),
						},
						"domain_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: descriptions["assume_role_domain_name"],
							DefaultFunc: schema.EnvDefaultFunc("HW_ASSUME_ROLE_DOMAIN_NAME", nil),
						},
					},
				},
			},

			"project_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["project_id"],
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HW_PROJECT_ID",
					"OS_PROJECT_ID",
				}, nil),
			},

			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["project_name"],
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HW_PROJECT_NAME",
					"OS_PROJECT_NAME",
				}, nil),
			},

			"tenant_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["tenant_id"],
				DefaultFunc: schema.EnvDefaultFunc("OS_TENANT_ID", ""),
			},

			"tenant_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["tenant_name"],
				DefaultFunc: schema.EnvDefaultFunc("OS_TENANT_NAME", ""),
			},

			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["token"],
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HW_AUTH_TOKEN",
					"OS_AUTH_TOKEN",
				}, ""),
			},

			"insecure": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: descriptions["insecure"],
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HW_INSECURE",
					"OS_INSECURE",
				}, false),
			},

			"cacert_file": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_CACERT", ""),
				Description: descriptions["cacert_file"],
			},

			"cert": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_CERT", ""),
				Description: descriptions["cert"],
			},

			"key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_KEY", ""),
				Description: descriptions["key"],
			},

			"agency_name": {
				Type:         schema.TypeString,
				Optional:     true,
				DefaultFunc:  schema.EnvDefaultFunc("OS_AGENCY_NAME", nil),
				Description:  descriptions["agency_name"],
				RequiredWith: []string{"agency_domain_name"},
			},

			"agency_domain_name": {
				Type:         schema.TypeString,
				Optional:     true,
				DefaultFunc:  schema.EnvDefaultFunc("OS_AGENCY_DOMAIN_NAME", nil),
				Description:  descriptions["agency_domain_name"],
				RequiredWith: []string{"agency_name"},
			},

			"delegated_project": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_DELEGATED_PROJECT", ""),
				Description: descriptions["delegated_project"],
			},

			"auth_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["auth_url"],
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HW_AUTH_URL",
					"OS_AUTH_URL",
				}, nil),
			},

			"cloud": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["cloud"],
				DefaultFunc: schema.EnvDefaultFunc("HW_CLOUD", defaultCloud),
			},

			"endpoints": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: descriptions["endpoints"],
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"shared_config_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["shared_config_file"],
				DefaultFunc: schema.EnvDefaultFunc("HW_SHARED_CONFIG_FILE", ""),
			},

			"profile": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["profile"],
				DefaultFunc: schema.EnvDefaultFunc("HW_PROFILE", ""),
			},

			"enterprise_project_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["enterprise_project_id"],
				DefaultFunc: schema.EnvDefaultFunc("HW_ENTERPRISE_PROJECT_ID", ""),
			},

			"max_retries": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: descriptions["max_retries"],
				DefaultFunc: schema.EnvDefaultFunc("HW_MAX_RETRIES", 5),
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"huaweicloud_antiddos": dataSourceAntiDdosV1(),

			"huaweicloud_apig_environments": apig.DataSourceEnvironments(),

			"huaweicloud_as_configurations": as.DataSourceASConfigurations(),
			"huaweicloud_as_groups":         as.DataSourceASGroups(),

			"huaweicloud_availability_zones": DataSourceAvailabilityZones(),

			"huaweicloud_bms_flavors": bms.DataSourceBmsFlavors(),
			"huaweicloud_cbr_vaults":  cbr.DataSourceCbrVaultsV3(),

			"huaweicloud_cbh_instances": cbh.DataSourceCbhInstances(),

			"huaweicloud_cce_addon_template": cce.DataSourceCCEAddonTemplateV3(),
			"huaweicloud_cce_cluster":        cce.DataSourceCCEClusterV3(),
			"huaweicloud_cce_clusters":       cce.DataSourceCCEClusters(),
			"huaweicloud_cce_node":           cce.DataSourceCCENodeV3(),
			"huaweicloud_cce_nodes":          cce.DataSourceCCENodes(),
			"huaweicloud_cce_node_pool":      cce.DataSourceCCENodePoolV3(),
			"huaweicloud_cci_namespaces":     cci.DataSourceCciNamespaces(),

			"huaweicloud_cdm_flavors": DataSourceCdmFlavorV1(),

			"huaweicloud_cdn_domain_statistics": cdn.DataSourceStatistics(),

			"huaweicloud_cfw_firewalls":        cfw.DataSourceFirewalls(),
			"huaweicloud_compute_flavors":      ecs.DataSourceEcsFlavors(),
			"huaweicloud_compute_instance":     ecs.DataSourceComputeInstance(),
			"huaweicloud_compute_instances":    ecs.DataSourceComputeInstances(),
			"huaweicloud_compute_servergroups": ecs.DataSourceComputeServerGroups(),

			"huaweicloud_cph_server_flavors": cph.DataSourceServerFlavors(),
			"huaweicloud_cph_phone_flavors":  cph.DataSourcePhoneFlavors(),
			"huaweicloud_cph_phone_images":   cph.DataSourcePhoneImages(),

			"huaweicloud_csbs_backup":        dataSourceCSBSBackupV1(),
			"huaweicloud_csbs_backup_policy": dataSourceCSBSBackupPolicyV1(),

			"huaweicloud_csms_secret_version": dew.DataSourceDewCsmsSecret(),
			"huaweicloud_css_flavors":         css.DataSourceCssFlavors(),

			"huaweicloud_dcs_flavors":        dcs.DataSourceDcsFlavorsV2(),
			"huaweicloud_dcs_maintainwindow": dcs.DataSourceDcsMaintainWindow(),
			"huaweicloud_dcs_instances":      dcs.DataSourceDcsInstance(),

			"huaweicloud_dds_flavors":   dds.DataSourceDDSFlavorV3(),
			"huaweicloud_dds_instances": dds.DataSourceDdsInstance(),

			"huaweicloud_dms_kafka_flavors":   dms.DataSourceKafkaFlavors(),
			"huaweicloud_dms_kafka_instances": dms.DataSourceDmsKafkaInstances(),
			"huaweicloud_dms_product":         dms.DataSourceDmsProduct(),
			"huaweicloud_dms_maintainwindow":  dms.DataSourceDmsMaintainWindow(),

			"huaweicloud_dms_rocketmq_broker":    dms.DataSourceDmsRocketMQBroker(),
			"huaweicloud_dms_rocketmq_instances": dms.DataSourceDmsRocketMQInstances(),

			"huaweicloud_enterprise_project": eps.DataSourceEnterpriseProject(),

			"huaweicloud_er_route_tables": er.DataSourceRouteTables(),

			"huaweicloud_evs_volumes":      evs.DataSourceEvsVolumesV2(),
			"huaweicloud_fgs_dependencies": fgs.DataSourceFunctionGraphDependencies(),

			"huaweicloud_gaussdb_cassandra_dedicated_resource": gaussdb.DataSourceGeminiDBDehResource(),
			"huaweicloud_gaussdb_cassandra_flavors":            gaussdb.DataSourceCassandraFlavors(),
			"huaweicloud_gaussdb_nosql_flavors":                gaussdb.DataSourceGaussDBNoSQLFlavors(),
			"huaweicloud_gaussdb_cassandra_instance":           gaussdb.DataSourceGeminiDBInstance(),
			"huaweicloud_gaussdb_cassandra_instances":          gaussdb.DataSourceGeminiDBInstances(),
			"huaweicloud_gaussdb_opengauss_instance":           gaussdb.DataSourceOpenGaussInstance(),
			"huaweicloud_gaussdb_opengauss_instances":          gaussdb.DataSourceOpenGaussInstances(),
			"huaweicloud_gaussdb_mysql_configuration":          gaussdb.DataSourceGaussdbMysqlConfigurations(),
			"huaweicloud_gaussdb_mysql_dedicated_resource":     gaussdb.DataSourceGaussDBMysqlDehResource(),
			"huaweicloud_gaussdb_mysql_flavors":                gaussdb.DataSourceGaussdbMysqlFlavors(),
			"huaweicloud_gaussdb_mysql_instance":               gaussdb.DataSourceGaussDBMysqlInstance(),
			"huaweicloud_gaussdb_mysql_instances":              gaussdb.DataSourceGaussDBMysqlInstances(),
			"huaweicloud_gaussdb_redis_instance":               gaussdb.DataSourceGaussRedisInstance(),

			"huaweicloud_identity_role":        iam.DataSourceIdentityRoleV3(),
			"huaweicloud_identity_custom_role": iam.DataSourceIdentityCustomRole(),
			"huaweicloud_identity_group":       iam.DataSourceIdentityGroup(),
			"huaweicloud_identity_projects":    iam.DataSourceIdentityProjects(),
			"huaweicloud_identity_users":       iam.DataSourceIdentityUsers(),

			"huaweicloud_iec_bandwidths":     dataSourceIECBandWidths(),
			"huaweicloud_iec_eips":           dataSourceIECNetworkEips(),
			"huaweicloud_iec_flavors":        dataSourceIecFlavors(),
			"huaweicloud_iec_images":         dataSourceIecImages(),
			"huaweicloud_iec_keypair":        dataSourceIECKeypair(),
			"huaweicloud_iec_network_acl":    dataSourceIECNetworkACL(),
			"huaweicloud_iec_port":           DataSourceIECPort(),
			"huaweicloud_iec_security_group": dataSourceIECSecurityGroup(),
			"huaweicloud_iec_server":         dataSourceIECServer(),
			"huaweicloud_iec_sites":          dataSourceIecSites(),
			"huaweicloud_iec_vpc":            DataSourceIECVpc(),
			"huaweicloud_iec_vpc_subnets":    DataSourceIECVpcSubnets(),

			"huaweicloud_images_image":  ims.DataSourceImagesImageV2(),
			"huaweicloud_images_images": ims.DataSourceImagesImages(),

			"huaweicloud_kms_key":      dew.DataSourceKmsKey(),
			"huaweicloud_kms_data_key": DataSourceKmsDataKeyV1(),
			"huaweicloud_kps_keypairs": dew.DataSourceKeypairs(),

			"huaweicloud_lb_listeners":    lb.DataSourceListeners(),
			"huaweicloud_lb_loadbalancer": lb.DataSourceELBV2Loadbalancer(),
			"huaweicloud_lb_certificate":  lb.DataSourceLBCertificateV2(),
			"huaweicloud_lb_pools":        lb.DataSourcePools(),

			"huaweicloud_elb_certificate": elb.DataSourceELBCertificateV3(),
			"huaweicloud_elb_flavors":     elb.DataSourceElbFlavorsV3(),
			"huaweicloud_elb_pools":       elb.DataSourcePools(),

			"huaweicloud_nat_gateway": nat.DataSourcePublicGateway(),

			"huaweicloud_networking_port":      vpc.DataSourceNetworkingPortV2(),
			"huaweicloud_networking_secgroup":  DataSourceNetworkingSecGroup(),
			"huaweicloud_networking_secgroups": vpc.DataSourceNetworkingSecGroups(),

			"huaweicloud_modelarts_datasets":         modelarts.DataSourceDatasets(),
			"huaweicloud_modelarts_dataset_versions": modelarts.DataSourceDatasetVerions(),
			"huaweicloud_modelarts_notebook_images":  modelarts.DataSourceNotebookImages(),

			"huaweicloud_obs_buckets":       obs.DataSourceObsBuckets(),
			"huaweicloud_obs_bucket_object": obs.DataSourceObsBucketObject(),

			"huaweicloud_rds_flavors":         rds.DataSourceRdsFlavor(),
			"huaweicloud_rds_engine_versions": rds.DataSourceRdsEngineVersionsV3(),
			"huaweicloud_rds_instances":       rds.DataSourceRdsInstances(),
			"huaweicloud_rds_backups":         rds.DataSourceBackup(),
			"huaweicloud_rds_storage_types":   rds.DataSourceStoragetype(),

			"huaweicloud_rms_policy_definitions": rms.DataSourcePolicyDefinitions(),

			"huaweicloud_servicestage_component_runtimes": servicestage.DataSourceComponentRuntimes(),

			"huaweicloud_smn_topics": smn.DataSourceTopics(),

			"huaweicloud_sms_source_servers": sms.DataSourceServers(),

			"huaweicloud_scm_certificates": scm.DataSourceCertificates(),

			"huaweicloud_sfs_file_system": DataSourceSFSFileSystemV2(),
			"huaweicloud_sfs_turbos":      sfs.DataSourceTurbos(),

			"huaweicloud_vbs_backup_policy": dataSourceVBSBackupPolicyV2(),
			"huaweicloud_vbs_backup":        dataSourceVBSBackupV2(),

			"huaweicloud_vpc_bandwidth": eip.DataSourceBandWidth(),
			"huaweicloud_vpc_eip":       eip.DataSourceVpcEip(),
			"huaweicloud_vpc_eips":      eip.DataSourceVpcEips(),

			"huaweicloud_vpc":                    vpc.DataSourceVpcV1(),
			"huaweicloud_vpcs":                   vpc.DataSourceVpcs(),
			"huaweicloud_vpc_ids":                vpc.DataSourceVpcIdsV1(),
			"huaweicloud_vpc_peering_connection": vpc.DataSourceVpcPeeringConnectionV2(),
			"huaweicloud_vpc_route_table":        vpc.DataSourceVPCRouteTable(),
			"huaweicloud_vpc_subnet":             vpc.DataSourceVpcSubnetV1(),
			"huaweicloud_vpc_subnets":            vpc.DataSourceVpcSubnets(),
			"huaweicloud_vpc_subnet_ids":         vpc.DataSourceVpcSubnetIdsV1(),

			"huaweicloud_vpcep_public_services": vpcep.DataSourceVPCEPPublicServices(),

			"huaweicloud_waf_certificate":         waf.DataSourceWafCertificateV1(),
			"huaweicloud_waf_policies":            waf.DataSourceWafPoliciesV1(),
			"huaweicloud_waf_dedicated_instances": waf.DataSourceWafDedicatedInstancesV1(),
			"huaweicloud_waf_reference_tables":    waf.DataSourceWafReferenceTablesV1(),
			"huaweicloud_waf_instance_groups":     waf.DataSourceWafInstanceGroups(),
			"huaweicloud_dws_flavors":             dws.DataSourceDwsFlavors(),

			// Legacy
			"huaweicloud_images_image_v2":        ims.DataSourceImagesImageV2(),
			"huaweicloud_networking_port_v2":     vpc.DataSourceNetworkingPortV2(),
			"huaweicloud_networking_secgroup_v2": DataSourceNetworkingSecGroup(),

			"huaweicloud_kms_key_v1":      dew.DataSourceKmsKey(),
			"huaweicloud_kms_data_key_v1": DataSourceKmsDataKeyV1(),

			"huaweicloud_rds_flavors_v3":     rds.DataSourceRdsFlavor(),
			"huaweicloud_sfs_file_system_v2": DataSourceSFSFileSystemV2(),

			"huaweicloud_vpc_v1":                    vpc.DataSourceVpcV1(),
			"huaweicloud_vpc_ids_v1":                vpc.DataSourceVpcIdsV1(),
			"huaweicloud_vpc_peering_connection_v2": vpc.DataSourceVpcPeeringConnectionV2(),
			"huaweicloud_vpc_subnet_v1":             vpc.DataSourceVpcSubnetV1(),
			"huaweicloud_vpc_subnet_ids_v1":         vpc.DataSourceVpcSubnetIdsV1(),

			"huaweicloud_cce_cluster_v3": cce.DataSourceCCEClusterV3(),
			"huaweicloud_cce_node_v3":    cce.DataSourceCCENodeV3(),

			"huaweicloud_csbs_backup_v1":        dataSourceCSBSBackupV1(),
			"huaweicloud_csbs_backup_policy_v1": dataSourceCSBSBackupPolicyV1(),

			"huaweicloud_dms_product_v1":        dms.DataSourceDmsProduct(),
			"huaweicloud_dms_maintainwindow_v1": dms.DataSourceDmsMaintainWindow(),

			"huaweicloud_vbs_backup_policy_v2": dataSourceVBSBackupPolicyV2(),
			"huaweicloud_vbs_backup_v2":        dataSourceVBSBackupV2(),

			"huaweicloud_antiddos_v1": dataSourceAntiDdosV1(),

			"huaweicloud_dcs_maintainwindow_v1": dcs.DataSourceDcsMaintainWindow(),

			"huaweicloud_dds_flavors_v3":   dds.DataSourceDDSFlavorV3(),
			"huaweicloud_identity_role_v3": iam.DataSourceIdentityRoleV3(),
			"huaweicloud_cdm_flavors_v1":   DataSourceCdmFlavorV1(),

			"huaweicloud_ddm_engines":        ddm.DataSourceDdmEngines(),
			"huaweicloud_ddm_flavors":        ddm.DataSourceDdmFlavors(),
			"huaweicloud_ddm_instance_nodes": ddm.DataSourceDdmInstanceNodes(),
			"huaweicloud_ddm_instances":      ddm.DataSourceDdmInstances(),

			// Deprecated ongoing (without DeprecationMessage), used by other providers
			"huaweicloud_vpc_route":        vpc.DataSourceVpcRouteV2(),
			"huaweicloud_vpc_route_ids":    vpc.DataSourceVpcRouteIdsV2(),
			"huaweicloud_vpc_route_v2":     vpc.DataSourceVpcRouteV2(),
			"huaweicloud_vpc_route_ids_v2": vpc.DataSourceVpcRouteIdsV2(),

			// Deprecated
			"huaweicloud_compute_availability_zones_v2": dataSourceComputeAvailabilityZonesV2(),
			"huaweicloud_networking_network_v2":         dataSourceNetworkingNetworkV2(),
			"huaweicloud_networking_subnet_v2":          dataSourceNetworkingSubnetV2(),
			"huaweicloud_cts_tracker":                   deprecated.DataSourceCTSTrackerV1(),
			"huaweicloud_dcs_az":                        deprecated.DataSourceDcsAZV1(),
			"huaweicloud_dcs_az_v1":                     deprecated.DataSourceDcsAZV1(),
			"huaweicloud_dcs_product":                   deprecated.DataSourceDcsProductV1(),
			"huaweicloud_dcs_product_v1":                deprecated.DataSourceDcsProductV1(),
			"huaweicloud_dms_az":                        deprecated.DataSourceDmsAZ(),
			"huaweicloud_dms_az_v1":                     deprecated.DataSourceDmsAZ(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"huaweicloud_aad_forward_rule": aad.ResourceForwardRule(),

			"huaweicloud_antiddos_basic": antiddos.ResourceCloudNativeAntiDdos(),

			"huaweicloud_aom_alarm_rule":             aom.ResourceAlarmRule(),
			"huaweicloud_aom_service_discovery_rule": aom.ResourceServiceDiscoveryRule(),

			"huaweicloud_rf_stack": rf.ResourceStack(),

			"huaweicloud_api_gateway_api":   ResourceAPIGatewayAPI(),
			"huaweicloud_api_gateway_group": ResourceAPIGatewayGroup(),

			"huaweicloud_apig_api":                         apig.ResourceApigAPIV2(),
			"huaweicloud_apig_api_publishment":             apig.ResourceApigApiPublishment(),
			"huaweicloud_apig_application":                 apig.ResourceApigApplicationV2(),
			"huaweicloud_apig_custom_authorizer":           apig.ResourceApigCustomAuthorizerV2(),
			"huaweicloud_apig_environment":                 apig.ResourceApigEnvironmentV2(),
			"huaweicloud_apig_group":                       apig.ResourceApigGroupV2(),
			"huaweicloud_apig_instance":                    apig.ResourceApigInstanceV2(),
			"huaweicloud_apig_response":                    apig.ResourceApigResponseV2(),
			"huaweicloud_apig_throttling_policy_associate": apig.ResourceThrottlingPolicyAssociate(),
			"huaweicloud_apig_throttling_policy":           apig.ResourceApigThrottlingPolicyV2(),
			"huaweicloud_apig_vpc_channel":                 apig.ResourceApigVpcChannelV2(),

			"huaweicloud_as_configuration":    as.ResourceASConfiguration(),
			"huaweicloud_as_group":            as.ResourceASGroup(),
			"huaweicloud_as_lifecycle_hook":   as.ResourceASLifecycleHook(),
			"huaweicloud_as_notification":     as.ResourceAsNotification(),
			"huaweicloud_as_policy":           as.ResourceASPolicy(),
			"huaweicloud_as_bandwidth_policy": as.ResourceASBandWidthPolicy(),

			"huaweicloud_bms_instance": bms.ResourceBmsInstance(),
			"huaweicloud_bcs_instance": resourceBCSInstanceV2(),

			"huaweicloud_cbr_policy": cbr.ResourceCBRPolicyV3(),
			"huaweicloud_cbr_vault":  cbr.ResourceVault(),

			"huaweicloud_cbh_instance": cbh.ResourceCBHInstance(),

			"huaweicloud_cc_connection":       cc.ResourceCloudConnection(),
			"huaweicloud_cc_network_instance": cc.ResourceNetworkInstance(),

			"huaweicloud_cce_cluster":     cce.ResourceCCEClusterV3(),
			"huaweicloud_cce_node":        cce.ResourceCCENodeV3(),
			"huaweicloud_cce_node_attach": cce.ResourceCCENodeAttachV3(),
			"huaweicloud_cce_addon":       cce.ResourceCCEAddonV3(),
			"huaweicloud_cce_node_pool":   cce.ResourceCCENodePool(),
			"huaweicloud_cce_namespace":   cce.ResourceCCENamespaceV1(),
			"huaweicloud_cce_pvc":         cce.ResourceCcePersistentVolumeClaimsV1(),

			"huaweicloud_cts_tracker":      cts.ResourceCTSTracker(),
			"huaweicloud_cts_data_tracker": cts.ResourceCTSDataTracker(),
			"huaweicloud_cts_notification": cts.ResourceCTSNotification(),
			"huaweicloud_cci_namespace":    cci.ResourceCciNamespace(),
			"huaweicloud_cci_network":      cci.ResourceCciNetworkV1(),
			"huaweicloud_cci_pvc":          ResourceCCIPersistentVolumeClaimV1(),

			"huaweicloud_cdm_cluster": cdm.ResourceCdmCluster(),
			"huaweicloud_cdm_job":     cdm.ResourceCdmJob(),
			"huaweicloud_cdm_link":    cdm.ResourceCdmLink(),

			"huaweicloud_cdn_domain":    resourceCdnDomainV1(),
			"huaweicloud_ces_alarmrule": ces.ResourceAlarmRule(),

			"huaweicloud_cfw_protection_rule": cfw.ResourceProtectionRule(),

			"huaweicloud_cloudtable_cluster": cloudtable.ResourceCloudTableCluster(),

			"huaweicloud_codehub_repository": codehub.ResourceRepository(),

			"huaweicloud_compute_instance":         ResourceComputeInstanceV2(),
			"huaweicloud_compute_interface_attach": ResourceComputeInterfaceAttachV2(),
			"huaweicloud_compute_keypair":          ResourceComputeKeypairV2(),
			"huaweicloud_compute_servergroup":      ResourceComputeServerGroupV2(),
			"huaweicloud_compute_eip_associate":    ResourceComputeFloatingIPAssociateV2(),
			"huaweicloud_compute_volume_attach":    ecs.ResourceComputeVolumeAttach(),

			"huaweicloud_cph_server": cph.ResourceCphServer(),

			"huaweicloud_csbs_backup":        resourceCSBSBackupV1(),
			"huaweicloud_csbs_backup_policy": resourceCSBSBackupPolicyV1(),

			"huaweicloud_cse_microservice":          cse.ResourceMicroservice(),
			"huaweicloud_cse_microservice_engine":   cse.ResourceMicroserviceEngine(),
			"huaweicloud_cse_microservice_instance": cse.ResourceMicroserviceInstance(),

			"huaweicloud_csms_secret": dew.ResourceCsmsSecret(),

			"huaweicloud_css_cluster":   css.ResourceCssCluster(),
			"huaweicloud_css_snapshot":  css.ResourceCssSnapshot(),
			"huaweicloud_css_thesaurus": css.ResourceCssthesaurus(),

			"huaweicloud_dbss_instance": dbss.ResourceInstance(),

			"huaweicloud_dc_virtual_gateway":   dc.ResourceVirtualGateway(),
			"huaweicloud_dc_virtual_interface": dc.ResourceVirtualInterface(),

			"huaweicloud_dcs_instance": dcs.ResourceDcsInstance(),

			"huaweicloud_dds_database_role": dds.ResourceDatabaseRole(),
			"huaweicloud_dds_database_user": dds.ResourceDatabaseUser(),
			"huaweicloud_dds_instance":      dds.ResourceDdsInstanceV3(),

			"huaweicloud_ddm_instance": ddm.ResourceDdmInstance(),
			"huaweicloud_ddm_schema":   ddm.ResourceDdmSchema(),

			"huaweicloud_dis_stream": dis.ResourceDisStream(),

			"huaweicloud_dli_database":     dli.ResourceDliSqlDatabaseV1(),
			"huaweicloud_dli_package":      dli.ResourceDliPackageV2(),
			"huaweicloud_dli_queue":        dli.ResourceDliQueue(),
			"huaweicloud_dli_spark_job":    dli.ResourceDliSparkJobV2(),
			"huaweicloud_dli_sql_job":      dli.ResourceSqlJob(),
			"huaweicloud_dli_table":        dli.ResourceDliTable(),
			"huaweicloud_dli_flinksql_job": dli.ResourceFlinkSqlJob(),
			"huaweicloud_dli_flinkjar_job": dli.ResourceFlinkJarJob(),
			"huaweicloud_dli_permission":   dli.ResourceDliPermission(),

			"huaweicloud_dms_kafka_user":        dms.ResourceDmsKafkaUser(),
			"huaweicloud_dms_kafka_permissions": dms.ResourceDmsKafkaPermissions(),
			"huaweicloud_dms_kafka_instance":    dms.ResourceDmsKafkaInstance(),
			"huaweicloud_dms_kafka_topic":       dms.ResourceDmsKafkaTopic(),
			"huaweicloud_dms_rabbitmq_instance": dms.ResourceDmsRabbitmqInstance(),

			"huaweicloud_dms_rocketmq_instance":       dms.ResourceDmsRocketMQInstance(),
			"huaweicloud_dms_rocketmq_consumer_group": dms.ResourceDmsRocketMQConsumerGroup(),
			"huaweicloud_dms_rocketmq_topic":          dms.ResourceDmsRocketMQTopic(),
			"huaweicloud_dms_rocketmq_user":           dms.ResourceDmsRocketMQUser(),

			"huaweicloud_dns_ptrrecord": ResourceDNSPtrRecordV2(),
			"huaweicloud_dns_recordset": ResourceDNSRecordSetV2(),
			"huaweicloud_dns_zone":      ResourceDNSZoneV2(),

			"huaweicloud_drs_job":     drs.ResourceDrsJob(),
			"huaweicloud_dws_cluster": dws.ResourceDwsCluster(),

			"huaweicloud_elb_certificate":     elb.ResourceCertificateV3(),
			"huaweicloud_elb_l7policy":        elb.ResourceL7PolicyV3(),
			"huaweicloud_elb_l7rule":          elb.ResourceL7RuleV3(),
			"huaweicloud_elb_listener":        elb.ResourceListenerV3(),
			"huaweicloud_elb_loadbalancer":    elb.ResourceLoadBalancerV3(),
			"huaweicloud_elb_monitor":         elb.ResourceMonitorV3(),
			"huaweicloud_elb_ipgroup":         elb.ResourceIpGroupV3(),
			"huaweicloud_elb_pool":            elb.ResourcePoolV3(),
			"huaweicloud_elb_member":          elb.ResourceMemberV3(),
			"huaweicloud_elb_logtank":         elb.ResourceLogTank(),
			"huaweicloud_elb_security_policy": elb.ResourceSecurityPolicy(),

			"huaweicloud_enterprise_project": eps.ResourceEnterpriseProject(),

			"huaweicloud_er_association":    er.ResourceAssociation(),
			"huaweicloud_er_instance":       er.ResourceInstance(),
			"huaweicloud_er_propagation":    er.ResourcePropagation(),
			"huaweicloud_er_route_table":    er.ResourceRouteTable(),
			"huaweicloud_er_vpc_attachment": er.ResourceVpcAttachment(),

			"huaweicloud_evs_snapshot": ResourceEvsSnapshotV2(),
			"huaweicloud_evs_volume":   evs.ResourceEvsVolume(),

			"huaweicloud_fgs_dependency": fgs.ResourceFgsDependency(),
			"huaweicloud_fgs_function":   fgs.ResourceFgsFunctionV2(),
			"huaweicloud_fgs_trigger":    fgs.ResourceFunctionGraphTrigger(),

			"huaweicloud_ga_accelerator":    ga.ResourceAccelerator(),
			"huaweicloud_ga_listener":       ga.ResourceListener(),
			"huaweicloud_ga_endpoint_group": ga.ResourceEndpointGroup(),
			"huaweicloud_ga_endpoint":       ga.ResourceEndpoint(),
			"huaweicloud_ga_health_check":   ga.ResourceHealthCheck(),

			"huaweicloud_gaussdb_cassandra_instance": gaussdb.ResourceGeminiDBInstanceV3(),
			"huaweicloud_gaussdb_mysql_instance":     gaussdb.ResourceGaussDBInstance(),
			"huaweicloud_gaussdb_mysql_proxy":        gaussdb.ResourceGaussDBProxy(),
			"huaweicloud_gaussdb_opengauss_instance": gaussdb.ResourceOpenGaussInstance(),
			"huaweicloud_gaussdb_redis_instance":     gaussdb.ResourceGaussRedisInstanceV3(),
			"huaweicloud_gaussdb_influx_instance":    gaussdb.ResourceGaussDBInfluxInstanceV3(),
			"huaweicloud_gaussdb_mongo_instance":     gaussdb.ResourceGaussDBMongoInstanceV3(),

			"huaweicloud_ges_graph": ResourceGesGraphV1(),

			"huaweicloud_hss_host_group": hss.ResourceHostGroup(),

			"huaweicloud_identity_access_key":       iam.ResourceIdentityKey(),
			"huaweicloud_identity_acl":              iam.ResourceIdentityACL(),
			"huaweicloud_identity_agency":           iam.ResourceIAMAgencyV3(),
			"huaweicloud_identity_group":            iam.ResourceIdentityGroupV3(),
			"huaweicloud_identity_group_membership": iam.ResourceIdentityGroupMembershipV3(),
			"huaweicloud_identity_project":          iam.ResourceIdentityProjectV3(),
			"huaweicloud_identity_role":             iam.ResourceIdentityRole(),
			"huaweicloud_identity_role_assignment":  iam.ResourceIdentityRoleAssignmentV3(),
			"huaweicloud_identity_user":             iam.ResourceIdentityUserV3(),
			"huaweicloud_identity_provider":         iam.ResourceIdentityProvider(),
			"huaweicloud_identity_password_policy":  iam.ResourceIdentityPasswordPolicy(),

			"huaweicloud_iec_eip":                 resourceIecNetworkEip(),
			"huaweicloud_iec_keypair":             resourceIecKeypair(),
			"huaweicloud_iec_network_acl":         resourceIecNetworkACL(),
			"huaweicloud_iec_network_acl_rule":    resourceIecNetworkACLRule(),
			"huaweicloud_iec_security_group":      resourceIecSecurityGroup(),
			"huaweicloud_iec_security_group_rule": resourceIecSecurityGroupRule(),
			"huaweicloud_iec_server":              resourceIecServer(),
			"huaweicloud_iec_vip":                 resourceIecVipV1(),
			"huaweicloud_iec_vpc":                 ResourceIecVpc(),
			"huaweicloud_iec_vpc_subnet":          resourceIecSubnet(),

			"huaweicloud_images_image": ims.ResourceImsImage(),

			"huaweicloud_iotda_space":               iotda.ResourceSpace(),
			"huaweicloud_iotda_product":             iotda.ResourceProduct(),
			"huaweicloud_iotda_device":              iotda.ResourceDevice(),
			"huaweicloud_iotda_device_group":        iotda.ResourceDeviceGroup(),
			"huaweicloud_iotda_dataforwarding_rule": iotda.ResourceDataForwardingRule(),
			"huaweicloud_iotda_amqp":                iotda.ResourceAmqp(),
			"huaweicloud_iotda_device_certificate":  iotda.ResourceDeviceCertificate(),
			"huaweicloud_iotda_device_linkage_rule": iotda.ResourceDeviceLinkageRule(),

			"huaweicloud_kms_key":     dew.ResourceKmsKey(),
			"huaweicloud_kps_keypair": dew.ResourceKeypair(),
			"huaweicloud_kms_grant":   dew.ResourceKmsGrant(),

			"huaweicloud_lb_certificate":  lb.ResourceCertificateV2(),
			"huaweicloud_lb_l7policy":     lb.ResourceL7PolicyV2(),
			"huaweicloud_lb_l7rule":       lb.ResourceL7RuleV2(),
			"huaweicloud_lb_listener":     lb.ResourceListenerV2(),
			"huaweicloud_lb_loadbalancer": lb.ResourceLoadBalancerV2(),
			"huaweicloud_lb_member":       lb.ResourceMemberV2(),
			"huaweicloud_lb_monitor":      lb.ResourceMonitorV2(),
			"huaweicloud_lb_pool":         lb.ResourcePoolV2(),
			"huaweicloud_lb_whitelist":    lb.ResourceWhitelistV2(),

			"huaweicloud_live_domain":          live.ResourceDomain(),
			"huaweicloud_live_recording":       live.ResourceRecording(),
			"huaweicloud_live_record_callback": live.ResourceRecordCallback(),
			"huaweicloud_live_transcoding":     live.ResourceTranscoding(),

			"huaweicloud_lts_group":  ResourceLTSGroupV2(),
			"huaweicloud_lts_stream": ResourceLTSStreamV2(),

			"huaweicloud_mapreduce_cluster": mrs.ResourceMRSClusterV2(),
			"huaweicloud_mapreduce_job":     mrs.ResourceMRSJobV2(),

			"huaweicloud_meeting_admin_assignment": meeting.ResourceAdminAssignment(),
			"huaweicloud_meeting_conference":       meeting.ResourceConference(),
			"huaweicloud_meeting_user":             meeting.ResourceUser(),

			"huaweicloud_modelarts_dataset":                modelarts.ResourceDataset(),
			"huaweicloud_modelarts_dataset_version":        modelarts.ResourceDatasetVersion(),
			"huaweicloud_modelarts_notebook":               modelarts.ResourceNotebook(),
			"huaweicloud_modelarts_notebook_mount_storage": modelarts.ResourceNotebookMountStorage(),

			"huaweicloud_dataarts_studio_instance": dataarts.ResourceStudioInstance(),

			"huaweicloud_mpc_transcoding_template":       mpc.ResourceTranscodingTemplate(),
			"huaweicloud_mpc_transcoding_template_group": mpc.ResourceTranscodingTemplateGroup(),

			"huaweicloud_mrs_cluster": ResourceMRSClusterV1(),
			"huaweicloud_mrs_job":     ResourceMRSJobV1(),

			"huaweicloud_nat_dnat_rule": nat.ResourcePublicDnatRule(),
			"huaweicloud_nat_gateway":   nat.ResourcePublicGateway(),
			"huaweicloud_nat_snat_rule": nat.ResourcePublicSnatRule(),

			"huaweicloud_network_acl":              ResourceNetworkACL(),
			"huaweicloud_network_acl_rule":         ResourceNetworkACLRule(),
			"huaweicloud_networking_port":          ResourceNetworkingPortV2(),
			"huaweicloud_networking_secgroup":      ResourceNetworkingSecGroup(),
			"huaweicloud_networking_secgroup_rule": ResourceNetworkingSecGroupRule(),
			"huaweicloud_networking_vip":           vpc.ResourceNetworkingVip(),
			"huaweicloud_networking_vip_associate": vpc.ResourceNetworkingVIPAssociateV2(),

			"huaweicloud_obs_bucket":        obs.ResourceObsBucket(),
			"huaweicloud_obs_bucket_object": obs.ResourceObsBucketObject(),
			"huaweicloud_obs_bucket_policy": obs.ResourceObsBucketPolicy(),

			"huaweicloud_oms_migration_task": oms.ResourceMigrationTask(),

			"huaweicloud_rds_account":               rds.ResourceRdsAccount(),
			"huaweicloud_rds_database":              rds.ResourceRdsDatabase(),
			"huaweicloud_rds_database_privilege":    rds.ResourceRdsDatabasePrivilege(),
			"huaweicloud_rds_instance":              rds.ResourceRdsInstance(),
			"huaweicloud_rds_parametergroup":        rds.ResourceRdsConfiguration(),
			"huaweicloud_rds_read_replica_instance": rds.ResourceRdsReadReplicaInstance(),
			"huaweicloud_rds_backup":                rds.ResourceBackup(),

			"huaweicloud_rms_policy_assignment": rms.ResourcePolicyAssignment(),

			"huaweicloud_servicestage_application":                 servicestage.ResourceApplication(),
			"huaweicloud_servicestage_component_instance":          servicestage.ResourceComponentInstance(),
			"huaweicloud_servicestage_component":                   servicestage.ResourceComponent(),
			"huaweicloud_servicestage_environment":                 servicestage.ResourceEnvironment(),
			"huaweicloud_servicestage_repo_token_authorization":    servicestage.ResourceRepoTokenAuth(),
			"huaweicloud_servicestage_repo_password_authorization": servicestage.ResourceRepoPwdAuth(),

			"huaweicloud_servicestage_v3_application":                 servicestagev3.ResourceApplication(),
			"huaweicloud_servicestage_v3_component":                   servicestagev3.ResourceComponent(),
			"huaweicloud_servicestage_v3_environment":                 servicestagev3.ResourceEnvironment(),
			"huaweicloud_servicestage_v3_repo_token_authorization":    servicestagev3.ResourceRepoTokenAuth(),
			"huaweicloud_servicestage_v3_repo_password_authorization": servicestagev3.ResourceRepoPwdAuth(),

			"huaweicloud_sfs_access_rule": ResourceSFSAccessRuleV2(),
			"huaweicloud_sfs_file_system": ResourceSFSFileSystemV2(),
			"huaweicloud_sfs_turbo":       ResourceSFSTurbo(),

			"huaweicloud_smn_topic":        smn.ResourceTopic(),
			"huaweicloud_smn_subscription": smn.ResourceSubscription(),

			"huaweicloud_sms_server_template": sms.ResourceServerTemplate(),
			"huaweicloud_sms_task":            sms.ResourceMigrateTask(),

			"huaweicloud_swr_organization":             swr.ResourceSWROrganization(),
			"huaweicloud_swr_organization_permissions": swr.ResourceSWROrganizationPermissions(),
			"huaweicloud_swr_repository":               swr.ResourceSWRRepository(),
			"huaweicloud_swr_repository_sharing":       swr.ResourceSWRRepositorySharing(),

			"huaweicloud_tms_tags": tms.ResourceTmsTag(),

			"huaweicloud_vbs_backup":        resourceVBSBackupV2(),
			"huaweicloud_vbs_backup_policy": resourceVBSBackupPolicyV2(),

			"huaweicloud_vod_media_asset":                vod.ResourceMediaAsset(),
			"huaweicloud_vod_media_category":             vod.ResourceMediaCategory(),
			"huaweicloud_vod_transcoding_template_group": vod.ResourceTranscodingTemplateGroup(),
			"huaweicloud_vod_watermark_template":         vod.ResourceWatermarkTemplate(),

			"huaweicloud_vpc_bandwidth":     eip.ResourceVpcBandWidthV2(),
			"huaweicloud_vpc_eip":           eip.ResourceVpcEIPV1(),
			"huaweicloud_vpc_eip_associate": eip.ResourceEIPAssociate(),

			"huaweicloud_vpc_peering_connection":          vpc.ResourceVpcPeeringConnectionV2(),
			"huaweicloud_vpc_peering_connection_accepter": vpc.ResourceVpcPeeringConnectionAccepterV2(),
			"huaweicloud_vpc_route_table":                 vpc.ResourceVPCRouteTable(),
			"huaweicloud_vpc_route":                       vpc.ResourceVPCRouteTableRoute(),
			"huaweicloud_vpc":                             vpc.ResourceVirtualPrivateCloudV1(),
			"huaweicloud_vpc_subnet":                      vpc.ResourceVpcSubnetV1(),
			"huaweicloud_vpc_address_group":               vpc.ResourceVpcAddressGroup(),
			"huaweicloud_vpc_flow_log":                    vpc.ResourceVpcFlowLog(),

			"huaweicloud_vpcep_approval": vpcep.ResourceVPCEndpointApproval(),
			"huaweicloud_vpcep_endpoint": vpcep.ResourceVPCEndpoint(),
			"huaweicloud_vpcep_service":  vpcep.ResourceVPCEndpointService(),

			"huaweicloud_vpn_gateway":          vpn.ResourceGateway(),
			"huaweicloud_vpn_customer_gateway": vpn.ResourceCustomerGateway(),
			"huaweicloud_vpn_connection":       vpn.ResourceConnection(),

			"huaweicloud_scm_certificate": scm.ResourceScmCertificate(),

			"huaweicloud_waf_certificate":                waf.ResourceWafCertificateV1(),
			"huaweicloud_waf_cloud_instance":             waf.ResourceCloudInstance(),
			"huaweicloud_waf_domain":                     waf.ResourceWafDomainV1(),
			"huaweicloud_waf_policy":                     waf.ResourceWafPolicyV1(),
			"huaweicloud_waf_rule_blacklist":             waf.ResourceWafRuleBlackListV1(),
			"huaweicloud_waf_rule_data_masking":          waf.ResourceWafRuleDataMaskingV1(),
			"huaweicloud_waf_rule_web_tamper_protection": waf.ResourceWafRuleWebTamperProtectionV1(),
			"huaweicloud_waf_dedicated_instance":         waf.ResourceWafDedicatedInstance(),
			"huaweicloud_waf_dedicated_domain":           waf.ResourceWafDedicatedDomainV1(),
			"huaweicloud_waf_instance_group":             waf.ResourceWafInstanceGroup(),
			"huaweicloud_waf_instance_group_associate":   waf.ResourceWafInstGroupAssociate(),
			"huaweicloud_waf_reference_table":            waf.ResourceWafReferenceTableV1(),

			"huaweicloud_workspace_desktop": workspace.ResourceDesktop(),
			"huaweicloud_workspace_service": workspace.ResourceService(),
			"huaweicloud_workspace_user":    workspace.ResourceUser(),

			"huaweicloud_cpts_project": cpts.ResourceProject(),
			"huaweicloud_cpts_task":    cpts.ResourceTask(),

			// devCloud
			"huaweicloud_projectman_project": projectman.ResourceProject(),

			"huaweicloud_dsc_instance":  dsc.ResourceDscInstance(),
			"huaweicloud_dsc_asset_obs": dsc.ResourceAssetObs(),

			// internal only
			"huaweicloud_apm_aksk":                apm.ResourceApmAkSk(),
			"huaweicloud_aom_alarm_policy":        aom.ResourceAlarmPolicy(),
			"huaweicloud_aom_prometheus_instance": aom.ResourcePrometheusInstance(),

			"huaweicloud_aom_application":                 cmdb.ResourceAomApplication(),
			"huaweicloud_aom_component":                   cmdb.ResourceAomComponent(),
			"huaweicloud_aom_cmdb_resource_relationships": cmdb.ResourceCiRelationships(),
			"huaweicloud_aom_environment":                 cmdb.ResourceAomEnvironment(),

			"huaweicloud_lts_access_rule":     lts.ResourceAomMappingRule(),
			"huaweicloud_lts_dashboard":       lts.ResourceLtsDashboard(),
			"huaweicloud_elb_log":             lts.ResourceLtsElb(),
			"huaweicloud_lts_struct_template": lts.ResourceLtsStruct(),

			// Legacy
			"huaweicloud_networking_eip_associate": eip.ResourceEIPAssociate(),

			"huaweicloud_compute_instance_v2":         ResourceComputeInstanceV2(),
			"huaweicloud_compute_interface_attach_v2": ResourceComputeInterfaceAttachV2(),
			"huaweicloud_compute_keypair_v2":          ResourceComputeKeypairV2(),
			"huaweicloud_compute_servergroup_v2":      ResourceComputeServerGroupV2(),
			"huaweicloud_compute_volume_attach_v2":    ecs.ResourceComputeVolumeAttach(),

			"huaweicloud_dns_ptrrecord_v2": ResourceDNSPtrRecordV2(),
			"huaweicloud_dns_recordset_v2": ResourceDNSRecordSetV2(),
			"huaweicloud_dns_zone_v2":      ResourceDNSZoneV2(),

			"huaweicloud_dcs_instance_v1": dcs.ResourceDcsInstance(),
			"huaweicloud_dds_instance_v3": dds.ResourceDdsInstanceV3(),

			"huaweicloud_fw_firewall_group_v2": resourceFWFirewallGroupV2(),
			"huaweicloud_fw_policy_v2":         resourceFWPolicyV2(),
			"huaweicloud_fw_rule_v2":           resourceFWRuleV2(),

			"huaweicloud_kms_key_v1": dew.ResourceKmsKey(),

			"huaweicloud_lb_certificate_v2":  lb.ResourceCertificateV2(),
			"huaweicloud_lb_loadbalancer_v2": lb.ResourceLoadBalancerV2(),
			"huaweicloud_lb_listener_v2":     lb.ResourceListenerV2(),
			"huaweicloud_lb_pool_v2":         lb.ResourcePoolV2(),
			"huaweicloud_lb_member_v2":       lb.ResourceMemberV2(),
			"huaweicloud_lb_monitor_v2":      lb.ResourceMonitorV2(),
			"huaweicloud_lb_l7policy_v2":     lb.ResourceL7PolicyV2(),
			"huaweicloud_lb_l7rule_v2":       lb.ResourceL7RuleV2(),
			"huaweicloud_lb_whitelist_v2":    lb.ResourceWhitelistV2(),

			"huaweicloud_mrs_cluster_v1": ResourceMRSClusterV1(),
			"huaweicloud_mrs_job_v1":     ResourceMRSJobV1(),

			"huaweicloud_networking_port_v2":          ResourceNetworkingPortV2(),
			"huaweicloud_networking_secgroup_v2":      ResourceNetworkingSecGroup(),
			"huaweicloud_networking_secgroup_rule_v2": ResourceNetworkingSecGroupRule(),

			"huaweicloud_smn_topic_v2":        smn.ResourceTopic(),
			"huaweicloud_smn_subscription_v2": smn.ResourceSubscription(),

			"huaweicloud_rds_instance_v3":       rds.ResourceRdsInstance(),
			"huaweicloud_rds_parametergroup_v3": rds.ResourceRdsConfiguration(),

			"huaweicloud_nat_dnat_rule_v2": nat.ResourcePublicDnatRule(),
			"huaweicloud_nat_gateway_v2":   nat.ResourcePublicGateway(),
			"huaweicloud_nat_snat_rule_v2": nat.ResourcePublicSnatRule(),

			"huaweicloud_sfs_access_rule_v2": ResourceSFSAccessRuleV2(),
			"huaweicloud_sfs_file_system_v2": ResourceSFSFileSystemV2(),

			"huaweicloud_iam_agency":    iam.ResourceIAMAgencyV3(),
			"huaweicloud_iam_agency_v3": iam.ResourceIAMAgencyV3(),

			"huaweicloud_vpc_bandwidth_v2":                   eip.ResourceVpcBandWidthV2(),
			"huaweicloud_vpc_eip_v1":                         eip.ResourceVpcEIPV1(),
			"huaweicloud_vpc_peering_connection_v2":          vpc.ResourceVpcPeeringConnectionV2(),
			"huaweicloud_vpc_peering_connection_accepter_v2": vpc.ResourceVpcPeeringConnectionAccepterV2(),
			"huaweicloud_vpc_route_v2":                       vpc.ResourceVPCRouteV2(),
			"huaweicloud_vpc_v1":                             vpc.ResourceVirtualPrivateCloudV1(),
			"huaweicloud_vpc_subnet_v1":                      vpc.ResourceVpcSubnetV1(),

			"huaweicloud_cce_cluster_v3": cce.ResourceCCEClusterV3(),
			"huaweicloud_cce_node_v3":    cce.ResourceCCENodeV3(),

			"huaweicloud_as_configuration_v1": as.ResourceASConfiguration(),
			"huaweicloud_as_group_v1":         as.ResourceASGroup(),
			"huaweicloud_as_policy_v1":        as.ResourceASPolicy(),

			"huaweicloud_csbs_backup_v1":        resourceCSBSBackupV1(),
			"huaweicloud_csbs_backup_policy_v1": resourceCSBSBackupPolicyV1(),
			"huaweicloud_vbs_backup_policy_v2":  resourceVBSBackupPolicyV2(),
			"huaweicloud_vbs_backup_v2":         resourceVBSBackupV2(),

			"huaweicloud_maas_task":    resourceMaasTaskV1(),
			"huaweicloud_maas_task_v1": resourceMaasTaskV1(),

			"huaweicloud_identity_project_v3":          iam.ResourceIdentityProjectV3(),
			"huaweicloud_identity_role_assignment_v3":  iam.ResourceIdentityRoleAssignmentV3(),
			"huaweicloud_identity_user_v3":             iam.ResourceIdentityUserV3(),
			"huaweicloud_identity_group_v3":            iam.ResourceIdentityGroupV3(),
			"huaweicloud_identity_group_membership_v3": iam.ResourceIdentityGroupMembershipV3(),
			"huaweicloud_identity_provider_conversion": iam.ResourceIAMProviderConversion(),

			"huaweicloud_cdm_cluster_v1": cdm.ResourceCdmCluster(),
			"huaweicloud_ges_graph_v1":   ResourceGesGraphV1(),
			"huaweicloud_css_cluster_v1": css.ResourceCssCluster(),
			"huaweicloud_dis_stream_v2":  dis.ResourceDisStream(),

			"huaweicloud_dli_queue_v1":                dli.ResourceDliQueue(),
			"huaweicloud_networking_vip_v2":           vpc.ResourceNetworkingVip(),
			"huaweicloud_networking_vip_associate_v2": vpc.ResourceNetworkingVIPAssociateV2(),
			"huaweicloud_fgs_function_v2":             fgs.ResourceFgsFunctionV2(),
			"huaweicloud_cdn_domain_v1":               resourceCdnDomainV1(),

			// Deprecated
			"huaweicloud_blockstorage_volume_v2":          resourceBlockStorageVolumeV2(),
			"huaweicloud_networking_network_v2":           resourceNetworkingNetworkV2(),
			"huaweicloud_networking_subnet_v2":            resourceNetworkingSubnetV2(),
			"huaweicloud_networking_floatingip_v2":        resourceNetworkingFloatingIPV2(),
			"huaweicloud_networking_router_v2":            resourceNetworkingRouterV2(),
			"huaweicloud_networking_router_interface_v2":  resourceNetworkingRouterInterfaceV2(),
			"huaweicloud_networking_router_route_v2":      resourceNetworkingRouterRouteV2(),
			"huaweicloud_ecs_instance_v1":                 resourceEcsInstanceV1(),
			"huaweicloud_compute_secgroup_v2":             ResourceComputeSecGroupV2(),
			"huaweicloud_compute_floatingip_v2":           ResourceComputeFloatingIPV2(),
			"huaweicloud_compute_floatingip_associate_v2": ResourceComputeFloatingIPAssociateV2(),
			"huaweicloud_oms_task":                        resourceMaasTaskV1(),

			"huaweicloud_images_image_v2": deprecated.ResourceImagesImageV2(),

			"huaweicloud_dms_instance":    deprecated.ResourceDmsInstancesV1(),
			"huaweicloud_dms_instance_v1": deprecated.ResourceDmsInstancesV1(),
			"huaweicloud_dms_group":       deprecated.ResourceDmsGroups(),
			"huaweicloud_dms_group_v1":    deprecated.ResourceDmsGroups(),
			"huaweicloud_dms_queue":       deprecated.ResourceDmsQueues(),
			"huaweicloud_dms_queue_v1":    deprecated.ResourceDmsQueues(),

			"huaweicloud_cs_cluster":            deprecated.ResourceCsClusterV1(),
			"huaweicloud_cs_cluster_v1":         deprecated.ResourceCsClusterV1(),
			"huaweicloud_cs_route":              deprecated.ResourceCsRouteV1(),
			"huaweicloud_cs_route_v1":           deprecated.ResourceCsRouteV1(),
			"huaweicloud_cs_peering_connect":    deprecated.ResourceCsPeeringConnectV1(),
			"huaweicloud_cs_peering_connect_v1": deprecated.ResourceCsPeeringConnectV1(),

			"huaweicloud_vpnaas_ipsec_policy_v2":    deprecated.ResourceVpnIPSecPolicyV2(),
			"huaweicloud_vpnaas_service_v2":         deprecated.ResourceVpnServiceV2(),
			"huaweicloud_vpnaas_ike_policy_v2":      deprecated.ResourceVpnIKEPolicyV2(),
			"huaweicloud_vpnaas_endpoint_group_v2":  deprecated.ResourceVpnEndpointGroupV2(),
			"huaweicloud_vpnaas_site_connection_v2": deprecated.ResourceVpnSiteConnectionV2(),

			"huaweicloud_vpnaas_endpoint_group":  deprecated.ResourceVpnEndpointGroupV2(),
			"huaweicloud_vpnaas_ike_policy":      deprecated.ResourceVpnIKEPolicyV2(),
			"huaweicloud_vpnaas_ipsec_policy":    deprecated.ResourceVpnIPSecPolicyV2(),
			"huaweicloud_vpnaas_service":         deprecated.ResourceVpnServiceV2(),
			"huaweicloud_vpnaas_site_connection": deprecated.ResourceVpnSiteConnectionV2(),
		},
	}

	provider.ConfigureContextFunc = func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		terraformVersion := provider.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11 cc
			terraformVersion = "0.11+compatible"
		}

		return configureProvider(ctx, d, terraformVersion)
	}

	return provider
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"auth_url": "The Identity authentication URL.",

		"region": "The HuaweiCloud region to connect to.",

		"user_name": "Username to login with.",

		"user_id": "User ID to login with.",

		"project_id": "The ID of the project to login with.",

		"project_name": "The name of the project to login with.",

		"tenant_id": "The ID of the Tenant (Identity v2) to login with.",

		"tenant_name": "The name of the Tenant (Identity v2) to login with.",

		"password": "Password to login with.",

		"token": "Authentication token to use as an alternative to username/password.",

		"domain_id": "The ID of the Domain to scope to.",

		"domain_name": "The name of the Domain to scope to.",

		"access_key":     "The access key of the HuaweiCloud to use.",
		"secret_key":     "The secret key of the HuaweiCloud to use.",
		"security_token": "The security token to authenticate with a temporary security credential.",

		"insecure": "Trust self-signed certificates.",

		"cacert_file": "A Custom CA certificate.",

		"cert": "A client certificate to authenticate with.",

		"key": "A client private key to authenticate with.",

		"agency_name": "The name of agency",

		"agency_domain_name": "The name of domain who created the agency (Identity v3).",

		"delegated_project": "The name of delegated project (Identity v3).",

		"assume_role_agency_name": "The name of agency for assume role.",

		"assume_role_domain_name": "The name of domain for assume role.",

		"cloud": "The endpoint of cloud provider, defaults to myhuaweicloud.com",

		"endpoints": "The custom endpoints used to override the default endpoint URL.",

		"shared_config_file": "The path to the shared config file. If not set, the default is ~/.hcloud/config.json.",

		"profile": "The profile name as set in the shared config file.",

		"max_retries": "How many times HTTP connection should be retried until giving up.",

		"enterprise_project_id": "enterprise project id",
	}
}

func configureProvider(_ context.Context, d *schema.ResourceData, terraformVersion string) (interface{},
	diag.Diagnostics) {
	var tenantName, tenantID, delegatedProject, identityEndpoint string
	region := d.Get("region").(string)
	cloud := getCloudDomain(d.Get("cloud").(string), region)

	// project_name is prior to tenant_name
	// if neither of them was set, use region as the default project
	if v, ok := d.GetOk("project_name"); ok && v.(string) != "" {
		tenantName = v.(string)
	} else if v, ok := d.GetOk("tenant_name"); ok && v.(string) != "" {
		tenantName = v.(string)
	} else {
		tenantName = region
	}

	// project_id is prior to tenant_id
	if v, ok := d.GetOk("project_id"); ok && v.(string) != "" {
		tenantID = v.(string)
	} else {
		tenantID = d.Get("tenant_id").(string)
	}

	// Use region as delegated_project if it's not set
	if v, ok := d.GetOk("delegated_project"); ok && v.(string) != "" {
		delegatedProject = v.(string)
	} else {
		delegatedProject = region
	}

	// use auth_url as identityEndpoint if provided
	if v, ok := d.GetOk("auth_url"); ok {
		identityEndpoint = v.(string)
	} else {
		// use cloud as basis for identityEndpoint
		if isDefaultHWCloudDomain(cloud) {
			identityEndpoint = fmt.Sprintf("https://iam.%s:443/v3", cloud)
		} else {
			identityEndpoint = fmt.Sprintf("https://iam.%s.%s:443/v3", region, cloud)
		}
	}

	config := config.Config{
		AccessKey:           d.Get("access_key").(string),
		SecretKey:           d.Get("secret_key").(string),
		CACertFile:          d.Get("cacert_file").(string),
		ClientCertFile:      d.Get("cert").(string),
		ClientKeyFile:       d.Get("key").(string),
		DomainID:            d.Get("domain_id").(string),
		DomainName:          d.Get("domain_name").(string),
		IdentityEndpoint:    identityEndpoint,
		Insecure:            d.Get("insecure").(bool),
		Password:            d.Get("password").(string),
		Token:               d.Get("token").(string),
		SecurityToken:       d.Get("security_token").(string),
		Region:              region,
		TenantID:            tenantID,
		TenantName:          tenantName,
		Username:            d.Get("user_name").(string),
		UserID:              d.Get("user_id").(string),
		AgencyName:          d.Get("agency_name").(string),
		AgencyDomainName:    d.Get("agency_domain_name").(string),
		DelegatedProject:    delegatedProject,
		Cloud:               cloud,
		MaxRetries:          d.Get("max_retries").(int),
		EnterpriseProjectID: d.Get("enterprise_project_id").(string),
		SharedConfigFile:    d.Get("shared_config_file").(string),
		Profile:             d.Get("profile").(string),
		TerraformVersion:    terraformVersion,
		RegionProjectIDMap:  make(map[string]string),
		RPLock:              new(sync.Mutex),
		SecurityKeyLock:     new(sync.Mutex),
	}

	// get assume role
	assumeRoleList := d.Get("assume_role").([]interface{})
	if len(assumeRoleList) == 1 {
		assumeRole := assumeRoleList[0].(map[string]interface{})
		config.AssumeRoleAgency = assumeRole["agency_name"].(string)
		config.AssumeRoleDomain = assumeRole["domain_name"].(string)
	}

	// get custom endpoints
	endpoints, err := flattenProviderEndpoints(d)
	if err != nil {
		return nil, diag.FromErr(err)
	}
	config.Endpoints = endpoints

	if err := config.LoadAndValidate(); err != nil {
		return nil, diag.FromErr(err)
	}

	return &config, nil
}

func flattenProviderEndpoints(d *schema.ResourceData) (map[string]string, error) {
	endpoints := d.Get("endpoints").(map[string]interface{})
	epMap := make(map[string]string)

	for key, val := range endpoints {
		endpoint := strings.TrimSpace(val.(string))
		// check empty string
		if endpoint == "" {
			return nil, fmt.Errorf("the value of customer endpoint %s must be specified", key)
		}

		// add prefix "https://" and suffix "/"
		if !strings.HasPrefix(endpoint, "http") {
			endpoint = fmt.Sprintf("https://%s", endpoint)
		}
		if !strings.HasSuffix(endpoint, "/") {
			endpoint = fmt.Sprintf("%s/", endpoint)
		}
		epMap[key] = endpoint
	}

	// unify the endpoint which has multiple versions
	for key := range endpoints {
		ep, ok := epMap[key]
		if !ok {
			continue
		}

		multiKeys := config.GetServiceDerivedCatalogKeys(key)
		for _, k := range multiKeys {
			epMap[k] = ep
		}
	}

	log.Printf("[DEBUG] customer endpoints: %+v", epMap)
	return epMap, nil
}

func getCloudDomain(cloud, region string) string {
	// the regions are named as eu-west-1xx in Europe
	if cloud == defaultCloud && strings.HasPrefix(region, "eu-west-1") {
		return defaultEuropeCloud
	}
	return cloud
}

func isDefaultHWCloudDomain(domain string) bool {
	if domain == defaultCloud || domain == defaultEuropeCloud {
		return true
	}

	return false
}
